// Package millipede provides a framework for creating millipedes.
// millipede is designed to be easy to understand and write, the most simple application can be written as follow:
//   func main() {
//     fmt.Println(millipede.Millipede(20))
//   }
package millipede

import (
	"strings"
)

// Millipede returns a string representing a millipede of the specified size
func Millipede(size uint64) string {
	paddingOffsets := []string{"  ", " ", "", " ", "  ", "   ", "    ", "    ", "   "}

	bodyLines := []string{"    ╚⊙ ⊙╝"}
	var x uint64
	for x = 0; x < size; x++ {
		line := paddingOffsets[x%9] + "╚═(███)═╝"
		bodyLines = append(bodyLines, line)
	}

	return strings.Join(bodyLines, "\n")
}
