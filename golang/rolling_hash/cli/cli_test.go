// Package cli
package cli

import (
	"testing"

	g "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestCLI(t *testing.T) {
	spec.Run(t, "CLI", func(t *testing.T, _ spec.G, it spec.S) {
		var (
			Expect = g.NewWithT(t).Expect
		)

		it("should validate a correct syntax", func() {
			// arrange
			args := []string{"../test/test_files/original.txt", "../test/test_files/updated.txt"}
			window := 8

			// act
			file1, file2, err := ValidateArgs(args, window)

			// assert
			Expect(err).ToNot(g.HaveOccurred())
			Expect(file1).To(g.Equal("../test/test_files/original.txt"))
			Expect(file2).To(g.Equal("../test/test_files/updated.txt"))
		})

		it("should return error if a file is missing", func() {
			// arrange
			args := []string{"original.txt"}
			window := 8

			// act
			_, _, err := ValidateArgs(args, window)

			// assert
			Expect(err).To(g.HaveOccurred())
			Expect(err.Error()).To(g.ContainSubstring("invalid syntax"))
		})

	}, spec.Report(report.Terminal{}))
}
