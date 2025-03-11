package service

import (
	"TODOapi"
	"TODOapi/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateList(userId int, list TODOapi.TodoList) (int, error) {
	return s.repo.CreateList(userId, list)
}

func (s *TodoListService) GetAllLists(userId int) ([]TODOapi.TodoList, error) {
	return s.repo.GetAllLists(userId)
}
