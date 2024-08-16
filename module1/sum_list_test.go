package module1

import (
	"fmt"
	"testing"
)

func TestSumList(t *testing.T) {
	tests := []struct {
		list []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-11, 2, 3, 4}, -2},
		{[]int{}, 0},
		{[]int{-1, -1}, -2},
		{[]int{-1, -1, 0, 0, 0}, -2},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, 0, 1}, 0},
		{[]int{1, 0, -1}, 0},
		{[]int{43, 67, 65, 42, 74, 95, 100, 21}, 507},
		{[]int{26, 26, 94, 55, 71, 57, 7, 72}, 408},
		{[]int{78, 11, 48, 91, 13, 26, 74, 99}, 440},
		{[]int{31, 96, 10, 55, 41, 13, 92, 5}, 343},
		{[]int{14, 39, 18, 39, 3, 29, 8, 29}, 179},
		{[]int{-95, -46, -65, -63, 10}, -259},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.list), func(t *testing.T) {
			ans := SumListWithRecursive(tc.list)
			if ans != tc.want {
				t.Fatalf("NumInList() = %v; want %v", ans, tc.want)
			}
		})

	}

}
