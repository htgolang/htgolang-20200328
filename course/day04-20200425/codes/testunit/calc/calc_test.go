package calc

import "testing"

func TestAdd01(t *testing.T) {
	if 3 != Add(1, 2) {
		t.Error("1 + 2 != 3")
	}
}

func TestFlag02(t *testing.T) {
	if -1 != Flag(-1000) {
		t.Error("flag -1000 != -1")
	}
}
