package main

import "fmt"

func main() {

	fmt.Println("hello world")
	fmt.Println(factorial(3))
}

func factorial(num int) int {
	ans := 1
	for i := 1; i <= num; i++ {
		ans *= i
	}
	return ans
}
