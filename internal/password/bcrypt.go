package password

import "golang.org/x/crypto/bcrypt"

type BcryptService struct {
}

func (s *BcryptService) Validate(hash, pswd []byte) error {
	return bcrypt.CompareHashAndPassword(hash, pswd)
}
