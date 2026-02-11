package main

type Transaction struct {
	From string
	To string
	Sum float64
}

func BalanceFor(ts []Transaction, customer string) float64 {
	calcBal := func(acc float64, t Transaction) float64 {
		if t.From == customer {
			return acc - t.Sum
		}
		if (t.To == customer) {
			return acc + t.Sum
		}
		return acc
	}
	bal := Reduce(ts, calcBal, 0.0)
	return bal
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

func Reduce[A,B any](collection []A, f func(B, A) B, init B) B {
	var res = init
	for _, item := range collection {
		res = f(res, item)
	}

	return res
}
