package helpers

import "golang.org/x/crypto/bcrypt"

func GenerateHash(val string) string {
	result, err := bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(result)
}

func CompareHash(val string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(val))

	return err == nil
}
