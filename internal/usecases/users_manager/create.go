package users_manager

import (
	"go-mux-mongodb-user-manager-api/internal/domain"
)

type UseCaseRepository struct {
	Collection     domain.MongoRepository
	HashingService domain.HashingService
}

func NewUseCaseRepository(collection domain.MongoRepository, hashingService domain.HashingService) *UseCaseRepository {
	return &UseCaseRepository{
		Collection:     collection,
		HashingService: hashingService,
	}
}

func (r *UseCaseRepository) ExecCreate(model *domain.User) (UserCreateOutput, error) {

	newHash, err := r.HashingService.HashPassword(model.Password)
	if err != nil {
		return UserCreateOutput{}, err
	}
	model.PasswordHash = newHash

	err = r.Collection.Create(model)
	if err != nil {
		return UserCreateOutput{}, err
	}

	response := UserCreateOutput{
		Name:  model.Name,
		Email: model.Email,
	}

	return response, nil
}
