package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type CommandFlags struct {
	Limit int
}

var ErrHelpRequested = errors.New("flag: help requested")

func parseCommandFlags(command string, args []string) (CommandFlags, error) {
	fs := flag.NewFlagSet(command, flag.ContinueOnError)

	var flags CommandFlags

	switch command {
	case errorsCommand, topCommand:
		fs.IntVar(&flags.Limit, "limit", 0, "limit number of results")
	}

	fs.SetOutput(os.Stdout)

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return flags, ErrHelpRequested
		}
		return CommandFlags{}, fmt.Errorf("parse flags: %w", err)
	}

	return flags, nil
}
