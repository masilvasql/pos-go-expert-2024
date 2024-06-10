package database

import "github.com/masilvasql/pos-go-expert-2024/APIs/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
