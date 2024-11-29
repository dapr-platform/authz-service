package api

import (
	"authz-service/entity"
	"authz-service/model"
	"authz-service/service"
	"net/http"
	"strconv"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func initUserExtHandler(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/user/current", userCurrentHandler)
	r.Post(common.BASE_CONTEXT+"/user/login", userLoginHandler)
	r.Post(common.BASE_CONTEXT+"/user/status", SetUserStatusHandler)
	r.Post(common.BASE_CONTEXT+"/user/change-password", changePasswordHandler)
}


// @Summary 修改密码
// @Description 修改密码
// @Tags User
// @Accept       json
// @Param item body entity.ChangePasswordInfo true "密码信息"
// @Produce  json
// @Success 200 {object} common.Response "用户"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user/change-password [post]
func changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var para entity.ChangePasswordInfo
	err := common.ReadRequestBody(r, &para)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg("数据错误"))
		return
	}
	err = service.ChangeUserPassword(r.Context(), para)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	common.HttpResult(w, common.OK)
}

// @Summary 设置用户状态
// @Description 设置用户状态, 1:正常，2：停用
// @Tags User
// @Param status query int true "用户状态"
// @Param item body []string true "用户id集合"
// @Produce  json
// @Success 200 {object} common.Response{data=model.User} "用户"
// @Failure 500 {object} common.Response "错误code和错误信息"
// @Router /user/status [post]
func SetUserStatusHandler(w http.ResponseWriter, r *http.Request) {
	statusStr := r.URL.Query().Get("status")
	if statusStr == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("status is blank"))
		return
	}
	status, _ := strconv.Atoi(statusStr)

	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.Logger.Error("read body error ", err)
	}
	errStr := ""
	for _, id := range ids {

		info := make(map[string]any, 0)
		info[model.User_FIELD_NAME_id] = id
		info[model.User_FIELD_NAME_status] = status
		err = common.DbUpsert[map[string]any](r.Context(), common.GetDaprClient(), info, model.UserTableInfo.Name, "id")
		if err != nil {
			errStr += id + " dbupsert error " + err.Error() + "\n"
		} else {
			//service.(r.Context(), id)
		}
	}
	if errStr != "" {
		common.HttpResult(w, common.ErrService.AppendMsg("db dbupsert error").AppendMsg(errStr))
		return
	}
	common.HttpResult(w, common.OK)

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
