package service

import (
	myconfig "authz-service/config"
	"authz-service/dapr_store"
	"authz-service/eventpub"
	"authz-service/model"
	"context"
	"github.com/dapr-platform/common"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

var OauthServer *server.Server

func init() {
	manager := manage.NewDefaultManager()
	config := &manage.Config{AccessTokenExp: time.Second * time.Duration(common.USER_EXPIRED_SECONDS), RefreshTokenExp: time.Second * time.Duration(common.USER_EXPIRED_SECONDS) * 12 * 10, IsGenerateRefresh: true}
	manager.SetPasswordTokenCfg(config)
	manager.MustTokenStorage(dapr_store.NewDaprTokenStore())
	clientStore := store.NewClientStore()
	err := clientStore.Set(myconfig.CLIENT_ID, &models.Client{ID: myconfig.CLIENT_ID, Secret: myconfig.CLIENT_SECRET})
	if err != nil {
		common.Logger.Error("Set client store error", err)
		panic(err)
	}

	manager.MapClientStorage(clientStore)
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	serverConfig := &server.Config{
		TokenType:            "Bearer",
		AllowedResponseTypes: []oauth2.ResponseType{oauth2.Token},
		AllowedGrantTypes: []oauth2.GrantType{
			oauth2.PasswordCredentials,
			oauth2.Refreshing,
		},
	}
	OauthServer = server.NewServer(serverConfig, manager)
	OauthServer.SetPasswordAuthorizationHandler(passwordAuthHandler)
	OauthServer.SetClientInfoHandler(func(r *http.Request) (clientID, clientSecret string, err error) {
		return clientID, clientSecret, nil
	})

}
func ValidToken(ctx context.Context, tokenInfo model.TokenInfo) (valid bool, err error) {
	_, err = OauthServer.Manager.LoadAccessToken(ctx, tokenInfo.AccessToken)
	if err == nil {
		valid = true
	}
	return
}
func passwordAuthHandler(ctx context.Context, clientId, id, password string) (userID string, err error) {

	user, err := GetUserByIdAndPassword(ctx, id, password)
	if err != nil {
		common.Logger.Error("GetUserByIdAndPassword error", err)
		return "", errors.Wrap(err, "GetUserByIdAndPassword error")
	}
	if user == nil {
		common.Logger.Error("GetUserByIdAndPassword not found" + id)
		return "", nil
	}
	msg := &common.InternalMessage{
		common.INTERNAL_MESSAGE_KEY_TYPE:    common.INTERNAL_MESSAGE_TYPE_USER_LOGIN,
		common.INTERNAL_MESSAGE_KEY_USER_ID: id,
	}
	err = eventpub.PublishInternalMessage(ctx, msg)
	if err != nil {
		common.Logger.Error(err)
	}
	return id, nil

}
