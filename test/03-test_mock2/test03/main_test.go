package main

import (
	"fmt"
	"testing"

	. "bou.ke/monkey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	guard := Patch(Hello, func(_ string) string {
		fmt.Println("testing of Patch is running")
		return "Wangwu"
	})
	defer guard.Unpatch()
	Convey("test Hello", t, func() {
		res := Hello("Lisi")
		So(res, ShouldEqual, "Wangwu")
	})
}