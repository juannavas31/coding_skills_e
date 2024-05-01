// Package compute provides a rolling hash based file diffing function implementation.
package compute

// RollingHashTable represents the object used for comparing two files using a rolling hash table
type RollingHashTable struct {
	hashSlice []uint64       // stores the rolling hash values of each sliding window
	hashMap   map[uint64]int // stores the hash value and the sliding window in the file
	data      []byte         // stores the data of the file
}

// Creates a new RollingHashTable object
func NewRollingHashTable(data []byte, window int) *RollingHashTable {
	if len(data) < window {
		return nil
	}

	hashSlice := make([]uint64, len(data)-window+1)
	hasher := NewRollingHash(data[:window])
	hashSlice[0] = hasher.hash
	// create a map to store the hash value and the sliding window
	hashMap := make(map[uint64]int)
	// add the first hash value to the map
	hashMap[hasher.hash] = 0 // position of the first character in the first window
	for i := window; i < len(data); i++ {
		hasher.Roll(data[i], data[i-window])
		hashSlice[i-window+1] = hasher.hash
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
			delta, newI, newJ := CreateDelta(i, j, h, other)
			i, j = newI, newJ

			deltaList.AddDelta(delta)
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
	var retI, retJ int = i, j
	deletedFlag := other.hashMap[h.hashSlice[i]] == 0

	// check for deleted data in the first file
	for auxI := i + 1; auxI < len(h.hashSlice); auxI++ {
		if other.hashMap[h.hashSlice[auxI]] != 0 {
			delta = Delta{
				Operation: Delete,
				Start:     i,
				End:       auxI,
				Literal:   h.data[i:auxI],
			}
			retI = auxI
			break
		}
		if auxI >= len(h.hashSlice)-1 {
			delta = Delta{
				Operation: Delete,
				Start:     i,
				End:       len(h.hashSlice),
				Literal:   h.data[i:],
			}
			retI = auxI
		}
	}
	// check for inserted data in the second file
	if h.hashMap[other.hashSlice[j]] == 0 {
		for auxJ := j + 1; auxJ < len(other.hashSlice); auxJ++ {
			if h.hashMap[other.hashSlice[auxJ]] != 0 {
				// found end of inserted data, update the delta object
				operation := Insert
				if deletedFlag {
					operation = Replace
				}
				delta = Delta{
					Operation: operation,
					Start:     i,
					End:       retI,
					Literal:   other.data[j:auxJ],
				}
				retJ = auxJ
				break
			} else if auxJ >= len(other.hashSlice)-1 {
				operation := Insert
				if deletedFlag {
					operation = Replace
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
