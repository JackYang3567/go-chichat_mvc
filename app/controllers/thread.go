package controllers

import (
	"fmt"
	//"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	. "chitchat_mvc/app/utils"
	"chitchat_mvc/app/models"	
	"net/http"
)

// GET /threads/new
// Show the new thread form page
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err :=  Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		 GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// Create the user account
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err :=  Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			 Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			 Danger(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		request.ParseForm() 
		if len(request.Form["topic"][0])==0{
			//为空的处理
		
			Danger(err, "Topic Cannot empty!")
		}else{

		
			if _, err := user.CreateThread(topic); err != nil {
				Danger(err, "Cannot create thread")
			}
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		 Error_message(writer, request, "Cannot read thread")
	} else {
		_, err :=  Session(writer, request)
		if err != nil {
			 GenerateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			 GenerateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
// Create the post
func PostThread(writer http.ResponseWriter, request *http.Request) {
	sess, err :=  Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			 Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			 Danger(err, "Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := models.ThreadByUUID(uuid)
		if err != nil {
			 Error_message(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			 Danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
