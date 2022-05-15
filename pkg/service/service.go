package service

import (
	"todo"
	"todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
