package api

import (
	"authz-service/entity"
	"authz-service/service"
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func initUserExtHandler(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/user/current", userCurrentHandler)
	r.Post(common.BASE_CONTEXT+"/user/login", userLoginHandler)
}

// @Summary get user info
// @Description get user info
// @Tags User
// @Produce  json
// @Success 200 {object} common.Response{data=entity.UserInfo} "UserInfo"
// @Failure 500 {object} common.Response ""
// @Router /user/current [get]
func userCurrentHandler(w http.ResponseWriter, r *http.Request) {
	sub, err := common.ExtractUserSub(r)
	if err != nil {
		common.Logger.Error(err)
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	user, err := service.GetCurrentUserInfo(r.Context(), sub)
	if err != nil {
		common.Logger.Error(err)
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.Logger.Debug("user:", user)
	common.HttpResult(w, common.OK.WithData(user))

}

// @Summary login by identity and password
// @Description login  by identity and password
// @Tags User
// @Param req body entity.UserLoginReq true "req"
// @Produce  json
// @Success 200 {object} common.Response{data=entity.UserInfo} "UserInfo"
// @Failure 500 {object} common.Response ""
// @Router /user/login [post]
func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	var req entity.UserLoginReq
	err := common.ReadRequestBody(r, &req)
	if err != nil {
		common.Logger.Error(err)
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	user, err := service.GetUserInfoByIdentityAndPassword(r.Context(), req.Identity, req.Password)
	common.Logger.Debug("user:", user)
	common.HttpResult(w, common.OK.WithData(user))

}
