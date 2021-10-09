package main

import (
	"appointy/controllers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

//Configuration for handling of routes.
func main() {

	RouteHandler := httprouter.New()
	UserController := controllers.NewUserController(getSession())
	PostController := controllers.NewPostController(getSession())
	RouteHandler.GET("/user/:id", UserController.GetUser)
	RouteHandler.POST("/user", UserController.CreateUser)
	RouteHandler.GET("/post/:id/user", UserController.UserPosts)
	RouteHandler.GET("/post/:id", PostController.GetPost)
	RouteHandler.POST("/post", PostController.CreatePost)
	http.ListenAndServe("localhost:9000", RouteHandler)

}

func getSession() *mgo.Session {
	//Connecting to MongoDB to use it as a datastore
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	log.Println("Firing the application---")
	return session

}
