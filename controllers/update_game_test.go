package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
	"github.com/jphillips2121/games-api/cloud"
	"github.com/jphillips2121/games-api/dao"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitHandleUpdateGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("Request Body Empty", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)

		req, _ := http.NewRequest(http.MethodPut, "/games", nil)
		w := httptest.NewRecorder()

		mockGamesService.HandleUpdateGame(w, req)
		So(w.Code, ShouldEqual, 400)
	})

	Convey("Request Body Invalid", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)

		reqBody, _ := json.Marshal(invalidBodyType())
		req, _ := http.NewRequest(http.MethodPut, "/games", bytes.NewReader(reqBody))
		w := httptest.NewRecorder()

		mockGamesService.HandleUpdateGame(w, req)
		So(w.Code, ShouldEqual, 400)
	})

	Convey("Request Body Missing Required Fields", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)

		reqBody, _ := json.Marshal(emptyRequestBody())
		req, _ := http.NewRequest(http.MethodPut, "/games", bytes.NewReader(reqBody))
		w := httptest.NewRecorder()

		mockGamesService.HandleUpdateGame(w, req)
		So(w.Code, ShouldEqual, 400)
	})

	Convey("Error Checking If Developer Is Valid To Make Change", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().UpdateGame("12345678", gomock.Any()).Return(false, false, fmt.Errorf("error"))

		w := httptest.NewRecorder()
		reqBody, _ := json.Marshal(validRequestBody())
		req := httptest.NewRequest(http.MethodPut, "/games", bytes.NewReader(reqBody))
		req = mux.SetURLVars(req, map[string]string{"game_id": "12345678"})

		mockGamesService.HandleUpdateGame(w, req)
		So(w.Code, ShouldEqual, 500)
	})

	Convey("Developer Is Not Valid To Make Change", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().UpdateGame("12345678", gomock.Any()).Return(false, true, nil)

		w := httptest.NewRecorder()
		reqBody, _ := json.Marshal(validRequestBody())
		req := httptest.NewRequest(http.MethodPut, "/games", bytes.NewReader(reqBody))
		req = mux.SetURLVars(req, map[string]string{"game_id": "12345678"})

		mockGamesService.HandleUpdateGame(w, req)
		So(w.Code, ShouldEqual, 401)
	})

	Convey("ID Is Not Present To Make Change", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().UpdateGame("12345678", gomock.Any()).Return(true, false, nil)

		w := httptest.NewRecorder()
		reqBody, _ := json.Marshal(validRequestBody())
		req := httptest.NewRequest(http.MethodPut, "/games", bytes.NewReader(reqBody))
		req = mux.SetURLVars(req, map[string]string{"game_id": "12345678"})

		mockGamesService.HandleUpdateGame(w, req)
		So(w.Code, ShouldEqual, 404)
	})

	Convey("Valid Game Update", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().UpdateGame("12345678", gomock.Any()).Return(true, true, nil)

		w := httptest.NewRecorder()
		reqBody, _ := json.Marshal(validRequestBody())
		req := httptest.NewRequest(http.MethodPut, "/games", bytes.NewReader(reqBody))
		req = mux.SetURLVars(req, map[string]string{"game_id": "12345678"})

		mockGamesService.HandleUpdateGame(w, req)
		So(w.Code, ShouldEqual, 204)
	})
}
