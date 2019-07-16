package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleGetListGames returns a list of all games.
func (service *GamesService) HandleGetListGames(w http.ResponseWriter, req *http.Request) {

	// Get gameResponseList from JSON.
	gameResponseList, err := service.Dao.GetListGames()
	if err != nil {
		fmt.Println(fmt.Errorf("error retrieving list of games from dao: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if gameResponseList.TotalResults < 1 {
		fmt.Println(fmt.Errorf("no items are present in the dao to return"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Return gameResponseList
	err = json.NewEncoder(w).Encode(gameResponseList)
	if err != nil {
		fmt.Println(fmt.Errorf("error encoding gameResponseList to response: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("\nSuccessfully Returned list of all games\n")
}
