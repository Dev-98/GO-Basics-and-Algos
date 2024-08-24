package module1

import (
	// "fmt"
	"math"
)

var decimapper = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"A": 10,
	"B": 11,
	"C": 12,
	"D": 13,
	"E": 14,
	"F": 15,
}

func BaseToDecimal(num string, base int) (ans int) {
	n := len(num)
	for i := range n {
		val_int := decimapper[string(num[n-i-1])]
		// expo := int(math.Pow(float64(base), float64(i)))
		// fmt.Printf("this is val_int : %v \tThis is expo : %v\n", val_int, expo)
		ans += val_int * int(math.Pow(float64(base), float64(i)))
	}
	return ans
}
