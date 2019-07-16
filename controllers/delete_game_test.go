package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/jphillips2121/games-api/cloud"
	"github.com/jphillips2121/games-api/dao"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitHandleDeleteGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("Error Deleting Game", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().DeleteGame("12345678", "developer-name").Return(false, false, fmt.Errorf("error"))

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/games", nil)
		req = mux.SetURLVars(req, map[string]string{"developer": "developer-name", "game_id": "12345678"})

		mockGamesService.HandleDeleteGame(w, req)
		So(w.Code, ShouldEqual, 500)
	})

	Convey("Developer Is Not Valid To Delete Game", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().DeleteGame("12345678", "developer-name").Return(false, true, nil)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/games", nil)
		req = mux.SetURLVars(req, map[string]string{"developer": "developer-name", "game_id": "12345678"})

		mockGamesService.HandleDeleteGame(w, req)
		So(w.Code, ShouldEqual, 401)
	})

	Convey("ID Does Not Exist For Game To Be Deleted", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().DeleteGame("12345678", "developer-name").Return(true, false, nil)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/games", nil)
		req = mux.SetURLVars(req, map[string]string{"developer": "developer-name", "game_id": "12345678"})

		mockGamesService.HandleDeleteGame(w, req)
		So(w.Code, ShouldEqual, 404)
	})

	Convey("Valid Deletion Of Game", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().DeleteGame("12345678", "developer-name").Return(true, true, nil)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/games", nil)
		req = mux.SetURLVars(req, map[string]string{"developer": "developer-name", "game_id": "12345678"})

		mockGamesService.HandleDeleteGame(w, req)
		So(w.Code, ShouldEqual, 204)
	})
}
