package log

import "testing"

func TestLog(t *testing.T) {
	Log().Println("hello printf")
	Log().Infoln("hello info")
	Log().Errorln("hello error")
}
