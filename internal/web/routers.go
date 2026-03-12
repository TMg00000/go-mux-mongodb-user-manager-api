package web

import (
	"fmt"
	"go-mux-mongodb-user-manager-api/pkg/configs"
	"net/http"

	"github.com/gorilla/mux"
)

func Routers(usecase *UserUseCasesRepository) error {
	r := mux.NewRouter()

	r.HandleFunc("/api/user", usecase.CreateNewUser).Methods("POST")
	r.HandleFunc("/api/user", usecase.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/user/login", usecase.LoginUser).Methods("POST")
	r.HandleFunc("/api/user/update_name", usecase.UpdateName).Methods("PUT")
	r.HandleFunc("/api/user/update_email", usecase.UpdateEmail).Methods("PUT")
	r.HandleFunc("/api/user/update_password", usecase.UpdatePassword).Methods("PUT")
	r.HandleFunc("/api/user/delete_user", usecase.DeleteUserByEmail).Methods("DELETE")

	if err := http.ListenAndServe(":"+configs.Env.Port, r); err != nil {
		return err
	}

	fmt.Println("server running on port " + configs.Env.Port)
	fmt.Println("server running on address http://localhost:" + configs.Env.Port)

	return nil
}
