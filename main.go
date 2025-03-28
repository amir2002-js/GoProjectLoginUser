package main

import (
	"net/http"
)

type NewUser struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

var users = map[string]NewUser{}

func main() {
	http.HandleFunc("/register", HandlerRegister)
	err := http.ListenAndServe(":8980", nil)
	if err != nil {
		panic(err)
	}
}
