// Package compute provides a rolling hash based file diffing function implementation.
package compute

// Delta represents a change in the file
type Delta struct {
	Operation string
	Start     int
	End       int
	Literal   []byte
}

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
