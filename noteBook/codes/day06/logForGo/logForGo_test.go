package logForGo

import (
	"testing"
)

func TestInfo(t *testing.T) {
	l := NewLogForGo()
	_, err := l.Info("%s\n", "I am chinese")
	if err != nil {
		t.Fatal("format error: ", err)
	}
}
