package ping

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPing(t *testing.T) {
	Convey("Ping", t, func() {
		//func Ping(address string, timeout int) bool
		So(Ping("8.8.8.8",1), ShouldBeTrue)
		So(Ping("1.2.3.4",1), ShouldBeFalse)
	})

}