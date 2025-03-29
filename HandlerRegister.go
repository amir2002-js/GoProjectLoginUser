package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HandlerRegister(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	//AI way
	//origin := r.Header.Get("Origin")
	//w.Header().Set("Access-Control-Allow-Origin", origin) // echo back the requesting origin
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	//w.Header().Set("Access-Control-Allow-Credentials", "true")

	//my way
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	//w.Header().Set("Access-Control-Allow-Credentials", "true")

	//If our method was "POST" the program execution would continue
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Creating an instance of new user
	var user NewUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
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

	_, exist := users[user.Email]
	if !exist {
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
