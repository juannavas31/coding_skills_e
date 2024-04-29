// Package main provides a simple example of using the rolling hash function.
package main

import (
	"fmt"
	"rolling_hash/compute"
)

func main() {
	file1 := "original.txt"
	file2 := "updated.txt"
	windowSize := 32

	deltaList, err := compute.DiffFiles(file1, file2, windowSize)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(deltaList.DiffList) == 0 {
		fmt.Println("Files are identical")
		return
	}

	fmt.Println("Delta for upgrade:")
	for _, d := range deltaList.DiffList {
		fmt.Println(d)
	}
}
