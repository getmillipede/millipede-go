package millipede

import (
	"fmt"
	"testing"
)

func ExampleNew() {
	millipede := New(20)
	fmt.Println(millipede)
	// Output:
	//   ╚⊙ ⊙╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
}

func ExampleConstructor() {
	millipede := &Millipede{
		Size:     20,
		Reverse:  false,
		Skin:     "default",
		Opposite: false,
		Width:    3,
		Curve:    4,
	}
	fmt.Println(millipede)
	// Output:
	//   ╚⊙ ⊙╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(20)
	}
}

func BenchmarkRendering(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s", New(20))
	}
}
