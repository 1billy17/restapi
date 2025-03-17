package service

import (
	"TODOapi"
	"TODOapi/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) CreateItem(userId, listId int, input TODOapi.TodoItem) (int, error) {
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.CreateItem(listId, input)
}

func (s *TodoItemService) GetAllItems(userId, listId int) ([]TODOapi.TodoItem, error) {
	return s.repo.GetAllItems(userId, listId)
}
