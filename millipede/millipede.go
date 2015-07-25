// Package millipede provides a framework for creating millipedes.
// millipede is designed to be easy to understand and write, the most simple
// application can be written as follow:
//   func main() {
//     fmt.Println(millipede.New(20))
//   }
package millipede

import (
	"fmt"
)

type Pede string
type Millipede []Pede

func (m Millipede) Sprint() string {
  str := ""

  for _, pede := range m {
    tmp := fmt.Sprintln(pede)
    str = fmt.Sprint(str, tmp)
  }

  return str
}

func (m Millipede) Println() {
  fmt.Print(m.Sprint())
}

// return a new Millipede Reversed
// m.Reverse().Reverse().Reverse()
func (m Millipede) Reverse() Millipede {
  milli := make(Millipede, len(m))

  v := 0
  for i := len(m) - 1; i >= 0; i-- {
    milli[i] = m[v]
    v++
  }

  return milli
}

// Millipede returns a string representing a millipede of the specified size
func New(size uint64) Millipede {
	paddingOffsets := []string{"  ", " ", "", " ", "  ", "   ", "    ", "    ", "   "}

	milli := Millipede{"    ╚⊙ ⊙╝"}
	var x uint64
	for x = 0; x < size; x++ {
		line := paddingOffsets[x%9] + "╚═(███)═╝"
		milli = append(milli, Pede(line))
	}

	return milli
}
