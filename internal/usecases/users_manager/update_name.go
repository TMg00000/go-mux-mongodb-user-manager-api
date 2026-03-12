package users_manager

import (
	"go-mux-mongodb-user-manager-api/internal/domain"
)

type UpdateNameServices struct {
	Collection      domain.MongoRepository
	HashingServices domain.HashingRepository
}

func NewUpdateNameServices(collection domain.MongoRepository, hashingServices domain.HashingRepository) *UpdateNameServices {
	return &UpdateNameServices{
		Collection:      collection,
		HashingServices: hashingServices,
	}
}

func (r *UpdateNameServices) ExecUpdateName(inputDto UserUpdateNameInput) (map[string]interface{}, error) {
	model := domain.NewUpdateName(
		inputDto.Email,
		inputDto.Password,
		inputDto.NewName,
	)

	result, err := r.Collection.UpdateName(model.Name, model.Email)
	if err != nil {
		return map[string]interface{}{}, err
	}

	err = r.HashingServices.ComparePassword(result.Password, model.Password)
	if err != nil {
		return map[string]interface{}{}, err
	}

	response := UserUpdateNameResponse{
		Id:      result.Id,
		Email:   result.Email,
		NewName: model.Name,
	}

	newResponse := map[string]interface{}{
		"update_success": map[string]interface{}{
			"user": map[string]interface{}{
				"id":    response.Id.Hex(),
				"email": response.Email,
				"new_name": map[string]string{
					"name": response.NewName,
				},
			},
		},
	}

	return newResponse, nil
}
