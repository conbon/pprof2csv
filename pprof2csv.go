package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// parsePprofOutput parses the input pprof text (-text or -top output) into a 2D slice for CSV output.
func parsePprofOutput(pprofOutput string) ([][]string, error) {
	var data [][]string
	header := []string{"flat", "flat%", "sum%", "cum", "cum%", "function"}
	data = append(data, header)

	// Regular expression to match pprof output lines with fields.
	re := regexp.MustCompile(`^\s*([0-9\.]+(?:ms|MB|KB|B)?)\s+([0-9\.]+%)\s+([0-9\.]+%)\s+([0-9\.]+(?:ms|MB|KB|B)?)\s+([0-9\.]+%)\s+(.+)$`)

	scanner := bufio.NewScanner(strings.NewReader(pprofOutput))
	for scanner.Scan() {
		line := scanner.Text()

		// Skip lines that do not match the profile data format
		if !re.MatchString(line) {
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) == 7 {
			// Add the matched groups (excluding the full match itself) to the CSV data.
			data = append(data, matches[1:])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

// writeCSV writes the parsed data into a CSV file.
func writeCSV(data [][]string, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// Define and parse command-line flags
	inputFile := flag.String("input", "", "Path to the input pprof -text or -top output file")
	outputFile := flag.String("output", "output.csv", "Path to the output CSV file")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Usage: pprof2csv -input <pprof_text_output.txt> -output <output.csv>")
		os.Exit(1)
	}

	// Read the input file
	pprofData, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Parse the pprof output into a CSV-compatible format
	data, err := parsePprofOutput(string(pprofData))
	if err != nil {
		fmt.Printf("Error parsing pprof output: %v\n", err)
		os.Exit(1)
	}

	// Write the parsed data into a CSV file
	err = writeCSV(data, *outputFile)
	if err != nil {
		fmt.Printf("Error writing CSV: %v\n", err)
		os.Exit(1)
	}
}
