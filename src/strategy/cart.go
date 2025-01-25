package cart

import "fmt"

type Item struct {
	Id   int
	Name string
	Cost int
}
type Cart struct {
	paymentMethod Payment
	items         []Item
}

func New(paymentMethod Payment) *Cart {
	c := Cart{
		paymentMethod: paymentMethod,
		items:         make([]Item, 0),
	}
	return &c
}

func (c *Cart) AddItems(items ...Item) {
	(*c).items = append((*c).items, items...)
}

func (c *Cart) MakePayment() {
	total := 0
	for i := 0; i < len((*c).items); i++ {
		total += (*c).items[i].Cost
	}
	resp, err := (*c).paymentMethod.Pay(total)
	fmt.Println(resp, err)
}
