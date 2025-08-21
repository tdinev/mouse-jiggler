# 🖱️ Mouse Jiggler for Windows (Rust)

A lightweight, portable Rust application that prevents your screen saver or auto-lock from activating by gently nudging your mouse cursor every minute.
It runs silently in the background and lives in your system tray with options to pause/resume or quit.

## ✨ Features

* ✅ Moves mouse by 1 pixel every minute to simulate activity
* ✅ Wiggles back and forth to avoid cursor drift
* ✅ System tray icon with menu
* ✅ Pause/Resume toggle
* ✅ Graceful exit
* ✅ No installation required — just run the `.exe`

## 🛠️ Build Instructions

### Prerequisites

- [Rust](https://rustup.rs) installed (`cargo` and `rustc`)
- Windows OS (tested on Windows 10/11)

### Clone and Build

```bash
git clone https://github.com/tdinev/mouse-jiggler.git
cd mouse-jiggler

# Add dependencies
cargo build --release
```

The compiled executable will be located at:

```
target\release\mouse-jiggler.exe
```

You can now run this `.exe` directly — no installation required.

## 📦 Dependencies

Add these to your Cargo.toml:

```toml
[dependencies]
windows = "0.48.0"
tray-item = "0.6.0"
anyhow = "1.0"
```

## 🚀 Usage

1. Run `mouse-jiggler.exe`
2. The app will minimize to your system tray
3. Right-click the tray icon to:
   * Pause/Resume mouse movement
   * Quit the application

The mouse will move by ±1 pixel every 60 seconds while active.

## 🧙‍♂️ Stealth Mode (Optional)

To hide the console window for a cleaner experience, compile with the Windows GUI subsystem:

```bash
cargo rustc --release -- -C link-args="/SUBSYSTEM:WINDOWS"
```

## 📜 License

Apache 2.0 License.
See [LICENSE](LICENSE) file for details.