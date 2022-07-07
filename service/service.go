package service

import (
	"encoding/json"
	"fmt"

	"jwt-practice/controller"
	"jwt-practice/database"
	"jwt-practice/entity"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	defer database.Closedatabase(db)

	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	var dbuser entity.User
	db.Where("email=?", user.Email).First(&dbuser)
	//Checking if the email already exists or not
	if dbuser.Email != "" {
		var err error
		err1 := fmt.Sprintf("%q", err)
		//err = SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err1)
		return

	}
	var err2 error
	user.Password, err2 = GeneratehashPassword(user.Password)
	if err2 != nil {
		fmt.Println("Error in generating password")
	}

	//Insert user details in database
	db.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

//Generating hash Password
func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

func Login(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	defer database.Closedatabase(db)

	var auth entity.Authentication
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	var authuser entity.User
	db.Where("email=?", authuser.Email).First(&authuser)
	if authuser.Email == "" {
		if err != nil {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(err)
		}
	}

	check := CheckPasswordHash(auth.Password, authuser.Password)

	if !check {
		var err error
		err1 := fmt.Sprintf("%q", err)
		//err = SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err1)
		return
	}
	validToken, err := controller.GenerateJWT(authuser.Email, authuser.Role)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	var token entity.Token
	token.Role = authuser.Role
	token.Email = authuser.Email
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SetError(err entity.Error, message string) entity.Error {
	err.IsError = true
	err.Message = message
	return err
}
