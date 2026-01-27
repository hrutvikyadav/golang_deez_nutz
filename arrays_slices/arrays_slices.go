package main

func Sum(nums []int) (int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}


// INFO: variadic function to accept variable number of arguments
func SumAll(slicesToSum ...[]int) []int {
	lengthOfArgs := len(slicesToSum)
	summedSlices := make([]int, lengthOfArgs)

	for i, numbers := range slicesToSum {
		summedSlices[i] = Sum(numbers)
	}
	return summedSlices
}
