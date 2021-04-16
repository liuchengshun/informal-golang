package foo

import (
	"fmt"
	"testing"

	"github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

// func TestMockFunc(t *testing.T) {
// 	Convey("TestMockFunc1", t, func() {
// 		var p1 = gomonkey.ApplyFunc(netWorkFunc, func(a, b int) (int, error) {
// 			fmt.Println("in mock function")
// 			return a + b, nil
// 		})
// 		defer p1.Reset()

// 		sum, err := logicFunc(10, 20)
// 		So(sum, ShouldEqual, 30)
// 		So(err, ShouldBeNil)
// 	})
// }

func TestLogicFunc(t *testing.T) {
	Convey("test logic func\n", t, func() {
		var p1 = gomonkey.ApplyFunc(insertData, func(a, b int) int {
			fmt.Println("insertData mock is running")
			return a * 100
		})
		defer p1.Reset()

		var p2 = gomonkey.ApplyFunc(netWorkFunc, func(a, b int) (int, error) {
			fmt.Println("netWorkFunc of mock is running")
			return a + b, nil
		})
		defer p2.Reset()

		sum, err := logicFunc(10, 20)
		So(sum, ShouldEqual, 30)
		So(err, ShouldBeNil)
	})
}
