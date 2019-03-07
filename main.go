package main

import (
	"net/http"
	"time"
	. "chitchat_mvc/app/utils"
	. "chitchat_mvc/app/controllers"
)

func main() {
	
	Info("===", GetCurrentDirectory(),Config.DbDriverName)
	 P("ChitChat",  Version(), "started at",  Config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(  Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", Index)
	// error
	mux.HandleFunc("/err", Err)

	// defined in route_auth.go
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/logout", Logout)
	mux.HandleFunc("/signup", Signup)
	mux.HandleFunc("/signup_account", SignupAccount)
	mux.HandleFunc("/authenticate", Authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", NewThread)
	mux.HandleFunc("/thread/create", CreateThread)
	mux.HandleFunc("/thread/post", PostThread)
	mux.HandleFunc("/thread/read", ReadThread)

	// starting up the server
	server := &http.Server{
		Addr:            Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(  Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(  Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
	
}
