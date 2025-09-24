package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	processor "github.com/haesbaerteduardo/CompareJson/src/processor"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <file1.json> <file2.json>")
		os.Exit(1)
	}

	file1Path := os.Args[1]
	file2Path := os.Args[2]

	ids1, err := processor.ExtractIDs(file1Path)
	if err != nil {
		log.Fatalf("Error processing %s: %v", file1Path, err)
	}

	ids2, err := processor.ExtractIDs(file2Path)
	if err != nil {
		log.Fatalf("Error processing %s: %v", file2Path, err)
	}

	result := processor.CompareIDs(ids1, ids2)
	printResults(result, file1Path, file2Path)
}

func printResults(result processor.ComparisonResult, file1, file2 string) {
	var sb strings.Builder

	sb.WriteString("=== JSON ID Comparison Results ===\n")
	fmt.Fprintf(&sb, "File 1: %s (%d unique IDs)\n", file1, result.File1Count)
	fmt.Fprintf(&sb, "File 2: %s (%d unique IDs)\n\n", file2, result.File2Count)

	writeSection(&sb, "IDs found in both files", result.InBothFiles, "✓")
	writeSection(&sb, fmt.Sprintf("IDs only in %s", file1), result.OnlyInFile1, "-")
	writeSection(&sb, fmt.Sprintf("IDs only in %s", file2), result.OnlyInFile2, "+")

	fmt.Fprintf(&sb, "\nSummary:\n  Common IDs: %d\n  Missing from file2: %d\n  Missing from file1: %d\n",
		len(result.InBothFiles), len(result.OnlyInFile1), len(result.OnlyInFile2))

	fmt.Print(sb.String())
}

func writeSection(sb *strings.Builder, title string, ids []string, prefix string) {
	fmt.Fprintf(sb, "%s (%d):\n", title, len(ids))
	for _, id := range ids {
		fmt.Fprintf(sb, "  %s %s\n", prefix, id)
	}
	sb.WriteString("\n")
}
