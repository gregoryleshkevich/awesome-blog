package controllers

import (
	"awesome-blog/internal/models"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

var postDB models.PostDB = models.PostDB{}

type PostController struct {
}

//Create post action
func (pc PostController) Create(w http.ResponseWriter, r *http.Request) {
	var Data struct {
		Title string
		Body  string
	}

	json.NewDecoder(r.Body).Decode(&Data)

	response, _ := json.Marshal(postDB.Create(&models.Post{Title: Data.Title, Body: Data.Body}))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// Show post action
func (pc PostController) Show(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	
	response, _ := json.Marshal(postDB.GetByID(id))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// Shows all posts 
func (pc PostController) Shows(w http.ResponseWriter, r *http.Request)  {
	response, _ := json.Marshal(postDB.GetAll())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
