package millipede

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Direction constants
const (
	DirectionUp = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

// Skins is the global collection of registered skins
var Skins SkinCollection

//////////////////////////////

// SkinDirection defines a skin for a direction (up, down, left, right)
type SkinDirection struct {
	// Head is used by the millipede to think about its life
	Head string
	// Pede are what make this arthropod so special
	Pede []string

	// Tail is what make a millipede looking like a snake
	//Tail string
}

// Skin defines the different parts of a millipede body
type Skin struct {
	Up   *SkinDirection
	Down *SkinDirection

	name             string
	currentDirection *SkinDirection
	directionChecked bool
	currentI         int
	width            int
	maxWidth         int
}

// SkinCollection represents a collection of skins
type SkinCollection struct {
	db map[string]Skin
}

func (sc *SkinCollection) AddSkin(name string, skin Skin) {
	skin.name = name
	sc.db[name] = skin
}

//////////////////////////////

func init() {
	Skins.db = make(map[string]Skin)

	Skins.AddSkin("default", Skin{
		Up: &SkinDirection{
			Head: "  ╚⊙ ⊙╝  ",
			Pede: []string{"╚═(███)═╝"},
		},
		Down: &SkinDirection{
			Head: "  ╔⊙ ⊙╗  ",
			Pede: []string{"╔═(███)═╗"},
		},
	})

	Skins.AddSkin("frozen", Skin{
		Up: &SkinDirection{
			Head: "  ╚⊙ ⊙╝  ",
			Pede: []string{"╚═(❄❄❄)═╝"},
		},
		Down: &SkinDirection{
			Head: "  ╔⊙ ⊙╗  ",
			Pede: []string{"╔═(❄❄❄)═╗"},
		},
	})

	Skins.AddSkin("diagonals", Skin{
		Up: &SkinDirection{
			Head: "  ╚⊙ ⊙╝  ",
			Pede: []string{"\\/(███)\\/", "/\\(███)/\\"},
		},
		Down: &SkinDirection{
			Head: "  ╔⊙ ⊙╗  ",
			Pede: []string{"/\\(███)/\\", "\\/(███)\\/"},
		},
	})

	Skins.AddSkin("corporate", Skin{
		Up: &SkinDirection{
			Head: "  ╚⊙ ⊙╝  ",
			Pede: []string{"╚═(©©©)═╝"},
		},
		Down: &SkinDirection{
			Head: "  ╔⊙ ⊙╗  ",
			Pede: []string{"╔═(©©©)═╗"},
		},
	})

	Skins.AddSkin("love", Skin{
		Up: &SkinDirection{
			Head: "  ╚⊙ ⊙╝  ",
			Pede: []string{"╚═(♥♥♥)═╝"},
		},
		Down: &SkinDirection{
			Head: "  ╔⊙ ⊙╗  ",
			Pede: []string{"╔═(♥♥♥)═╗"},
		},
	})

	Skins.AddSkin("musician", Skin{
		Up: &SkinDirection{
			Head: "  ╚⊙ ⊙╝  ",
			Pede: []string{"╚═(♫♩♬)═╝"},
		},
		Down: &SkinDirection{
			Head: "  ╔⊙ ⊙╗  ",
			Pede: []string{"╔═(♫♩♬)═╗"},
		},
	})

	Skins.AddSkin("bocal", Skin{
		Up: &SkinDirection{
			Head: "  ╚⊙ ⊙╝  ",
			Pede: []string{"╚═(🐟🐟🐟)═╝"},
		},
		Down: &SkinDirection{
			Head: "  ╔⊙ ⊙╗  ",
			Pede: []string{"╔═(🐟🐟🐟)═╗"},
		},
	})

	Skins.AddSkin("ascii", Skin{
		Up: &SkinDirection{
			Head: "  \\o o/  ",
			Pede: []string{"|=(###)=|"},
		},
		Down: &SkinDirection{
			Head: "  /o o\\  ",
			Pede: []string{"|=(###)=|"},
		},
	})

	Skins.AddSkin("inception", Skin{
		Up: &SkinDirection{
			Head: "    👀    ",
			Pede: []string{"╚═(🐛🐛🐛)═╝"},
		},
		Down: &SkinDirection{
			Head: "    👀    ",
			Pede: []string{"╔═(🐛🐛🐛)═╗"},
		},
	})

	Skins.AddSkin("humancentipede", Skin{
		Up: &SkinDirection{
			Head: "    👀    ",
			Pede: []string{"╚═(😷😷😷)═╝"},
		},
		Down: &SkinDirection{
			Head: "    👀    ",
			Pede: []string{"╔═(😷😷😷)═╗"},
		},
	})

	Skins.AddSkin("finger", Skin{
		Up: &SkinDirection{
			Head: "    👀    ",
			Pede: []string{"👈~~~  ~~~👉"},
		},
		Down: &SkinDirection{
			Head: "    👀    ",
			Pede: []string{"👈~~~~~~~~👉"},
		},
	})

	Skins.AddSkin("handy", Skin{
		Up: &SkinDirection{
			Head: "    👀    ",
			Pede: []string{"╚═(👌👌👌)═╝"},
		},
		Down: &SkinDirection{
			Head: "    👀    ",
			Pede: []string{"╔═(👌👌👌)═╗"},
		},
	})

	Skins.AddSkin("human", Skin{
		Up: &SkinDirection{
			Head: " ( ͡° ͜ʖ ͡°)",
			Pede: []string{"👆 ......👆"},
		},
		Down: &SkinDirection{
			Head: " ( ͡° ͜ʖ ͡°)",
			Pede: []string{"👇 ......👇"},
		},
	})
}

//////////////////////////////

// GetByName returns a skin based on its name or an error
func (sc *SkinCollection) GetByName(name string) (*Skin, error) {
	if skin, ok := sc.db[name]; ok {
		return &skin, nil
	}
	return nil, fmt.Errorf("no such skin: '%s'", name)
}

func (s *Skin) adaptWidth(line string) string {
	if s.width > s.maxWidth {
		slice := stringToRuneSlice(line)

		left := string(slice[:s.maxWidth/2])
		middle := strings.Repeat(string(slice[s.maxWidth/2:s.maxWidth/2+1]), int(s.width-s.maxWidth+1))
		right := string(slice[s.maxWidth/2+1:])

		return left + middle + right
	}
	return line
}

// GetHead returns a the head of the millipede
func (s *Skin) GetHead() string {
	s.checkDirection()
	s.Width()

	head := s.currentDirection.Head
	head = s.adaptWidth(head)
	return strings.TrimRight(head, " ")
}

// GetPede returns the pede of the millipede
func (s *Skin) GetPede(i int) string {
	s.checkDirection()
	s.Width()

	pede := s.currentDirection.Pede[i%len(s.currentDirection.Pede)]
	pede = s.adaptWidth(pede)
	return strings.TrimRight(pede, " ")
}

// NextPede returns GetPede with an incremental i
func (s *Skin) NextPede() string {
	return s.GetPede(s.nextI())
}

func (s *Skin) nextI() int {
	defer func() {
		s.currentI++
	}()
	return s.currentI
}

// Width returns the max width of any element
func (s *Skin) Width() int {
	s.checkDirection()

	if s.maxWidth == 0 {
		s.maxWidth = utf8.RuneCountInString(s.currentDirection.Head)

		/*
			if s.currentDirection.Tail != "" {
				tailWidth := utf8.RuneCountInString(s.currentDirection.Tail)
				if tailWidth > s.maxWidth {
					s.maxWidth = tailWidth
				}
			}
		*/

		for _, pedeType := range s.currentDirection.Pede {
			pedeLen := utf8.RuneCountInString(pedeType)
			if pedeLen > s.maxWidth {
				s.maxWidth = pedeLen
			}
		}
	}

	return s.maxWidth
}

func (s *Skin) checkDirection() {
	if s.directionChecked {
		return
	}
	if s.currentDirection == nil {
		s.currentDirection = s.Up
	}
	s.directionChecked = true
}

func (s *Skin) reset() {
	s.directionChecked = false
	s.width = 0
}

// SetDirection sets the current direction of the skin
func (s *Skin) SetDirection(newDirection int) error {
	s.reset()
	switch newDirection {
	case DirectionUp:
		s.currentDirection = s.Up
	case DirectionDown:
		s.currentDirection = s.Down
	default:
		s.currentDirection = nil
		return fmt.Errorf("invalid direction")
	}
	// s.checkDirection()
	// s.Width()
	return nil
}

// SetWidth set wanted width
func (s *Skin) SetWidth(width int) error {
	if width < s.Width() {
		s.width = s.Width()
		//return fmt.Errorf("this skin minimum width is '%d'", s.Width())
		return nil
	}
	s.width = width
	return nil
}
