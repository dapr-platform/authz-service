package api

import (
	"authz-service/dapr_store"
	"encoding/json"
	"github.com/dapr-platform/common"
	"github.com/dchest/captcha"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

var captchaHandler http.Handler

func init() {
	captcha.SetCustomStore(&dapr_store.DaprCaptchaStore{
		Expiration: time.Minute,
	})

	captchaHandler = captcha.Server(captcha.StdWidth, captcha.StdHeight)
}

func initCaptchaRoute(r chi.Router) {
	r.Get(common.BASE_CONTEXT+"/captcha-gen", captchaGenHandler)
	r.HandleFunc(common.BASE_CONTEXT+"/captcha/{file}", captchaHandler.ServeHTTP)
}

// @Summary get Captcha
// @Description get Captcha
// @Tags Oauth2
// @Produce  json
// @Success 200 {object} common.Response{data=string} "captcha"
// @Failure 500 {object} string ""
// @Router /captcha-gen [get]
func captchaGenHandler(w http.ResponseWriter, r *http.Request) {
	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}
	ret := common.OK.WithData(d.CaptchaId)
	bytes, _ := json.Marshal(ret)
	w.Write(bytes)
	return
}
