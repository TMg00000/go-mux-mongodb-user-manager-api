package users_manager

import "go-mux-mongodb-user-manager-api/internal/domain"

type GetAllServices struct {
	Collection domain.MongoRepository
}

func NewGetAllServices(collection domain.MongoRepository) *GetAllServices {
	return &GetAllServices{
		Collection: collection,
	}
}

func (r *GetAllServices) ExecGetAll() ([]UserGetAllResponse, error) {

	result, err := r.Collection.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]UserGetAllResponse, len(result))
	for i, val := range result {
		response[i] = UserGetAllResponse{
			Id:    val.Id,
			Name:  val.Name,
			Email: val.Email,
		}
	}

	return response, nil
}
