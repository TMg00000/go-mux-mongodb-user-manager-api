package users_manager

import (
	"go-mux-mongodb-user-manager-api/internal/domain"
)

type CreateUserServices struct {
	Collection        domain.MongoRepository
	HashingRepository domain.HashingRepository
}

func NewCreateUserServices(collection domain.MongoRepository, hashingRepository domain.HashingRepository) *CreateUserServices {
	return &CreateUserServices{
		Collection:        collection,
		HashingRepository: hashingRepository,
	}
}

func (r *CreateUserServices) ExecCreate(model *domain.User) (UserCreateResponse, error) {

	newHash, err := r.HashingRepository.HashPassword(model.Password)
	if err != nil {
		return UserCreateResponse{}, err
	}
	model.PasswordHash = newHash

	err = r.Collection.Create(model)
	if err != nil {
		return UserCreateResponse{}, err
	}

	response := UserCreateResponse{
		Name:  model.Name,
		Email: model.Email,
	}

	return response, nil
}
