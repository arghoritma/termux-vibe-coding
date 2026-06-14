#!/usr/bin/env bash
#
# termux-vibe-coding — install script
# Usage: curl -fsSL https://raw.githubusercontent.com/rozaq/termux-vibe-coding/main/install.sh | bash
#
# Installs termux-vibe-coding via go install (Go required).
# Supports Termux (Android), macOS, Linux.

set -e

REPO="github.com/rozaq/termux-vibe-coding"
BINDIR="${GOPATH:-$HOME/go}/bin"
CMD="$REPO/cmd/termux-vibe-coding@latest"

# Colors
BOLD='\033[1m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo ""
echo -e "${BOLD}┌──────────────────────────────────────────┐${NC}"
echo -e "${BOLD}│       ✨ termux-vibe-coding installer ✨        │${NC}"
echo -e "${BOLD}│   Vibe Coding di Mana Aja dari HP        │${NC}"
echo -e "${BOLD}└──────────────────────────────────────────┘${NC}"
echo ""

# ── Detect OS ──
OS="$(uname -s)"
ARCH="$(uname -m)"
IS_TERMUX=false

if [ -n "$PREFIX" ] && echo "$PREFIX" | grep -q "com.termux"; then
    IS_TERMUX=true
fi

echo -e "  📱 OS:     ${OS} ${ARCH}"
if $IS_TERMUX; then
    echo -e "  📱 Mode:   Termux (Android)"
fi
echo ""

# ── Check Go ──
if ! command -v go &>/dev/null; then
    echo -e "  ${YELLOW}⚠️  Go belum terinstall.${NC}"
    if $IS_TERMUX; then
        echo -e "  ${YELLOW}   Install Go dulu: pkg install golang${NC}"
        echo ""
        read -rp "  Install Go sekarang? [Y/n]: " yn
        yn="${yn:-Y}"
        if [[ "$yn" =~ ^[Yy]$ ]]; then
            pkg install golang -y
        else
            echo -e "  ${RED}❌ Batal. Install Go manual lalu jalankan ulang script ini.${NC}"
            exit 1
        fi
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        echo -e "  ${YELLOW}   Install Go dulu: brew install go${NC}"
        echo ""
        read -rp "  Install Go sekarang? [Y/n]: " yn
        yn="${yn:-Y}"
        if [[ "$yn" =~ ^[Yy]$ ]]; then
            brew install go
        else
            echo -e "  ${RED}❌ Batal. Install Go manual lalu jalankan ulang script ini.${NC}"
            exit 1
        fi
    else
        echo -e "  ${YELLOW}   Install Go: https://go.dev/dl/${NC}"
        exit 1
    fi
fi

# ── Install via go install ──
echo -e "  🔧 Building from source via go install..."
echo -e "     go install ${CMD}"
echo ""

go install "$CMD"

# ── Verify ──
BINARY="$BINDIR/termux-vibe-coding"
if [ -f "$BINARY" ]; then
    echo ""
    echo -e "  ${GREEN}✅  Installed!${NC}"
    echo -e "     Binary: ${BINARY}"
    echo -e "     Version: $("$BINARY" version 2>/dev/null || echo "0.1.0")"
else
    # Try $GOPATH/bin
    GOPATH_BIN="$(go env GOPATH)/bin/termux-vibe-coding"
    if [ -f "$GOPATH_BIN" ]; then
        BINARY="$GOPATH_BIN"
        echo ""
        echo -e "  ${GREEN}✅  Installed!${NC}"
        echo -e "     Binary: ${BINARY}"
    else
        echo -e "  ${RED}❌ Binary tidak ditemukan di $BINDIR atau $(go env GOPATH)/bin${NC}"
        exit 1
    fi
fi

# ── Ensure in PATH ──
if ! command -v termux-vibe-coding &>/dev/null; then
    echo ""
    echo -e "  ${YELLOW}⚠️  termux-vibe-coding belum di PATH${NC}"
    echo -e "     Tambahkan ini ke ~/.bashrc atau ~/.zshrc:"
    echo -e "     export PATH=\"\$PATH:$(dirname "$BINARY")\""
fi

# ── Create config dir ──
CONFIG_DIR="${HOME}/.config/termux-vibe-coding"
mkdir -p "$CONFIG_DIR"
echo ""
echo -e "  📁 Config dir: ${CONFIG_DIR}"

# ── Instructions ──
echo ""
echo -e "  ${BOLD}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "  ${BOLD}✨  termux-vibe-coding SIAP DIGUNAKAN!${NC}"
echo -e "  ${BOLD}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo ""
echo -e "  ${GREEN}1. Set API Key${NC}"
echo -e "     export OPENROUTER_API_KEY=\"sk-or-...\""
echo ""
if $IS_TERMUX; then
    echo -e "  ${GREEN}2. Jalankan${NC}"
    echo -e "     termux-vibe-coding"
    echo ""
    echo -e "  ${GREEN}3. Mulai chat${NC}"
    echo -e "     > buat aplikasi catatan harian"
else
    echo -e "  ${GREEN}2. Jalankan${NC}"
    echo -e "     termux-vibe-coding"
fi
echo ""
echo -e "  ${BOLD}Butuh bantuan?${NC}"
echo -e "  Chat: /help  |  Menu: Ctrl+M  |  Provider: /provider"
echo ""
