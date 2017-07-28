package goauth

import (
	"encoding/json"
	"net/http"

	auth "github.com/peiqi/goauth/oauth"
	provider "github.com/peiqi/goauth/provider"
)

var Config map[string]provider.OAuthConfig
var AuthHandles map[string]*auth.AuthHandler

type ResHandler func(w http.ResponseWriter, r *http.Request, user provider.Result)

func init() {
	AuthHandles = make(map[string]*auth.AuthHandler)
	for name, provider := range provider.OAuthProvider {
		AuthHandles[name] = &auth.AuthHandler{Client: provider}
	}
}

func Auth(jsonStr string) error {
	json.Unmarshal([]byte(jsonStr), &Config)
	for name, _ := range Config {
		if AuthHandles[name] != nil {
			clientData := (*AuthHandles[name].GetClient()).GetData()
			clientData.ClientId = Config[name].ClientId
			clientData.ClientSecret = Config[name].ClientSecret
			clientData.CallbackUri = Config[name].CallbackUri
			clientData.Scope = Config[name].Scope
			clientData.AuthResUri = Config[name].AuthResUri
			http.Handle(Config[name].CallbackUri, AuthHandles[name])
		}
	}
	return nil
}

func ResHandlerFunc(userHandler ResHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userHandler(w, r, auth.ResultData)
	}
}
