package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jphillips2121/games-api/controllers"
	"net/http"
)

func main() {
	// Create router
	mainRouter := mux.NewRouter()
	controllers.Register(mainRouter)

	fmt.Println("Starting games-api")

	err := http.ListenAndServe(":8080", mainRouter)
	if err != nil {
		fmt.Println(err)
	}
}
