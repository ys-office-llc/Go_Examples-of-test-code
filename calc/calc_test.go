package calc

import (
	"testing"
)

func TestMaxY(t *testing.T) {
	got := Max(1, 2)
	want := 2
	if got != want {
		t.Errorf("Max(1, 2) == %d, want %d", got, want)
	}
}

func TestMaxX(t *testing.T) {
	got := Max(2, 1)
	want := 2
	if got != want {
		t.Errorf("Max(2, 1) == %d, want %d", got, want)
	}
}
