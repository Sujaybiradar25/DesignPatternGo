package product

type ProductInterface interface {
	AddSubscriber(o Observer)
	DeleteSubscriber(o Observer)
	UpdatePrice(price int)
	NotifyAll(message string)
}
