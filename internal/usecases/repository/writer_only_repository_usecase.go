package repository

import (
	"go-mux-mongodb-user-manager-api/internal/domain"
	"go-mux-mongodb-user-manager-api/internal/usecases/users_manager"
)

type WriterOnlyRepositoryUseCases interface {
	ExecCreate(model *domain.User) (users_manager.UserCreateOutput, error)
}
