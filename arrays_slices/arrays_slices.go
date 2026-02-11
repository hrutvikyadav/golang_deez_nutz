package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
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
