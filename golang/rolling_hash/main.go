// Package main provides a simple example of using the rolling hash function.
package main

import (
	"flag"
	"fmt"

	"rolling_hash/cli"
	"rolling_hash/compute"
)

func main() {
	// Define flags for the command line arguments
	windowPtr := flag.Int("window", 6, "Window size for the rolling hash function")
	flag.Parse()
	// Get non-flag arguments
	args := flag.Args()

	file1, file2, err := cli.ValidateArgs(args, *windowPtr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	deltaList, err := compute.DiffFiles(file1, file2, *windowPtr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cli.PrintResult(deltaList)
}
