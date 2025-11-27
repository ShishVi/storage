package storage

type Storage struct {
	years int
	name  string
}

func NewStorage() *Storage {
	return &Storage{}
}
