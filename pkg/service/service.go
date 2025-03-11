package service

import (
	"TODOapi"
	"TODOapi/pkg/repository"
)

type Authorization interface {
	CreateUser(user TODOapi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateList(userId int, list TODOapi.TodoList) (int, error)
	GetAllLists(userId int) ([]TODOapi.TodoList, error)
	GetListById(userId, listId int) (TODOapi.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
