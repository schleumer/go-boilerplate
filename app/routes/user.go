package routes

import (
	"net/http"

	"../database"
	"../models"
	"../utils"
	"golang.org/x/crypto/bcrypt"
)

/*
ShowUser is the route to show user information
*/
func ShowUser(user models.User, rs http.ResponseWriter, rq *http.Request) {
	bag := utils.NewBag(user)

	utils.JsonResp(rs, bag)
}

type signInUserModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signedUserModel struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

/*
SignInUser is the route for sign in users
*/
func SignInUser(rs http.ResponseWriter, rq *http.Request) {
	input := signInUserModel{}

	utils.JsonReq(rq, &input)

	user := models.User{}

	database.LocalDb.First(&user, "username = ?", input.Username)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		bag := utils.NewBag(nil)

		bag.AddError("", "Login ou senha invalidos", "username")

		bag.WriteTo(rs)

	} else {
		utils.NewBag(&signedUserModel{
			user,
			user.SignToken(),
		}).WriteTo(rs)
	}
}
