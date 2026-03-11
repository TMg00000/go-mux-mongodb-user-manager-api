package users_manager

import "go-mux-mongodb-user-manager-api/internal/domain"

type UserGetAllRepository struct {
	Collection domain.MongoRepository
}

func NewUserGetAllRepository(collection domain.MongoRepository) *UserGetAllRepository {
	return &UserGetAllRepository{
		Collection: collection,
	}
}

func (r *UserGetAllRepository) ExecGetAll() ([]UserGetAllOutput, error) {

	result, err := r.Collection.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]UserGetAllOutput, len(result))
	for i, val := range result {
		response[i] = UserGetAllOutput{
			Id:           val.Id,
			Name:         val.Name,
			Email:        val.Email,
			PasswordHash: val.PasswordHash,
		}
	}

	return response, nil
}
