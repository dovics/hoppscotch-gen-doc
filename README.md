# Hoppscotch to Markdown Generator

A Go CLI tool that converts Hoppscotch JSON collections into Markdown API documentation.

## Features

- Parses Hoppscotch JSON collection files
- Generates well-formatted Markdown documentation
- Multi-level folder structure support
- Includes:
  - Table of Contents
  - HTTP methods with visual badges (🟢 GET, 🟡 POST, 🔴 DELETE)
  - Request headers
  - Query parameters
  - Request bodies (with formatted JSON)
  - Authentication information
  - Full description support

## Installation

### Using Make (Recommended)

```bash
# Clone the repository
git clone https://github.com/dovics/hoppscotch-gen-doc.git
cd hoppscotch-gen-doc

# Build using make
make build

# Or install directly to GOPATH/bin
make install
```

### From source

```bash
# Build
go build -o hoppscotch-gen-doc

# Install to GOPATH/bin
go install
```

### Using Go install

```bash
go install github.com/dovics/hoppscotch-gen-doc@latest
```

## Development

### Make commands

```bash
# Show all available commands
make help

# Build the application
make build

# Install to GOPATH/bin
make install

# Clean build files
make clean

# Run tests
make test

# Run tests with coverage
make test-coverage

# Format code
make fmt

# Tidy go modules
make tidy

# Run linter (requires golangci-lint)
make lint

# Build release binaries for multiple platforms
make release

# Generate example documentation
make generate
```

## Usage

### Show help

```bash
# Show main command help
hoppscotch-gen-doc --help

# Show generate command help
hoppscotch-gen-doc generate --help
```

### Generate documentation

```bash
# Generate to file
hoppscotch-gen-doc generate -i example.json -o API.md

# Generate to stdout
hoppscotch-gen-doc generate -i example.json

# Using long flag names
hoppscotch-gen-doc generate --input example.json --output API.md

# Using make
make generate
```

### Command-line arguments

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--input` | `-i` | Path to Hoppscotch JSON file | Yes |
| `--output` | `-o` | Path to output Markdown file (optional, defaults to stdout) | No |
| `--help` | `-h` | Show help message | No |

## Example

Given a Hoppscotch JSON file like `example.json`, running:

```bash
hoppscotch-gen-doc generate -i example.json -o API.md
```

Will generate a Markdown file with:

- API collection name as title
- Hierarchical Table of Contents (organized by folders)
- Folder groupings with descriptions
- Detailed documentation for each request including:
  - HTTP method with visual badge
  - Endpoint URL
  - Description
  - Headers table
  - Query parameters table
  - Request body (formatted JSON)
  - Authentication details

## Project Structure

```
.
├── cmd/
│   ├── root.go         # Root command
│   └── generate.go     # Generate subcommand
├── internal/
│   └── generator/
│       └── generator.go # Documentation generation logic
├── main.go             # Entry point
├── Makefile            # Build automation
├── go.mod
├── go.sum
├── README.md
├── README_zh.md
├── example.json        # Example input file
└── .gitignore
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License

## Technologies Used

- [Cobra](https://github.com/spf13/cobra) - Powerful CLI applications for Go
- Go standard library - encoding/json, fmt, strings, os
