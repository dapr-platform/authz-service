package api

import (
	"authz-service/config"
	"authz-service/model"
	"authz-service/service"
	"github.com/dapr-platform/common"
	"github.com/dchest/captcha"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func initOauthRoute(r chi.Router) {
	r.HandleFunc(common.BASE_CONTEXT+"/oauth/token", tokenHandler)
	r.HandleFunc(common.BASE_CONTEXT+"/oauth/token-by-field", tokenByFieldHandler)
	r.HandleFunc(common.BASE_CONTEXT+"/oauth/token-valid", tokenValidHandler)
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
	if field == "" {
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
				common.HttpError(w, common.ErrParam.AppendMsg("captcha key or value error"), http.StatusBadRequest)
				return
			}
			if !captcha.VerifyString(vKey, vVal) {
				common.HttpError(w, common.ErrParam.AppendMsg("captcha key or value error"), http.StatusNotAcceptable)
				return
			}
		}
		user, err := service.GetUserByFieldName(r.Context(), "identity", identity)
		if err != nil {
			common.HttpError(w, common.ErrParam.AppendMsg(err.Error()), http.StatusInternalServerError)
			return
		}
		if user == nil {
			common.Logger.Error("user is nil for ", identity)
			common.HttpError(w, common.ErrParam.AppendMsg("user is nil"), http.StatusBadRequest)
			return
		}
		if user.Status != 1 {
			common.HttpError(w, common.ErrParam.AppendMsg("user is forbidden"), http.StatusBadRequest)
			return
		}
		r.Form.Set("username", user.ID)
		err = service.OauthServer.HandleTokenRequest(w, r)
		if err != nil {
			common.HttpError(w, common.ErrService.AppendMsg(err.Error()), http.StatusInternalServerError)
		}
	} else if grantType == "refresh_token" {
		err := service.OauthServer.HandleTokenRequest(w, r)
		if err != nil {
			common.HttpError(w, common.ErrParam.AppendMsg(err.Error()), http.StatusInternalServerError)
		}
	} else {
		common.HttpError(w, common.ErrParam.AppendMsg("grant_type "+grantType+" is not support "), http.StatusBadRequest)
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
