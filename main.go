package main

import (
	"fmt"
	"github.com/rs/cors"
	"log"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/register", HandlerRegister)

	// تنظیمات CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // for greater security you can enter the site address : "http://localhost:5174" (react project default)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// CORS Middleware
	handler := corsHandler.Handler(mux)

	fmt.Println("Server is running on port 8980...")
	log.Fatal(http.ListenAndServe(":8980", handler))

}
