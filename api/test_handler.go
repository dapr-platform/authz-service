package api

import (
	"authz-service/entity"
	"authz-service/model"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func initTestHandler(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/test", testHandler)
}

// @Summary test
// @Description test
// @Tags Test
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /test [get]
func testHandler(w http.ResponseWriter, r *http.Request) {
	user, err := common.DbGetOne[entity.UserInfo](r.Context(), common.GetDaprClient(), model.User_with_menuTableInfo.Name, "id=admin")
	if err != nil {
		common.Logger.Error("db query error", err)
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpResult(w, common.OK.WithData(user))
}
