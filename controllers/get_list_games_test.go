package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jphillips2121/games-api/cloud"
	"github.com/jphillips2121/games-api/dao"
	"github.com/jphillips2121/games-api/models"

	. "github.com/smartystreets/goconvey/convey"
)

func emptyGameResponseListResponse() *models.GameResponseList {
	return &models.GameResponseList{}
}

func validGameResponseListResponse() *models.GameResponseList {
	return &models.GameResponseList{
		ItemsPerPage: 1,
		StartIndex:   1,
		TotalResults: 1,
		Items: []models.GameResponse{
			*validGameResponse(),
		},
	}
}

func TestUnitHandleGetListGames(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	Convey("Error retrieving list of games from dao", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().GetListGames().Return(nil, fmt.Errorf("error"))

		req, _ := http.NewRequest(http.MethodGet, "/games", nil)
		w := httptest.NewRecorder()

		mockGamesService.HandleGetListGames(w, req)
		So(w.Code, ShouldEqual, 500)
	})

	Convey("No items present in DAO", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().GetListGames().Return(emptyGameResponseListResponse(), nil)

		req, _ := http.NewRequest(http.MethodGet, "/games", nil)
		w := httptest.NewRecorder()

		mockGamesService.HandleGetListGames(w, req)
		So(w.Code, ShouldEqual, 404)
	})

	Convey("valid get list games request", t, func() {
		mockDao := dao.NewMockDAO(mockCtrl)
		mockCloud := cloud.NewMockCLOUD(mockCtrl)
		mockGamesService := createMockGamesService(mockDao, mockCloud)
		mockDao.EXPECT().GetListGames().Return(validGameResponseListResponse(), nil)

		req, _ := http.NewRequest(http.MethodGet, "/games", nil)
		w := httptest.NewRecorder()

		mockGamesService.HandleGetListGames(w, req)
		So(w.Code, ShouldEqual, 200)
	})
}
