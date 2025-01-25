package product

type Observer interface {
	Update(string)
	GetId() int
}
