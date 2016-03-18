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

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	groups := []models.Group{}

	database.DB.Preload("Users").Preload("Clients").Preload("Abilities").Find(&groups)

	json.NewEncoder(w).Encode(groups)
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	group := models.Group{}

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		database.DB.Preload("Users").Preload("Clients").Preload("Abilities").First(&group, id)

		if group.ID == 0 {
			message := models.APIMessage{"Group not found"}
			json.NewEncoder(w).Encode(message)
		} else {
			json.NewEncoder(w).Encode(group)
		}
	}
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var newGroup models.NewGroup

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048676))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &newGroup); err != nil {
		message := models.APIMessage{"Input format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		var group models.Group
		group.Name = newGroup.Name
		group.Description = newGroup.Description

		database.DB.Create(&group)

		if database.DB.NewRecord(group) {
			message := models.APIMessage{"Error creating group"}

			json.NewEncoder(w).Encode(message)
		} else {
			json.NewEncoder(w).Encode(group)
		}
	}
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	group := models.Group{}
	var newGroup models.NewGroup

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {

		database.DB.First(&group, id)

		if group.ID == 0 {
			message := models.APIMessage{"Group not found"}
			json.NewEncoder(w).Encode(message)
		} else {

			body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048676))

			if err != nil {
				panic(err)
			}

			if err := r.Body.Close(); err != nil {
				panic(err)
			}

			if err := json.Unmarshal(body, &newGroup); err != nil {
				message := models.APIMessage{"Input format invalid"}
				json.NewEncoder(w).Encode(message)
			} else {
				group.Name = newGroup.Name
				group.Description = newGroup.Description

				database.DB.Save(&group)

				json.NewEncoder(w).Encode(group)
			}
		}
	}
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	group := models.Group{}

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		database.DB.First(&group, id)

		if group.ID == 0 {
			message := models.APIMessage{"Group not found"}
			json.NewEncoder(w).Encode(message)
		} else {
			database.DB.Delete(&group)
			message := models.APIMessage{"Group successful deleted"}
			json.NewEncoder(w).Encode(message)
		}
	}
}
