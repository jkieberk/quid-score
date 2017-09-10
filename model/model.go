package model

type db interface {
	SelectPeople() ([]*Person, error)
}

type Model struct {
	db
}

func New(db db) *Model {
	return &Model{
		db: db,
	}
}

func (m *Model) people() ([]*Person, error) {
	return m.SelectPeople()
}

type Person struct {
	id int64
	Fist, Last string
}
