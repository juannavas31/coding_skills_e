// Pakcage compute_test tests the compute package.
package compute

import (
	"testing"

	g "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func TestDiffFiles(t *testing.T) {
	spec.Run(t, "DiffFiles", func(t *testing.T, _ spec.G, it spec.S) {
		var (
			Expect = g.NewWithT(t).Expect
		)

		it("should return delta list for two different files", func() {
			// arrange
			file1 := "../test/test_files/original.txt"
			file2 := "../test/test_files/updated.txt"
			windowSize := 6

			// act
			deltaList, err := DiffFiles(file1, file2, windowSize)

			// assert
			Expect(err).ToNot(g.HaveOccurred())
			Expect(deltaList).ToNot(g.BeNil())
			Expect(deltaList.DiffList).ToNot(g.BeEmpty())
		})

		it("should return empty delta list for two identical files", func() {
			// arrange
			file1 := "../test/test_files/original.txt"
			file2 := "../test/test_files/original.txt"
			windowSize := 6
			expectedDeltaList := &DeltaList{
				DiffList: []Delta{},
			}

			// act
			deltaList, err := DiffFiles(file1, file2, windowSize)

			// assert
			Expect(err).ToNot(g.HaveOccurred())
			Expect(deltaList).To(g.Equal(expectedDeltaList))
		})
	})
}
