package main

import "fmt"

/* Interface */
type Coin interface {
	// Abtract method
	getBalance() float64
}

/* Type implement interface */
type Bitcoin struct {
	price, amount float64
}

/* Type implement interface */
type Ripple struct {
	price, amount float64
}

/* Method of Bitcoin type */
func (b Bitcoin) getBalance() float64 {
	total := b.amount * b.price
	return total
}

/* Method of Ripple type */
func (r Ripple) getBalance() float64 {
	total := r.amount * r.price
	return total
}

/* Go though all types are implementing by interface */
func totalBalance(c ...Coin) {
	var total float64
	for _, v := range c {
		total += v.getBalance()
	}
	fmt.Println("Total balance is: $", total)
}

func main() {

	b := Bitcoin{3500, 2}
	r := Ripple{0.2, 1000}

	fmt.Println("Bitcoin balance is: $", b.getBalance())
	fmt.Println("Ripple balance is: $", r.getBalance())

	totalBalance(b, r)

}
