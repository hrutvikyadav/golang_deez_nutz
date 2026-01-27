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
	var summedSlices []int

	for _, numbers := range slicesToSum {
		summedSlices = append(summedSlices, Sum(numbers))
	}
	return summedSlices
}
