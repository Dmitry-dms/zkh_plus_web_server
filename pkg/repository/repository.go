package repository

type Authorization interface {
}
type UserList interface {
}

type Repository struct {
	Authorization
	UserList
}

func NewRepository() *Repository {
	return &Repository{}
}
