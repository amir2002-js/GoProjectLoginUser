package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HandlerRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user NewUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user.Username = strings.TrimSpace(user.Username)
	user.Password = strings.TrimSpace(user.Password)
	user.ConfirmPassword = strings.TrimSpace(user.ConfirmPassword)
	user.Email = strings.TrimSpace(user.Email)

	if len(user.Email) < 6 {
		http.Error(w, "Email too short", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "wrong password  should be +6 length ", http.StatusBadRequest)
		return
	}

	if user.ConfirmPassword != user.Password {
		http.Error(w, "wrong password noy match confirm password ", http.StatusBadRequest)
		return
	}

	if len(user.Username) < 2 {
		http.Error(w, " username is wrong ", http.StatusBadRequest)
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
