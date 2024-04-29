// Package compute provides a rolling hash based file diffing function implementation.
package compute

import (
	"os"
)

const (
	defaultWindowSize = 32 // Size of the window for the rolling hash function
)

// DiffFiles compares two files and generates a delta for upgrade
func DiffFiles(file1, file2 string, inputWindowSize int) (*DeltaList, error) {
	data1, err := os.ReadFile(file1)
	if err != nil {
		return nil, err
	}
	data2, err := os.ReadFile(file2)
	if err != nil {
		return nil, err
	}

	windowSize := inputWindowSize
	if windowSize == 0 {
		windowSize = defaultWindowSize
	}

	file1HashTable := NewRollingHashTable(data1, windowSize)
	file2HashTable := NewRollingHashTable(data2, windowSize)

	deltaList := file1HashTable.Compare(file2HashTable)

	return deltaList, nil
}
