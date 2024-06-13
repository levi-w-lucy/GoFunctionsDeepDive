package functionsAsValues

type TransformFunc func(int) int

func GetTransformerFunction(numbers *[]int) TransformFunc {
	if (*numbers)[0] == 1 {
		return Double
	}
	return Triple
}

func Double(number int) int {
	return number * 2
}

func Triple(number int) int {
	return number * 3
}
