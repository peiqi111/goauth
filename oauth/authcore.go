package oauth

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	provider "github.com/peiqi/goauth/provider"
	utils "github.com/peiqi/goauth/utils"
)

type AuthHandler struct {
	Client provider.Client
}

var state string
var ResultData provider.Result

func (ah *AuthHandler) GetClient() *provider.Client {
	return &ah.Client
}

func (ah *AuthHandler) RedirectNeed(r *http.Request) bool {
	return r.URL.Query().Get("code") == ""
}

func (ah *AuthHandler) RedirectError(r *http.Request) bool {
	return r.URL.Query().Get("error") != ""
}

func (ah *AuthHandler) Redirect(w http.ResponseWriter, r *http.Request) error {
	if ah.RedirectError(r) {
		return errors.New("error report by http")
	} else if ah.RedirectNeed(r) {
		var redictUrl string
		reg, _ := regexp.Compile("(http://|https://)?([^/]*)")
		ah.Client.GetData().Referer = reg.FindString(r.Referer())
		redictUrl, state = ah.Client.GetCodeUrl()
		http.Redirect(w, r, redictUrl, http.StatusFound)
		return nil
	} else {
		err := ah.GetAccessToken(w, r)
		if err != nil {
			return err
		}
		user, err := ah.GetUser(w, r)
		if err != nil {
			return err
		}
		ResultData.User = *user
	}
	return nil
}

func (ah *AuthHandler) GetAccessToken(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Query().Get("state") == state {
		ah.Client.GetData().Code = r.URL.Query().Get("code")
		AccessTokenStr, err := utils.GetUrlData(ah.Client.GetTokenUrl())
		if err != nil {
			return err
		}
		err = json.Unmarshal(AccessTokenStr, &ah.Client.GetData().Token)
		if err != nil {
			return err
		}
		if len(ah.Client.GetData().Token.AccessToken) == 0 {
			return errors.New(string(AccessTokenStr))
		}
	} else {
		return errors.New("The state code we recevied is not matched with the state code we sent to the Github.")
	}
	return nil
}

func (ah *AuthHandler) GetUser(w http.ResponseWriter, r *http.Request) (*provider.User, error) {
	UserStr, err := utils.GetUrlData(ah.Client.GetUserUrl())
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(UserStr, ah.Client.GetUser())
	if err != nil {
		return nil, err
	}
	return ah.Client.GetUser(), nil
}

func (ah *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := ah.Redirect(w, r)
	if err != nil {
		jsonErr := json.Unmarshal([]byte(err.Error()), ah.Client.GetError())
		if jsonErr != nil {
			if ah.RedirectError(r) {
				(*ah.Client.GetError()).HttpErrToJson(r)
			} else {
				(*ah.Client.GetError()).SetError(err.Error())
			}
		}
		ResultData.Error = *ah.Client.GetError()
	}
	if ResultData.User != nil || ResultData.Error != nil {
		http.Redirect(w, r, ah.Client.GetData().AuthResUri, http.StatusSeeOther)
	}
}
