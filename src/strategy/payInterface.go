package cart

type Payment interface {
	Pay(amount int) (string, error)
}
