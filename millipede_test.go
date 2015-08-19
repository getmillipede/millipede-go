package millipede

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleNewWithSize() {
	millipede := NewWithSize(20)
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

func ExampleNew() {
	millipede := New()
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
	millipede := NewWithSize(20)
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
	millipede := NewWithSize(20)
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
	millipede := NewWithSize(20)
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
	millipede := NewWithSize(20)
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
	millipede := NewWithSize(20)
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
	millipede := NewWithSize(20)
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
	millipede := NewWithSize(20)
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
	millipede := NewWithSize(20)
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
	millipede := NewWithSize(20)
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
	millipede := NewWithSize(20)
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
	millipede := NewWithSize(20)
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
	millipede := NewWithSize(20)
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

func TestNewWithSize(t *testing.T) {
	Convey("Testing NewWithSize()", t, func() {
		Convey("Size = 42", func() {
			millipede := NewWithSize(42)

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
		})

		Convey("Size = 43", func() {
			millipede := NewWithSize(43)
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
	})
}

func TestNew(t *testing.T) {
	Convey("Testing New()", t, func() {
		Convey("Default values", func() {
			millipede := New()

			So(millipede, ShouldNotBeNil)
			So(millipede.Size, ShouldEqual, 20)
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
		Convey("Compare with NewWithSize()", func() {
			So(New(), ShouldResemble, NewWithSize(20))
		})
	})
}

func TestMillipede_Draw(t *testing.T) {
	Convey("Testing Millipede.Draw", t, func() {
		millipede := New()
		str, err := millipede.Draw()
		So(err, ShouldBeNil)
		So(str, ShouldNotBeNil)
		expected := `
  ╚⊙ ⊙╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
    ╚═(███)═╝
   ╚═(███)═╝
  ╚═(███)═╝
 ╚═(███)═╝
╚═(███)═╝
 ╚═(███)═╝
  ╚═(███)═╝
   ╚═(███)═╝
`
		So(fmt.Sprintf("\n%s\n", str), ShouldEqual, expected)
	})
}

func TestMillipede_String_zalgo(t *testing.T) {
	// FIXME: find a better test
	millipede := NewWithSize(20)
	millipede.Zalgo = true
	millipede.String()
}

func TestMillipede_String_smallwidth(t *testing.T) {
	millipede := NewWithSize(20)
	millipede.Width = 2
	// FIXME: check if it exits
	//millipede.String()
}

func TestMillipede_Draw_nosuchskin(t *testing.T) {
	Convey("Testing Millipede.Draw() with invalid skin", t, func() {
		millipede := New()
		millipede.Skin = "idontexist"
		str, err := millipede.Draw()
		So(str, ShouldEqual, "")
		So(err, ShouldNotBeNil)
		So(fmt.Sprintf("%s", err), ShouldEqual, "no such skin: 'idontexist'")
	})
}

func TestMillipede_String_nosuchskin(t *testing.T) {
	Convey("Testing Millipede.String() with invalid skin", t, func() {
		millipede := New()
		millipede.Skin = "idontexist"
		str := millipede.String()
		So(str, ShouldEqual, "")
	})
}

func TestMillipede_Draw_smallwidth(t *testing.T) {
	Convey("Testing Millipede.Draw() with valid width", t, func() {
		millipede := New()
		millipede.Width = 5
		str, err := millipede.Draw()
		So(str, ShouldNotEqual, "")
		So(err, ShouldBeNil)
	})
	Convey("Testing Millipede.Draw() with invalid width", t, func() {
		millipede := New()
		millipede.Width = 2
		str, err := millipede.Draw()
		So(str, ShouldEqual, "")
		So(err, ShouldNotBeNil)
		So(fmt.Sprintf("%s", err), ShouldEqual, "millipede cannot have a width < 3")
	})
}

func ExampleMillipede_String_complex() {
	millipede := NewWithSize(20)
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

func BenchmarkNewWithSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewWithSize(20)
	}
}

func BenchmarkMillipede_String(b *testing.B) {
	millipede := NewWithSize(20)
	for i := 0; i < b.N; i++ {
		millipede.String()
	}
}
