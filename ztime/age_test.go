package ztime

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseBirthday(t *testing.T) {
	Convey("测试生日是否通过", t, func() {
		Convey("分割符号为,且不规范", func() {
			y, m, d, err := ParseBirthday("1992,0101", ",")
			So(y, ShouldEqual, "1992")
			So(m, ShouldEqual, "01")
			So(d, ShouldEqual, "01")
			So(err, ShouldBeNil)
		})
		Convey("分割符号为-,规范", func() {
			y, m, d, err := ParseBirthday("1992-01-01", "-")
			So(y, ShouldEqual, "1992")
			So(m, ShouldEqual, "01")
			So(d, ShouldEqual, "01")
			So(err, ShouldBeNil)
		})

		Convey("分割符号为空", func() {
			y, m, d, err := ParseBirthday("19940409", "")
			So(y, ShouldEqual, "1994")
			So(m, ShouldEqual, "04")
			So(d, ShouldEqual, "09")
			So(err, ShouldBeNil)
		})

		Convey("生日位数不足", func() {
			_, _, _, err := ParseBirthday("199449", "")
			So(err, ShouldBeError)
		})
	})
}

func TestParseParseBirthdayInt(t *testing.T) {
	Convey("测试生日是否通过", t, func() {
		Convey("分割符号为,且不规范", func() {
			y, m, d, err := ParseBirthdayInt("1992,0101", ",")
			So(y, ShouldEqual, 1992)
			So(m, ShouldEqual, 1)
			So(d, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
		Convey("分割符号为-,规范", func() {
			y, m, d, err := ParseBirthdayInt("1992-01-01", "-")
			So(y, ShouldEqual, 1992)
			So(m, ShouldEqual, 1)
			So(d, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("分割符号为空", func() {
			y, m, d, err := ParseBirthdayInt("19940412", "")
			So(y, ShouldEqual, 1994)
			So(m, ShouldEqual, 4)
			So(d, ShouldEqual, 12)
			So(err, ShouldBeNil)
		})

		Convey("生日位数不足", func() {
			_, _, _, err := ParseBirthdayInt("199449", "")
			So(err, ShouldBeError)
		})
	})
}
