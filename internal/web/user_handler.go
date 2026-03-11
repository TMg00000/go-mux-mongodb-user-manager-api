package web

import (
	"encoding/json"
	"go-mux-mongodb-user-manager-api/internal/domain"
	"go-mux-mongodb-user-manager-api/internal/usecases/repository"
	"go-mux-mongodb-user-manager-api/internal/usecases/users_manager"
	"net/http"
)

type UserHandlerRepository struct {
	usecaseWriterOnly repository.WriterOnlyRepositoryUseCases
	usecaseReadOnly   repository.ReadOnlyRepositoryUseCases
}

func NewUserHandlerRepository(writerOnly repository.WriterOnlyRepositoryUseCases, readOnly repository.ReadOnlyRepositoryUseCases) *UserHandlerRepository {
	return &UserHandlerRepository{
		usecaseWriterOnly: writerOnly,
		usecaseReadOnly:   readOnly,
	}
}

func (repo *UserHandlerRepository) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var inputDto users_manager.UserCreateInput

	if err := json.NewDecoder(r.Body).Decode(&inputDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := users_manager.UserCreateInputValidate(inputDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model := domain.NewUser(
		inputDto.Name,
		inputDto.Email,
		inputDto.Password,
	)

	response, err := repo.usecaseWriterOnly.ExecCreate(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"new_user": response,
	})
}

func (repo *UserHandlerRepository) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	response, err := repo.usecaseReadOnly.ExecGetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users": response,
	})
}
