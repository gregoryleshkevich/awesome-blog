package controllers

import (
	"awesome-blog/internal/models"
	"awesome-blog/internal/services"
	"encoding/json"
	"net/http"
)

type TokenResponce struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type AuthController struct{}

func (ac AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var userDB = models.UserDB{}
	var userService = models.UserService{}
	var jwtService = services.JWTService{}

	var Data struct {
		Email    string
		Password string
	}

	json.NewDecoder(r.Body).Decode(&Data)

	user := userDB.GetByEmail(Data.Email)

	if userService.CheckPasswordHash(Data.Password, user.Password) == false {
		//TODO
	}

	token, exTime := jwtService.GetToken(user.Email)

	response, _ := json.Marshal(TokenResponce{AccessToken: token, TokenType: "bearer", ExpiresIn: exTime.Second()})

	jsonResponse(w, response)
}
