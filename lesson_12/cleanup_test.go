package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestExampleCleanup(t *testing.T) {
	x := 0
	Convey("A", t, func() {
		x++
		Convey("AB", func() {
			x++
			Convey("ABC1", func() {
				So(x, ShouldEqual, 2)
			})
			Convey("ABC2", func() {
				So(x, ShouldEqual, 4)
			})
		})
		Reset(func() {
			t.Log("finish")
		})
	})

}
