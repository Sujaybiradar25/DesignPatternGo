package main

import (
	product "DesignPatternGo/src/observer"
	"DesignPatternGo/src/observer/emailNotifier"
	"DesignPatternGo/src/observer/sms_notifier"
)

func main() {
	pen := product.New(10, "pen")
	book := product.New(20, "book")

	bookSubscriber1 := emailNotifier.New()
	bookSubscriber2 := sms_notifier.New()

	penSubscriber1 := sms_notifier.New()

	pen.AddSubscriber(penSubscriber1)
	pen.UpdatePrice(15)

	book.AddSubscriber(bookSubscriber1)
	book.AddSubscriber(bookSubscriber2)
	book.UpdatePrice(25)
	book.DeleteSubscriber(bookSubscriber1)
	book.UpdatePrice(20)
}
