package millipede

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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

func ExampleMillipede() {
	millipede := &Millipede{
		Size:      20,
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
	fmt.Printf("%s\n", millipede)
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

func ExampleMillipede_String() {
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

func ExampleMillipede_String_reverse() {
	millipede := New(20)
	millipede.Reverse = true

	fmt.Println(millipede)
	// Output:
	//    ╔═(███)═╗
	//   ╔═(███)═╗
	//  ╔═(███)═╗
	// ╔═(███)═╗
	//  ╔═(███)═╗
	//   ╔═(███)═╗
	//    ╔═(███)═╗
	//     ╔═(███)═╗
	//    ╔═(███)═╗
	//   ╔═(███)═╗
	//  ╔═(███)═╗
	// ╔═(███)═╗
	//  ╔═(███)═╗
	//   ╔═(███)═╗
	//    ╔═(███)═╗
	//     ╔═(███)═╗
	//    ╔═(███)═╗
	//   ╔═(███)═╗
	//  ╔═(███)═╗
	// ╔═(███)═╗
	//   ╔⊙ ⊙╗
}

func ExampleMillipede_String_center() {
	millipede := New(20)
	millipede.Center = true

	// FIXME: if possible make a better test

	output := fmt.Sprint(millipede)
	fmt.Println(output)
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

func ExampleMillipede_String_opposite() {
	millipede := New(20)
	millipede.Opposite = true

	fmt.Println(millipede)
	// Output:
	//       ╚⊙ ⊙╝
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
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
}

func ExampleMillipede_String_skin() {
	millipede := New(20)
	millipede.Skin = "bocal"

	fmt.Println(millipede)
	// Output:
	//   ╚⊙ ⊙╝
	// ╚═(🐟🐟🐟)═╝
	//  ╚═(🐟🐟🐟)═╝
	//   ╚═(🐟🐟🐟)═╝
	//    ╚═(🐟🐟🐟)═╝
	//     ╚═(🐟🐟🐟)═╝
	//    ╚═(🐟🐟🐟)═╝
	//   ╚═(🐟🐟🐟)═╝
	//  ╚═(🐟🐟🐟)═╝
	// ╚═(🐟🐟🐟)═╝
	//  ╚═(🐟🐟🐟)═╝
	//   ╚═(🐟🐟🐟)═╝
	//    ╚═(🐟🐟🐟)═╝
	//     ╚═(🐟🐟🐟)═╝
	//    ╚═(🐟🐟🐟)═╝
	//   ╚═(🐟🐟🐟)═╝
	//  ╚═(🐟🐟🐟)═╝
	// ╚═(🐟🐟🐟)═╝
	//  ╚═(🐟🐟🐟)═╝
	//   ╚═(🐟🐟🐟)═╝
	//    ╚═(🐟🐟🐟)═╝
}

func ExampleMillipede_String_width() {
	millipede := New(20)
	millipede.Width = 6

	fmt.Println(millipede)
	// Output:
	//   ╚⊙    ⊙╝
	// ╚═(██████)═╝
	//  ╚═(██████)═╝
	//   ╚═(██████)═╝
	//    ╚═(██████)═╝
	//     ╚═(██████)═╝
	//    ╚═(██████)═╝
	//   ╚═(██████)═╝
	//  ╚═(██████)═╝
	// ╚═(██████)═╝
	//  ╚═(██████)═╝
	//   ╚═(██████)═╝
	//    ╚═(██████)═╝
	//     ╚═(██████)═╝
	//    ╚═(██████)═╝
	//   ╚═(██████)═╝
	//  ╚═(██████)═╝
	// ╚═(██████)═╝
	//  ╚═(██████)═╝
	//   ╚═(██████)═╝
	//    ╚═(██████)═╝
}

func ExampleMillipede_String_rainbow() {
	millipede := New(20)
	millipede.Rainbow = true

	fmt.Println(millipede)
	// Output:
	//   [30;41m╚⊙ ⊙╝[0m
	// [30;42m╚═(███)═╝[0m
	//  [30;43m╚═(███)═╝[0m
	//   [30;44m╚═(███)═╝[0m
	//    [30;45m╚═(███)═╝[0m
	//     [30;46m╚═(███)═╝[0m
	//    [30;47m╚═(███)═╝[0m
	//   [30;41m╚═(███)═╝[0m
	//  [30;42m╚═(███)═╝[0m
	// [30;43m╚═(███)═╝[0m
	//  [30;44m╚═(███)═╝[0m
	//   [30;45m╚═(███)═╝[0m
	//    [30;46m╚═(███)═╝[0m
	//     [30;47m╚═(███)═╝[0m
	//    [30;41m╚═(███)═╝[0m
	//   [30;42m╚═(███)═╝[0m
	//  [30;43m╚═(███)═╝[0m
	// [30;44m╚═(███)═╝[0m
	//  [30;45m╚═(███)═╝[0m
	//   [30;46m╚═(███)═╝[0m
	//    [30;47m╚═(███)═╝[0m
}

func ExampleMillipede_String_chameleon() {
	millipede := New(20)
	millipede.Chameleon = true

	fmt.Println(millipede)
	// Output:
	//   [30m╚⊙ ⊙╝[0m
	// [30m╚═(███)═╝[0m
	//  [30m╚═(███)═╝[0m
	//   [30m╚═(███)═╝[0m
	//    [30m╚═(███)═╝[0m
	//     [30m╚═(███)═╝[0m
	//    [30m╚═(███)═╝[0m
	//   [30m╚═(███)═╝[0m
	//  [30m╚═(███)═╝[0m
	// [30m╚═(███)═╝[0m
	//  [30m╚═(███)═╝[0m
	//   [30m╚═(███)═╝[0m
	//    [30m╚═(███)═╝[0m
	//     [30m╚═(███)═╝[0m
	//    [30m╚═(███)═╝[0m
	//   [30m╚═(███)═╝[0m
	//  [30m╚═(███)═╝[0m
	// [30m╚═(███)═╝[0m
	//  [30m╚═(███)═╝[0m
	//   [30m╚═(███)═╝[0m
	//    [30m╚═(███)═╝[0m
}

func ExampleMillipede_String_chameleon_rainbow() {
	millipede := New(20)
	millipede.Chameleon = true
	millipede.Rainbow = true

	fmt.Println(millipede)
	// Output:
	//   [30;41m╚⊙ ⊙╝[0m
	// [30;42m╚═(███)═╝[0m
	//  [30;43m╚═(███)═╝[0m
	//   [30;44m╚═(███)═╝[0m
	//    [30;45m╚═(███)═╝[0m
	//     [30;46m╚═(███)═╝[0m
	//    [30;47m╚═(███)═╝[0m
	//   [30;41m╚═(███)═╝[0m
	//  [30;42m╚═(███)═╝[0m
	// [30;43m╚═(███)═╝[0m
	//  [30;44m╚═(███)═╝[0m
	//   [30;45m╚═(███)═╝[0m
	//    [30;46m╚═(███)═╝[0m
	//     [30;47m╚═(███)═╝[0m
	//    [30;41m╚═(███)═╝[0m
	//   [30;42m╚═(███)═╝[0m
	//  [30;43m╚═(███)═╝[0m
	// [30;44m╚═(███)═╝[0m
	//  [30;45m╚═(███)═╝[0m
	//   [30;46m╚═(███)═╝[0m
	//    [30;47m╚═(███)═╝[0m
}

func TestMillipede_String_padright(t *testing.T) {
	millipede := New(20)
	millipede.PadRight = true

	output := fmt.Sprint(millipede)

	expected := "  ╚⊙ ⊙╝    \n" +
		"╚═(███)═╝    \n" +
		" ╚═(███)═╝    \n" +
		"  ╚═(███)═╝    \n" +
		"   ╚═(███)═╝    \n" +
		"    ╚═(███)═╝    \n" +
		"   ╚═(███)═╝    \n" +
		"  ╚═(███)═╝    \n" +
		" ╚═(███)═╝    \n" +
		"╚═(███)═╝    \n" +
		" ╚═(███)═╝    \n" +
		"  ╚═(███)═╝    \n" +
		"   ╚═(███)═╝    \n" +
		"    ╚═(███)═╝    \n" +
		"   ╚═(███)═╝    \n" +
		"  ╚═(███)═╝    \n" +
		" ╚═(███)═╝    \n" +
		"╚═(███)═╝    \n" +
		" ╚═(███)═╝    \n" +
		"  ╚═(███)═╝    \n" +
		"   ╚═(███)═╝    "

	if output != expected {
		t.Fatalf("Invalid output for millipede with PadRight=true")
	}
}

func ExampleMillipede_String_curve() {
	millipede := New(20)
	millipede.Curve = 6

	fmt.Println(millipede)
	// Output:
	//   ╚⊙ ⊙╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//      ╚═(███)═╝
	//       ╚═(███)═╝
	//      ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//      ╚═(███)═╝
	//       ╚═(███)═╝
	//      ╚═(███)═╝
}

func ExampleMillipede_String_curve_zero() {
	millipede := New(20)
	millipede.Curve = 0

	fmt.Println(millipede)
	// Output:
	//   ╚⊙ ⊙╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
	// ╚═(███)═╝
}

func TestNew(t *testing.T) {
	Convey("Testing New()", t, func() {
		millipede := New(42)

		So(millipede, ShouldNotBeNil)
		So(millipede.Size, ShouldEqual, 42)
		So(millipede.Reverse, ShouldEqual, false)
		So(millipede.Skin, ShouldEqual, "default")
		So(millipede.Opposite, ShouldEqual, false)
		So(millipede.Width, ShouldEqual, 3)
		So(millipede.Curve, ShouldEqual, 4)
		So(millipede.Chameleon, ShouldEqual, false)
		So(millipede.Rainbow, ShouldEqual, false)
		So(millipede.Zalgo, ShouldEqual, false)
		So(millipede.Steps, ShouldEqual, 0)
		So(millipede.Center, ShouldEqual, false)

		millipede = New(43)
		So(millipede, ShouldNotBeNil)
		So(millipede.Size, ShouldEqual, 43)
		So(millipede.Reverse, ShouldEqual, false)
		So(millipede.Skin, ShouldEqual, "default")
		So(millipede.Opposite, ShouldEqual, false)
		So(millipede.Width, ShouldEqual, 3)
		So(millipede.Curve, ShouldEqual, 4)
		So(millipede.Chameleon, ShouldEqual, false)
		So(millipede.Rainbow, ShouldEqual, false)
		So(millipede.Zalgo, ShouldEqual, false)
		So(millipede.Steps, ShouldEqual, 0)
		So(millipede.Center, ShouldEqual, false)
	})
}

func TestMillipede_String_zalgo(t *testing.T) {
	// FIXME: find a better test
	millipede := New(20)
	millipede.Zalgo = true
	millipede.String()
}

func TestMillipede_String_smallwidth(t *testing.T) {
	millipede := New(20)
	millipede.Width = 2
	// FIXME: check if it exits
	//millipede.String()
}

func ExampleMillipede_String_complex() {
	millipede := New(20)
	millipede.Size = 42
	millipede.Reverse = true
	millipede.Skin = "bocal"
	millipede.Opposite = true
	millipede.Width = 6
	millipede.Curve = 10

	fmt.Println(millipede)
	// Output:
	//          ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//           ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//          ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//         ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//        ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//       ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//      ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//     ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//    ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//   ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//  ╔═(🐟🐟🐟🐟🐟🐟)═╗
	// ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//  ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//   ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//    ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//     ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//      ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//       ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//        ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//         ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//          ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//           ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//          ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//         ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//        ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//       ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//      ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//     ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//    ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//   ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//  ╔═(🐟🐟🐟🐟🐟🐟)═╗
	// ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//  ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//   ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//    ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//     ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//      ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//       ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//        ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//         ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//          ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//           ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//             ╔⊙    ⊙╗
}

func TestStringToRuneSlice(t *testing.T) {
	Convey("Testing StringToRunSlice()", t, func() {
		output := StringToRuneSlice("╔═(🐟🐟🐟🐟🐟🐟)═╗")
		So(output, ShouldResemble, []rune{9556, 9552, 40, 128031, 128031, 128031, 128031, 128031, 128031, 41, 9552, 9559})
		So(len(output), ShouldEqual, 12)

		output = StringToRuneSlice("HELLO WORLD !")
		So(output, ShouldResemble, []rune{72, 69, 76, 76, 79, 32, 87, 79, 82, 76, 68, 32, 33})
		So(len(output), ShouldEqual, 13)
	})
}

func ExampleStringToRuneSlice() {
	input := "╔═(🐟🐟🐟🐟🐟🐟)═╗"
	output := StringToRuneSlice(input)
	fmt.Println(output)
	// Output: [9556 9552 40 128031 128031 128031 128031 128031 128031 41 9552 9559]
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(20)
	}
}

func BenchmarkMillipede_String(b *testing.B) {
	millipede := New(20)
	for i := 0; i < b.N; i++ {
		millipede.String()
	}
}
