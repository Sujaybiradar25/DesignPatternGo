package main

import (
	"DesignPatternGo/src/strategy"
	"DesignPatternGo/src/strategy/upi"
)

func main() {
	paymentMethod := upi.UPI{}
	c := cart.New(&paymentMethod)
	item1 := cart.Item{
		Id:   1,
		Name: "Hello",
		Cost: 1,
	}
	item2 := cart.Item{
		Id:   2,
		Name: "World",
		Cost: 10,
	}
	c.AddItems(item1, item2)
	c.MakePayment()
}
