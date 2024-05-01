// Package compute provides a rolling hash based file diffing function implementation.
package compute

import (
	"testing"

	g "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func TestDeltaList(t *testing.T) {
	spec.Run(t, "DeltaList", func(t *testing.T, _ spec.G, it spec.S) {
		var (
			Expect = g.NewWithT(t).Expect
			deltas *DeltaList
		)

		it.Before(func() {
			deltas = NewDeltaList()
		})

		it("should add a new delta to the list", func() {
			// arrange
			delta := Delta{
				Operation: Insert,
				Start:     0,
				End:       5,
				Literal:   []byte("Hello"),
			}

			// act
			deltas.AddDelta(delta)

			// assert
			Expect(deltas.GetDeltas()).To(g.ContainElement(delta))
			Expect(deltas.GetDeltas()).To(g.HaveLen(1))
			Expect(deltas.GetDeltas()[0]).To(g.Equal(delta))
		})

		it("should return an empty list of deltas", func() {
			Expect(deltas.GetDeltas()).To(g.BeEmpty())
		})
	})
}
