package api

import (
	"authz-service/config"
	"authz-service/model"
	"authz-service/service"
	"net/http"

	"github.com/dapr-platform/common"
	"github.com/dchest/captcha"
	"github.com/go-chi/chi/v5"
)

func initOauthRoute(r chi.Router) {
	r.HandleFunc(common.BASE_CONTEXT+"/oauth/token", tokenHandler)
	r.HandleFunc(common.BASE_CONTEXT+"/oauth/token-by-field", tokenByFieldHandler)
	r.HandleFunc(common.BASE_CONTEXT+"/oauth/token-valid", tokenValidHandler)
	r.Post(common.BASE_CONTEXT+"/oauth/sso-token", ssoTokenHandler)
	r.Post(common.BASE_CONTEXT+"/oauth/sso-logout", ssoLogoutHandler)
}

// @Summary 根据用户字段 get token
// @Description 根据用户字段 get token,例如根据手机号,只限server端使用，
//
//	需要带有client_id,client_secret, 这两个字段部署时可通过环境变量设置,避免安全问题
//
// @Tags Oauth2
// @Param field formData string true "field"
// @Param value formData string true "value"
// @Param client_id formData string true "client_id"
// @Param client_secret formData string false "client_secret"
// @Produce  json
// @Success 200 {object} model.TokenInfo "token info"
// @Failure 500 {object} string ""
// @Router /oauth/token-by-field [post]
func tokenByFieldHandler(w http.ResponseWriter, r *http.Request) {
	field := r.FormValue("field")
	if field == "" {
		common.Logger.Error("field is required")
		http.Error(w, "field is required", http.StatusBadRequest)
		return
	}
	value := r.FormValue("value")
	if value == "" {
		common.Logger.Error("value is required")
		http.Error(w, "value is required", http.StatusBadRequest)
		return
	}
	client_id := r.FormValue("client_id")
	if client_id == "" {
		common.Logger.Error("client_id is required")
		http.Error(w, "client_id is required", http.StatusBadRequest)
		return
	}
	client_secret := r.FormValue("client_secret")
	if client_secret == "" {
		common.Logger.Error("client_secret is required")
		http.Error(w, "client_secret is required", http.StatusBadRequest)
		return
	}
	if client_id != config.CLIENT_ID || client_secret != config.CLIENT_SECRET {
		common.HttpError(w, common.ErrParam.AppendMsg("client_id or client_secret error"), http.StatusBadRequest)
		return
	}
	user, err := service.GetUserByFieldName(r.Context(), field, value)
	if err != nil {
		common.HttpError(w, common.ErrParam.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}
	if user == nil {
		common.HttpError(w, common.ErrParam.AppendMsg("user is nil"), http.StatusBadRequest)
		return
	}
	if user.Status != 1 {
		common.HttpError(w, common.ErrParam.AppendMsg("user is forbidden"), http.StatusBadRequest)
		return
	}
	r.Form.Set("username", user.ID)
	r.Form.Set("password", user.Password)
	r.Form.Set("grant_type", "password")
	err = service.OauthServer.HandleTokenRequest(w, r)
	if err != nil {
		common.HttpError(w, common.ErrService.AppendMsg(err.Error()), http.StatusInternalServerError)
	}

}

// @Summary get token
// @Description get token
// @Tags Oauth2
// @Param identity formData string true "identity"
// @Param password formData string true "password"
// @Param grant_type formData string true "grant_type"
// @Param captcha_key formData string false "captcha_key"
// @Param captcha_value formData string false "captcha_value"
// @Param refresh_token formData string false "captcha_value"
// @Produce  json
// @Success 200 {object} model.TokenInfo "token info"
// @Failure 500 {object} string ""
// @Router /oauth/token [post]
func tokenHandler(w http.ResponseWriter, r *http.Request) {
	grantType := r.FormValue("grant_type")
	if grantType == "" {
		common.Logger.Error("grant_type is required")
		http.Error(w, "grant_type is required", http.StatusBadRequest)
		return
	}
	if grantType == "password" {
		identity := r.FormValue("identity")
		if config.VERIFY_CAPTCHA {
			vKey := r.FormValue("captcha_key")
			vVal := r.FormValue("captcha_value")
			if vKey == "" || vVal == "" {
				common.Logger.Error("captcha key or value error")
				common.HttpError(w, common.ErrParam.AppendMsg("captcha key or value error"), http.StatusBadRequest)
				return
			}
			if !captcha.VerifyString(vKey, vVal) {
				common.Logger.Error("captcha key or value error")
				common.HttpError(w, common.ErrParam.AppendMsg("captcha key or value error"), http.StatusNotAcceptable)
				return
			}
		}
		user, err := service.GetUserByFieldName(r.Context(), "identity", identity)
		if err != nil {
			common.Logger.Error("get user by field error: " + err.Error())
			common.HttpError(w, common.ErrParam.AppendMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		if user == nil {
			common.Logger.Error("user is nil for ", identity)
			common.HttpError(w, common.ErrParam.AppendMsg("user is nil"), http.StatusBadRequest)
			return
		}
		if user.Status != 1 {
			common.Logger.Error("user is forbidden for ", identity)
			common.HttpError(w, common.ErrParam.AppendMsg("user is forbidden"), http.StatusBadRequest)
			return
		}
		r.Form.Set("username", user.ID)
		err = service.OauthServer.HandleTokenRequest(w, r)
		if err != nil {
			common.Logger.Error("handle token request error: " + err.Error())
			common.HttpError(w, common.ErrService.AppendMsg(err.Error()), http.StatusInternalServerError)
		}
	} else if grantType == "refresh_token" {
		err := service.OauthServer.HandleTokenRequest(w, r)
		if err != nil {
			common.Logger.Error("handle refresh token request error: " + err.Error())
			common.HttpError(w, common.ErrParam.AppendMsg(err.Error()), http.StatusInternalServerError)
		}
	} else {
		common.Logger.Error("grant_type " + grantType + " is not support")
		common.HttpError(w, common.ErrParam.AppendMsg("grant_type "+grantType+" is not support"), http.StatusBadRequest)
	}

}

// @Summary  token valid
// @Description  token valid
// @Tags Oauth2
// @Param info body model.TokenInfo true "tokenInfo"
// @Produce  json
// @Success 200 {object} model.TokenInfo "token info"
// @Failure 500 {object} string ""
// @Router /oauth/token-valid [post]
func tokenValidHandler(w http.ResponseWriter, r *http.Request) {
	var tokenInfo model.TokenInfo
	err := common.ReadRequestBody(r, &tokenInfo)
	if err != nil {
		common.Logger.Error("tokenInfo read error: " + err.Error())
		http.Error(w, "tokenInfo read error: "+err.Error(), http.StatusBadRequest)
		return
	}
	valid, err := service.ValidToken(r.Context(), tokenInfo)
	if valid {
		common.HttpResult(w, common.OK)
	} else {
		http.Error(w, "tokenInfo read error: "+err.Error(), http.StatusForbidden)
	}
}

// @Summary SSO单点登录获取token
// @Description 通过中台SSO ticket换取本地JWT token
// @Tags Oauth2
// @Accept json
// @Param req body service.SSORestoreRequest true "SSO ticket"
// @Produce json
// @Success 200 {object} model.TokenInfo "token info"
// @Failure 400 {object} string ""
// @Failure 500 {object} string ""
// @Router /oauth/sso-token [post]
func ssoTokenHandler(w http.ResponseWriter, r *http.Request) {
	if !config.SSO_ENABLED {
		common.HttpError(w, common.ErrParam.AppendMsg("SSO功能未启用"), http.StatusBadRequest)
		return
	}

	var req service.SSORestoreRequest
	err := common.ReadRequestBody(r, &req)
	if err != nil {
		common.Logger.Error("读取SSO请求体失败: " + err.Error())
		common.HttpError(w, common.ErrParam.AppendMsg(err.Error()), http.StatusBadRequest)
		return
	}
	if req.Ticket == "" {
		common.HttpError(w, common.ErrParam.AppendMsg("ticket不能为空"), http.StatusBadRequest)
		return
	}

	ssoUser, err := service.SSORestoreTicket(req.Ticket)
	if err != nil {
		common.Logger.Error("SSO ticket验证失败: " + err.Error())
		common.HttpError(w, common.ErrService.AppendMsg("SSO ticket验证失败: "+err.Error()), http.StatusInternalServerError)
		return
	}

	user, err := service.GetUserByFieldName(r.Context(), "identity", ssoUser.LoginName)
	if err != nil {
		common.Logger.Error("查询本地用户失败: " + err.Error())
		common.HttpError(w, common.ErrService.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}
	if user == nil {
		common.Logger.Error("SSO用户在本地不存在: " + ssoUser.LoginName)
		common.HttpError(w, common.ErrParam.AppendMsg("用户不存在，请先同步用户数据"), http.StatusBadRequest)
		return
	}
	if user.Status != 1 {
		common.HttpError(w, common.ErrParam.AppendMsg("用户已被禁用"), http.StatusBadRequest)
		return
	}

	r.Form = make(map[string][]string)
	r.Form.Set("username", user.ID)
	r.Form.Set("password", user.Password)
	r.Form.Set("grant_type", "password")
	err = service.OauthServer.HandleTokenRequest(w, r)
	if err != nil {
		common.Logger.Error("签发token失败: " + err.Error())
		common.HttpError(w, common.ErrService.AppendMsg(err.Error()), http.StatusInternalServerError)
	}
}

// SSOLogoutReq SSO登出请求体
type SSOLogoutReq struct {
	Code string `json:"code"`
}

// @Summary SSO单点登出
// @Description 调用中台SSO登出接口，根据用户编号一键退出
// @Tags Oauth2
// @Accept json
// @Param req body SSOLogoutReq true "用户编号"
// @Produce json
// @Success 200 {object} common.Response "成功"
// @Failure 500 {object} common.Response "失败"
// @Router /oauth/sso-logout [post]
func ssoLogoutHandler(w http.ResponseWriter, r *http.Request) {
	if !config.SSO_ENABLED {
		common.HttpError(w, common.ErrParam.AppendMsg("SSO功能未启用"), http.StatusBadRequest)
		return
	}

	var req SSOLogoutReq
	err := common.ReadRequestBody(r, &req)
	if err != nil {
		common.HttpError(w, common.ErrParam.AppendMsg(err.Error()), http.StatusBadRequest)
		return
	}
	if req.Code == "" {
		common.HttpError(w, common.ErrParam.AppendMsg("code不能为空"), http.StatusBadRequest)
		return
	}

	err = service.SSORevokeByCode(req.Code)
	if err != nil {
		common.Logger.Error("SSO登出失败: " + err.Error())
		common.HttpError(w, common.ErrService.AppendMsg(err.Error()), http.StatusInternalServerError)
		return
	}

	common.HttpResult(w, common.OK)
}
