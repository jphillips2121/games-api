package controllers

import (
	"github.com/gorilla/mux"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRegister(t *testing.T) {
	Convey("Register routes", t, func() {
		router := mux.NewRouter()
		Register(router)
		So(router.GetRoute("createGame"), ShouldNotBeNil)
		So(router.GetRoute("getListGames"), ShouldNotBeNil)
		So(router.GetRoute("getGame"), ShouldNotBeNil)
		So(router.GetRoute("updateGame"), ShouldNotBeNil)
		So(router.GetRoute("deleteGame"), ShouldNotBeNil)
	})
}
