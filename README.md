# Kei

Kubernetes TUI :D

## Prerequisites

- Go 1.25.4 or later
- [just](https://github.com/casey/just)

### With Nix

- Install [direnv](https://direnv.net/)
- `direnv allow`

OR

- `just shell`

### Without nix

- `just deps-tooling`

## Development

### Quick Start

```bash
# Install dependencies
just deps

# Run the application
just run # or `just dev` for live reload
```

### Tooling

Run `just` without arguments to see all available commands:

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run `just check-all` to ensure quality
5. Submit a pull request

### Before Committing

Always run the quality checks:

- `just check-all`

## Credits

- Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) by Charm.
- Project structure referenced by [Superfile](https://github.com/yorukot/superfile) by Yorukot.
