package main

import "fmt"

func addTwoNums(son1, son2 int) (summa int) {
	summa = son1 + son2
	return summa
}
func getMaxNumber(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}
func getFactorial(N int) (factorial int) {
	factorial = 1
	for i := 1; i <= N; i++ {
		factorial = factorial * i
	}
	return factorial
}
func main() {
	fmt.Println("addTwoNums -", addTwoNums(5, 3))
	fmt.Println("getMaxNumber -", getMaxNumber(6, 5, 3))
	fmt.Println("getFactorial -", getFactorial(4))
}
