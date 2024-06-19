package main

import "fmt"

func main() {
	sum := sumup(1, 2, 3)
	fmt.Println(sum)
}

func sumup(numbers ...int) (sum int) {
	sum = 0
	for _, val := range numbers {
		sum += val
	}

	return
}
