# sshm (SSH Manager)

A lightweight command-line tool to manage SSH config entries within a delimited block in your `~/.ssh/config` file. It allows you to easily add, remove, and list SSH host configurations while preserving any existing manual configurations outside its managed block.

## Features

- Add new SSH host configurations
- Remove existing host configurations
- List all managed hosts
- Preserves existing SSH config entries outside the managed block
- Automatically creates the config file and directories if they don't exist
- Maintains proper file permissions (0600 for config file, 0700 for directories)

## Installation

### Prerequisites

- Go 1.16 or later
- Make (optional, for using Makefile commands)

### Building from Source

```bash
# Clone the repository
git clone https://github.com/will-wright-eng/sshm.git
cd sshm

# Build the project
make build

# Install system-wide (optional)
make install
```

## Usage

```
A tool to manage SSH config entries in a controlled block.

Usage:
  sshm [command]

Available Commands:
  add         Add a new SSH host configuration
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List all managed SSH hosts
  remove      Remove an SSH host configuration

Flags:
  -h, --help   help for sshm

Use "sshm [command] --help" for more information about a command.
```

### Adding a Host

```bash
sshm add -n myserver -H 192.168.1.100 -u admin -i ~/.ssh/id_rsa -p 22
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

### Shell Completion

Generate shell completion scripts for your preferred shell:

```bash
# Bash
sshm completion bash > ~/.bash_completion.d/sshm

# Zsh
sshm completion zsh > "${fpath[1]}/_sshm"

# Fish
sshm completion fish > ~/.config/fish/completions/sshm.fish
```

[Rest of README remains the same...]
```

I've updated the usage section to:
1. Include the exact command output from the tool
2. Add information about the shell completion feature
3. Streamline the command examples
4. Make the flags section more consistent with the actual CLI output

Would you like me to:
1. Add example output for each command?
2. Include more detailed completion setup instructions?
3. Add advanced usage scenarios?
4. Include error messages and their solutions?