package mongodb

import (
	"context"
	"go-mux-mongodb-user-manager-api/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository(col *mongo.Collection) *MongoRepository {
	return &MongoRepository{
		Collection: col,
	}
}

func IndexUnique(col *mongo.Collection) error {
	model := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := col.Indexes().CreateOne(context.Background(), model)
	return err
}

func (repo *MongoRepository) Create(u *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc := bson.M{
		"name":     u.Name,
		"email":    u.Email,
		"password": u.PasswordHash,
	}

	_, err := repo.Collection.InsertOne(ctx, doc)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return err
		}
		return err
	}

	return nil
}

func (repo *MongoRepository) GetAll() ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var result []domain.User
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
