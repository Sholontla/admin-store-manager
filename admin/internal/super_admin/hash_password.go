package super_admin

import (
	"fmt"
	"log"
)

func HashPasswowrd(s string) {
	hash, err := SetPassword(s)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(hash)
}
