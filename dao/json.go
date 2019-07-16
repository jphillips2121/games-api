package dao

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jphillips2121/games-api/models"
	"github.com/twinj/uuid"
)

// JSON struct provides a service to be accessed from the dao.
type JSON struct{}

var fileName = "games.json"

// CreateNewGame creates a new game in the json.
func (j *JSON) CreateNewGame(game *models.Game) (*models.GameResponse, error) {

	// Load Games from JSON file
	games, err := load()
	if err != nil {
		return nil, err
	}

	// Append new GameResponse to the existing games.
	gameResponse := CreateGameResponse(uuid.NewV4().String(), *game)
	games.Games = append(games.Games, gameResponse)

	// Save new Games to JSON file
	err = save(games)
	if err != nil {
		return nil, err
	}

	return &gameResponse, nil

}

// GetListGames returns a list of all games.
func (j *JSON) GetListGames() (*models.GameResponseList, error) {

	// Load Games from JSON file
	games, err := load()
	if err != nil {
		return nil, err
	}

	// Map Games to GameResponseList
	GameResponseList := models.GameResponseList{
		ItemsPerPage: len(games.Games),
		StartIndex:   1,
		TotalResults: len(games.Games),
		Items:        games.Games,
	}

	return &GameResponseList, nil
}

// GetGame returns a game specified by an ID.
func (j *JSON) GetGame(id string) (*models.GameResponse, error) {
	// Load Games from JSON file
	games, err := load()
	if err != nil {
		return nil, err
	}

	// Loop through Games to find one matching the id.
	for _, game := range games.Games {
		if game.ID == id {
			return &game, nil
		}
	}

	// If it reaches here no game with that id exists.
	return nil, nil
}

// UpdateGame updates a game specified by an ID.
// First bool represents if developer is valid, second bool represents if id is present
func (j *JSON) UpdateGame(id string, newGame *models.Game) (bool, bool, error) {

	// Load Games from JSON file
	games, err := load()
	if err != nil {
		return false, false, err
	}

	// Convert type Game to type GameResponse
	newGameResponse := CreateGameResponse(id, *newGame)

	// Loop through Games to find one matching the ID.
	for index, game := range games.Games {
		if game.ID == id {
			if newGame.Developer == game.Developer {
				games.Games[index] = newGameResponse

				// Save new Games to JSON file
				err = save(games)
				if err != nil {
					return false, false, err
				}

				return true, true, nil
			}

			// If developer is not valid return first bool as false
			return false, true, nil
		}
	}

	// If no ID is found return second bool as false
	return true, false, nil

}

// DeleteGame deletes a game specified by an ID.
// First bool represents if developer is valid, second bool represents if id is present
func (j *JSON) DeleteGame(id, developer string) (bool, bool, error) {

	// Load Games from JSON file
	games, err := load()
	if err != nil {
		return false, false, err
	}

	// Loop through Games to find one matching the ID.
	for index, game := range games.Games {
		if game.ID == id {
			if developer == game.Developer {
				// Delete game from array
				games.Games = append(games.Games[:index], games.Games[index+1:]...)

				// Save new Games to JSON file
				err = save(games)
				if err != nil {
					return false, false, err
				}

				return true, true, nil
			}

			// If developer is not valid return first bool as false
			return false, true, nil

		}
	}

	// If no ID is found return second bool as false
	return true, false, nil

}

// CreateGameResponse converts type Game to type GameResponse
func CreateGameResponse(id string, game models.Game) models.GameResponse {

	GameResponse := models.GameResponse{
		ID:          id,
		Title:       game.Title,
		ReleaseDate: game.ReleaseDate,
		Genres:      game.Genres,
		Developer:   game.Developer,
	}

	return GameResponse
}

func load() (*models.GameList, error) {
	//Read the JSON file
	gamesFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Unmarshal bytes into an array of GameResponse (GameList)
	games := &models.GameList{}
	err = json.Unmarshal(gamesFile, games)
	if err != nil {
		return nil, err
	}

	return games, nil
}

func save(games *models.GameList) error {
	newGameJSON, err := json.MarshalIndent(games, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, newGameJSON, 0644)
}
