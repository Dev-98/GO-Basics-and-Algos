package module1

import "fmt"
import "testing"

func TestNumList(t *testing.T) {

	tests := []struct {
		list []int
		num  int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, 1, true},
		{[]int{1, 2, 3, 4, 5}, 2, true},
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 4, true},
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{1, 2, 3, 4, 5}, 0, false},
		{[]int{1, 2, 3, 4, -1}, -1, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 123, true},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 321, false},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 32, false},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 7, false},
		{[]int{8, 2, 5, 4, 1, 8, 9, 3, 0, 88, 23, 44, 123}, 6, false},
		{[]int{-1, -1, -1, -1, -1, -1, -1, -1}, -1, true},
		{[]int{-1, -1, -1, -1, -1, -1, -1, -1}, 1, false},
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{3, 3, 3, 3, 3}, 5, false},
		{[]int{3, 5, 3, 5, 3}, 5, true},
		{[]int{4, 2, 22, -10, 8}, -10, true},

		// empty lists!
		{nil, 5, false},
		{[]int{}, 5, false},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v,%v)", tc.list, tc.num), func(t *testing.T) {
			got := NumInList(tc.list, tc.num)
			if got != tc.want {
				t.Fatalf("NumInList() = %v; want %v", got, tc.want)
			}
		})
	}

}
