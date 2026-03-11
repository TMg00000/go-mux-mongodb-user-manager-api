package security

import "golang.org/x/crypto/bcrypt"

type HashingService struct{}

func NewHashingService() *HashingService {
	return &HashingService{}
}

func (h *HashingService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *HashingService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
