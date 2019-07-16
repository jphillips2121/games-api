package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jphillips2121/games-api/models"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/jphillips2121/games-api/cloud"
	"github.com/jphillips2121/games-api/dao"

	. "github.com/smartystreets/goconvey/convey"
)

func validGameResponse() *models.GameResponse {
	return &models.GameResponse{
		ID:          "123-456-789",
		Title:       "Game",
		Developer:   "Valid Developer",
		ReleaseDate: "2019-07-16",
		Genres:      []string{"Test", "Test Test"},
	}
}

func TestUnitHandleGetGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("Error retrieving list of games from dao", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().GetGame("12345678").Return(nil, fmt.Errorf("error"))

		req := httptest.NewRequest(http.MethodGet, "/games", nil)
		req = mux.SetURLVars(req, map[string]string{"game_id": "12345678"})

		w := httptest.NewRecorder()

		mockGamesService.HandleGetGame(w, req)
		So(w.Code, ShouldEqual, 500)
	})

	Convey("Game does not exist", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().GetGame("12345678").Return(nil, nil)

		req := httptest.NewRequest(http.MethodGet, "/games", nil)
		req = mux.SetURLVars(req, map[string]string{"game_id": "12345678"})

		w := httptest.NewRecorder()

		mockGamesService.HandleGetGame(w, req)
		So(w.Code, ShouldEqual, 404)
	})

	Convey("valid get game request", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().GetGame("12345678").Return(validGameResponse(), nil)

		req := httptest.NewRequest(http.MethodGet, "/games", nil)
		req = mux.SetURLVars(req, map[string]string{"game_id": "12345678"})

		w := httptest.NewRecorder()

		mockGamesService.HandleGetGame(w, req)
		So(w.Code, ShouldEqual, 200)
	})
}
