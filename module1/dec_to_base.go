package module1

import "strconv"

var Mapper = map[int]string{
	10: "A",
	11: "B",
	12: "C",
	13: "D",
	14: "E",
	15: "F",
}

func DecToBase(dec, base int) (ans string) {

	for n := dec; n > 0; n /= base {
		remainder := n % base
		if remainder >= 10 {
			ans = string(Mapper[remainder]) + ans
		} else {
			ans = strconv.Itoa(remainder) + ans
		}
	}
	return ans

}
