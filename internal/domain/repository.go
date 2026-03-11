package domain

type MongoRepository interface {
	Create(u *User) error
	GetAll() ([]User, error)
}

type HashingService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
