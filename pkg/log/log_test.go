package logback

import "testing"

func TestInfo(t *testing.T) {

	l, _ := NewLogBack(".", INFO)
	l.Info("你好")
}
