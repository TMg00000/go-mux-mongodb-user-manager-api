package configs

import (
	"context"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Config struct {
	MongoURI      string `envconfig:"MONGO_URI"`
	Port          string `envconfig:"PORT"`
	DbConfigName  string `envconfig:"DB_CONFIG_NAME"`
	ColConfigName string `envconfig:"COL_CONFIG_NAME"`
	DbName        string
	DbCol         string
}

var Env Config

func Connect() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := envconfig.Process("", &Env); err != nil {
		return err
	}

	return nil
}

func LoadEnv(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cfg := client.Database(Env.DbConfigName).Collection(Env.ColConfigName)

	id := bson.M{
		"id": "user_db_config",
	}

	var result bson.M
	err := cfg.FindOne(ctx, id).Decode(&result)
	if err != nil {
		return err
	}

	Env.DbName = result["DB_NAME"].(string)
	Env.DbCol = result["COLLECTION"].(string)

	return nil
}
