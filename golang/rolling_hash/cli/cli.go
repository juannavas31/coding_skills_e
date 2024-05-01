// Package cli provides a simple command line interface for validating the comamnd argiments
package cli

import (
	"errors"
	"fmt"
	"os"

	"rolling_hash/compute"
)

// function to test if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// validateArgs validates the command line arguments
func ValidateArgs(args []string, window int) (string, string, error) {
	if len(args) != 2 {
		return "", "", errors.New("invalid syntax, expected rolling_hash --window=<n> <old_file> <new_file>")
	}

	if window < 2 {
		return "", "", errors.New("invalid window size, it must be greater than 1")
	}

	file1 := args[0]
	file2 := args[1]
	if !fileExists(file1) {
		return "", "", errors.New(file1 + " does not exist")
	}

	if !fileExists(file2) {
		return "", "", errors.New(file2 + " does not exist")
	}

	return file1, file2, nil
}

// PrintResult prints the result of the file comparison
func PrintResult(deltaList *compute.DeltaList) {
	if len(deltaList.DiffList) == 0 {
		fmt.Println("Files are identical")
		return
	}

	fmt.Println("Delta for original file update:")
	for _, d := range deltaList.DiffList {
		fmt.Println("Delta change detected")
		fmt.Println("Operation: ", d.Operation)
		fmt.Println("Start: ", d.Start, " End: ", d.End)
		fmt.Println("Data: ", string(d.Literal))
	}
}
