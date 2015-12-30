package millipede

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSkin_SetDirection(t *testing.T) {
	Convey("Testing Skin.SetDirection", t, func() {
		skin, err := Skins.GetByName("default")
		So(err, ShouldBeNil)
		So(skin.SetDirection, ShouldNotBeNil)

		err = skin.SetDirection(DirectionUp)
		So(err, ShouldBeNil)
		So(skin.currentDirection, ShouldEqual, skin.Up)

		err = skin.SetDirection(DirectionDown)
		So(err, ShouldBeNil)
		So(skin.currentDirection, ShouldEqual, skin.Down)

		err = skin.SetDirection(DirectionLeft)
		So(err, ShouldNotBeNil)
		So(skin.currentDirection, ShouldBeNil)

		err = skin.SetDirection(DirectionRight)
		So(err, ShouldNotBeNil)
		So(skin.currentDirection, ShouldBeNil)
	})
}
