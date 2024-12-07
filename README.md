# sshm (SSH Manager)

A lightweight command-line tool to manage SSH config entries within a delimited block in your `~/.ssh/config` file. It allows you to easily add, remove, and list SSH host configurations while preserving any existing manual configurations outside its managed block.

## Features

- Add new SSH host configurations
- Remove existing host configurations
- List all managed hosts
- Preserves existing SSH config entries outside the managed block
- Automatically creates the config file and directories if they don't exist
- Maintains proper file permissions (0600 for config file, 0700 for directories)

## Project Structure

```
.
├── Makefile
├── README.md
├── cmd/                    # CLI commands
│   ├── add.go             # Add command implementation
│   ├── list.go            # List command implementation
│   ├── remove.go          # Remove command implementation
│   └── root.go            # Root command and CLI setup
├── dist/                   # Build artifacts
│   └── sshm
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksum
├── internal/              # Internal packages
│   ├── config/            # SSH config handling
│   │   └── ssh_config.go
│   ├── manager/           # Host management logic
│   │   └── host_manager.go
│   └── models/            # Data models
│       └── host_entry.go
└── main.go                # Application entry point
```

## Installation

### Prerequisites

- Go 1.16 or later
- Make (optional, for using Makefile commands)

### Building from Source

1. Clone the repository:
```bash
git clone https://github.com/will-wright-eng/sshm.git
cd sshm
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
# Using make
make build

# Or using go directly
go build -o dist/sshm
```

4. Install system-wide (optional):
```bash
# Using make
make install

# Or manually
sudo cp dist/sshm /usr/local/bin/
```

## Usage

### Adding a Host

```bash
sshm add -n myserver \
         -H 192.168.1.100 \
         -u admin \
         -i ~/.ssh/id_rsa \
         -p 22
```

Flags:
- `-n, --name`: Host name (required)
- `-H, --hostname`: Remote hostname or IP (required)
- `-u, --user`: Username (required)
- `-i, --identity`: Path to identity file (required)
- `-p, --port`: SSH port (optional, defaults to 22)

### Removing a Host

```bash
sshm remove -n myserver
```

Flags:
- `-n, --name`: Host name to remove (required)

### Listing All Managed Hosts

```bash
sshm list
```

## File Structure

The tool manages SSH configurations within a delimited block in your `~/.ssh/config` file:

```
# Your existing SSH config entries remain here

### BEGIN MANAGED BLOCK ###
Host myserver
  Hostname 192.168.1.100
  User admin
  IdentityFile ~/.ssh/id_rsa
  Port 22
### END MANAGED BLOCK ###

# Your existing SSH config entries remain here
```

## Development

### Running Tests

```bash
# Run all tests
make test

# Or using go directly
go test ./...
```

### Linting

```bash
# Using make
make lint

# Or using golangci-lint directly
golangci-lint run
```

### Making Changes

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

## Permissions

The tool automatically:
- Creates the `.ssh` directory with 0700 permissions if it doesn't exist
- Creates or modifies the config file with 0600 permissions

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. When contributing:

1. Write tests for new features
2. Update documentation
3. Follow Go best practices and style guidelines
4. Use clear commit messages

## License

MIT License - see LICENSE file for details

## Troubleshooting

### Common Issues

1. **Permission denied**
   ```bash
   # Fix directory permissions
   chmod 700 ~/.ssh

   # Fix config file permissions
   chmod 600 ~/.ssh/config
   ```

2. **Build errors**
   ```bash
   # Clean and rebuild
   make clean
   make build

   # Check dependencies
   go mod tidy
   ```

3. **Command not found**
   ```bash
   # Ensure the binary is in your PATH
   echo $PATH

   # Reinstall
   make install
   ```

## Roadmap

- [ ] Add support for additional SSH options
- [ ] Add backup functionality
- [ ] Add import/export features
- [ ] Add configuration validation
- [ ] Add support for multiple managed blocks
- [ ] Add automated releases
- [ ] Add Docker support

## Security Notes

- The tool automatically enforces secure file permissions
- Identity file paths are preserved as specified
- No sensitive information is logged or stored outside the SSH config file
