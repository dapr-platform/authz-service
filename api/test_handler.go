package api

import (
	"authz-service/config"
	"authz-service/entity"
	"authz-service/model"
	"authz-service/service"
	"fmt"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func initTestHandler(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/test", testHandler)
	r.Post(common.BASE_CONTEXT+"/sso/sync-members", ssoSyncMembersHandler)
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

// @Summary 手动触发SSO用户同步
// @Description 从中台全量同步人员数据到本地用户表
// @Tags SSO
// @Produce json
// @Success 200 {object} common.Response "同步结果"
// @Failure 500 {object} common.Response "错误信息"
// @Router /sso/sync-members [post]
func ssoSyncMembersHandler(w http.ResponseWriter, r *http.Request) {
	if !config.SSO_ENABLED {
		common.HttpResult(w, common.ErrParam.AppendMsg("SSO功能未启用"))
		return
	}
	syncCount, err := service.SSOSyncMembers(r.Context())
	if err != nil {
		common.Logger.Error("SSO用户同步失败: ", err)
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpResult(w, common.OK.WithData(fmt.Sprintf("同步完成，共同步 %d 个用户", syncCount)))
}
