package controllers

import (
	"awesome-blog/internal/models"
	"encoding/json"
	"net/http"
)

var userService models.UserService = models.UserService{}
var userDB models.UserDB = models.UserDB{}

type UserController struct{}

func (uc UserController) Create(w http.ResponseWriter, r *http.Request) {
	var Data struct {
		Name string
		Email  string
		Password string
	}

	json.NewDecoder(r.Body).Decode(&Data)

	password, _ := userService.HashPassword(Data.Password)

	response, _ := json.Marshal(userDB.Create(&models.User{Name: Data.Name, Password: password, Email: Data.Email}))

	jsonResponse(w, response)
}
