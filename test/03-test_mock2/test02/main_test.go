package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/agiledragon/gomonkey"
)

func TestDoubleRight(t *testing.T) {
	gomonkey.ApplyFunc(Add, func(a,b int) int {
		fmt.Println("applyfunc is running")
		return a * 2
	})
	// defer patch.Reset()
	fmt.Println(GetDouble(2))
	Convey("test 2 x 2", t, func() {
		So(GetDouble(2), ShouldEqual, 4)
	})
}
