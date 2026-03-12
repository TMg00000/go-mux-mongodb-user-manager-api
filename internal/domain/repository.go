package domain

type MongoRepository interface {
	Create(u *User) error
	GetAll() ([]User, error)
	GetByEmail(email string) (*User, error)
}

type HashingRepository interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hash string) error
}
