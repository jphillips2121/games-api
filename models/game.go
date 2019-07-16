// Package models contains the required data structures.
package models

// Game contains the expected incoming data of a new/updated game.
type Game struct {
	Title       string   `json:"title" validate:"required"`
	ReleaseDate string   `json:"release_date"`
	Genres      []string `json:"genres"`
	Developer   string   `json:"developer" validate:"required"`
}

// GameResponse contains the returned data from a new game, as well as how the game is saved to the database.
type GameResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title" validate:"required"`
	ReleaseDate string   `json:"release_date"`
	Genres      []string `json:"genres"`
	Developer   string   `json:"developer" validate:"required"`
}

// GameList contains how the GameResponse's are stored in the games.json file.
type GameList struct {
	Games []GameResponse `json:"games"`
}

// GameResponseList contains the returned data from a list of all games.
type GameResponseList struct {
	ItemsPerPage int            `json:"items_per_page"`
	StartIndex   int            `json:"start_index"`
	TotalResults int            `json:"total_results"`
	Items        []GameResponse `json:"items"`
}
