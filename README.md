# loganalyzer

![Go Version](https://img.shields.io/badge/go-1.18+-blue)
![License](https://img.shields.io/badge/license-MIT-green)

loganalyzer is a command-line tool for analyzing application log files.

It provides a simple and predictable interface for inspecting logs, extracting basic statistics, identifying error entries, and finding the most frequently occurring log messages.
## Table of Contents

- [Features](#features)
- [Quality](#quality)
- [Installation](#installation)
- [How to Use](#how-to-use)
- [Project Status](#project-status)
- [License](#license)

## Features

- Count total log entries and group them by log level (`INFO`, `WARN`, `ERROR`)
- Display only error log entries
- Show the most frequently occurring log messages
- Process log files line by line without loading the entire file into memory
- Clear and predictable CLI interface

## Quality

- Unit tests for core analysis logic
- Integration tests for file system operations

## Supported log format

The tool is designed to work with plain text log files where each log entry is stored on a separate line and log levels are indicated by prefixes such as `[INFO]`, `[WARN]`, and `[ERROR]`.

Each log line must start with a log level prefix followed by a message:

```text
[INFO] Application started
[WARN] Slow database query
[ERROR] Failed to save record
```

Log lines that do not match any known prefix are counted as Total entries but are not included in any specific log level.

## Installation

### Build from source

To build the binary yourself, you need Go installed locally.

#### Requirements

- Go 1.18 or newer

#### Build

Clone the repository and build the binary locally.

```bash
git clone https://github.com/Malanavi/loganalyzer.git
cd loganalyzer
go build -o loganalyzer ./cmd/loganalyzer
./loganalyzer help
```

The resulting `loganalyzer` binary can then be moved to any directory included in your PATH.

## How to Use

### Show statistics

Display total number of log entries and a breakdown by log level.

```bash
loganalyzer stats app.log
```

Example output:

```
Statistics:
-------------------------
Total   : 15230
INFO    : 12050
WARN    : 2140
ERROR   : 1040
-------------------------
```

### Show error entries

Display only log lines with the `ERROR` level.

```bash
loganalyzer errors app.log
```

Limit the number of displayed entries:

```bash
loganalyzer errors app.log --limit=10
```

### Show most frequent messages

Display the most frequently occurring log messages.

```bash
loganalyzer top app.log --limit=5
```

## Project Status

* Current version: 0.1.0
* Status: Stable
* The project is feature-complete for the current scope and maintained as needed

## License

This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file for details.
