package main

import (
	"fmt"

	"github.com/malanavi/loganalyzer/internal/analyzer"
)

func printInfo() {
	fmt.Println("loganalyzer is a command-line tool for analyzing application log files.")
	fmt.Println()
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  loganalyzer <command> <filepath> [--option=value ...]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Printf("  %-12s show statistics\n", statsCommand)
	fmt.Printf("  %-12s show only error lines\n", errorsCommand)
	fmt.Printf("  %-12s show most frequent messages\n", topCommand)
	fmt.Printf("  %-12s show version\n", versionCommand)
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --limit=N    limit number of displayed entries")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  loganalyzer stats app.log")
	fmt.Println("  loganalyzer errors app.log --limit=10")
	fmt.Println("  loganalyzer top app.log --limit=5")
}

func printHelp() {
	printInfo()
	printUsage()
}

func printStats(lines []string) {
	result := analyzer.Stats(lines)

	fmt.Println("Statistics:")
	fmt.Println("-------------------------")
	fmt.Printf("%-8s: %v\n", "Total", result.Total)
	fmt.Printf("%-8s: %v\n", "INFO", result.Info)
	fmt.Printf("%-8s: %v\n", "WARN", result.Warn)
	fmt.Printf("%-8s: %v\n", "ERROR", result.Error)
	fmt.Println("-------------------------")
}

func printErrors(lines []string, limit int) {
	result := analyzer.Errors(lines)

	if len(result) == 0 {
		fmt.Println("No errors found.")
		return
	}

	if limit <= 0 || limit > len(result) {
		limit = len(result)
	}

	fmt.Println("Errors:")
	fmt.Println("-------------------------")
	for i := 0; i < limit; i++ {
		fmt.Println(result[i])
	}
	fmt.Println("-------------------------")
}

func printTop(lines []string, limit int) {
	result := analyzer.Top(lines)

	if limit <= 0 || limit > len(result) {
		limit = len(result)
	}

	fmt.Println("Top repeated messages:")
	fmt.Println("-------------------------")
	for i := 0; i < limit; i++ {
		fmt.Printf(
			"%-8s: %d\n",
			result[i].Log,
			result[i].Times,
		)
	}
	fmt.Println("-------------------------")
}
