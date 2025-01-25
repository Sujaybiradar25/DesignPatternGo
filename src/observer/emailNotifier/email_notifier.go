package emailNotifier

import "fmt"

type EmailNotifier struct {
	id int
}

var id int

func New() *EmailNotifier {
	e := EmailNotifier{}
	e.id = id
	id++
	return &e
}

func (e *EmailNotifier) GetId() int {
	return (*e).id
}

func (e *EmailNotifier) Update(update string) {
	fmt.Println("updated email notifier", update)
}
