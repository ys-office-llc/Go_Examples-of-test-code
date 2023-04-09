package calc_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ys-office-llc/Go_Examples-of-test-code/calc"
)

// 前処理、後処理追加

func TestMain(t *testing.M) {
	fmt.Println("ここに前処理を書く")
	code := t.Run()
	fmt.Println("ここに後処理を書く")
	os.Exit(code)
}

// テーブルテスト
func TestMaxTableDrivenTests(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"正の数", 1, 2, 2},
		{"負の数", -1, -2, -1},
		{"両方", 1, -2, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Max(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("Max(%d, %d) == %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
