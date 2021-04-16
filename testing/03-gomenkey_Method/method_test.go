package method

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMockMethod(t *testing.T) {
	Convey("TestMockMethod", t, func() {
		var p *myType
		fmt.Printf("method num:%d\n", reflect.TypeOf(p).NumMethod())
		p1 := gomonkey.ApplyMethod(reflect.TypeOf(p), "NetWorkFunc", func(_ *myType, a, b int) (int, error) {
			if a < 0 && b < 0 {
				errmsg := "a<0 && b<0"
				return 0, fmt.Errorf("%v", errmsg)
			}
			fmt.Println("mock method is running")
			return a * b, nil
		})
		defer p1.Reset()

		var m myType
		sum, err := m.logicFunc(10, 20)
		So(sum, ShouldEqual, 200)
		So(err, ShouldBeNil)
	})
}
