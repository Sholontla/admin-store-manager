package super_admin

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SuperAdmin struct {
	ID          uuid.UUID
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	Role        bool   `json:"role"`
	Permissions bool   `json:"permission"`
}

func SetPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Panic("Error ")
	}
	return string(hashedPassword), nil
}

func ComparePassword(s string, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(s), []byte(password))
}
