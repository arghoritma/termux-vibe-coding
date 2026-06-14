package tui

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rozaq/termux-vibe-coding/internal/agent"
	"github.com/rozaq/termux-vibe-coding/internal/ai"
	"github.com/rozaq/termux-vibe-coding/internal/config"
)

// session stores one chat conversation.
type session struct {
	title    string
	messages []chatMsg
}

type chatMsg struct {
	role    string // "user", "agent", "tool", "system"
	content string
}

// AgentMsg is sent from the agent goroutine back to the TUI.
type AgentMsg struct {
	text   string
	err    error
	done   bool
	total  ai.Usage
}

// ApproveRequest is sent when a dangerous tool needs user approval.
type ApproveRequest struct {
	name    string
	args    string
	respond chan<- bool
}

// eventCh is the shared channel between agent goroutine and TUI.
type eventCh struct {
	agent    chan AgentMsg
	approval chan ApproveRequest
}

// Model is the Bubbletea model.
type Model struct {
	cfg    *config.Config
	ctx    context.Context
	cancel context.CancelFunc

	// state
	provider ai.Provider
	agent    *agent.Agent

	// UI state
	ready         bool
	width         int
	height        int
	theme         Theme
	s             styles
	spinner       spinner.Model
	loading       bool
	thinkingDots  int

	// chat
	viewport      viewport.Model
	input         textinput.Model
	messages      []chatMsg

	// sessions
	sessions      []session
	currentSess   int

	// menus
	menuOpen      bool
	menuItems     []menuItem
	menuSelected  int
	menuTitle     string
	submenuOpen   bool
	submenuItems  []menuItem
	submenuTitle  string

	// approval
	awaitingApproval bool
	approvalName     string
	approvalArgs     string
	approvalResp     chan<- bool

	// events
	events        eventCh
}

type menuItem struct {
	label string
	desc  string
	action func(*Model) tea.Cmd
}

// NewModel creates the initial TUI model.
func NewModel(cfg *config.Config) (*Model, error) {
	prov, err := ai.Build(cfg, cfg.Provider)
	if err != nil {
		return nil, err
	}
	_, modelName := cfg.Current()

	ctx, cancel := context.WithCancel(context.Background())

	sp := spinner.New()
	sp.Spinner = spinner.Dot

	ti := textinput.New()
	ti.Placeholder = "Chat di sini... (ketik / untuk menu, Ctrl+M menu utama)"
	ti.Focus()
	ti.CharLimit = 2000
	ti.Width = 60

	return &Model{
		cfg:    cfg,
		ctx:    ctx,
		cancel: cancel,

		provider: prov,
		agent:    agent.New(cfg, prov, modelName, "."),

		spinner:  sp,
		input:    ti,
		messages: []chatMsg{},

		sessions:    []session{{title: "Sesi 1"}},
		currentSess: 0,

		events: eventCh{
			agent:    make(chan AgentMsg, 100),
			approval: make(chan ApproveRequest, 5),
		},
	}, nil
}

// Init implements tea.Model.
func (m *Model) Init() tea.Cmd {
	m.theme = themeByName(m.cfg.TUI.Theme)
	m.s = newStyles(m.theme)

	m.spinner.Style = lipgloss.NewStyle().Foreground(m.theme.Primary)

	// Welcome message
	welcome := fmt.Sprintf("✨ Selamat datang di %s!\nKetik / untuk menu cepat, Ctrl+M menu utama, atau langsung chat mau bikin apa~", m.cfg.Provider)
	m.appendMsg("system", welcome)

	return tea.Batch(m.spinner.Tick, textinput.Blink, m.waitForAgent())
}

// appendMsg adds a chat message and scrolls down.
func (m *Model) appendMsg(role, content string) {
	m.messages = append(m.messages, chatMsg{role: role, content: content})
	m.sessions[m.currentSess].messages = append(m.sessions[m.currentSess].messages, chatMsg{role: role, content: content})
	if m.ready {
		m.viewport.SetContent(m.renderMessages())
		m.viewport.GotoBottom()
	}
}

// waitForAgent returns a command that listens for agent events.
func (m *Model) waitForAgent() tea.Cmd {
	return func() tea.Msg {
		select {
		case msg := <-m.events.agent:
			return msg
		case req := <-m.events.approval:
			return req
		default:
			return nil
		}
	}
}

// Update implements tea.Model.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true

		headerH := 2
		inputH := 3
		statusH := 1
		vpHeight := m.height - headerH - inputH - statusH - 2
		if vpHeight < 10 {
			vpHeight = 10
		}

		m.viewport = viewport.New(m.width-4, vpHeight)
		m.viewport.SetContent(m.renderMessages())

		m.input.Width = m.width - 10

		return m, nil

	case tea.KeyMsg:
		return m.handleKeyMsg(msg)

	case spinner.TickMsg:
		if m.loading {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}
		return m, nil

	case AgentMsg:
		m.loading = false
		if msg.text != "" {
			m.appendMsg("agent", msg.text)
		}
		if msg.err != nil {
			m.appendMsg("tool", fmt.Sprintf("⚠️ Error: %v", msg.err))
		}
		if msg.done {
			m.loading = false
		}
		return m, m.waitForAgent()

	case ApproveRequest:
		m.awaitingApproval = true
		m.approvalName = msg.name
		m.approvalArgs = msg.args
		m.approvalResp = msg.respond
		content := fmt.Sprintf("⚡ Izinkan tool: %s\nArgs: %s\n[Y]es / [N]o?", msg.name, msg.args)
		m.appendMsg("system", content)
		return m, nil
	}

	return m, nil
}

// handleKeyMsg processes keyboard input.
func (m *Model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Global keybindings
	switch msg.Type {
	case tea.KeyCtrlC:
		m.cancel()
		return m, tea.Quit

	case tea.KeyEsc:
		if m.submenuOpen {
			m.submenuOpen = false
			m.submenuItems = nil
			return m, nil
		}
		if m.menuOpen {
			m.menuOpen = false
			return m, nil
		}
		return m, nil

	case tea.KeyTab:
		// Auto-complete: not fully implemented yet — just toggle input focus
		return m, nil
	}

	// If menu is open, handle menu navigation
	if m.menuOpen {
		return m.handleMenuNav(msg)
	}
	if m.submenuOpen {
		return m.handleSubmenuNav(msg)
	}

	// Approval mode
	if m.awaitingApproval {
		switch msg.String() {
		case "y", "Y":
			m.awaitingApproval = false
			m.approvalResp <- true
			close(m.approvalResp)
			m.appendMsg("system", "✅ Diizinkan. Melanjutkan...")
			return m, m.waitForAgent()
		case "n", "N":
			m.awaitingApproval = false
			m.approvalResp <- false
			close(m.approvalResp)
			m.appendMsg("system", "❌ Dibatalkan.")
			return m, m.waitForAgent()
		}
		return m, nil
	}

	// Ctrl+M — toggle main menu
	if msg.Type == tea.KeyCtrlM {
		m.toggleMainMenu()
		return m, nil
	}

	// Input handling
	switch msg.Type {
	case tea.KeyEnter:
		text := m.input.Value()
		if text == "" {
			return m, nil
		}
		m.input.SetValue("")
		m.input.Blur()

		// Check for slash command
		if strings.HasPrefix(text, "/") {
			cmd := m.handleSlashCommand(text)
			m.input.Focus()
			return m, cmd
		}

		// Normal chat — send to agent
		m.appendMsg("user", text)
		m.loading = true
		go m.runAgent(text)
		m.input.Focus()
		return m, tea.Batch(m.spinner.Tick, m.waitForAgent())

	case tea.KeyCtrlS:
		m.appendMsg("system", "💾 Sesi disimpan (local).")
		return m, nil

	case tea.KeyCtrlL:
		m.appendMsg("system", "📂 Muat sesi... (belum diimplementasi penuh)")
		return m, nil

	case tea.KeyCtrlT:
		m.newSession()
		return m, nil
	}

	// Update input
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

// runAgent starts the agent loop in a goroutine.
func (m *Model) runAgent(text string) {
	h := agent.Handler{
		OnText: func(t string) {
			m.events.agent <- AgentMsg{text: t}
		},
		OnTool: func(name, args string) {
			m.events.agent <- AgentMsg{text: fmt.Sprintf("🛠 Menggunakan tool: %s", name)}
		},
		OnResult: func(name, result string, err error) {
			if err != nil {
				m.events.agent <- AgentMsg{err: err}
			} else {
				short := result
				if len(short) > 200 {
					short = short[:200] + "..."
				}
				m.events.agent <- AgentMsg{text: fmt.Sprintf("✅ %s: %s", name, short)}
			}
		},
		Approve: func(name, args string) bool {
			ch := make(chan bool, 1)
			m.events.approval <- ApproveRequest{name: name, args: args, respond: ch}
			return <-ch
		},
	}

	err := m.agent.Send(m.ctx, text, h)
	m.events.agent <- AgentMsg{done: true, err: err, total: m.agent.TotalUsage}
}

// toggleMainMenu opens/closes the main menu overlay.
func (m *Model) toggleMainMenu() {
	if m.menuOpen {
		m.menuOpen = false
		return
	}
	m.menuOpen = true
	m.menuSelected = 0
	m.menuTitle = "📋 MENU UTAMA"
	m.menuItems = []menuItem{
		{label: "🤖 Ganti Provider", desc: "OpenRouter, OpenAI, Claude, dll", action: (*Model).menuOpenProviders},
		{label: "🧠 Ganti Model", desc: "Pilih model AI", action: (*Model).menuOpenModels},
		{label: "📁 Info Project", desc: "Lihat project saat ini", action: (*Model).menuProjectInfo},
		{label: "🚀 Deploy", desc: "Deploy aplikasi", action: (*Model).menuDeploy},
		{label: "🔒 Scan Keamanan", desc: "OWASP security scan", action: (*Model).menuScan},
		{label: "🎨 Ganti Tema", desc: "Catppuccin, Dracula, Nord...", action: (*Model).menuOpenThemes},
		{label: "💾 Simpan Sesi", desc: "Simpan percakapan", action: (*Model).menuSaveSession},
		{label: "🆘 Bantuan", desc: "Panduan cepat", action: (*Model).menuHelp},
	}
}

// menuOpenProviders shows provider selection.
func (m *Model) menuOpenProviders() tea.Cmd {
	m.menuOpen = false
	m.submenuOpen = true
	m.submenuTitle = "🤖 PILIH PROVIDER"
	m.submenuItems = nil
	for _, name := range m.cfg.ProviderNames() {
		name := name
		label := name
		if name == m.cfg.Provider {
			label += " ◀"
		}
		switch name {
		case "openrouter": label += " — 200+ model"
		case "openai": label += " — ChatGPT"
		case "anthropic": label += " — Claude"
		case "gemini": label += " — gratis"
		case "ollama": label += " — offline"
		}
		m.submenuItems = append(m.submenuItems, menuItem{
			label: label,
			action: func(m2 *Model) tea.Cmd {
				return m2.switchProvider(name)
			},
		})
	}
	m.submenuItems = append(m.submenuItems, menuItem{
		label: "🔙 Kembali",
		action: func(m2 *Model) tea.Cmd {
			m2.submenuOpen = false
			return nil
		},
	})
	return nil
}

func (m *Model) switchProvider(name string) tea.Cmd {
	m.cfg.Provider = name
	prov, err := ai.Build(m.cfg, name)
	if err != nil {
		m.appendMsg("tool", fmt.Sprintf("❌ Gagal ganti provider: %v", err))
		return nil
	}
	m.provider = prov
	_, model := m.cfg.Current()
	m.agent.SetProvider(prov, model)
	m.submenuOpen = false
	m.appendMsg("system", fmt.Sprintf("✅ Provider diganti ke: %s (model: %s)", name, model))
	return nil
}

// menuOpenModels shows model selection.
func (m *Model) menuOpenModels() tea.Cmd {
	m.menuOpen = false
	m.submenuOpen = true
	m.submenuTitle = "🧠 PILIH MODEL"
	m.submenuItems = nil
	pc, current := m.cfg.Current()
	for _, model := range pc.Models {
		model := model
		label := model
		if model == current {
			label += " ◀"
		}
		m.submenuItems = append(m.submenuItems, menuItem{
			label: label,
			action: func(m2 *Model) tea.Cmd {
				return m2.switchModel(model)
			},
		})
	}
	m.submenuItems = append(m.submenuItems, menuItem{
		label: "🔙 Kembali",
		action: func(m2 *Model) tea.Cmd {
			m2.submenuOpen = false
			return nil
		},
	})
	return nil
}

func (m *Model) switchModel(model string) tea.Cmd {
	m.cfg.Model = model
	m.agent.SetProvider(m.provider, model)
	m.submenuOpen = false
	m.appendMsg("system", fmt.Sprintf("✅ Model diganti ke: %s", model))
	return nil
}

// menuOpenThemes shows theme selection.
func (m *Model) menuOpenThemes() tea.Cmd {
	m.menuOpen = false
	m.submenuOpen = true
	m.submenuTitle = "🎨 PILIH TEMA"
	m.submenuItems = nil
	for _, name := range themeNames() {
		name := name
		label := name
		if name == m.cfg.TUI.Theme {
			label += " ◀"
		}
		m.submenuItems = append(m.submenuItems, menuItem{
			label: label,
			action: func(m2 *Model) tea.Cmd {
				return m2.switchTheme(name)
			},
		})
	}
	m.submenuItems = append(m.submenuItems, menuItem{
		label: "🔙 Kembali",
		action: func(m2 *Model) tea.Cmd {
			m2.submenuOpen = false
			return nil
		},
	})
	return nil
}

func (m *Model) switchTheme(name string) tea.Cmd {
	m.cfg.TUI.Theme = name
	m.theme = themeByName(name)
	m.s = newStyles(m.theme)
	m.submenuOpen = false
	m.appendMsg("system", fmt.Sprintf("🎨 Tema diganti ke: %s", themes[name].Name))
	return nil
}

func (m *Model) menuProjectInfo() tea.Cmd {
	m.menuOpen = false
	wd, _ := os.Getwd()
	m.appendMsg("system", fmt.Sprintf("📁 Working directory: %s\n🤖 Provider: %s\n🧠 Model: %s\n⚡ Token: %d", wd, m.provider.Name(), m.agent.Model(), m.agent.TotalUsage.TotalTokens))
	return nil
}

func (m *Model) menuDeploy() tea.Cmd {
	m.menuOpen = false
	m.appendMsg("system", "🚀 Deploy: ketik 'deploy ke [platform]' di chat. Contoh: 'deploy ke vercel' atau 'deploy ke railway'")
	return nil
}

func (m *Model) menuScan() tea.Cmd {
	m.menuOpen = false
	m.appendMsg("system", "🔒 Scan: ketik 'scan keamanan [url]'. Contoh: 'scan keamanan https://example.com'")
	return nil
}

func (m *Model) menuSaveSession() tea.Cmd {
	m.menuOpen = false
	m.appendMsg("system", "💾 Sesi akan disimpan ke ~/.config/termux-vibe-coding/sessions/")
	return nil
}

func (m *Model) menuHelp() tea.Cmd {
	m.menuOpen = false
	help := `🆘 BANTUAN CEPAT

💬 CHAT BIASA:
  "buat aplikasi catatan"
  "scan keamanan https://situs.com"
  "deploy ke vercel"

/ PERINTAH:
  /model     ganti model AI
  /provider  ganti penyedia AI
  /theme     ganti tema
  /skill     pilih template
  /export    export project
  /deploy    deploy aplikasi
  /scan      scan keamanan
  /clear     bersihkan layar
  /help      bantuan ini

⌨️ TOMBOL:
  Ctrl+M    menu utama
  Ctrl+T    sesi baru
  Ctrl+S    simpan sesi
  Esc       tutup menu`
	m.appendMsg("system", help)
	return nil
}

// handleMenuNav processes arrow keys and enter in the main menu.
func (m *Model) handleMenuNav(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyUp:
		if m.menuSelected > 0 {
			m.menuSelected--
		}
	case tea.KeyDown:
		if m.menuSelected < len(m.menuItems)-1 {
			m.menuSelected++
		}
	case tea.KeyEnter:
		item := m.menuItems[m.menuSelected]
		if item.action != nil {
			return m, item.action(m)
		}
	case tea.KeyEsc:
		m.menuOpen = false
	}
	return m, nil
}

// handleSubmenuNav processes navigation in submenus.
func (m *Model) handleSubmenuNav(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyUp:
		if m.menuSelected > 0 {
			m.menuSelected--
		}
	case tea.KeyDown:
		if m.menuSelected < len(m.submenuItems)-1 {
			m.menuSelected++
		}
	case tea.KeyEnter:
		item := m.submenuItems[m.menuSelected]
		if item.action != nil {
			return m, item.action(m)
		}
	case tea.KeyEsc:
		m.submenuOpen = false
	}
	return m, nil
}

// handleSlashCommand processes slash commands.
func (m *Model) handleSlashCommand(text string) tea.Cmd {
	parts := strings.Fields(text)
	cmd := strings.ToLower(parts[0])

	switch cmd {
	case "/model":
		return m.menuOpenModelsCmd()
	case "/provider":
		return m.menuOpenProvidersCmd()
	case "/theme":
		return m.menuOpenThemesCmd()
	case "/help":
		return m.menuHelp()
	case "/clear":
		m.messages = nil
		m.appendMsg("system", "✨ Chat dibersihkan.")
		return nil
	case "/project":
		return m.menuProjectInfo()
	case "/deploy":
		return m.menuDeploy()
	case "/scan":
		return m.menuScan()
	case "/export":
		m.appendMsg("system", "📦 Export: ketik 'export project' untuk mengexport ke ZIP.")
		return nil
	case "/session":
		m.appendMsg("system", "💾 Gunakan Ctrl+S untuk simpan, Ctrl+T untuk sesi baru.")
		return nil
	}

	// Unknown command — show available
	m.appendMsg("system", fmt.Sprintf("❌ Perintah '%s' tidak dikenal. Ketik /help untuk bantuan.", text))
	return nil
}

func (m *Model) menuOpenModelsCmd() tea.Cmd {
	m.menuOpenModels()
	// Render submenu immediately in view
	m.appendMsg("system", "🧠 Pilih model dari menu (Ctrl+M).")
	return nil
}

func (m *Model) menuOpenProvidersCmd() tea.Cmd {
	m.menuOpenProviders()
	m.appendMsg("system", "🤖 Pilih provider dari menu (Ctrl+M).")
	return nil
}

func (m *Model) menuOpenThemesCmd() tea.Cmd {
	m.menuOpenThemes()
	m.appendMsg("system", "🎨 Pilih tema dari menu (Ctrl+M).")
	return nil
}

// newSession starts a fresh chat session.
func (m *Model) newSession() {
	m.sessions = append(m.sessions, session{title: fmt.Sprintf("Sesi %d", len(m.sessions)+1)})
	m.currentSess = len(m.sessions) - 1
	m.messages = nil
	m.appendMsg("system", fmt.Sprintf("📄 %s dimulai.", m.sessions[m.currentSess].title))
}

// renderMessages builds the full chat viewport content.
func (m *Model) renderMessages() string {
	var b strings.Builder
	for _, msg := range m.messages {
		switch msg.role {
		case "system":
			b.WriteString(m.s.muted.Render(msg.content) + "\n\n")
		case "user":
			b.WriteString(m.s.user.Render("🧑 Kamu:") + "\n")
			b.WriteString(m.s.agent.Render(msg.content) + "\n\n")
		case "agent":
			b.WriteString(m.s.agent.Render("🤖 " + m.provider.Name() + ":") + "\n")
			b.WriteString(m.s.agent.Render(msg.content) + "\n\n")
		case "tool":
			b.WriteString(m.s.tool.Render(msg.content) + "\n\n")
		}
	}
	if m.loading {
		b.WriteString(m.s.tool.Render(fmt.Sprintf("⏳ %s ", m.spinner.View())))
	}
	return b.String()
}

// View implements tea.Model.
func (m *Model) View() string {
	if !m.ready {
		return "\n\n  🚀 Loading termux-vibe-coding..."
	}

	// Render full layout
	var b strings.Builder

	// Header
	header := fmt.Sprintf("◉ termux-vibe-coding v0.1.0  ⚡ %s:%s",
		m.provider.Name(),
		m.agent.Model())
	b.WriteString(m.s.header.Render(header) + "\n")
	b.WriteString(strings.Repeat("─", m.width-2) + "\n")

	// Viewport (chat area)
	vpHeight := m.height - 6
	if vpHeight < 5 {
		vpHeight = 5
	}
	m.viewport.Height = vpHeight
	b.WriteString(m.viewport.View())
	b.WriteString("\n")

	// Approval bar
	if m.awaitingApproval {
		b.WriteString(m.s.warn.Render(fmt.Sprintf("❓ Izinkan tool '%s'? [Y]es / [N]o: ", m.approvalName)))
		b.WriteString("\n")
	}

	// Input
	inputText := m.input.View()
	b.WriteString(m.s.box.Render(inputText) + "\n")

	// Status bar + quick actions
	usage := m.agent.TotalUsage
	status := fmt.Sprintf("⚡ Token: %d  |  %s:%s  |  [Ctrl+M:Menu] [/:Cmd]",
		usage.TotalTokens, m.provider.Name(), m.agent.Model())
	b.WriteString(m.s.status.Render(status))

	// Menu overlay (if open)
	if m.menuOpen {
		menuContent := m.renderMenu()
		overlay := m.s.menu.Render(menuContent)
		// Simple overlay on top
		return overlay + "\n\n" + b.String()
	}
	if m.submenuOpen {
		menuContent := m.renderSubmenu()
		overlay := m.s.menu.Render(menuContent)
		return overlay + "\n\n" + b.String()
	}

	return b.String()
}

func (m *Model) renderMenu() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("\n  %s\n\n", m.menuTitle))
	for i, item := range m.menuItems {
		line := fmt.Sprintf("  %s %s — %s", icon(i), item.label, item.desc)
		if i == m.menuSelected {
			b.WriteString(m.s.selected.Render("▸ " + line) + "\n")
		} else {
			b.WriteString("  " + line + "\n")
		}
	}
	b.WriteString("\n  [↑/↓: pilih] [Enter: buka] [Esc: tutup]")
	return b.String()
}

func (m *Model) renderSubmenu() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("\n  %s\n\n", m.submenuTitle))
	for i, item := range m.submenuItems {
		if i == m.menuSelected {
			b.WriteString(m.s.selected.Render("▸ " + item.label) + "\n")
		} else {
			b.WriteString("  " + item.label + "\n")
		}
	}
	b.WriteString("\n  [↑/↓: pilih] [Enter: pilih] [Esc: kembali]")
	return b.String()
}

func icon(i int) string {
	icons := []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟"}
	if i < len(icons) {
		return icons[i]
	}
	return "•"
}
