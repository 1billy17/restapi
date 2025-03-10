package repository

import (
	"TODOapi"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user TODOapi.User) (int, error)
	GenerateToken(username, password string) (TODOapi.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
