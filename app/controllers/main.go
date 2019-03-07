package controllers

import (
	//"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	. "chitchat_mvc/app/utils"
	"chitchat_mvc/app/models"	
	
	"net/http"
)

// GET /err?msg=
// shows the error message page
func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err :=  Session(writer, request)
	if err != nil {
		 GenerateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		 GenerateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := models.Threads()
	if err != nil {
		 Error_message(writer, request, "Cannot get threads")
	} else {
		_, err :=  Session(writer, request)
		if err != nil {
			 GenerateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			 GenerateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
