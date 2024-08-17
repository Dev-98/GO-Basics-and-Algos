package module1

import "fmt"

func FizzBuzz(num int) {
	for i := 1; i <= num; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Printf("Fizz Buzz\t")
		} else if i%5 == 0 {
			fmt.Printf("Buzz\t")
		} else if i%3 == 0 {
			fmt.Printf("Fizz\t")
		} else {
			fmt.Printf("%v\t", i)
		}
	}
}
