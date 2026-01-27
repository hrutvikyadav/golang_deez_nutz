package main

type Wallet struct {
	balance int // private field with lowercase (can be accessed only with receiver methods)
}

// WARN: in order to modify the intended struct correctly, we need to pass by reference
// passing by value will create a copy and the method will act upon a copy.
// THIS WILL only cause issues if we are mutating the data, for reading only, pass by value will still work
// however it is convention to write all method receivers with pointers.
func (w *Wallet) Deposit(amount int) {
	w.balance += amount // struct pointers are automatically dereferenced;
	// we can also do it explicitly like `(*w).balance`
}

func (w *Wallet) Balance() (currentBalance int) {
	currentBalance = w.balance
	return
}
