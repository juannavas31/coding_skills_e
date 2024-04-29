// Package rolling-hash provides a rolling hash function implementation.
package compute

const (
	prime = 31      // A large prime number for better hash distribution
	base  = 128     // Base for the hash function (size of the alphabet)
	mod   = 1e9 + 9 // Modulo to keep the hash value within a specific range
)

// RollingHash represents a rolling hash object
type RollingHash struct {
	hash int
	pow  int // Precomputed power of base
}

// NewRollingHash initializes a RollingHash object
func NewRollingHash(window []byte) *RollingHash {
	hash := 0
	pow := 1

	for _, char := range window {
		hash = (hash*base + int(char)) % mod
		pow = (pow * base) % mod
	}

	return &RollingHash{
		hash: hash,
		pow:  pow,
	}
}

// Roll updates the hash value for a new character entering the window
func (h *RollingHash) Roll(char byte) {
	oldChar := (h.hash - int(char)*h.pow + mod) % mod
	h.hash = (oldChar*base + int(char)) % mod
}
