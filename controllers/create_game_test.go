package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jphillips2121/games-api/cloud"
	"github.com/jphillips2121/games-api/dao"

	. "github.com/smartystreets/goconvey/convey"
)

func createMockGamesService(t *testing.T) GamesService {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	return GamesService{
		Dao:   dao.NewMockDAO(mockCtrl),
		Cloud: cloud.NewMockCLOUD(mockCtrl),
	}
}

func TestUnitHandleCreateGame(t *testing.T) {
	Convey("Request Body Empty", t, func() {
		mockGamesService := createMockGamesService(t)
		req, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()

		mockGamesService.HandleCreateGame(w, req)
		So(w.Code, ShouldEqual, 400)
	})

}
