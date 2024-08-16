package module1

// import "fmt"

func ReverseString(s string) (ans string) {
	last := len(s) - 1
	for i := range s {

		ans += string(s[last-i])
	}

	return ans
}

func Reverse(s string) (ans string) {
	for _, r := range s {
		ans = string(r) + ans
	}
	return ans
}
