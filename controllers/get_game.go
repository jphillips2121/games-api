package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleGetGame returns a game specified by an ID.
func (service *GamesService) HandleGetGame(w http.ResponseWriter, req *http.Request) {

	// Check for game ID in request
	vars := mux.Vars(req)
	id := vars["game_id"]

	// Get game with that ID from the JSON
	gameResponse, err := service.Dao.GetGame(id)
	if err != nil {
		fmt.Println(fmt.Errorf("error retrieving list of games from dao: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if gameResponse == nil {
		fmt.Println(fmt.Errorf("game with ID of [%v] does not exist", id))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Return gameResponseList
	err = json.NewEncoder(w).Encode(gameResponse)
	if err != nil {
		fmt.Println(fmt.Errorf("error encoding gameResponse to response: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("\nSuccessfully returned game: [%v]\n", gameResponse.ID)
}
