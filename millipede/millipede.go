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

	// Reverse is the flag that indicates the direction
	Reverse bool
}

type Skin struct {
	Head string
	Pede string

	Reverse *Skin
}

// String returns a string representing a millipede
func (m *Millipede) String() string {
	paddingOffsets := []string{"  ", " ", "", " ", "  ", "   ", "    ", "    ", "   "}

	skin := Skin{
		Head: "  ╚⊙ ⊙╝  ",
		Pede: "╚═(███)═╝",
		Reverse: &Skin{
			Head: "  ╔⊙ ⊙╗  ",
			Pede: "╔═(███)═╗",
		},
	}

	if m.Reverse {
		skin = *skin.Reverse
	}

	body := []string{paddingOffsets[0] + skin.Head}
	var x uint64
	for x = 0; x < m.Size; x++ {
		line := paddingOffsets[x%9] + skin.Pede
		body = append(body, line)
	}

	if m.Reverse {
		for i, j := 0, len(body)-1; i < j; i, j = i+1, j-1 {
			body[i], body[j] = body[j], body[i]
		}
	}

	return strings.Join(body, "\n")
}

// New returns a millipede
func New(size uint64) *Millipede {
	return &Millipede{
		Size: size,
	}
}
