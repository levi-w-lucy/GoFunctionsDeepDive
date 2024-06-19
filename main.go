package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	sum := sumup(numbers...)
	fmt.Println(sum)
}

func sumup(numbers ...int) (sum int) {
	sum = 0
	for _, val := range numbers {
		sum += val
	}

	return
}
