package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jphillips2121/games-api/models"
)

//HandleUpdateGame updates a game on the database.
func (service *GamesService) HandleUpdateGame(w http.ResponseWriter, req *http.Request) {

	// Check for game ID in request
	vars := mux.Vars(req)
	id := vars["game_id"]

	// Checks if incoming request body is empty
	if req.Body == nil {
		fmt.Println(fmt.Errorf("request body is empty"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Decodes incoming request body to set model
	requestDecoder := json.NewDecoder(req.Body)
	var game models.Game
	err := requestDecoder.Decode(&game)
	if err != nil {
		fmt.Println(fmt.Errorf("request body is invalid"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate incoming request to game model
	err = validateGame(game)
	if err != nil {
		fmt.Println(fmt.Errorf("request body is invalid: [%v]", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Attempt to update game
	isValidDeveloper, isIDPresent, err := service.Dao.UpdateGame(id, &game)
	if err != nil {
		fmt.Println(fmt.Errorf("error updating data: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isValidDeveloper {
		fmt.Println(fmt.Errorf("developer is not authorized to update this game: [%v], [%v]", game.Developer, game.Title))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !isIDPresent {
		fmt.Println(fmt.Errorf("game with ID of [%v] does not exist in dao", id))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	fmt.Printf("\nSuccessfully updated game: [%v]\n", game.Title)
}
