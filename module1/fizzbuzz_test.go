package module1

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{2, "tac"},
		{3, "allam"},
		{5, "god"},
		{12, "ved"},
		{30, "tebahpla"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.num), func(t *testing.T) {
			FizzBuzz(tc.num)
			// if ans != tc.want {
			// 	t.Fatalf("Reverse() = %v; want %v", ans, tc.want)
			// }
		})

	}

}
