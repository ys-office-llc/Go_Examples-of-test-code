package calc_test

import (
	"testing"

	"github.com/ys-office-llc/Go_Examples-of-test-code/calc"
)

// サブテスト
func TestMaxSubtests(t *testing.T) {
	t.Run("正の数", func(t *testing.T) {
		got := calc.Max(1, 2)
		want := 2
		if got != want {
			t.Errorf("Max(1, 2) == %d, want %d", got, want)
		}
	})
	t.Run("負の数", func(t *testing.T) {
		got := calc.Max(-1, -2)
		want := -1
		if got != want {
			t.Errorf("Max(-1, -2) == %d, want %d", got, want)
		}
	})
	t.Run("両方", func(t *testing.T) {
		got := calc.Max(1, -2)
		want := 1
		if got != want {
			t.Errorf("Max(1, -2) == %d, want %d", got, want)
		}
	})
}
