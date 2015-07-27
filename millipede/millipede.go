// Package millipede provides a framework for creating millipedes.
// millipede is designed to be easy to understand and write, the most simple
// application can be written as follow:
//   func main() {
//     fmt.Println(millipede.New(20))
//   }
package millipede

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"

	"github.com/getmillipede/millipede-go/vendor/github.com/Sirupsen/logrus"
	"github.com/getmillipede/millipede-go/vendor/github.com/kortschak/zalgo"
	"github.com/getmillipede/millipede-go/vendor/github.com/mgutz/ansi"
)

type Millipede struct {
	// Size is the amount of feet pairs
	Size uint64
	// Reverse is the flag that indicates the direction (up/down)
	Reverse bool
	// Skin is the current millipede skin (template)
	Skin string
	// Opposite is the flag that indicates the direction (left/right)
	Opposite bool
	// Width is the width of the millipede (depending on its age and the food it consumes)
	Width uint64
	// Curve is the size of the curve
	Curve uint64
	// Chameleon is the flag that indicates the millipede share its environment color
	Chameleon bool
	// Rainbow is the flag that indicates the millipede live with care bears
	Rainbow bool
	// Zalgo is the flag that invoke the hive-mind representing chaos
	Zalgo bool
	// Step is the amount of steps done by the millipede (useful for animations)
	Steps uint64
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
	paddingOffsets := []string{""}
	for n := uint64(1); n < m.Curve*2; n++ {
		size := int(math.Min(float64(n%(m.Curve*2)), float64(m.Curve*2-n%(m.Curve*2))))
		paddingOffsets = append(paddingOffsets, strings.Repeat(" ", size))
	}

	// --opposite support
	if m.Opposite {
		paddingOffsets = append(paddingOffsets[m.Curve:], paddingOffsets[:m.Curve]...)
	}

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
		"love": {
			Head: "  ╚⊙ ⊙╝  ",
			Pede: "╚═(♥♥♥)═╝",
			Reverse: &Skin{
				Head: "  ╔⊙ ⊙╗  ",
				Pede: "╔═(♥♥♥)═╗",
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
		"inception": {
			Head: "    👀    ",
			Pede: "╚═(🐛🐛🐛)═╝",
			Reverse: &Skin{
				Head: "    👀    ",
				Pede: "╔═(🐛🐛🐛)═╗",
			},
		},
		"humancentipede": {
			Head: "    👀    ",
			Pede: "╚═(😷😷😷)═╝",
			Reverse: &Skin{
				Head: "    👀    ",
				Pede: "╔═(😷😷😷)═╗",
			},
		},
		"finger": {
			Head: "    👀    ",
			Pede: "👈~~~  ~~~👉",
			Reverse: &Skin{
				Head: "    👀    ",
				Pede: "👈~~~~~~~~👉",
			},
		},
	}

	// --skin support
	skin := skins[m.Skin]
	if skin.Head == "" {
		logrus.Fatalf("no such skin: '%s'", m.Skin)
	}

	// --reverse support
	if m.Reverse && skin.Reverse != nil && skin.Reverse.Head != "" {
		skin = *skin.Reverse
	}

	// --width support
	if m.Width < 3 {
		logrus.Fatalf("millipede cannot have a witch < 3")
	}
	if m.Width > 3 {
		w := utf8.RuneCountInString(skin.Head)
		head := StringToRuneSlice(skin.Head)
		skin.Head = string(head[:w/2]) + strings.Repeat(string(head[w/2:w/2+1]), int(m.Width-2)) + string(head[w/2+1:])
		pede := StringToRuneSlice(skin.Pede)
		skin.Pede = string(pede[:w/2]) + strings.Repeat(string(pede[w/2:w/2+1]), int(m.Width-2)) + string(pede[w/2+1:])
	}

	// build the millipede body
	body := []string{paddingOffsets[m.Steps%(m.Curve*2)] + strings.TrimRight(skin.Head, " ")}
	var x uint64
	for x = 0; x < m.Size; x++ {
		var line string
		if m.Curve > 0 {
			line = paddingOffsets[(m.Steps+x)%(m.Curve*2)] + skin.Pede
		} else {
			line = "" + skin.Pede
		}
		body = append(body, line)
	}

	// --reverse support
	if m.Reverse {
		for i, j := 0, len(body)-1; i < j; i, j = i+1, j-1 {
			body[i], body[j] = body[j], body[i]
		}
	}

	// --chameleon and --rainbow support
	for idx, line := range body {
		colors := []string{"red", "green", "yellow", "blue", "magenta", "cyan", "white"}

		fgColor := ""
		bgColor := ""

		// --chameleon support
		if m.Chameleon {
			fgColor = "black"
		}

		// --rainbow
		if m.Rainbow {
			bgColor = colors[idx%len(colors)]
			if m.Chameleon {
				fgColor = bgColor
				fgColor = "black"
			}
		}

		if fgColor != "" || bgColor != "" {
			paddingSize := len(line) - len(strings.TrimSpace(line))
			line = strings.Repeat(" ", paddingSize) + ansi.Color(line[paddingSize:], fgColor+":"+bgColor)
		}

		body[idx] = line
	}

	output := strings.Join(body, "\n")

	// --zalgo support
	if m.Zalgo {
		buf := new(bytes.Buffer)

		z := zalgo.NewCorrupter(buf)
		z.Zalgo = func(n int, r rune, z *zalgo.Corrupter) bool {
			if string(r) == " " || r == 10 {
				z.Up = 0
				z.Middle = 0
				z.Down = 0
			} else {
				if z.Up == 0 {
					z.Up = complex(0, 0.2)
					z.Middle = complex(0, 0.2)
					z.Down = complex(0.001, 0.3)
				}
				z.Up += 0.1
				z.Middle += complex(0.1, 0.2)
				z.Down += 0.1
			}
			return false
		}
		fmt.Fprintln(z, output)
		output = buf.String()
	}

	return output
}

// New returns a millipede
func New(size uint64) *Millipede {
	return &Millipede{
		Size:      size,
		Reverse:   false,
		Skin:      "default",
		Opposite:  false,
		Width:     3,
		Curve:     4,
		Chameleon: false,
		Rainbow:   false,
		Zalgo:     false,
		Steps:     0,
	}
}

// StringToRuneSlice converts a string to a slice of runes
func StringToRuneSlice(input string) []rune {
	output := make([]rune, utf8.RuneCountInString(input))
	n := 0
	for _, r := range input {
		output[n] = r
		n++
	}
	return output
}
