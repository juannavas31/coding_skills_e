// Package compute provides a rolling hash based file diffing function implementation.
package compute

// Define an enum-like type for the operation type
// OperationType represents the type of operation in the delta
type OperationType string

func (s OperationType) String() string {
	return string(s)
}

// enum values for operation type
const (
	Insert  OperationType = "insert"
	Delete  OperationType = "delete"
	Replace OperationType = "replace"
)

// Delta represents a change in the original file compared to the new file
type Delta struct {
	Operation OperationType
	Start     int    // Start position of the change in the original file
	End       int    // End position of the change in the original file
	Literal   []byte // Data that was inserted or deleted
}

// DeltaList represents a list of deltas
type DeltaList struct {
	DiffList []Delta
}

// Return an empty DeltaList object
func NewDeltaList() *DeltaList {
	return &DeltaList{
		DiffList: make([]Delta, 0),
	}
}

// Add a new delta to the list
func (d *DeltaList) AddDelta(delta Delta) {
	d.DiffList = append(d.DiffList, delta)
}

// Get the list of deltas
func (d *DeltaList) GetDeltas() []Delta {
	return d.DiffList
}
