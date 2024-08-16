package module1

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"cat", "tac"},
		{"malla", "allam"},
		{"dog", "god"},
		{"dev", "ved"},
		{"alphabet", "tebahpla"},
		{"日本語", "語本日"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.word), func(t *testing.T) {
			ans := Reverse(tc.word)
			if ans != tc.want {
				t.Fatalf("NumInList() = %v; want %v", ans, tc.want)
			}
		})

	}

}
