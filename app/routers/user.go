package routes

import (
	"net/http"
	"time"
	. "chitchat_mvc/app/utils"
	. "chitchat_mvc/app/controllers"
)

func init() {
	
	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(  Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))


	// defined in routes user
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/logout", Logout)
	mux.HandleFunc("/signup", Signup)
	mux.HandleFunc("/signup_account", SignupAccount)
	mux.HandleFunc("/authenticate", Authenticate)

	
	
}
