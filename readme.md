# Brancher

A powerful terminal-based Git branch manager that provides a fast and intuitive interface for managing Git branches.

![Go Version](https://img.shields.io/badge/Go-1.23.4+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

## Features

- 🚀 **Fast Navigation**: Quickly browse and switch between branches
- 🎯 **One-Key Operations**: Checkout, create, delete, rename, and pull branches with single keystrokes
- 💻 **Terminal UI**: Clean and responsive terminal interface built with tview
- 🔐 **SSH Support**: Seamless integration with SSH keys for remote operations
- 📊 **Status Display**: Real-time status messages and branch information

## Installation

### Option 1: Install via Go (Recommended)

```bash
go install github.com/Harichandra-Prasath/Brancher@latest
```

### Option 2: Build from Source

```bash
git clone https://github.com/Harichandra-Prasath/Brancher.git
cd Brancher
make build
```

### Option 3: Install from Source

```bash
git clone https://github.com/Harichandra-Prasath/Brancher.git
cd Brancher
go install .
```

## Usage

### Basic Usage

Run Brancher in any Git repository:

```bash
Brancher
```

### SSH Authentication

For operations requiring authentication (like pull), specify your SSH private key:

```bash
PV_KEY_FILE=id_rsa Brancher
```

The `PV_KEY_FILE` should be the name of your private SSH key file located in `$HOME/.ssh/`.

## Interface

Brancher provides a terminal-based user interface (TUI) that displays:

- **Branch List**: All local and remote branches
- **Status Bar**: Real-time operation feedback and messages
- **Interactive Controls**: Keyboard-driven navigation and operations

## Key Bindings

| Key | Action | Description |
|-----|--------|-------------|
| `c` | Checkout | Switch to the selected branch |
| `d` | Delete | Delete the selected branch |
| `n` | New | Create a new branch |
| `r` | Rename | Rename the selected branch |
| `p` | Pull | Pull updates for the selected branch |
| `q` | Quit | Exit the application |
| `↑/↓` | Navigate | Move up and down through the branch list |

## Requirements

- **Go**: 1.23.4 or higher
- **Git**: Any recent version
- **Operating System**: Linux, macOS, or Windows (with WSL)

## Dependencies

Brancher uses the following Go modules:
- `github.com/rivo/tview` - Terminal UI framework
- `github.com/go-git/go-git/v5` - Git operations
- `github.com/gdamore/tcell/v2` - Terminal interface

## Development

### Building

```bash
make build
```

### Running

```bash
make run
```

### Project Structure

```
Brancher/
├── main.go              # Application entry point
├── brancher/            # Git operations core
│   ├── core.go          # Branch management logic
│   └── operations.go    # Git command operations
├── ui/                  # User interface
│   ├── core.go          # UI initialization and setup
│   └── operations.go    # UI event handlers
├── Makefile             # Build automation
├── go.mod               # Go module definition
└── readme.md            # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Troubleshooting

### SSH Key Issues
Ensure your SSH private key is properly configured in `$HOME/.ssh/` and the correct filename is specified via `PV_KEY_FILE`.

### Git Repository Errors
Brancher must be run from within a Git repository. Make sure you're in a directory with `.git/` folder.

### Terminal Display Issues
If the TUI doesn't display correctly, ensure your terminal supports ANSI colors and has sufficient dimensions.

---

**Made with ❤️ for Git workflow optimization**  
 

