package main

func Sum(nums []int) (int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}


// INFO: variadic function to accept variable number of arguments
func SumAllTails(slicesToSum ...[]int) []int {
	var summedSlices []int

	for _, numbers := range slicesToSum {
		if len(numbers) == 0 {
			summedSlices = append(summedSlices, 0)
		} else {
			tail := numbers[1:]
			summedSlices = append(summedSlices, Sum(tail))
		}
	}
	return summedSlices
}
