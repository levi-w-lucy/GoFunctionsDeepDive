package main

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := make([]int, 4)
	for i := 0; i < 4; i++ {
		numbers[i] = i + 1
	}

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
