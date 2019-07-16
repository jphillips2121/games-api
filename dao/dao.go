// Package dao handles integration with database.
package dao

import "github.com/jphillips2121/games-api/models"

// DAO interface declares how to interact with database regardless of which database type is used.
type DAO interface {
	CreateNewGame(game *models.Game) (*models.GameResponse, error)
	GetListGames() (*models.GameResponseList, error)
	GetGame(id string) (*models.GameResponse, error)
	UpdateGame(id string, newGame *models.Game) (bool, bool, error)
	DeleteGame(id, developer string) (bool, bool, error)
}
