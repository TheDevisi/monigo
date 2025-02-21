# 🤖 Monigo - Telegram Server Monitoring Bot

A powerful Telegram bot for monitoring your server's vital statistics in real-time. Built with Go, this bot provides an elegant way to check CPU usage, RAM consumption, disk space, and system uptime.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-1.24-blue.svg)

## ✨ Features

- 🌐 Bilingual support in progress (English/Russian)
- 💻 Real-time CPU monitoring
- 💾 RAM usage tracking
- 💿 Disk space analysis
- ⏱️ System uptime monitoring
- 🔒 Secure owner-only access
- ⌨️ Intuitive keyboard interface

## 🚀 Getting Started

### Prerequisites

- Go 1.24 or higher
- Telegram Bot Token (from [@BotFather](https://t.me/BotFather))
- Your Telegram User ID (from [@userinfobot](https://t.me/userinfobot))

### Installation

1. Clone the repository:
```bash
git clone https://github.com/TheDevisi/monigo.git
cd monigo
```

2. Create a `.env` file in the project root:
```env
BOT_TOKEN=your_bot_token_here
OWNER_ID=your_telegram_id_here(get it on @userinfobot)
```

3. Install dependencies:
```bash
go mod download
```

4. Build and run:

#### Development Build
```bash
go build -v -o monigo
./monigo
```

#### Production Build
```bash
# Optimized build with reduced binary size
go build -v -ldflags="-s -w" -o monigo
./monigo
```

#### Cross Compilation
For different architectures:
```bash
# For ARM (e.g., Raspberry Pi)
GOOS=linux GOARCH=arm go build -o monigo-arm

# For 64-bit Windows
GOOS=windows GOARCH=amd64 go build -o monigo.exe

# For macOS
GOOS=darwin GOARCH=amd64 go build -o monigo-mac
```

## 🎮 Usage

1. Start the bot by sending `/start`
2. Choose your preferred language (🇬🇧 English or 🇷🇺 Russian)
3. Use the keyboard interface to:
   - Check CPU status
   - Monitor RAM usage
   - View disk space
   - Check system uptime
   - Switch language

## 🤝 Commands

| Command | Description |
|---------|------------|
| `/start` | Initialize the bot and select language |
| `💻 CPU` | Show CPU information |
| `💾 RAM` | Display memory usage |
| `💿 Disk` | Show disk space stats |
| `⏱ Uptime` | Display system uptime |
| `🔄 Change language` | Switch interface language |

## 🔒 Security

- Bot access is restricted to the owner specified in `OWNER_ID`
- Environment variables are used for sensitive data
- Unauthorized access attempts are logged

## 🛠️ Technical Details

### Dependencies

- [github.com/go-telegram/bot](https://github.com/go-telegram/bot) - Telegram Bot API
- [github.com/shirou/gopsutil](https://github.com/shirou/gopsutil) - System statistics
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - Environment configuration
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) - Logging

### Project Structure

```
monigo/
├── config/
│   └── config.go  # Configuration management
│       
├── translations/
│   ├── translations.go # Work still in progress.
│
├── telegram/
│   ├── keyboard.go  # Keyboard layouts
│   ├── send_cpu.go      # CPU monitoring
│   ├── send_ram.go      # RAM monitoring
│   ├── send_disk_usage.go # Disk monitoring
│   ├── send_uptime.go   # Uptime monitoring
│   └── start.go         # Bot initialization
├── utils/
│   ├── disk_usage.go    # Disk utilities
│   ├── get_cpu_info.go  # CPU utilities
│   ├── get_memory.go    # Memory utilities
│   └── uptime.go        # Uptime utilities
├── main.go              # Entry point
```

## 📝 TODO

- [ ] SSH login monitoring
- [ ] Additional system metrics
- [ ] Custom alert thresholds
- [ ] Make it more user-friendly
- [ ] Full server control via bot
- [ ] Make it systemd service
- [ ] Add Windows Server support 

## ❓ FAQ

### General Questions

**Q: How secure is this bot?**
A: The bot implements owner-only access through Telegram User ID verification. Only the user specified in `OWNER_ID` can access the monitoring features.

**Q: Can I monitor multiple servers?**
A: Currently, one bot instance monitors one server. For multiple servers, you'll need to run separate instances with different bot tokens.

**Q: Does it work on Windows?**
A: While the bot can be compiled for Windows, it's primarily tested on Linux systems. Windows support is on my roadmap.

### Installation & Setup

**Q: Why can't I connect to my bot?**
A: Common issues include:
- Incorrect `BOT_TOKEN` in `.env`
- Bot not running/terminated
- Firewall blocking connections

**Q: How do I get my User ID?**
A: Send a message to [@userinfobot](https://t.me/userinfobot) on Telegram. It will reply with your User ID.

**Q: How do I run the bot in the background?**
A: You can:
1. Use systemd (recommended for Linux):
```bash
sudo systemctl start monigo # if you set up it on your own
```
2. Use screen or tmux:
```bash
screen -S monigo
./monigo
# Press Ctrl+A+D to detach
```

### Technical Questions

**Q: What's the resource usage of the bot?**
A: The bot is lightweight, typically using:
- CPU: < 1%
- RAM: ~20-30MB
- Disk: ~10MB

**Q: How often does it update statistics?**
A: Stats are fetched on-demand when you request them through commands.

**Q: Can I customize the monitoring intervals?**
A: Currently, monitoring is on-demand. Automated interval monitoring is planned for future releases.

**Q: How do I update the bot?**
A: Pull the latest changes and rebuild:
```bash
git pull origin main
go build

```

### Troubleshooting

**Q: Why are some metrics showing as "Error"?**
A: This usually happens when:
- The bot lacks necessary permissions
- System calls are restricted
- The monitored resource is unavailable

### Contributing

**Q: How can I contribute?**
A: You can:
- Submit issues for bugs or feature requests
- Fork the repository and create pull requests
- Improve documentation
- Add translations

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/thedevisi/monigo/issues).

## 👤 Author

**Denis Sinyutin** (TheDevisi)

---

⭐️ Star this project if you find it useful!