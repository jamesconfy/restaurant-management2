package service

import "golang.org/x/crypto/bcrypt"

var Crypto cryptoSrv

type cryptoSrv struct{}

func (c cryptoSrv) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (c cryptoSrv) comparePassword(hashed, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}
