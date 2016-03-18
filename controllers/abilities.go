package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/efrenfuentes/go-authentication/database"
	"github.com/efrenfuentes/go-authentication/models"
	"github.com/gorilla/mux"
)

func GetAllAbilities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	abilities := []models.Ability{}

	database.Database.Preload("Groups.Users").Preload("Groups.Clients").Preload("Groups.Abilities").Find(&abilities)

	json.NewEncoder(w).Encode(abilities)
}

func GetAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	ability := models.Ability{}

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		database.Database.Preload("Groups.Users").Preload("Groups.Clients").Preload("Groups.Abilities").First(&ability, id)

		if ability.ID == 0 {
			message := models.APIMessage{"Ability not found"}
			json.NewEncoder(w).Encode(message)
		} else {
			json.NewEncoder(w).Encode(ability)
		}
	}
}

func CreateAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var newAbility models.NewAbility

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048676))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &newAbility); err != nil {
		message := models.APIMessage{"Input format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		var ability models.Ability
		ability.Name = newAbility.Name

		database.Database.Create(&ability)

		if database.Database.NewRecord(ability) {
			message := models.APIMessage{"Error creating ability"}

			json.NewEncoder(w).Encode(message)
		} else {
			json.NewEncoder(w).Encode(ability)
		}
	}
}

func UpdateAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	ability := models.Ability{}
	var newAbility models.NewAbility

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {

		database.Database.First(&ability, id)

		if ability.ID == 0 {
			message := models.APIMessage{"Ability not found"}
			json.NewEncoder(w).Encode(message)
		} else {

			body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048676))

			if err != nil {
				panic(err)
			}

			if err := r.Body.Close(); err != nil {
				panic(err)
			}

			if err := json.Unmarshal(body, &newAbility); err != nil {
				message := models.APIMessage{"Input format invalid"}
				json.NewEncoder(w).Encode(message)
			} else {
				ability.Name = newAbility.Name

				database.Database.Save(&ability)

				json.NewEncoder(w).Encode(ability)
			}
		}
	}
}

func DeleteAbility(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	ability := models.Ability{}

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		database.Database.First(&ability, id)

		if ability.ID == 0 {
			message := models.APIMessage{"Ability not found"}
			json.NewEncoder(w).Encode(message)
		} else {
			database.Database.Delete(&ability)
			message := models.APIMessage{"Ability successful deleted"}
			json.NewEncoder(w).Encode(message)
		}
	}
}
