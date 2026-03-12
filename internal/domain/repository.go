package domain

type MongoRepository interface {
	Create(*User) error
	GetAll() ([]User, error)
	GetByEmail(email string) (*User, error)
	UpdateName(name string, email string) (model *User, err error)
}

type HashingRepository interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hash string) error
}
