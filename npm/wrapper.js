#!/usr/bin/env node

/**
 * termux-vibe-coding — npm wrapper
 *
 * Priority binary sources:
 * 1. Bundled binary (already in npm/bin/)
 * 2. Download from GitHub Releases
 * 3. Build from source via `go install`
 */

const { existsSync, chmodSync } = require("fs");
const { join } = require("path");
const { platform, arch } = require("os");
const { spawnSync, execSync } = require("child_process");

const PKG = join(__dirname, "package.json");
const BIN = join(__dirname, "bin", "termux-vibe-coding");
const VERSION = require(PKG).version || "0.1.0";
const REPO = "github.com/rozaq/termux-vibe-coding";

const TARGETS = {
  "android-arm64": "termux-vibe-coding-android-arm64",
  "linux-arm64": "termux-vibe-coding-linux-arm64",
  "linux-x64": "termux-vibe-coding-linux-amd64",
  "darwin-x64": "termux-vibe-coding-darwin-amd64",
  "darwin-arm64": "termux-vibe-coding-darwin-arm64",
  "win32-x64": "termux-vibe-coding-windows-amd64.exe",
};

function detectAsset() {
  let os = platform();
  let cpu = arch();
  // Detect Termux on Android (PREFIX env var is always set by Termux)
  if (process.env.PREFIX && process.env.PREFIX.includes("com.termux")) {
    os = "android";
  }
  const key = `${os}-${cpu}`;
  return TARGETS[key] || null;
}

function installViaBundled() {
  // Check if there's a binary already bundled (from `make npm-release`)
  const asset = detectAsset();
  if (!asset) return false;
  const bundled = join(__dirname, "bin", asset);
  if (existsSync(bundled)) {
    console.log(`📦 Using bundled binary: ${asset}`);
    // Symlink or copy to standard name
    if (bundled !== BIN) {
      const { copyFileSync } = require("fs");
      copyFileSync(bundled, BIN);
    }
    chmodSync(BIN, 0o755);
    return true;
  }
  // Check if BIN already exists (from previous install)
  if (existsSync(BIN)) return true;
  return false;
}

function installViaGo() {
  console.log("🔧 Building from source via 'go install'...");
  console.log(`   Running: go install ${REPO}/cmd/termux-vibe-coding@latest`);
  try {
    execSync(`go install ${REPO}/cmd/termux-vibe-coding@latest`, {
      stdio: "inherit",
    });
    // Copy from GOPATH/bin to here
    const { execSync: exec } = require("child_process");
    const gopath = exec("go env GOPATH").toString().trim();
    const src = join(gopath, "bin", "termux-vibe-coding");
    if (existsSync(src)) {
      const { copyFileSync } = require("fs");
      copyFileSync(src, BIN);
      chmodSync(BIN, 0o755);
      console.log(`✅  Built from source: ${BIN}`);
      return true;
    }
  } catch (e) {
    console.error(`   go install gagal: ${e.message}`);
  }
  return false;
}

function installViaDownload() {
  const asset = detectAsset();
  if (!asset) return false;
  const url = `https://github.com/rozaq/termux-vibe-coding/releases/download/v${VERSION}/${asset}`;
  console.log(`⬇️  Downloading termux-vibe-coding v${VERSION} for ${asset}...`);
  try {
    const { createWriteStream } = require("fs");
    const { get } = require("https");
    const dest = BIN;
    const file = createWriteStream(dest);
    return new Promise((resolve) => {
      get(url, (res) => {
        if (res.statusCode >= 400) {
          console.error(`   HTTP ${res.statusCode}: ${res.statusMessage}`);
          resolve(false);
          return;
        }
        res.pipe(file);
        file.on("finish", () => {
          chmodSync(dest, 0o755);
          console.log(`✅  Downloaded: ${dest}`);
          resolve(true);
        });
      }).on("error", (e) => {
        console.error(`   Download gagal: ${e.message}`);
        resolve(false);
      });
    });
  } catch (e) {
    console.error(`   Download gagal: ${e.message}`);
    return false;
  }
}

async function install() {
  // Try in order: bundled → download → go install
  if (installViaBundled()) {
    console.log(`✅  termux-vibe-coding v${VERSION} siap digunakan!`);
    return true;
  }

  const downloaded = await installViaDownload();
  if (downloaded) {
    console.log(`✅  termux-vibe-coding v${VERSION} siap digunakan!`);
    return true;
  }

  console.log("   Mencoba build dari source...");
  if (installViaGo()) {
    console.log(`✅  termux-vibe-coding v${VERSION} siap digunakan!`);
    return true;
  }

  console.error("");
  console.error("❌  Gagal menginstall binary termux-vibe-coding.");
  console.error("   Pastikan Go sudah terinstall, lalu jalankan:");
  console.error(`   go install ${REPO}/cmd/termux-vibe-coding@latest`);
  console.error(
    `   Atau download manual dari: https://github.com/rozaq/termux-vibe-coding/releases`,
  );
  return false;
}

function run() {
  if (!existsSync(BIN)) {
    console.log("⬇️  Binary belum tersedia. Download sekarang...");
    install().then((ok) => {
      if (ok) {
        const args = process.argv.slice(2);
        const result = spawnSync(BIN, args, { stdio: "inherit" });
        process.exit(result.status ?? 1);
      } else {
        process.exit(1);
      }
    });
    return;
  }
  const args = process.argv.slice(2);
  const result = spawnSync(BIN, args, { stdio: "inherit" });
  process.exit(result.status ?? 1);
}

// --- Main ---
const cmd = process.argv[2];
if (cmd === "install") {
  install().then((ok) => process.exit(ok ? 0 : 1));
} else if (cmd === "version") {
  console.log(VERSION);
} else {
  run();
}
