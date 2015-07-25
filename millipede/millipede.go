// Package millipede provides a framework for creating millipedes.
// millipede is designed to be easy to understand and write, the most simple
// application can be written as follow:
//   func main() {
//     fmt.Println(millipede.New(20))
//   }
package millipede

import (
	"strings"

	"github.com/Sirupsen/logrus"
)

type Millipede struct {
	// Size is the amount of feet pairs
	Size uint64

	// Reverse is the flag that indicates the direction
	Reverse bool

	// Skin is the current millipede skin (template)
	Skin string
}

type Skin struct {
	// Head is used by the millipede to think about its life
	Head string
	// Pede are what make this arthropod so special
	Pede string

	// Reverse is the reverse skin of the millipede
	Reverse *Skin
}

// String returns a string representing a millipede
func (m *Millipede) String() string {
	paddingOffsets := []string{"  ", " ", "", " ", "  ", "   ", "    ", "    ", "   "}

	skins := map[string]Skin{
		"default": {
			Head: "  ╚⊙ ⊙╝  ",
			Pede: "╚═(███)═╝",
			Reverse: &Skin{
				Head: "  ╔⊙ ⊙╗  ",
				Pede: "╔═(███)═╗",
			},
		},
		"frozen": {
			Head: "  ╚⊙ ⊙╝  ",
			Pede: "╚═(❄❄❄)═╝",
			Reverse: &Skin{
				Head: "  ╔⊙ ⊙╗  ",
				Pede: "╔═(❄❄❄)═╗",
			},
		},
		"corporate": {
			Head: "  ╚⊙ ⊙╝  ",
			Pede: "╚═(©©©)═╝",
			Reverse: &Skin{
				Head: "  ╔⊙ ⊙╗  ",
				Pede: "╔═(©©©)═╗",
			},
		},
		"musician": {
			Head: "  ╚⊙ ⊙╝  ",
			Pede: "╚═(♫♩♬)═╝",
			Reverse: &Skin{
				Head: "  ╔⊙ ⊙╗  ",
				Pede: "╔═(♫♩♬)═╗",
			},
		},
		"bocal": {
			Head: "  ╚⊙ ⊙╝  ",
			Pede: "╚═(🐟🐟🐟)═╝",
			Reverse: &Skin{
				Head: "  ╔⊙ ⊙╗  ",
				Pede: "╔═(🐟🐟🐟)═╗",
			},
		},
		"ascii": {
			Head: "  \\o o/  ",
			Pede: "|=(###)=|",
			Reverse: &Skin{
				Head: "  /o o\\  ",
				Pede: "|=(###)=|",
			},
		},
	}

	skin := skins[m.Skin]
	if skin.Head == "" {
		logrus.Fatalf("no such skin: '%s'", m.Skin)
	}

	if m.Reverse && skin.Reverse != nil && skin.Reverse.Head != "" {
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
		Size:    size,
		Reverse: false,
		Skin:    "default",
	}
}
