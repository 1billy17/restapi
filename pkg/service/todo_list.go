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
func (s *TodoListService) GetListById(userId, listId int) (TODOapi.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *TodoListService) DeleteListById(userId, listId int) error {
	return s.repo.DeleteListById(userId, listId)
}

func (s *TodoListService) UpdateList(userId, listId int, input TODOapi.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateList(userId, listId, input)
}
