// Package millipede provides a framework for creating millipedes.
// millipede is designed to be easy to understand and write, the most simple
// application can be written as follow:
//   func main() {
//     fmt.Println(millipede.New())
//   }
package millipede

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"

	"github.com/kortschak/zalgo"
	"github.com/mgutz/ansi"
)

// Millipede defines a millipede configuration
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
	// Center is the flag that centers the millipede in the middle of the screen
	Center bool

	// PadRight is the flag that indicates the need to add spacing on the right of the output
	PadRight bool

	paddingOffsets []string
}

// getPadding returns a left padding
func (m *Millipede) getPadding(i uint64) string {
	if len(m.paddingOffsets) == 0 {
		m.paddingOffsets = append(m.paddingOffsets, "")
		//m.paddingOffsets = make([]string)
		if m.Curve > 0 {
			for n := uint64(1); n < m.Curve*2; n++ {
				size := int(math.Min(float64(n%(m.Curve*2)), float64(m.Curve*2-n%(m.Curve*2))))
				m.paddingOffsets = append(m.paddingOffsets, strings.Repeat(" ", size))
			}

			// --opposite support
			if m.Opposite {
				m.paddingOffsets = append(m.paddingOffsets[m.Curve:], m.paddingOffsets[:m.Curve]...)
			}
		}
	}
	if len(m.paddingOffsets) > 0 && m.Curve > 0 {
		return m.paddingOffsets[(m.Steps+i)%(m.Curve*2)]
	}
	return ""
}

// Draw returns a string representing a millipede and an error if any
func (m *Millipede) Draw() (string, error) {
	// --skin support
	skin, err := Skins.GetByName(m.Skin)
	if err != nil {
		return "", err
	}

	// --reverse support
	if m.Reverse {
		err = skin.SetDirection(DirectionDown)
		if err != nil {
			return "", err
		}
	}

	// --width support
	err = skin.SetWidth(int(m.Width))
	if err != nil {
		return "", err
	}

	// build the millipede body
	var body []string
	body = []string{m.getPadding(0) + skin.GetHead()}
	for x := uint64(0); x < m.Size; x++ {
		body = append(body, m.getPadding(x)+skin.NextPede())
	}
	tail := skin.GetTail()
	if tail != "" {
		body = append(body, m.getPadding(m.Size-1)+tail)
	}

	// --reverse support
	if m.Reverse {
		for i, j := 0, len(body)-1; i < j; i, j = i+1, j-1 {
			body[i], body[j] = body[j], body[i]
		}
	}

	// --center support
	if m.Center {
		w, err := getSize()
		if err == nil {
			var maxWidth int
			for _, line := range body {
				if len(line) > maxWidth {
					maxWidth = utf8.RuneCount([]byte(line))
				}
			}
			if maxWidth < w {
				leftMarginSize := (w / 2) - (maxWidth / 2)
				margin := strings.Repeat(" ", leftMarginSize)
				for idx, line := range body {
					body[idx] = margin + line
				}
			}
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

		if m.PadRight {
			line = line + strings.Repeat(" ", int(m.Curve))
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

	return output, nil
}

// String returns a string representing a millipede and an empty string if an error occured
func (m *Millipede) String() string {
	string, err := m.Draw()
	if err != nil {
		return ""
	}
	return string
}

// New returns a millipede with default values
func New() *Millipede {
	return NewWithSize(20)
}

// NewWithSize returns a millipede with default values but a custom size
func NewWithSize(size uint64) *Millipede {
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
		Center:    false,
	}
}

// stringToRuneSlice converts a string to a slice of runes
func stringToRuneSlice(input string) []rune {
	output := make([]rune, utf8.RuneCountInString(input))
	n := 0
	for _, r := range input {
		output[n] = r
		n++
	}
	return output
}
