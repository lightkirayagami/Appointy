package controllers

import (
	"appointy/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}

}

func (UserController UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	UserId := p.ByName("id")

	if !bson.IsObjectIdHex(UserId) {
		w.WriteHeader(http.StatusNotFound)
	}

	ConvertedUserId := bson.ObjectIdHex(UserId)

	UserModelRef := models.User{}

	if err := UserController.session.DB("appointy").C("users").FindId(ConvertedUserId).One(&UserModelRef); err != nil {
		w.WriteHeader(404)
		return
	}

	response, err := json.Marshal(UserModelRef)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", response)
}

func (UserController UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	UserModelRef := models.User{}
	json.NewDecoder(r.Body).Decode(&UserModelRef) //decoding in form of data model
	UserModelRef.Id = bson.NewObjectId()          //creating id
	UserModelRef.Password = getHash([]byte(UserModelRef.Password))
	UserController.session.DB("appointy").C("users").Insert(UserModelRef)
	response, err := json.Marshal(UserModelRef)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", response)
}
func (UserController UserController) UserPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	UserModelRef := models.Post{}

	if err := UserController.session.DB("appointy").C("posts").Find(bson.M{"owner": p.ByName("id")}).One(&UserModelRef); err != nil {
		w.WriteHeader(300)
		return
	}

	response, err := json.Marshal(UserModelRef)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", response)
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
