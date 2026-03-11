package main

import (
	"context"
	"fmt"
	"go-mux-mongodb-user-manager-api/internal/infra/database/mongodb"
	"go-mux-mongodb-user-manager-api/internal/infra/security"
	"go-mux-mongodb-user-manager-api/internal/usecases/users_manager"
	"go-mux-mongodb-user-manager-api/internal/web"
	"go-mux-mongodb-user-manager-api/pkg/configs"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	err := configs.Connect()
	returnFatalError(err)

	col := initDatabase()
	mongorepo := mongodb.NewMongoRepository(col)

	createUseCase := users_manager.NewUseCaseRepository(mongorepo, security.NewHashingService())
	getUseCase := users_manager.NewUserGetAllRepository(mongorepo)

	usecase := web.NewUserHandlerRepository(
		createUseCase,
		getUseCase,
	)

	err = web.Routers(usecase)
	returnFatalError(err)

	fmt.Println("server running on port " + configs.Env.Port)
	fmt.Println("server running on address http://localhost:" + configs.Env.Port)
}

func initDatabase() *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongodb.NewMongoConnection(ctx)
	returnFatalError(err)

	configs.LoadEnv(client)

	col := mongodb.NewCollection(client, configs.Env.DbCol)

	errIndex := mongodb.IndexUnique(col)
	returnFatalError(errIndex)

	return col
}

func returnFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
