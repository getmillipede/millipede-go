// Package millipede provides a framework for creating millipedes.
// millipede is designed to be easy to understand and write, the most simple
// application can be written as follow:
//   func main() {
//     fmt.Println(millipede.New(20))
//   }
package millipede

import (
	"strings"
)

type Millipede struct {
	// Size is the amount of feet pairs
	Size uint64
}

// String returns a string representing a millipede
func (m *Millipede) String() string {
	paddingOffsets := []string{"  ", " ", "", " ", "  ", "   ", "    ", "    ", "   "}

	bodyLines := []string{"    ╚⊙ ⊙╝"}
	var x uint64
	for x = 0; x < m.Size; x++ {
		line := paddingOffsets[x%9] + "╚═(███)═╝"
		bodyLines = append(bodyLines, line)
	}

	return strings.Join(bodyLines, "\n")
}

// New returns a millipede
func New(size uint64) *Millipede {
	return &Millipede{
		Size: size,
	}
}
