package service

import (
	"github.com/Goganad/TodoList-REST-API/entities"
	"github.com/Goganad/TodoList-REST-API/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	ListRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, ListRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item entities.TodoItem) (int, error) {
	_, err := s.ListRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]entities.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (entities.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}
