package main

import (
	"errors"
	"net/http"
)

func CheckVal(userName, password, confirmPass, email string, w *http.ResponseWriter) (err error) {
	err = nil
	if len(email) < 6 {
		http.Error(*w, "Email too short", http.StatusBadRequest)
		err = errors.New("email too short")
		return err
	}

	if len(password) < 6 {
		http.Error(*w, "wrong password  should be +6 length ", http.StatusBadRequest)
		err = errors.New("password too short")
		return err
	}

	if confirmPass != password {
		http.Error(*w, "wrong password noy match confirm password ", http.StatusBadRequest)
		err = errors.New("wrong password")
		return err
	}

	if len(userName) < 2 {
		http.Error(*w, " username is wrong ", http.StatusBadRequest)
		err = errors.New("username is wrong ")
		return err
	}
	return err
}
