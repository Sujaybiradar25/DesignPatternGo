package product

import "fmt"

type Product struct {
	id          int
	description string
	price       int
	subscribers []Observer
}

func New(price int, description string) *Product {
	p := Product{}
	p.description = description
	p.price = price
	p.subscribers = make([]Observer, 0)
	return &p
}

func (p *Product) AddSubscriber(o Observer) {
	(*p).subscribers = append((*p).subscribers, o)
}

func (p *Product) DeleteSubscriber(o Observer) {
	for i := 0; i < len((*p).subscribers); i++ {
		if (*p).subscribers[i].GetId() == o.GetId() {
			suffix := (*p).subscribers[i+1 : len((*p).subscribers)]
			(*p).subscribers = (*p).subscribers[0:i]
			(*p).subscribers = append((*p).subscribers, suffix...)
			return
		}
	}
}

func (p *Product) UpdatePrice(price int) {
	(*p).price = price
	(*p).NotifyAll(fmt.Sprintf("Updated Price of %s is %d", (*p).description, price))
}

func (p *Product) NotifyAll(message string) {
	for i := 0; i < len((*p).subscribers); i++ {
		(*p).subscribers[i].Update(message)
	}
	return
}
