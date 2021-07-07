package safesql

import (
	"fmt"
	"strconv"
)

var user = struct {
	ID       int32
	Username string
	Password string
	Email    string
	FullName string
}{}

func InsertUser() error {
	q := fmt.Sprintf("INSERT INTO users (username, password, email, fullname) VALUES (%d, %q, %q, %q, %q) returning id", user.ID, user.Username, user.Password, user.Email, user.FullName)
	if err := instance.QueryRow(q).
		Scan(&user.ID); err != nil {
		return fmt.Errorf("error creating user, %v", err)
	}

	return nil
}

func assingStringId(id string) {
	convertedId, _ := strconv.Atoi(id)
	user.ID = int32(convertedId)
}

func main() {
	InsertUser()
}
