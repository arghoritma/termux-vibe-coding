# PRD: termux-vibe-coding — Vibe Coding di Mana Aja dari HP!

<p align="center">
  <strong>✨ Bikin aplikasi sesimpel chat. Dari HP. Untuk siapa aja. ✨</strong>
</p>

---

## 0. BRAND IDENTITY — termux-vibe-coding

### 0.1 Brand Essence

```
╔══════════════════════════════════════════════════╗
║                   termux-vibe-coding                    ║
║          "Vibe Coding di Mana Aja dari HP"        ║
╚══════════════════════════════════════════════════╝
```

| Elemen      | Definisi                                                                                                 |
| ----------- | -------------------------------------------------------------------------------------------------------- |
| **Nama**    | `termux-vibe-coding` — Terminal + Termux + AI Code Agent                                                 |
| **Tagline** | _Vibe Coding di Mana Aja dari HP_                                                                        |
| **Slogan**  | _Bikin aplikasi sesimpel chat._                                                                          |
| **Visi**    | Coding bukan lagi buat yang punya laptop. Semua orang bisa bikin aplikasi dari HP.                       |
| **Misi**    | Membuat AI coding agent yang mudah, murah, dan bisa dipakai siapa pun — cukup modal HP & kuota internet. |

### 0.2 Brand Personality

| Traits              | Deskripsi                                                  |
| ------------------- | ---------------------------------------------------------- |
| 🧠 **Cerdas**       | Tahu apa yang kamu mau, bahkan sebelum kamu selesai ngetik |
| 🤝 **Ramah**        | Ngomongnya pake bahasa sehari-hari, bukan jargon teknis    |
| ⚡ **Cepat**        | Instalasi 1 menit, langsung bisa dipake                    |
| 🎨 **Keren**        | Tampilan terminal warna-warni, nyaman dipandang            |
| 🔓 **Gratis**       | Open source, bawa API key sendiri, bebas provider          |
| 📱 **Mobile-first** | Dioptimalkan buat layar HP, portrait mode                  |

### 0.3 Brand Colors

```css
/* Palette termux-vibe-coding — Dark & Neon */
--primary: #00ff9d /* Hijau neon — terminal vibe, life, energy */
  --secondary: #7c3aed /* Ungu — AI, magic, kreativitas */ --accent: #ff6b6b
  /* Merah coral — semangat, action */ --bg-dark: #0a0a0f
  /* Hampir hitam — elegan, kontras */ --bg-surface: #1a1a2e
  /* Navy gelap — surface card */ --text: #e2e8f0
  /* Putih kebiruan — readability */ --muted: #64748b
  /* Abu-abu — secondary text */ --success: #00ff9d /* Hijau — success state */
  --warning: #fbbf24 /* Kuning — warning */ --error: #ff6b6b /* Merah — error */;
```

### 0.4 Logo Concept (ASCII untuk TUI)

```
      ████████  ██    ██  ██████   ██████  ██████  ███████
         ██     ██    ██  ██   ██  ██      ██   ██ ██
         ██     ██    ██  ██████   ██      ██   ██ █████
         ██     ██    ██  ██   ██  ██      ██   ██ ██
         ██      ██████   ██   ██   ██████ ██████  ███████

            Version 0.1.0  •  "vibe coding di mana aja"
```

### 0.5 Brand Voice & Tone

| Situasi         | Tone                | Contoh                                                          |
| --------------- | ------------------- | --------------------------------------------------------------- |
| **Sambutan**    | Hangat, excited     | "Halo! Mau bikin aplikasi apa hari ini? 😎"                     |
| **Instalasi**   | Simple, encouraging | "Tinggal 3 langkah doang, kok. Gampang!"                        |
| **Error**       | Empati, solutif     | "Waduh, ada yang salah nih. Coba cek koneksi internet kamu ya." |
| **Sukses**      | Proud, celebratory  | "🎉 Mantap! Aplikasi kamu udah jadi. Coba jalankan!"            |
| **Prompt user** | Santai, friendly    | "Mau bikin apa? Chat aja, kayak ngobrol sama temen."            |

---

## 1. Ringkasan Eksekutif

**termux-vibe-coding** adalah AI coding agent dari HP — install lewat npm di Termux, tinggal chat, langsung jadi aplikasi.

**Kok bisa sesimpel itu?**

1. Install Termux (dari F-Droid — gratis)
2. Ketik `npm install -g termux-vibe-coding`
3. Ketik `termux-vibe-coding`
4. Chat: _"buatkan aplikasi catatan harian"_ → langsung jadi

**Tanpa ribet:**

- ✅ Gak perlu jago coding
- ✅ Gak perlu beli domain/server
- ✅ Gak perlu laptop/mahal-mahal
- ✅ Gak perlu setup environment ribet

**Cukup: HP Android + Kuota internet + API key (gratis dari OpenRouter)**

---

## 2. Masalah & Solusi

### Masalah yang Dipecahkan

| Masalah                                      | Siapa                    | Selama Ini                           | Solusi termux-vibe-coding           |
| -------------------------------------------- | ------------------------ | ------------------------------------ | ----------------------------------- |
| **Mau belajar coding tapi gak punya laptop** | Pelajar, mahasiswa       | Harus pinjem laptop, pergi ke warnet | Coding dari HP langsung             |
| **Punya ide aplikasi tapi gak bisa coding**  | Non-teknis, founder, UKM | Harus cari developer, bayar mahal    | Chat ide → jadi aplikasi            |
| **Mau coding santai sambil rebahan**         | Developer, hobbyist      | Harus duduk di depan laptop          | Bikin aplikasi dari HP, di mana aja |
| **Mau audit keamanan website**               | Pentester, sysadmin      | Butuh tools complex, setup ribet     | Satu command, scan otomatis         |
| **Mau deploy aplikasi cepet**                | Freelancer, DevOps       | Setup server, config DNS, SSL        | Agent yang handle semuanya          |

### Solusi: Satu Tool untuk Semua

```
┌─────────────────────────────────────────────┐
│              termux-vibe-coding                     │
│                                              │
│  "Aku mau bikin..."                          │
│                                              │
│  📱 Aplikasi Catatan Harian                  │
│  🌐 Website Portfolio                        │
│  🛒 Toko Online                              │
│  🤖 Bot Telegram                             │
│  📊 Dashboard Admin                          │
│  🔐 Scanner Keamanan                         │
│  🚀 Deploy ke Server                         │
│                                              │
│  → Chat aja, sisanya biar AI yang handle     │
└─────────────────────────────────────────────┘
```

---

## 3. Siapa yang Butuh termux-vibe-coding?

### Persona 1: Andi — Pelajar SMA (16 tahun)

- **Situasi**: Mau belajar coding tapi cuma punya HP Android jadul
- **Penasaran**: Pingin bikin website buat portofolio tugas sekolah
- **Dengan termux-vibe-coding**:
  1. Install Termux + npm
  2. `termux-vibe-coding`
  3. Chat: _"buat portofolio pribadi dengan foto dan skill"_
  4. ✅ Jadi website HTML dalam 2 menit

### Persona 2: Bu Dewi — Pemilik Toko Kelontong (42 tahun)

- **Situasi**: Pengen punya aplikasi stok barang biar gak pake buku
- **Gak bisa coding**: Gak ngerti programming sama sekali
- **Dengan termux-vibe-coding**:
  1. Minta tolong anaknya install-in Termux
  2. Chat: _"buat aplikasi catatan stok barang, bisa nambah, edit, hapus barang"_
  3. ✅ Jadi aplikasi web yang bisa dibuka dari Chrome

### Persona 3: Raka — Developer Freelance (28 tahun)

- **Situasi**: Lagi ngopi di kafe, client minta revisi API
- **Bawa HP doang**: Laptop ketinggalan di rumah
- **Dengan termux-vibe-coding**:
  1. Buka Termux, `termux-vibe-coding`
  2. Chat: _"tambahin endpoint /api/products/search di project ini"_
  3. ✅ Revisi selesai sebelum kopi dingin

### Persona 4: Sari — Mahasiswa Teknik Informatika (20 tahun)

- **Situasi**: Lagi iseng di kosan, pengen bikin bot Telegram
- **Dengan termux-vibe-coding**:
  1. Chat: _"buat bot Telegram buat ngirim jadwal kuliah otomatis"_
  2. ✅ Bot siap dalam 5 menit

---

## 4. Cara Kerja — Simpel Banget

### 4.1 Instalasi (Tanpa Ribet)

**Langkah 1:** Install Termux dari F-Droid

```
Buka F-Droid → Cari "Termux" → Install
```

**Langkah 2:** Buka Termux, ketik 3 baris ini:

```
pkg update
pkg install nodejs
npm install -g termux-vibe-coding
```

_Tinggal copy-paste. Gak perlu mikir._

**Langkah 3:** Dapatkan API Key (gratis)

```
Buka openrouter.ai/keys → Login Google → Copy API Key
```

**Langkah 4:** Jalankan!

```
export OPENROUTER_API_KEY="sk-or-..."
termux-vibe-coding
```

⏱ **Total waktu: 2-3 menit.** Termasuk ngetik. Termasuk download.

### 4.2 Alternatif: Instalasi All-in-One (Paling Gampang)

```bash
curl -fsSL https://termux-vibe-coding.dev/install | bash
```

Satu baris aja. Semua otomatis: install Termux dependencies, npm, API key setup guide.

Bahkan sambil **mata terpejam pun bisa** — tinggal paste, enter, selesai.

### 4.3 Cara Pakai (Anak SD pun Bisa)

```
$ termux-vibe-coding

  ╔══════════════════════════════════════════════╗
  ║         ✨ termux-vibe-coding v0.1.0 ✨             ║
  ║     "Vibe Coding di Mana Aja dari HP"        ║
  ║                                              ║
  ║  Hai! Aku asisten coding kamu.                ║
  ║  Bilang aja mau bikin apa, nanti aku bikin.   ║
  ║                                              ║
  ║  Contoh: "buat aplikasi todo list"            ║
  ║          "scan keamanan website ini"          ║
  ║          "deploy app ke server"               ║
  ╚══════════════════════════════════════════════╝

> buat aplikasi catatan harian dengan tema pink

  ─────────────────────────────────────────────────
  ✅ Oke! Nanti kita bikin "Daily Journal" ✨
  📦 Bikin project structure...
  ✅ file index.html
  ✅ file style.css
  ✅ file app.js
  💾 Menambah fitur: simpan catatan
  💾 Menambah fitur: hapus catatan
  💾 Menambah fitur: tema pink
  💾 Menambah fitur: dark mode
  ─────────────────────────────────────────────────
  🎉 Selesai! Project di folder: ./daily-journal
  📁 Buka: cd daily-journal && npx serve .
  ─────────────────────────────────────────────────
  Mau tambahin fitur lain? Tinggal chat aja~
```

---

## 5. Menu System — Kendali Penuh dalam Genggaman

termux-vibe-coding punya **sistem menu kaya fitur** yang bisa diakses dengan 3 cara:

1. **Slash Commands** — ketik `/` di chat (mirip Discord/WhatsApp)
2. **Menu Popup** — tekan `Ctrl+M` buka menu visual
3. **Quick Action Bar** — shortcut di bagian bawah layar

### 5.1 Slash Commands (Ketik `/` di Chat)

Ketik `/` di kolom chat → langsung muncul daftar perintah:

```
> /                               ← cukup ketik slash

  ┌───────────────────────────────────────────┐
  │  /model        Ganti model AI           │
  │  /provider     Ganti provider AI        │
  │  /theme        Ganti tema tampilan      │
  │  /skill        Ambil skill/template     │
  │  /project      Info & kelola project    │
  │  /export       Export project ke ZIP    │
  │  /deploy       Deploy aplikasi          │
  │  /scan         Scan keamanan            │
  │  /session      Simpan/muat sesi         │
  │  /key          Ganti API key            │
  │  /compact      Mode hemat layar         │
  │  /status       Status sistem            │
  │  /clear        Bersihkan chat           │
  │  /help         Bantuan lengkap          │
  │  /undo         Batalkan aksi terakhir   │
  └───────────────────────────────────────────┘

> /model
  ┌───────────────────────────────────────────┐
  │  Pilih Model AI                            │
  │  ─────────────────────────────             │
  │  ◉ auto (rekomendasi) — otomatis terbaik   │
  │  ○ claude-sonnet-4 — paling jago coding   │
  │  ○ gpt-4o — serba bisa                    │
  │  ○ gemini-2.5-pro — gratis & kenceng      │
  │  ○ deepseek-chat — termurah               │
  │  ○ qwen2.5-coder — offline (Ollama)       │
  └───────────────────────────────────────────┘
  > Pilih: 1  (cukup tekan angka)
  ✅ Model diganti ke: claude-sonnet-4
```

**Contoh pemakaian slash command:**

```
> /provider
  ┌───────────────────────────────────────────┐
  │  Pilih Provider AI                         │
  │  ─────────────────────────────             │
  │  ◉ OpenRouter — 200+ model (rekomendasi)   │
  │  ○ OpenAI — ChatGPT                       │
  │  ○ Anthropic — Claude                     │
  │  ○ Google — Gemini                        │
  │  ○ Ollama — offline, gratis               │
  │  ○ Custom — provider sendiri              │
  └───────────────────────────────────────────┘
  > Pilih: 4
  ✅ Provider diganti ke: Google Gemini (gratis!)
  >> Sekarang kamu bisa pake AI gratis~

> /theme
  ┌───────────────────────────────────────────┐
  │  Pilih Tema                                │
  │  ─────────────────────────────             │
  │  ◉ Catppuccin Mocha — gelap keunguan       │
  │  ○ Dracula — gelap klasik                 │
  │  ○ Nord — biru elegan                     │
  │  ○ Tokyo Night — biru neon                │
  │  ○ Gruvbox — retro hangat                 │
  │  ○ Light — terang buat siang hari         │
  └───────────────────────────────────────────┘
  > Pilih: 3
  ✅ Tema diganti ke: Nord 🌙

> /skill
  ┌───────────────────────────────────────────┐
  │  Pilih Skill / Template                    │
  │  ─────────────────────────────             │
  │  1. 🌐 Website Portofolio                 │
  │  2. 🛒 Toko Online                        │
  │  3. 📱 Aplikasi Catatan                   │
  │  4. 🤖 Bot Telegram                       │
  │  5. 📊 Dashboard Admin                    │
  │  6. 🔐 Login System                       │
  │  7. 🎮 Game Sederhana                     │
  │  8. 🚀 Deploy ke Server                   │
  │  9. 🔒 Scan Keamanan                      │
  └───────────────────────────────────────────┘
  > Pilih: 2
  ✅ Skill "Toko Online" diaktifkan!
  >> Tinggal chat mau jualan apa~

> /help
  ┌───────────────────────────────────────────┐
  │              🆘 BANTUAN 🆘                 │
  ├───────────────────────────────────────────┤
  │                                           │
  │  💬 CHAT BIASA                             │
  │    "buat aplikasi catatan"                │
  │    "scan keamanan https://situs.com"      │
  │    "deploy ke Vercel"                     │
  │                                           │
  │  / PERINTAH CEPAT                         │
  │    /model     ganti AI model              │
  │    /provider  ganti penyedia AI           │
  │    /theme     ganti tema                  │
  │    /skill     pilih template              │
  │    /export    export project              │
  │    /deploy    deploy aplikasi             │
  │    /scan      scan keamanan               │
  │    /clear     bersihkan layar             │
  │                                           │
  │  ⌨ TOMBOL CEPAT                           │
  │    Ctrl+P    ganti provider               │
  │    Ctrl+M    buka menu                    │
  │    Ctrl+H    bantuan ini                  │
  │    Ctrl+S    simpan sesi                  │
  │    Tab       auto-lengkapi perintah       │
  └───────────────────────────────────────────┘
```

### 5.2 Menu Popup Visual (Tekan `Ctrl+M`)

Selain slash, user bisa tekan `Ctrl+M` untuk buka menu visual penuh:

```
╔═══════════════════════════════════════════════╗
║              📋 MENU UTAMA                     ║
║                                                ║
║  ┌──────────────┐  ┌──────────────────────┐   ║
║  │ 🤖 AI & Model │  │ 📁 PROJECT           │   ║
║  │───────────────│  │──────────────────────│   ║
║  │ Ganti Provider│  │ Info Project         │   ║
║  │ Ganti Model   │  │ Export ke ZIP        │   ║
║  │ Ganti API Key │  │ Buka Folder Project  │   ║
║  │ Status Token  │  │ Hapus Project        │   ║
║  └──────────────┘  └──────────────────────┘   ║
║                                                ║
║  ┌──────────────┐  ┌──────────────────────┐   ║
║  │ 🚀 DEPLOY    │  │ 🔒 KEAMANAN          │   ║
║  │───────────────│  │──────────────────────│   ║
║  │ Deploy App    │  │ Quick Scan           │   ║
║  │ Setup Domain  │  │ Full Security Audit  │   ║
║  │ Setup SSL     │  │ Lihat Laporan        │   ║
║  │ CI/CD Config  │  │ OWASP Checklist      │   ║
║  └──────────────┘  └──────────────────────┘   ║
║                                                ║
║  ┌──────────────┐  ┌──────────────────────┐   ║
║  │ 🎨 TAMPILAN  │  │ ⚙️ PENGATURAN        │   ║
║  │───────────────│  │──────────────────────│   ║
║  │ Ganti Tema    │  │ Bahasa (ID/EN)       │   ║
║  │ Mode Compact  │  │ Auto-save            │   ║
║  │ Ukuran Font   │  │ Reset Pengaturan     │   ║
║  │ Reset Layout  │  │ Tentang & Info       │   ║
║  └──────────────┘  └──────────────────────┘   ║
║                                                ║
║  ┌──────────────┐  ┌──────────────────────┐   ║
║  │ 💾 SESI      │  │ 🆘 BANTUAN           │   ║
║  │───────────────│  │──────────────────────│   ║
║  │ Simpan Sesi   │  │ Panduan Cepat        │   ║
║  │ Muat Sesi     │  │ Daftar Perintah      │   ║
║  │ Hapus Sesi    │  │ Tutorial             │   ║
║  │ Export Chat   │  │ Laporkan Masalah     │   ║
║  └──────────────┘  └──────────────────────┘   ║
║                                                ║
╚═══════════════════════════════════════════════╝
  [↑↓: pilih] [→: masuk submenu] [Esc: tutup]
```

### 5.3 Submenu Detail

Setiap kategori punya submenu yang lebih detail:

**Submenu: 🤖 AI & Model**

```
┌──────────────────────────────────────────────┐
│  🤖 PENGATURAN AI & MODEL                    │
├──────────────────────────────────────────────┤
│                                              │
│  Provider saat ini: ◉ OpenRouter            │
│  Model saat ini:    ◉ auto (claude-sonnet-4) │
│  Token bulan ini:   1.234 / 100.000 (1%)     │
│                                              │
│  ─── GANTI PROVIDER ───                      │
│  [1] OpenRouter     ⚡ 200+ model            │
│  [2] OpenAI         💰 mulai $0.01/request   │
│  [3] Anthropic      🧠 paling jago coding    │
│  [4] Google Gemini  🆓 gratis!               │
│  [5] Ollama         📴 offline               │
│                                              │
│  ─── GANTI MODEL ───                         │
│  [6] auto          推荐 otomatis terbaik     │
│  [7] claude-sonnet-4  🏆 recommended         │
│  [8] gpt-4o         ⚡ serba bisa            │
│  [9] gemini-2.5-pro 🆓 gratis & kenceng      │
│ [10] deepseek-chat  💰 termurah              │
│ [11] qwen2.5-coder  📴 offline (Ollama)      │
│                                              │
│  [G] Ganti API Key   [H] History biaya       │
│  [Esc] Kembali                               │
└──────────────────────────────────────────────┘
```

**Submenu: 🚀 DEPLOY**

```
┌──────────────────────────────────────────────┐
│  🚀 DEPLOY APLIKASI                          │
├──────────────────────────────────────────────┤
│                                              │
│  📦 Project: toko-online                     │
│  📁 Path: ./toko-online                      │
│                                              │
│  Mau deploy ke mana?                         │
│                                              │
│  [1] Vercel        🌐 best buat frontend     │
│  [2] Railway       ⚡ backend & fullstack    │
│  [3] GitHub Pages  🆓 gratis!               │
│  [4] VPS sendiri   🔧 full control           │
│  [5] Netlify       🆓 gratis!               │
│                                              │
│  Atau bilang aja di chat:                    │
│  > "deploy ke vercel"                        │
│                                              │
│  [Esc] Kembali                               │
└──────────────────────────────────────────────┘
```

**Submenu: 🔒 KEAMANAN**

```
┌──────────────────────────────────────────────┐
│  🔒 SCAN & KEAMANAN                          │
├──────────────────────────────────────────────┤
│                                              │
│  Pilih target:                               │
│  [1] Project saat ini (local scan)           │
│  [2] Website eksternal (URL scan)            │
│                                              │
│  Tipe scan:                                  │
│  [3] 🚀 Quick Scan — 30 detik                │
│      (port, SSL, header, basic vuln)         │
│  [4] 📋 Standard — 2 menit                   │
│      + SQLi, XSS, CORS, auth test            │
│  [5] 🔬 Full Audit — 5 menit                 │
│      + OWASP lengkap, API test, report       │
│                                              │
│  [R] Lihat laporan terakhir                  │
│  [O] OWASP checklist                         │
│  [Esc] Kembali                               │
└──────────────────────────────────────────────┘
```

**Submenu: 🎨 TAMPILAN**

```
┌──────────────────────────────────────────────┐
│  🎨 TEMA & TAMPILAN                          │
├──────────────────────────────────────────────┤
│                                              │
│  Tema saat ini: Catppuccin Mocha 🌙          │
│                                              │
│  Pilih tema:                                 │
│  [1] Catppuccin Mocha  🟣 gelap keunguan     │
│  [2] Dracula           🧛 gelap klasik       │
│  [3] Nord              🩵 biru elegan        │
│  [4] Tokyo Night       🌃 biru neon         │
│  [5] Gruvbox           🟠 retro hangat       │
│  [6] Light             ☀️ terang             │
│                                              │
│  Mode:                                       │
│  [C] Mode Compact — hemat layar HP           │
│  [F] Ukuran Font — kecil/sedang/besar        │
│  [S] Sidebar — tampil/sembunyi               │
│                                              │
│  [Esc] Kembali                               │
└──────────────────────────────────────────────┘
```

### 5.4 Quick Action Bar

Di bagian bawah layar, selalu ada shortcut yang bisa langsung diklik:

```
┌──────────────────────────────────────────────────┐
│ > _                                               │
├──────────────────────────────────────────────────┤
│  💬 Chat  │  📁 File  │  🤖 Model  │  ⚡ Deploy  │
│  🔒 Scan  │  💾 Save  │  🎨 Theme  │  🆘 Help    │
│   [Ctrl+P:Provider] [Ctrl+M:Menu] [Tab:Auto]     │
└──────────────────────────────────────────────────┘
```

User tinggal pencet angka/panah atau klik (via sentuhan di layar Termux).

### 5.5 Auto-Completion & Suggestions

Saat user ngetik, sistem akan auto-suggest:

```
> /                         → muncul daftar semua slash command
> /mod                      → auto-lengkapi: /model
> /p                        → auto-lengkapi: /provider
> dep                       → auto-lengkapi: "deploy ke ..."
> scan                      → "scan keamanan https://"
> bikin                     → "bikin aplikasi ..."
```

---

## 6. Fitur — Keren Tapi Tetap Simpel

### 6.1 Multi-Provider (Pilih AI Favorit Kamu)

Ganti AI kapan aja — baik via slash command `/provider` atau dari menu popup `Ctrl+M`:

```
> /provider
  ┌────────────────────────────────────┐
  │  Pilih Provider                    │
  │  ─────────────────────────────     │
  │  ◉ OpenRouter (rekomendasi)        │
  │  ○ OpenAI (ChatGPT)                │
  │  ○ Anthropic (Claude)              │
  │  ○ Google Gemini                   │
  │  ○ Ollama (offline, gratis)        │
  │  ○ Lainnya...                      │
  └────────────────────────────────────┘
```

**Mode "Otomatis"**: Default! Agent otomatis milih AI terbaik buat task kamu:

- Coding → pilih yang paling jago
- Chat santai → pilih yang paling murah
- Security → pilih yang paling teliti

**Gak punya API key?**

- OpenRouter: daftar gratis, dapat $1 credit
- Google Gemini: free tier 60 request/menit
- Ollama: gratis total, jalan di HP kamu sendiri

### 6.2 Multi-Bahasa (Ngobrol Pake Bahasa Kesukaan)

Kamu bisa ngobrol pake **Bahasa Indonesia** atau **Bahasa Inggris** — atau campur aduk. Agent paham semua.

```
> buat REST API untuk toko online
> tambahin login pake JWT
> deploy ke railway

--- atau ---

> make a tic-tac-toe game with HTML/CSS/JS
> add multiplayer support
> deploy to vercel
```

### 6.3 Bikin Aplikasi Apa Aja

| Ingin Membuat...      | Cukup Chat...                                          |
| --------------------- | ------------------------------------------------------ |
| 🌐 Website portofolio | `buat website portofolio dengan foto, skill, kontak`   |
| 🛒 Toko online        | `buat toko online dengan keranjang belanja`            |
| 📱 Aplikasi catatan   | `buat notes app dengan kategori warna`                 |
| 🤖 Bot Telegram       | `buat bot Telegram yang reply otomatis`                |
| 📊 Dashboard          | `buat dashboard penjualan dengan grafik`               |
| 🔐 Login system       | `buat halaman login dengan register dan lupa password` |
| 🎮 Game sederhana     | `buat game tebak angka`                                |
| 📝 Blog               | `buat blog dengan komentar dan admin panel`            |

### 6.4 Fitur Keamanan (Bonus!)

Bukan cuma bikin aplikasi, tapi juga bisa **ngecek keamanan aplikasi**:

```
> scan keamanan https://tokoku.com

  🔍 Scanning...
  ✅ SSL valid, aman
  ⚠️ CORS misconfig — bisa diperbaiki
  ❌ XSS vulnerability di halaman pencarian
  ✅ SQL injection: aman

  Mau aku perbaiki langsung? (y/N)
```

### 6.5 Deploy Aplikasi

```
> deploy ke Vercel

  📦 Siap-siap deploy...
  ✅ File konfigurasi vercel.json dibuat
  ✅ Git init
  ✅ Push ke GitHub (kalo ada)
  🚀 Tinggal klik deploy di vercel.com

  Atau mau aku handle semuanya? (y/N)
```

---

## 7. Tampilan TUI — Warna-warni, Gak Mbulet

Dengan **menu system** terintegrasi, tampilan TUI jadi pusat kendali semua fitur:

### 7.1 Tampilan Utama

```
┌──────────────────────────────────────────────────┐
│  ◉ termux-vibe-coding v0.1.0         ⚡ openrouter:auto │
├──────────────────────────────────────────────────┤
│                                                    │
│  Hai! Aku asisten coding kamu. 🔥                  │
│  Ketik / untuk menu cepat, atau chat aja~          │
│                                                    │
│  │  📁 Files                    💬 Chat            │
│  │  ┌──────────────┐  ┌──────────────────────┐   │
│  │  │ 📂 project   │  │ > /skill              │   │
│  │  │  ├ index.html│  │                      │   │
│  │  │  ├ style.css │  │ 📋 MENU SKILL        │   │
│  │  │  └ app.js    │  │ 1. Website Portfolio │   │
│  │  │              │  │ 2. Toko Online 🛒    │   │
│  │  │              │  │ 3. Aplikasi Catatan  │   │
│  │  │              │  │ 4. Bot Telegram 🤖   │   │
│  │  │              │  │ > Pilih: 2           │   │
│  │  │              │  │ ✅ Toko Online aktif!│   │
│  │  └──────────────┘  └──────────────────────┘   │
│                                                     │
│  ┌──────────────────────────────────────────────┐  │
│  │ > _                                           │  │
│  └──────────────────────────────────────────────┘  │
├──────────────────────────────────────────────────┤
│  💬 Chat  │  🤖 Model  │  ⚡ Deploy  │  🔒 Scan  │
│  [Ctrl+P:Provider] [Ctrl+M:Menu] [/ :Commands]   │
└──────────────────────────────────────────────────┘
```

### 7.2 Tampilan Menu Aktif (Ctrl+M)

```
╔═══════════════════════════════════════════════╗
║              📋 MENU UTAMA                     ║
║                                                ║
║  🤖 AI     📁 Project   🚀 Deploy   🔒 Scan   ║
║  🎨 Tema   💾 Session   ⚙️ Settings  🆘 Help   ║
║                                                ║
║  Pilih kategori pakai panah ← → lalu Enter    ║
║  Atau tekan Esc untuk balik ke chat            ║
╚═══════════════════════════════════════════════╝
```

### 7.3 Tampilan Mode Compact (untuk HP layar kecil)

```
┌──────────────────────────────┐
│ termux-vibe-coding  openrouter:auto│
├──────────────────────────────┤
│ > /skill                     │
│ 📋 1.🌐 2.🛒 3.📱 4.🤖     │
│ Pilih: 2 ✅ Toko Online     │
│                              │
│ > buat toko online baju      │
│ ✅ index.html                │
│ ✅ style.css                 │
│ ✅ app.js                    │
│ 🎉 Selesai!                  │
├──────────────────────────────┤
│ > _                          │
├──────────────────────────────┤
│ [P]rovider [M]enu [/]Cmd     │
└──────────────────────────────┘
```

### 7.4 Desain untuk HP:

- Layar portrait — gak perlu rotate HP
- Font besar — nyaman dibaca
- Warna kontras — di luar ruangan tetap kelihatan
- Navigasi keyboard minimal — cukup panah + enter + Esc
- Menu bisa dipake dengan **sekali sentuh** di layar Termux

---

## 8. Arsitektur Teknis (Untuk Developer)

### 9.1 Arsitektur Menu System

```
                     termux-vibe-coding
                    ──────────────
  ┌─────────────────────────────────────────────────┐
  │  TUI (Bubbletea + Lipgloss)                     │
  │  ┌───────────┐ ┌──────────┐ ┌────────────────┐ │
  │  │ Chat Panel│ │ FileTree │ │ Status Bar     │ │
  │  └───────────┘ └──────────┘ └────────────────┘ │
  ├─────────────────────────────────────────────────┤
  │  Agent Core (Think → Act → Observe loop)        │
  │  ├── System Prompt Builder                       │
  │  ├── Context Manager                             │
  │  └── Tool Router                                 │
  ├─────────────────────────────────────────────────┤
  │  Provider Layer                                  │
  │  ├── OpenRouter ◀── Primary (200+ model)         │
  │  ├── OpenAI                                      │
  │  ├── Anthropic                                   │
  │  ├── Google Gemini                               │
  │  ├── Ollama (lokal, offline)                     │
  │  └── Custom Provider                             │
  ├─────────────────────────────────────────────────┤
  │  Tool Layer                                      │
  │  read/write/edit/bash/grep/glob/npm/git/webfetch │
  │  portscan/sqli/xss/ssl-check/cors-check          │
  ├─────────────────────────────────────────────────┤
  │  Storage (bbolt)                                 │
  │  ├── ~/.config/termux-vibe-coding/config.yaml          │
  │  └── ~/.config/termux-vibe-coding/sessions/            │
  └─────────────────────────────────────────────────┘
```

**Tech Stack:**

| Layer          | Teknologi     | Kenapa                                       |
| -------------- | ------------- | -------------------------------------------- |
| **Bahasa**     | **Go 1.22+**  | Kenceng, binary kecil (ARM64), cocok buat HP |
| **TUI**        | **Bubbletea** | UI terminal modern, kaya fitur               |
| **Distribusi** | **npm**       | Udah ada di Termux, install gampang          |
| **Config**     | **YAML**      | Simple, gampang diedit                       |

---

## 9. Distribusi & Instalasi (Detail Teknis)

### 9.1 Cara Install

```bash
# Cara 1: npm (termudah)
npm install -g termux-vibe-coding

# Cara 2: npx (tanpa install permanen)
npx termux-vibe-coding

# Cara 3: One-liner (paling recommended buat non-teknis)
curl -fsSL https://termux-vibe-coding.dev/install | bash

# Update
npm update -g termux-vibe-coding
```

### 9.2 Binary Distribution

```
termux-vibe-coding/
├── bin/
│   ├── termux-vibe-coding-android-arm64    # Termux (target utama)
│   ├── termux-vibe-coding-linux-amd64
│   ├── termux-vibe-coding-darwin-arm64
│   └── termux-vibe-coding-windows-amd64
├── wrapper.js     # NPM entry: download binary → run
├── install.sh     # One-liner installer (curl pipe)
└── package.json
```

**Cara kerja npm wrapper:**

1. `npm install -g termux-vibe-coding`
2. `wrapper.js` jalan, deteksi OS/arch
3. Download binary yang cocok dari GitHub Releases
4. Simpan ke `~/.local/bin/termux-vibe-coding`
5. Selesai! Tinggal ketik `termux-vibe-coding`

---

## 10. Contoh Skenario Pemakaian

### Skenario 1: "Bikin Aplikasi Catatan" (5 Menit)

```
Pengguna:  "buat aplikasi catatan harian"
Agent:     ✅ Siap! Aku buatin yang simpel aja ya.
           📦 Membuat index.html
           📦 Membuat style.css
           📦 Membuat app.js
           ✨ Selesai! Bisa nyimpen catatan, hapus, edit.
           👉 cd catatan-harian && npx serve .

Pengguna:  "tambahin fitur kunci pake password"
Agent:     ✅ Oke, nanti ditambahin form login.
           💾 Memodifikasi index.html
           💾 Memodifikasi app.js
           ✨ Sekarang aplikasi kamu punya kunci password!
           Password default: 1234
```

### Skenario 2: "Bikin Toko Online" (15 Menit)

```
Pengguna:  "buat toko online baju"
Agent:     ✅ Siap! Aku buat dengan React + Tailwind.
           📦 npx create-react-app...
           📦 Install dependencies...
           📦 Membuat komponen:
               - Navbar
               - ProductCard
               - ShoppingCart
               - CheckoutForm
           ✨ Toko online siap!
           👉 cd toko-baju && npm start
```

### Skenario 3: "Cek Keamanan Website" (2 Menit)

```
Pengguna:  "cek keamanan https://toko-online-ku.com"
Agent:     🔍 Scanning...
           ✅ SSL valid
           ⚠️ [SEDANG] CORS terlalu longgar
           ❌ [KRITIS] XSS di halaman produk
           ❌ [KRITIS] SQL Injection di login
           📊 Report: laporan-keamanan.html

           Mau aku perbaiki yang kritis? (y/N) > y
           💾 Memperbaiki XSS...
           💾 Memperbaiki SQL injection...
           ✅ Semua kerentanan kritis udah diperbaiki!
```

---

## 11. Roadmap

### Fase 0: Rilis Cepat (Minggu 1) 🚀

- [ ] Setup Go project + npm wrapper
- [ ] Provider: OpenRouter (satu aja dulu)
- [ ] Agent: bisa chat + generate file
- [ ] TUI: chat panel + input box
- [ ] **MVP Rilis!** — Orang udah bisa cobain

### Fase 1: Bikin Aplikasi Beneran (Minggu 2) ⚡

- [ ] Tools: read, write, edit, bash, glob, grep
- [ ] Agent bisa baca project context
- [ ] Generate project structure
- [ ] File tree di TUI

### Fase 2: Banyak Provider (Minggu 3) 🔌

- [ ] OpenAI, Anthropic, Gemini, Ollama
- [ ] Ganti provider dari TUI (Ctrl+P)
- [ ] Ganti model dari TUI (Ctrl+M)
- [ ] Auto model selection

### Fase 3: Fitur Keren (Minggu 4) ✨

- [ ] Deploy (Vercel, Railway, VPS)
- [ ] Git integration (init, commit, push)
- [ ] Export project ke ZIP
- [ ] Session save/load
- [ ] Skill system (template YAML)

### Fase 4: Keamanan (Minggu 5) 🔒

- [ ] Port scanner + SQLi + XSS detector
- [ ] SSL/CORS/security headers checker
- [ ] OWASP report generator

### Fase 5: Polishing (Minggu 6) 🎨

- [ ] Error handling ramah
- [ ] Theming (catppuccin, dracula, nord)
- [ ] Dokumentasi + video tutorial
- [ ] Publikasi ke npm + GitHub

---

## 12. Perbandingan dengan Tool Lain

| Fitur                 | **termux-vibe-coding** 🚀 | opencode | aider       | cursor       |
| --------------------- | ------------------------- | -------- | ----------- | ------------ |
| **Bisa dipake di HP** | ✅ Ya, native             | ❌       | ❌          | ❌           |
| **Install gampang**   | ✅ npm 1 baris            | ❌       | ✅ npm      | ❌           |
| **Buat non-teknis**   | ✅ Chat aja               | ❌       | ❌          | ❌           |
| **Ganti provider**    | ✅ 6+ provider            | ❌       | ✅ terbatas | ❌           |
| **Gratis**            | ✅ Open source            | ✅       | ✅          | ❌ ($20/bln) |
| **Bahasa Indonesia**  | ✅ Support penuh          | ❌       | ❌          | ❌           |
| **TUI keren**         | ✅ Bubbletea              | ❌ basic | ❌ basic    | ❌ (GUI)     |
| **Offline**           | ✅ Ollama                 | ❌       | ❌          | ❌           |
| **Keamanan**          | ✅ OWASP built-in         | ❌       | ❌          | ❌           |
| **Size**              | ~15 MB                    | ~30 MB   | ~10 MB      | N/A          |

---

## 13. Brand & Marketing

### Taglines

| Penggunaan     | Tagline                                  |
| -------------- | ---------------------------------------- |
| **Utama**      | ✨ _Vibe Coding di Mana Aja dari HP_     |
| **Hero**       | _Bikin aplikasi sesimpel chat._          |
| **Fun**        | _Coding? Dari HP doang udah cukup._      |
| **Inspiratif** | _Gak punya laptop? Gak masalah._         |
| **Aksi**       | _Bikin aplikasi kamu sekarang, dari HP._ |

### Social Media Quotes

> "Dulu mau bikin aplikasi harus punya laptop, install software, belajar coding berbulan-bulan. Sekarang? Tinggal chat. Dari HP. Selesai." — **termux-vibe-coding**

> "Gak ada laptop bukan alasan buat gak berkarya. termux-vibe-coding ada buat kamu yang punya ide dan HP." — **termux-vibe-coding**

> "Umur 16 tahun, punya HP doang, bisa bikin website. Gapapa. Zaman sekarang, modal niat + kuota udah cukup." — **termux-vibe-coding**

### Target Audiens

```
┌──────────────────────────────────────────┐
│              SIAPA PAKE?                  │
├──────────────────────────────────────────┤
│                                          │
│  🧑‍🎓 Pelajar — Belajar coding dari HP    │
│  👩‍💼 Pemilik Toko — Bikin app stok barang │
│  🧑‍💻 Developer — Coding sambil ngopi     │
│  🔒 Pentester — Audit keamanan mobile    │
│  🚀 Founder — Bikin prototype cepet      │
│  🎨 Desainer — Pengen ngerti coding      │
│  👨‍👩‍👧 Ibu Rumah Tangga — Bikin app resep │
│  👴 Kakek-kakek — Iseng bikin blog       │
└──────────────────────────────────────────┘
  "Siapa pun bisa bikin aplikasi. Serius."
```

---

## 14. Success Metrics (Target 3 Bulan)

| Metrik                         | Target               |
| ------------------------------ | -------------------- |
| **npm downloads**              | 5,000+               |
| **GitHub stars**               | 500+                 |
| **Aplikasi dibuat via agent**  | 2,000+               |
| **Non-teknis users**           | 30% dari total users |
| **User retention (hari ke-7)** | 40%                  |
| **Rating (jika ada survey)**   | ⭐ 4.5 / 5           |

---

## 15. Risiko & Mitigasi

| Risiko                              | Dampak         | Mitigasi                                                                          |
| ----------------------------------- | -------------- | --------------------------------------------------------------------------------- |
| **Gaptek — bingung install Termux** | Gagal di awal  | Video tutorial step-by-step, one-liner installer                                  |
| **Gak punya API key**               | Gak bisa pake  | Default: Ollama (offline) atau Gemini (free), guide daftar OpenRouter             |
| **HP lemot/RAM kecil**              | Lambat         | Lightweight binary, compact mode, cache management                                |
| **Token abis (berbayar)**           | Agent berhenti | Notifikasi peringatan, mode hemat token, auto switch ke provider murah            |
| **Bingung mau bikin apa**           | User buntu     | Skill template built-in: "Bikin toko online", "Bikin blog", dll                   |
| **Error tiba-tiba**                 | Frustrasi      | Error message pake bahasa manusia: "Waduh, ada yang error nih. Coba ulangi lagi?" |

---

## 16. Monetization — Tetap Gratis Buat Semua

| Tier           | Harga           | Siapa          | Fitur                                              |
| -------------- | --------------- | -------------- | -------------------------------------------------- |
| **Community**  | **Gratis** 🎉   | Semua orang    | Semua fitur dasar, bawa API key sendiri            |
| **Pro**        | Rp 15.000/bulan | Power user     | Cloud session sync, custom theme, priority support |
| **Enterprise** | Kustom          | Tim/Perusahaan | On-premise, SSO, audit log, dedicated support      |

Prinsip: **Yang basic harus gratis selamanya.** Bayar cuma kalau mau fitur tambahan.

---

## 17. Cara Mulai (Dalam 3 Langkah)

```
╔══════════════════════════════════════════════════╗
║                                                  ║
║   ✨ 3 LANGKAH MULAI CODING DARI HP ✨           ║
║                                                  ║
║   1️⃣  Install Termux dari F-Droid                ║
║   2️⃣  Buka Termux, ketik:                        ║
║         npm install -g termux-vibe-coding               ║
║   3️⃣  Ketik: termux-vibe-coding                        ║
║         lalu chat mau bikin apa                   ║
║                                                  ║
║   ⏱ Cuma 3 menit. Serius.                        ║
║                                                  ║
║   📱 HP kamu → jadi mesin aplikasi!              ║
║                                                  ║
╚══════════════════════════════════════════════════╝
```

---

## 18. Next Steps

1. ✅ PRD approval
2. ⬜ **Fase 0 (Minggu ini):** Build prototype Go + OpenRouter + chat TUI
3. ⬜ **Rilis v0.1.0:** Biar orang bisa cobain (walaupun masih mentah)
4. ⬜ **Iterate:** Dengerin feedback user, perbaiki yang kurang
5. ⬜ **Rilis v1.0.0:** Fitur lengkap, dokumentasi, one-liner installer

---

```
╔══════════════════════════════════════════════════╗
║                                                  ║
║           termux-vibe-coding v0.1.0                     ║
║                                                  ║
║     "Vibe Coding di Mana Aja dari HP"            ║
║                                                  ║
║   📱 Android-ready                              ║
║   🎨 TUI Warna-warni                            ║
║   🤖 Multi-Provider AI                          ║
║   🔒 OWASP Pentest Ready                        ║
║   🚀 Deploy Satu Perintah                       ║
║   💯 100% Gratis & Open Source                  ║
║                                                  ║
║   "Bikin aplikasi sesimpel chat."                ║
║                                                  ║
╚══════════════════════════════════════════════════╝
```

---

_Dokumen PRD v3.0 — Brand Identity: termux-vibe-coding_  
_Tagline: "Vibe Coding di Mana Aja dari HP"_  
_Author: Rozak_  
_Lisensi: MIT_  
_Terakhir diperbarui: 14 Juni 2025_
