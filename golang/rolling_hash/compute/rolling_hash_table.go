// Package compute provides a rolling hash based file diffing function implementation.
package compute

type RollingHashTable struct {
	hashSlice []int       // stores the rolling hash values of each sliding window
	hashMap   map[int]int // stores the hash value and the sliding window in the file
	data      []byte      // stores the data of the file
}

// Creates a new RollingHashTable object
func NewRollingHashTable(data []byte, window int) *RollingHashTable {
	if len(data) < window {
		return nil
	}

	hashSlice := make([]int, len(data)-window+1)
	hasher := NewRollingHash(data[:window])
	hashSlice[0] = hasher.hash
	// create a map to store the hash value and the sliding window
	hashMap := make(map[int]int)
	// add the first hash value to the map
	hashMap[hasher.hash] = window - 1 // position of the last character in the window
	for i := window; i < len(data); i++ {
		hasher.Roll(data[i])
		hashSlice[i] = hasher.hash
		hashMap[hasher.hash] = i
	}
	return &RollingHashTable{
		hashSlice: hashSlice,
		hashMap:   hashMap,
		data:      data,
	}
}

// function to compare two RollingHashTable objects
func (h *RollingHashTable) Compare(other *RollingHashTable) *DeltaList {
	i, j := 0, 0
	deltaList := NewDeltaList()

	for i < len(h.hashSlice) && j < len(other.hashSlice) {
		if h.hashSlice[i] != other.hashSlice[j] {
			delta, i, j := CreateDelta(i, j, h, other)
			deltaList.AddDelta(delta)
			if i >= len(h.hashSlice) || j >= len(other.hashSlice) {
				break
			}
		} else {
			j++
			i++
		}
	}
	return deltaList
}

// function to create a delta object
// it takes the rolling hash tables of the two files as input, as well as the positions where there is a difference
// returns the delta object, the new i and j values, as the positions of the last character in the respective windows
// in the first and second file once they are equal again (i.e. the end of the deleted and inserted data),
func CreateDelta(i, j int, h, other *RollingHashTable) (Delta, int, int) {
	var delta Delta
	var retI, retJ int
	deletedFound := false
	// check for deleted data in the first file
	for auxI := i; auxI < len(h.hashSlice); auxI++ {
		if other.hashMap[h.hashSlice[auxI]] == 0 {
			deletedFound = true
		} else {
			// found end of deleted data, create a delta object
			delta = Delta{
				Operation: "delete",
				Start:     i,
				End:       auxI,
				Literal:   h.data[i:auxI],
			}
			retI = auxI
			break
		}
		if auxI == len(h.hashSlice)-1 {
			delta = Delta{
				Operation: "delete",
				Start:     i,
				End:       len(h.hashSlice),
				Literal:   h.data[i:],
			}
			retI = auxI
		}
	}
	// check for inserted data in the second file
	if h.hashMap[other.hashSlice[j]] == 0 {
		// found inserted data, find out how many and create a delta object
		for auxJ := j + 1; auxJ < len(other.hashSlice); auxJ++ {
			if h.hashMap[other.hashSlice[auxJ]] != 0 {
				// found end of inserted data, update the delta object
				operation := "insert"
				if deletedFound {
					operation = "replace"
				}
				delta = Delta{
					Operation: operation,
					Start:     i,
					End:       retI,
					Literal:   other.data[j:auxJ],
				}
				retJ = auxJ
				break
			} else if auxJ == len(other.hashSlice)-1 {
				operation := "insert"
				if deletedFound {
					operation = "replace"
				}
				delta = Delta{
					Operation: operation,
					Start:     i,
					End:       retI,
					Literal:   other.data[j:],
				}
				retJ = auxJ
			}
		}
	}
	return delta, retI, retJ
}
