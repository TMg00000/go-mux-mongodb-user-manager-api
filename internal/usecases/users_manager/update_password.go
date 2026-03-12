package users_manager

import (
	"go-mux-mongodb-user-manager-api/internal/domain"
)

type UpdatePasswordServices struct {
	Collection      domain.MongoRepository
	HashingServices domain.HashingRepository
}

func NewUpdatePasswordServices(collection domain.MongoRepository, hashingServices domain.HashingRepository) *UpdatePasswordServices {
	return &UpdatePasswordServices{
		Collection:      collection,
		HashingServices: hashingServices,
	}
}

func (r *UpdatePasswordServices) ExecUpdatePassword(inputDto UserUpdatePasswordInput) (map[string]interface{}, error) {
	model := domain.NewUpdatePassword(
		inputDto.Email,
		inputDto.Password,
		inputDto.NewPassword,
	)

	hash, err := r.HashingServices.HashPassword(model.NewPassword)
	if err != nil {
		return map[string]interface{}{}, err
	}
	model.NewPassword = hash

	result, err := r.Collection.GetByEmail(model.Email)
	if err != nil {
		return map[string]interface{}{}, err
	}

	err = r.HashingServices.ComparePassword(result.Password, model.Password)
	if err != nil {
		return map[string]interface{}{}, err
	}

	err = r.Collection.UpdatePassword(model.NewPassword, model.Email)
	if err != nil {
		return map[string]interface{}{}, err
	}

	response := UserUpdatePasswordResponse{
		Id:    result.Id,
		Email: result.Email,
	}

	newResponse := map[string]interface{}{
		"update_success": map[string]interface{}{
			"user": map[string]interface{}{
				"id":           response.Id.Hex(),
				"email":        response.Email,
				"new_password": "Success !",
			},
		},
	}

	return newResponse, nil
}
