package method

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMockMethod(t *testing.T) {
	Convey("test method seq", t, func() {
		var m *myType
		outputs := []gomonkey.OutputCell{
			{Values: gomonkey.Params{100, nil}, Times: 1},
			{Values: gomonkey.Params{200, nil}, Times: 1},
			{Values: gomonkey.Params{300, nil}, Times: 2},
		}
		var p1 = gomonkey.ApplyMethodSeq(reflect.TypeOf(m), "NetWorkFunc", outputs)
		defer p1.Reset()

		var mTest *myType
		sum, err := mTest.LogicFunc(2, 2)
		So(sum, ShouldEqual, 100)
		So(err, ShouldEqual, nil)

		sum, err = mTest.LogicFunc(2, 2)
		So(sum, ShouldEqual, 200)
		So(err, ShouldEqual, nil)

		sum, err = mTest.LogicFunc(2, 2)
		So(sum, ShouldEqual, 300)
		So(err, ShouldEqual, nil)

		sum, err = mTest.LogicFunc(2, 2)
		So(sum, ShouldEqual, 300)
		So(err, ShouldEqual, nil)
	})
}

func TestMockLogic(t *testing.T) {
	Convey("test method seq", t, func() {
		var m *myType
		outputs := []gomonkey.OutputCell{
			{Values: gomonkey.Params{100, nil}, Times: 1},
			{Values: gomonkey.Params{200, nil}, Times: 1},
			{Values: gomonkey.Params{300, nil}, Times: 2},
		}
		var p1 = gomonkey.ApplyMethodSeq(reflect.TypeOf(m), "SubLogicFunc", outputs)
		defer p1.Reset()

		var mTest *myType
		sum, err := mTest.SubLogicFunc(2, 2)
		So(sum, ShouldEqual, 100)
		So(err, ShouldEqual, nil)

		sum, err = mTest.SubLogicFunc(2, 2)
		So(sum, ShouldEqual, 200)
		So(err, ShouldEqual, nil)

		sum, err = mTest.SubLogicFunc(2, 2)
		So(sum, ShouldEqual, 300)
		So(err, ShouldEqual, nil)

		sum, err = mTest.SubLogicFunc(2, 2)
		So(sum, ShouldEqual, 300)
		So(err, ShouldEqual, nil)
	})
}
