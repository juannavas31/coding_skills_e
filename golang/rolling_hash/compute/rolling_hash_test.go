// Package compute implements a rolling hash based file diffing algorithm
// This test suite implements a white box testing for the rolling hash function
package compute

import (
	"testing"

	g "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func TestRollingHash(t *testing.T) {
	spec.Run(t, "RollingHash", func(t *testing.T, _ spec.G, it spec.S) {
		var (
			Expect = g.NewWithT(t).Expect
			hasher *RollingHash
		)

		it.Before(func() {
			hasher = NewRollingHash([]byte("abcde"))
		})

		it("should initialize a RollingHash object", func() {
			Expect(hasher).NotTo(g.BeNil())
			Expect(hasher.hash).To(g.Equal(uint64(6785445)))
		})

		it("should update the hash value when rolling", func() {
			hasher.Roll('f', 'a')
			Expect(hasher.hash).To(g.Equal(uint64(6855350)))
		})
	})
}
