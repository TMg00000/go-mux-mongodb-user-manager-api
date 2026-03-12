package users_manager

import "go-mux-mongodb-user-manager-api/internal/domain"

type DeleteUserByEmailServices struct {
	Collection      domain.MongoRepository
	HashingServices domain.HashingRepository
}

func NewDeleteUserByEmailServices(collection domain.MongoRepository, hashingServices domain.HashingRepository) *DeleteUserByEmailServices {
	return &DeleteUserByEmailServices{
		Collection:      collection,
		HashingServices: hashingServices,
	}
}

func (r *DeleteUserByEmailServices) ExecDeleteUserByEmail(inputDto UserDeleteUserByEmailInput) (UserDeleteUserByEmailResponse, error) {
	model := domain.NewDeleteUserByEmail(
		inputDto.Email,
		inputDto.Password,
		inputDto.ConfirmPassword,
	)

	result, err := r.Collection.GetByEmail(model.Email)
	if err != nil {
		return UserDeleteUserByEmailResponse{}, err
	}

	err = r.HashingServices.ComparePassword(result.Password, model.Password)
	if err != nil {
		return UserDeleteUserByEmailResponse{}, err
	}

	err = r.HashingServices.ComparePassword(result.Password, model.ConfirmPassword)
	if err != nil {
		return UserDeleteUserByEmailResponse{}, err
	}

	err = r.Collection.DeleteUserByEmail(model.Email)
	if err != nil {
		return UserDeleteUserByEmailResponse{}, err
	}

	response := UserDeleteUserByEmailResponse{
		Id:    result.Id,
		Name:  result.Name,
		Email: result.Email,
	}

	return response, nil
}
