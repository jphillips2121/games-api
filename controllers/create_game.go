package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jphillips2121/games-api/models"
	"gopkg.in/go-playground/validator.v9"
)

// HandleCreateGame adds a new game to the database.
func (service *GamesService) HandleCreateGame(w http.ResponseWriter, req *http.Request) {

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

	isValidDeveloper, err := service.Cloud.IsValidDeveloper(&game)
	if err != nil {
		fmt.Println(fmt.Errorf("error checking if developer is valid: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isValidDeveloper {
		fmt.Println(fmt.Errorf("developer is not authorized to add games: [%v]", game.Developer))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Add new game to DAO
	gameResponse, err := service.Dao.CreateNewGame(&game)
	if err != nil {
		fmt.Println(fmt.Errorf("error writing new game to dao: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Return new game
	err = json.NewEncoder(w).Encode(gameResponse)
	if err != nil {
		fmt.Println(fmt.Errorf("error encoding game json to response: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("\nSuccessfully added game: [%v]\n", game.Title)
}

func validateGame(game models.Game) error {
	validate := validator.New()
	err := validate.Struct(game)
	if err != nil {
		return err
	}

	return nil
}
