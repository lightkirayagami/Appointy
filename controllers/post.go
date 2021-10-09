package controllers

import (
	"appointy/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PostController struct {
	session *mgo.Session
}

func NewPostController(session *mgo.Session) *PostController {
	return &PostController{session}

}

func (PostController PostController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	PostId := p.ByName("id")

	if !bson.IsObjectIdHex(PostId) {
		w.WriteHeader(http.StatusNotFound)
	}

	ConvertedPostId := bson.ObjectIdHex(PostId)

	PostModelRef := models.Post{}

	if err := PostController.session.DB("appointy").C("posts").FindId(ConvertedPostId).One(&PostModelRef); err != nil {
		w.WriteHeader(404)
		return
	}

	response, err := json.Marshal(PostModelRef)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", response)
}

func (PostController PostController) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	PostModelRef := models.Post{}
	json.NewDecoder(r.Body).Decode(&PostModelRef) //decoding in form of data model
	PostModelRef.Id = bson.NewObjectId()          //creating id
	PostController.session.DB("appointy").C("posts").Insert(PostModelRef)
	response, err := json.Marshal(PostModelRef)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", response)
}
