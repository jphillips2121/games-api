package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jphillips2121/games-api/cloud"
	"github.com/jphillips2121/games-api/dao"

	. "github.com/smartystreets/goconvey/convey"
)

func createMockGamesService(dao *dao.MockDAO, cloud *cloud.MockCLOUD) GamesService {
	return GamesService{
		Dao:   dao,
		Cloud: cloud,
	}
}

func invalidBodyType() map[string]interface{} {
	return map[string]interface{}{
		"title": 123,
	}
}

func emptyRequestBody() map[string]interface{} {
	return map[string]interface{}{}
}

func validRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"title":        "Game",
		"developer":    "Valid Developer",
		"release_date": "2019-07-16",
		"genres":       [2]string{"Test", "Test Test"},
	}
}

func TestUnitHandleCreateGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("Request Body Empty", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)

		req, _ := http.NewRequest(http.MethodPost, "/games", nil)
		w := httptest.NewRecorder()

		mockGamesService.HandleCreateGame(w, req)
		So(w.Code, ShouldEqual, 400)
	})

	Convey("Request Body Invalid", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)

		reqBody, _ := json.Marshal(invalidBodyType())
		req, _ := http.NewRequest(http.MethodPost, "/games", bytes.NewReader(reqBody))
		w := httptest.NewRecorder()

		mockGamesService.HandleCreateGame(w, req)
		So(w.Code, ShouldEqual, 400)
	})

	Convey("Request Body Missing Required Fields", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)

		reqBody, _ := json.Marshal(emptyRequestBody())
		req, _ := http.NewRequest(http.MethodPost, "/games", bytes.NewReader(reqBody))
		w := httptest.NewRecorder()

		mockGamesService.HandleCreateGame(w, req)
		So(w.Code, ShouldEqual, 400)
	})

	Convey("Error Checking If Developer Is Valid", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockCloud.EXPECT().IsValidDeveloper(gomock.Any()).Return(false, fmt.Errorf("error"))

		w := httptest.NewRecorder()
		reqBody, _ := json.Marshal(validRequestBody())
		req, _ := http.NewRequest(http.MethodPost, "/games", bytes.NewReader(reqBody))

		mockGamesService.HandleCreateGame(w, req)
		So(w.Code, ShouldEqual, 500)
	})

	Convey("Developer is not valid", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockCloud.EXPECT().IsValidDeveloper(gomock.Any()).Return(false, nil)

		reqBody, _ := json.Marshal(validRequestBody())
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/games", bytes.NewReader(reqBody))

		mockGamesService.HandleCreateGame(w, req)
		So(w.Code, ShouldEqual, 401)
	})

	Convey("Error adding new game to DAO", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockCloud.EXPECT().IsValidDeveloper(gomock.Any()).Return(true, nil)
		mockDao.EXPECT().CreateNewGame(gomock.Any()).Return(nil, fmt.Errorf("error"))

		reqBody, _ := json.Marshal(validRequestBody())
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/games", bytes.NewReader(reqBody))

		mockGamesService.HandleCreateGame(w, req)
		So(w.Code, ShouldEqual, 500)
	})

	Convey("Valid create game request", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockCloud.EXPECT().IsValidDeveloper(gomock.Any()).Return(true, nil)
		mockDao.EXPECT().CreateNewGame(gomock.Any()).Return(validGameResponse(), nil)

		reqBody, _ := json.Marshal(validRequestBody())
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/games", bytes.NewReader(reqBody))

		mockGamesService.HandleCreateGame(w, req)
		So(w.Code, ShouldEqual, 201)
	})
}
