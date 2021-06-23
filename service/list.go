package service

import (
	"github.com/Goganad/TodoList-REST-API/entities"
	"github.com/Goganad/TodoList-REST-API/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list entities.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
