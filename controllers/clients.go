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

func GetAllClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	clients := []models.Client{}

	database.DB.Preload("Groups.Users").Preload("Groups.Clients").Preload("Groups.Abilities").Find(&clients)

	json.NewEncoder(w).Encode(clients)
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	client := models.Client{}

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		database.DB.Preload("Groups.Users").Preload("Groups.Clients").Preload("Groups.Abilities").First(&client, id)

		if client.ID == 0 {
			message := models.APIMessage{"Client not found"}
			json.NewEncoder(w).Encode(message)
		} else {
			json.NewEncoder(w).Encode(client)
		}
	}
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var newClient models.NewClient

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048676))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &newClient); err != nil {
		message := models.APIMessage{"Input format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		var client models.Client
		client.Name = newClient.Name
		client.GenerateKeys()

		database.DB.Create(&client)

		if database.DB.NewRecord(client) {
			message := models.APIMessage{"Error creating client"}

			json.NewEncoder(w).Encode(message)
		} else {
			json.NewEncoder(w).Encode(client)
		}
	}
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	client := models.Client{}
	var newClient models.NewClient

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {

		database.DB.First(&client, id)

		if client.ID == 0 {
			message := models.APIMessage{"Client not found"}
			json.NewEncoder(w).Encode(message)
		} else {

			body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048676))

			if err != nil {
				panic(err)
			}

			if err := r.Body.Close(); err != nil {
				panic(err)
			}

			if err := json.Unmarshal(body, &newClient); err != nil {
				message := models.APIMessage{"Input format invalid"}
				json.NewEncoder(w).Encode(message)
			} else {
				client.Name = newClient.Name

				database.DB.Save(&client)

				json.NewEncoder(w).Encode(client)
			}
		}
	}
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	urlParams := mux.Vars(r)
	strID := urlParams["id"]
	client := models.Client{}

	id, err := strconv.Atoi(strID)
	if err != nil {
		// handle error
		message := models.APIMessage{"Id format invalid"}
		json.NewEncoder(w).Encode(message)
	} else {
		database.DB.First(&client, id)

		if client.ID == 0 {
			message := models.APIMessage{"Client not found"}
			json.NewEncoder(w).Encode(message)
		} else {
			database.DB.Delete(&client)
			message := models.APIMessage{"Client successful deleted"}
			json.NewEncoder(w).Encode(message)
		}
	}
}
