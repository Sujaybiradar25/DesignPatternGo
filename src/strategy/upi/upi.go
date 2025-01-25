package upi

import "fmt"

type UPI struct {
}

func (u *UPI) Pay(amount int) (string, error) {
	fmt.Println("Payment of", amount, "received through UPI")
	return "", nil
}
