// Package millipede provides a framework for creating millipedes.
// millipede is designed to be easy to understand and write, the most simple application can be written as follow:
//   func main() {
//     fmt.Println(millipede.New(20))
//   }
package millipede

import (
	"strings"
)

type Millipede string

// Millipede returns a string representing a millipede of the specified size
func New(size uint64) Millipede {
	paddingOffsets := []string{"  ", " ", "", " ", "  ", "   ", "    ", "    ", "   "}

	bodyLines := []string{"    ╚⊙ ⊙╝"}
	var x uint64
	for x = 0; x < size; x++ {
		line := paddingOffsets[x%9] + "╚═(███)═╝"
		bodyLines = append(bodyLines, line)
	}

	return Millipede(strings.Join(bodyLines, "\n"))
}
