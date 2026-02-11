package main

type Transaction struct {
	From string
	To string
	Sum float64
}

func BalanceFor(ts []Transaction, customer string) float64 {
	var acc float64
	for _, t := range ts {
		if t.From == customer {
			acc -= t.Sum
		}
		if (t.To == customer) {
			acc += t.Sum
		}
	}
	return acc
}

func Sum(nums []int) (int) {
	addF := func(acc, elem int) int { return acc + elem }
	sum := Reduce(nums, addF, 0)

	return sum
}

func SumAllTails(slicesToSum ...[]int) []int {
	addTail := func(acc, elem []int) []int {
		if len(elem) == 0 {
			return append(acc, 0)
		} else {
			tail := elem[1:]
			return append(acc, Sum(tail))
		}
	}
	res := Reduce(slicesToSum, addTail, []int{})
	return res
}

func Reduce[A any](collection []A, f func(A, A) A, init A) A {
	var res = init
	for _, item := range collection {
		res = f(res, item)
	}

	return res
}
