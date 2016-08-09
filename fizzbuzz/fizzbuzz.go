package fizzbuzz

import "fmt"

func RunFizzbuzz(n int) {
	for i := 0; i < n; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}