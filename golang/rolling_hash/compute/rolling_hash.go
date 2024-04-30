// Package compute provides a rolling hash based file diffing function implementation.
// The Rabin-Karp algorithm is implemented for the rolling hash function.
package compute

const (
	base = 16      // Base for the hash function
	mod  = 1e9 + 7 // Modulo to keep the hash value within a specific range
)

// RollingHash represents a rolling hash object
type RollingHash struct {
	hash uint64
	pow  uint64 // Precomputed power of base
}

// NewRollingHash initializes a RollingHash object based on the Rabin-Karp algorithm
func NewRollingHash(window []byte) *RollingHash {
	var hash uint64 = 0
	var pow uint64 = 1

	for _, char := range window {
		hash = (hash*base + uint64(char)) % mod
		pow = (pow * base) % mod
	}
	pow = pow / base

	return &RollingHash{
		hash: hash,
		pow:  pow,
	}
}

// Roll updates the hash value for a new character entering the window and an old character leaving the window
func (h *RollingHash) Roll(charIn, charOut byte) {
	hashOut := h.hash - uint64(charOut)*h.pow
	h.hash = (hashOut*base + uint64(charIn)) % mod
}
