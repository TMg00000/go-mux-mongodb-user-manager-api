package repository

import (
	"go-mux-mongodb-user-manager-api/internal/usecases/users_manager"
)

type ReadOnlyRepositoryUseCases interface {
	ExecGetAll() ([]users_manager.UserGetAllOutput, error)
}
