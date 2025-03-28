package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HandlerRegister(w http.ResponseWriter, r *http.Request) {

	//If our method was "POST" the program execution would continue
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Creating a instance of new user
	var user NewUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//trim all value
	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	user.ConfirmPassword = strings.TrimSpace(user.ConfirmPassword)
	user.Email = strings.TrimSpace(user.Email)

	//check values
	err = CheckVal(user.Username, user.Password, user.ConfirmPassword, user.Email, &w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if (users[user.Email] == NewUser{}) {
		users[user.Email] = user

		w.WriteHeader(http.StatusCreated)
		fmt.Println()
		err = json.NewEncoder(w).Encode(map[string]string{"message": "its ok"})
		if err != nil {
			return
		}

	} else {
		w.WriteHeader(http.StatusConflict)
		err = json.NewEncoder(w).Encode(map[string]string{"message": "user already exists"})
		if err != nil {
			return
		}

	}

}
