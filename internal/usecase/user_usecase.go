package usecase

import (
	"cashier-backend/internal/domain/model"
	"cashier-backend/internal/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]model.User, error)
}

type userUsecaseImpl struct {
	userRepo repository.UserRepository
}

// Constructor
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepo: userRepo}
}

// Implementasi GetAllUsers
func (u *userUsecaseImpl) GetAllUsers() ([]model.User, error) {
	return u.userRepo.FetchAll()
}
