package ping

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"syscall"
)

func TestPing(t *testing.T) {
	// we try if we have setuid bit in Linux
	syscall.Setuid(0)
	Convey("Ping", t, func() {
		//func Ping(address string, timeout int) bool
		So(Ping("8.8.8.8",1000), ShouldBeTrue)
		So(Ping("1.2.3.4",1000), ShouldBeFalse)
	})

}