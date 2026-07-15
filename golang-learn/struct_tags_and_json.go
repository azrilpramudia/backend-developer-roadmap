package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID	int	`json:"id"`
 	Username	string	`json:"username"`
	Password	string	`json:"-"`
}

func main() {
	user := User{ID: 1, Username: "burhan", Password: "secret_password"}
	
	// convert struct to JSON
	data, _ := json.MarshalIndent(user, "", " ")

	fmt.Println(string(data))
}