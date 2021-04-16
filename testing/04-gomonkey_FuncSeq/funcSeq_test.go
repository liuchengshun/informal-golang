package funcSeq

import (
	"testing"

	"github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMockFunc(t *testing.T) {
	Convey("func seq", t, func() {
		outputs := []gomonkey.OutputCell{
			{Values: gomonkey.Params{2}, Times:1},
			{Values: gomonkey.Params{10}, Times:1},
			{Values: gomonkey.Params{50}, Times:1},
		}
		var p1 = gomonkey.ApplyFuncSeq(getInt, outputs)
		defer p1.Reset()

		So(getInt(), ShouldEqual, 2)
		So(getInt(), ShouldEqual, 10)
		So(getInt(), ShouldEqual, 50)
	})
}