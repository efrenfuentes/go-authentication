package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/efrenfuentes/go-authentication/core/jwt"
	"github.com/efrenfuentes/go-authentication/database"
	"github.com/efrenfuentes/go-authentication/models"
	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	users := []models.User{}

	database.DB.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	user := models.User{}

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		database.DB.Preload("Groups").Preload("Groups.Clients").First(&user, id)

		if user.ID == 0 {
			message := models.APIMessage{"User not found"}
			json.NewEncoder(w).Encode(message)
		} else {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var newUser models.NewUser

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048676))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &newUser); err != nil {
		message := models.APIMessage{"Input format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		var user models.User
		user.Name = newUser.Name
		user.SetEmail(newUser.Email)
		user.SetPassword(newUser.Password)

		database.DB.Create(&user)

		if database.DB.NewRecord(user) {
			message := models.APIMessage{"Error creating user"}

			json.NewEncoder(w).Encode(message)
		} else {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	user := models.User{}
	var newUser models.NewUser

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {

		database.DB.First(&user, id)

		if user.ID == 0 {
			message := models.APIMessage{"User not found"}
			json.NewEncoder(w).Encode(message)
		} else {

			body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048676))

			if err != nil {
				panic(err)
			}

			if err := r.Body.Close(); err != nil {
				panic(err)
			}

			if err := json.Unmarshal(body, &newUser); err != nil {
				message := models.APIMessage{"Input format invalid"}
				json.NewEncoder(w).Encode(message)
			} else {
				user.Name = newUser.Name
				user.SetEmail(newUser.Email)
				user.SetPassword(newUser.Password)

				database.DB.Save(&user)

				json.NewEncoder(w).Encode(user)
			}
		}
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	user := models.User{}

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		database.DB.First(&user, id)

		if user.ID == 0 {
			message := models.APIMessage{"User not found"}
			json.NewEncoder(w).Encode(message)
		} else {
			database.DB.Delete(&user)
			message := models.APIMessage{"User successful deleted"}
			json.NewEncoder(w).Encode(message)
		}
	}
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var authUser models.AuthUser

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048676))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &authUser); err != nil {
		message := models.APIMessage{"Input format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		var user models.User

		database.DB.Where(&models.User{Email: authUser.Email}).Preload("Groups.Clients").First(&user)

		if user.ID == 0 {
			message := models.APIMessage{"User not found"}

			json.NewEncoder(w).Encode(message)
		} else {
			if user.Authenticate(authUser.Email, authUser.Password, authUser.ClientSecret) {
				token := jwt.CreateToken(user)
				apiToken := jwt.APIToken{token}
				json.NewEncoder(w).Encode(apiToken)
			} else {
				message := models.APIMessage{"Client key, email or password invalid"}

				json.NewEncoder(w).Encode(message)
			}
		}

	}
}
