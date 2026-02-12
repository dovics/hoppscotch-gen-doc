# Hoppscotch to Markdown Generator

A Go CLI tool that converts Hoppscotch JSON collections into Markdown API documentation.

## Features

- Parses Hoppscotch JSON collection files
- Generates well-formatted Markdown documentation
- Multi-level folder structure support
- Execute GET requests and include actual responses in documentation
- Includes:
  - Table of Contents
  - HTTP methods with visual badges (🟢 GET, 🟡 POST, 🔴 DELETE)
  - Request headers
  - Query parameters
  - Request bodies (with formatted JSON)
  - Response data (status code, headers, body)
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
# Generate to file (without executing requests)
hoppscotch-gen-doc generate -i example.json -o API.md

# Replace server URL in documentation (--server mode)
hoppscotch-gen-doc generate -i example.json --server https://api.example.com -o API.md

# Generate to file with GET request execution
hoppscotch-gen-doc generate -i example.json -o API.md -x

# Replace server URL only when executing requests (--target-server mode)
# Documentation shows original URL, but requests go to target server
hoppscotch-gen-doc generate -i example.json --target-server https://api.example.com -x -o API.md

# Replace server URL in both documentation and requests
hoppscotch-gen-doc generate -i example.json --server https://api.example.com -x -o API.md

# Generate to stdout
hoppscotch-gen-doc generate -i example.json

# Execute GET requests with custom timeout (default: 10 seconds)
hoppscotch-gen-doc generate -i example.json -x -t 30 -o API.md

# Using long flag names
hoppscotch-gen-doc generate --input example.json --output API.md --execute

# Using make
make generate
```

### Command-line arguments

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--input` | `-i` | Path to Hoppscotch JSON file | Yes |
| `--output` | `-o` | Path to output Markdown file (optional, defaults to stdout) | No |
| `--server` | | Replace endpoint host in documentation only (requests sent to original URL) | No |
| `--target-server` | | Replace endpoint host only when executing requests (documentation shows original URL) | No |
| `--execute` | `-x` | Execute GET requests and include responses in documentation | No |
| `--timeout` | `-t` | Request timeout in seconds (default: 10) | No |
| `--help` | `-h` | Show help message | No |

## Example

### Basic documentation generation

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

### Replace server URLs

There are two modes for server URL replacement:

#### Mode 1: Documentation Replacement (`--server`)

Replace the server URL only in documentation, requests are sent to the original URL:

```bash
hoppscotch-gen-doc generate -i example.json --server https://api.example.com -o API.md
```

**Example:**

If your original endpoint is:
- `http://localhost:8080/api/v1/health`

After using `--server https://api.example.com`, the documentation shows:
- `https://api.example.com/api/v1/health`

But requests are sent to: `http://localhost:8080/api/v1/health` (original URL)

This is useful when:
- Your Hoppscotch collection uses a development server URL
- You want to generate documentation with production server URLs
- But requests should still be sent to the development server
- Or when used with `--execute`, you want to execute requests against a different server than shown in docs

#### Mode 2: Request Replacement (`--target-server`)

Replace the server URL only when executing requests, documentation keeps original URL:

```bash
hoppscotch-gen-doc generate -i example.json --target-server https://api.example.com -x -o API.md
```

**Example:**

If your original endpoint is:
- `http://localhost:8080/api/v1/health`

Using `--target-server https://api.example.com`, the documentation shows:
- `http://localhost:8080/api/v1/health` (original URL)

But requests are sent to: `https://api.example.com/api/v1/health` (replaced URL)

This is useful when:
- Your Hoppscotch collection uses a development server URL
- You want to generate documentation with the original URLs
- But you need to execute requests against a different server (e.g., production)
- You need to test APIs against different environments without changing the documentation

### Execute GET requests

To include actual API responses in your documentation, use the `--execute` flag:

```bash
hoppscotch-gen-doc generate -i example.json -x -o API.md
```

This will execute all GET requests and include:

- **Response Status Code**: HTTP status code and message
- **Response Headers**: Table of response headers
- **Response Body**: Formatted JSON or text response

Example output:

```markdown
### Health

**🟢 GET**

**Endpoint:** `https://api.example.com/health`

#### Response

**Status Code:** 200 200 OK

**Response Headers:**

| Key | Value |
|-----|-------|
| Content-Type | application/json |
| Server | nginx |

**Response Body:**

```json
{
  "status": "healthy"
}
```
```

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
- Go standard library - encoding/json, fmt, strings, os, net/http
