// Package compute implements a rolling hash based file diffing algorithm
// This test suite implements a white box testing for the rolling hash table functionality
package compute

import (
	"testing"

	g "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func TestRollingHashTable(t *testing.T) {
	spec.Run(t, "RollingHashTable", func(t *testing.T, _ spec.G, it spec.S) {
		var (
			Expect = g.NewWithT(t).Expect
		)

		it("should create a new RollingHashTable object", func() {
			// arrange
			data := []byte("my first and original example data")
			window := 8

			// act
			hashTable := NewRollingHashTable(data, window)

			// assert
			Expect(hashTable).ToNot(g.BeNil())
			Expect(hashTable.data).To(g.Equal(data))
			Expect(hashTable.hashSlice).ToNot(g.BeEmpty())
			Expect(hashTable.hashMap).ToNot(g.BeEmpty())
		})

		it("should compare two RollingHashTable objects", func() {
			// arrange
			data1 := []byte("my very first and original example diff data")
			data2 := []byte("my very second not original example diff data")
			window := 4

			hashTable1 := NewRollingHashTable(data1, window)
			hashTable2 := NewRollingHashTable(data2, window)

			// act
			deltaList := hashTable1.Compare(hashTable2)

			// assert
			Expect(deltaList).ToNot(g.BeNil())
			Expect(deltaList.GetDeltas()).ToNot(g.BeEmpty())
			Expect(deltaList.GetDeltas()[0].Operation).To((g.Equal(Replace)))
			Expect(deltaList.GetDeltas()[0].Start).To(g.Equal(5))
			Expect(deltaList.GetDeltas()[0].End).To(g.Equal(17))
		})

		it("should create a delta object", func() {
			// arrange
			data1 := []byte("first example data")
			data2 := []byte("second example data")
			window := 4

			hashTable1 := NewRollingHashTable(data1, window)
			hashTable2 := NewRollingHashTable(data2, window)

			// act
			delta, i, j := CreateDelta(0, 0, hashTable1, hashTable2)

			// assert
			Expect(delta).ToNot(g.BeNil())
			Expect(delta.Operation).To((g.Equal(Replace)))
			Expect(i).To(g.Equal(5))
			Expect(j).To(g.Equal(6))
		})
	})
}
