# What Time NTP

Simple Go application that fetches and displays accurate current time using NTP (Network Time Protocol) servers.

## Features

- Fetches accurate time from NTP servers
- Configurable NTP server and time format via environment variables
- Proper error handling with meaningful error messages
- Comprehensive test coverage

## Installation

### Prerequisites

- Go 1.21 or higher
- Internet connection (for NTP server access)

### Build from source

```bash
git clone https://github.com/rzmsq/what-time-ntp.git
cd what-time-ntp
go build -o what-time-ntp .
```

### Using Makefile

```bash
make build
```

## Usage

### Basic usage

```bash
./what-time-ntp
```

Output:
```
Current time: 2023-12-25T15:30:45Z
```

### Configuration

Configure the application using environment variables:

```bash
# Set custom NTP server
export NTP_SERVER="time.google.com"

# Set custom time format
export TIME_FORMAT="2006-01-02 15:04:05"

./what-time-ntp
```

Output:
```
Current time: 2023-12-25 15:30:45
```

### Available time formats

The application uses Go's time formatting syntax:

- `time.RFC3339` (default): `2006-01-02T15:04:05Z07:00`
- `2006-01-02 15:04:05`: `2023-12-25 15:30:45`
- `Jan 2, 2006 3:04:05 PM`: `Dec 25, 2023 3:30:45 PM`
- `15:04:05`: `15:30:45`

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `NTP_SERVER` | NTP server address | `0.beevik-ntp.pool.ntp.org` |
| `TIME_FORMAT` | Go time format string | `time.RFC3339` |

## Development

### Setup

```bash
# Install dependencies
make deps

# Install development tools
make install-tools
```

### Code Quality

```bash
# Format code
make fmt

# Run linters
make lint

# Run static analysis
make vet

# All code checks
make check
```

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-cover

# Run only unit tests (skip integration tests)
go test -short ./...
```

### Build

```bash
# Build for current platform
make build

# Full CI pipeline (format, lint, test, build)
make ci
```

## Project Structure

```
what-time-ntp/
├── main.go                    # Application entry point
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
├── Makefile                   # Build automation
├── .golangci.yml             # Linter configuration
├── .github/
│   └── workflows/
│       └── ci.yml            # GitHub Actions CI/CD
└── internal/
    ├── config/
    │   ├── config.go         # Configuration management
    │   └── config_test.go    # Configuration tests
    ├── ntpclient/
    │   ├── client.go         # NTP client implementation
    │   └── client_test.go    # NTP client tests
    └── timeformatter/
        ├── formatter.go      # Time formatting logic
        └── formatter_test.go # Time formatter tests
```

## Error Handling

The application handles various error scenarios:

- **Network connectivity issues**: Returns error code 1
- **Invalid NTP server response**: Returns error code 1
- **Server unreachable**: Returns error code 1

All errors are written to STDERR with descriptive messages.

## Dependencies

- [github.com/beevik/ntp](https://github.com/beevik/ntp) - NTP client library


## Examples

### Different time formats

```bash
# RFC3339 format (default)
./what-time-ntp
# Output: Current time: 2023-12-25T15:30:45Z

# Custom date format
TIME_FORMAT="Monday, January 2, 2006" ./what-time-ntp
# Output: Current time: Monday, December 25, 2023

# Time only
TIME_FORMAT="15:04:05" ./what-time-ntp
# Output: Current time: 15:30:45

# Unix timestamp
TIME_FORMAT="1136239445" ./what-time-ntp
# Output: Current time: 1703516245
```

### Different NTP servers

```bash
# Google's NTP server
NTP_SERVER="time.google.com" ./what-time-ntp

# Cloudflare's NTP server
NTP_SERVER="time.cloudflare.com" ./what-time-ntp

# NIST NTP server
NTP_SERVER="time.nist.gov" ./what-time-ntp
```