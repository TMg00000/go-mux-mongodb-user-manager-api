package users_manager

import (
	"go-mux-mongodb-user-manager-api/internal/domain"
)

type LoginUserServices struct {
	Collection      domain.MongoRepository
	HashingServices domain.HashingRepository
}

func NewLoginUserServices(collection domain.MongoRepository, hashingServices domain.HashingRepository) *LoginUserServices {
	return &LoginUserServices{
		Collection:      collection,
		HashingServices: hashingServices,
	}
}

func (r *LoginUserServices) ExecLogin(email, password string, model *domain.User) (UserGetByEmailResponse, error) {

	result, err := r.Collection.GetByEmail(email)
	if err != nil {
		return UserGetByEmailResponse{}, err
	}

	err = r.HashingServices.ComparePassword(result.Password, password)
	if err != nil {
		return UserGetByEmailResponse{}, err
	}

	response := UserGetByEmailResponse{
		Id:   result.Id,
		Name: result.Name,
	}

	return response, nil
}
