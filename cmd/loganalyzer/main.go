package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Malanavi/loganalyzer/internal/reader"
)

const version = "v0.1.0"

const (
	helpCommand    = "help"
	statsCommand   = "stats"
	errorsCommand  = "errors"
	topCommand     = "top"
	versionCommand = "version"
)

const (
	limitPrefix = "--limit="
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		printHelp()
		return errors.New("command is required")
	}

	command := os.Args[1]
	if command == helpCommand {
		printHelp()
		return nil
	}
	if command == versionCommand {
		fmt.Println(version)
		return nil
	}

	if len(os.Args) < 3 {
		return errors.New("missing log file path\nUsage: loganalyzer <command> <file>")
	}

	path := os.Args[2]

	flags, err := parseCommandFlags(command, os.Args[3:])
	if err != nil {
		if errors.Is(err, ErrHelpRequested) {
			return nil
		}
		return err
	}

	lines, err := reader.ReadLines(path)
	if err != nil {
		return fmt.Errorf("read log file %q: %w", path, err)
	}

	switch command {
	case statsCommand:
		printStats(lines)
	case errorsCommand:
		printErrors(lines, flags.Limit)
	case topCommand:
		printTop(lines, flags.Limit)
	default:
		return fmt.Errorf(
			"unknown command %q\nRun 'loganalyzer help' to see available commands",
			command,
		)
	}

	return nil
}
