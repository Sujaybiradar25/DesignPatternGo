package sms_notifier

import "fmt"

type SMSNotifier struct {
	id int
}

var id int

func New() *SMSNotifier {
	s := SMSNotifier{}
	s.id = id
	id++
	return &s
}

func (s *SMSNotifier) Update(update string) {
	fmt.Println("updated sms notifier", update)
}

func (s *SMSNotifier) GetId() int {
	return (*s).id
}
