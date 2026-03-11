package web

import (
	"go-mux-mongodb-user-manager-api/pkg/configs"
	"net/http"

	"github.com/gorilla/mux"
)

func Routers(usecase *UserHandlerRepository) error {
	r := mux.NewRouter()

	r.HandleFunc("/api/user", usecase.CreateNewUser).Methods("POST")
	r.HandleFunc("/api/user", usecase.GetAllUsers).Methods("GET")

	if err := http.ListenAndServe(":"+configs.Env.Port, r); err != nil {
		return err
	}

	return nil
}
