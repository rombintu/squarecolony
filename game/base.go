package game

type Base struct {
	Id       int
	Owner    int
	Capacity int
	Name     string
}

func newBase(id, owner, capacity int, name string) Base {
	return Base{id, owner, capacity, name}
}
