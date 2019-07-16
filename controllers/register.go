// Package controllers handles receiving and returning data from API calls.
package controllers

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/jphillips2121/games-api/cloud"
	"github.com/jphillips2121/games-api/dao"
)

// GamesService declares the DAO and CLOUD interfaces to use.
type GamesService struct {
	Dao   dao.DAO
	Cloud cloud.CLOUD
}

var gamesService *GamesService

// Register defines the route mappings for the API
func Register(mainRouter *mux.Router) {

	gamesService = &GamesService{
		Dao: &dao.JSON{},
		Cloud: &cloud.AWS{
			FileName:  os.Getenv("FILE_NAME"),
			AwsRegion: os.Getenv("AWS_REGION"),
			AwsID:     os.Getenv("AWS_ID"),
			AwsSecret: os.Getenv("AWS_SECRET"),
			AwsToken:  os.Getenv("AWS_TOKEN"),
			AwsBucket: os.Getenv("AWS_BUCKET"),
		},
	}

	mainRouter.HandleFunc("/games", gamesService.HandleCreateGame).Methods("POST").Name("createGame")
	mainRouter.HandleFunc("/games", gamesService.HandleGetListGames).Methods("GET").Name("getListGames")
	mainRouter.HandleFunc("/games/{game_id}", gamesService.HandleGetGame).Methods("GET").Name("getGame")
	mainRouter.HandleFunc("/games/{game_id}", gamesService.HandleUpdateGame).Methods("PUT").Name("updateGame")
	mainRouter.HandleFunc("/games/{game_id}", gamesService.HandleDeleteGame).Queries("developer", "{developer}").Methods("DELETE").Name("deleteGame")
}
