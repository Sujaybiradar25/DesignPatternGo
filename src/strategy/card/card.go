package card

import "fmt"

type Card struct {
}

func (c *Card) Pay(amount int) (string, error) {
	fmt.Println("Payment of", amount, "received through Card")
	return "", nil
}
