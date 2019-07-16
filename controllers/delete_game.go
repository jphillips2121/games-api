package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleDeleteGame removes a Game from the database.
func (service *GamesService) HandleDeleteGame(w http.ResponseWriter, req *http.Request) {

	// Check for game ID in request
	vars := mux.Vars(req)
	id := vars["game_id"]
	developer := vars["developer"]

	// Attempt to delete game
	isValidDeveloper, isIDPresent, err := service.Dao.DeleteGame(id, developer)
	if err != nil {
		fmt.Println(fmt.Errorf("error updating data: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isValidDeveloper {
		fmt.Println(fmt.Errorf("developer is not authorized to delete this game: [%v], [%v]", developer, id))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !isIDPresent {
		fmt.Println(fmt.Errorf("game with ID of [%v] does not exist", id))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	fmt.Printf("\nSuccessfully deleted game: [%v]\n", id)
}
