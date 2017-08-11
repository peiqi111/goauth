package provider

import (
	"net/http"
	"net/url"

	utils "github.com/peiqi/goauth/utils"
)

type GoogleClient struct {
	ClientData
	User
	Error
}

type GoogleUser struct {
	UserIss      string `json:"iss"`
	UserEmail    string `json:"email"`
	UserName     string `json:"name"`
	UserPicture  string `json:"picture"`
	UserProfile  string `json:"profile"`
	UserSub      string `json:"sub"`
	UserLocation string `json:"locale"`
}

func (u *GoogleUser) Id() string       { return u.UserSub }
func (u *GoogleUser) Provider() string { return u.UserIss }
func (u *GoogleUser) Name() string     { return u.UserName }
func (u *GoogleUser) Email() string    { return u.UserEmail }
func (u *GoogleUser) Picture() string  { return u.UserPicture }
func (u *GoogleUser) Link() string     { return u.UserProfile }
func (u *GoogleUser) Bio() string      { return "" }
func (u *GoogleUser) Location() string { return u.UserLocation }

type GoogleError struct {
	ErrorStr string `json:"error_description"`
}

func (e *GoogleError) GetError() string    { return e.ErrorStr }
func (e *GoogleError) SetError(err string) { e.ErrorStr = err }
func (e *GoogleError) HttpErrToJson(r *http.Request) {
	e.ErrorStr = r.URL.Query().Get(`error_description`)
}

func init() {
	client := &GoogleClient{
		ClientData: ClientData{
			CodeUri:          `https://accounts.google.com/o/oauth2/v2/auth?`,
			AccessTokenUri:   `https://www.googleapis.com/oauth2/v4/token?`,
			AuthorizationUri: `https://www.googleapis.com/oauth2/v3/userinfo?`,
			Token:            Token{},
		},
		User:  &GoogleUser{},
		Error: &GoogleError{},
	}
	Register(`google`, client)
	Register(`Google`, client)
}

func (this *GoogleClient) GetData() *ClientData { return &this.ClientData }

func (this *GoogleClient) GetUser() *User { return &this.User }

func (this *GoogleClient) GetError() *Error { return &this.Error }

func (this *GoogleClient) GetCodeUrl() (string, string) {
	state := utils.GetState(16)
	codeUrl := url.Values{}
	codeUrl.Set(`response_type`, `code`)
	codeUrl.Set(`client_id`, this.ClientData.ClientId)
	codeUrl.Set(`redirect_uri`, this.ClientData.Referer+this.ClientData.CallbackUri)
	codeUrl.Set(`scope`, this.ClientData.Scope)
	codeUrl.Set(`state`, state)
	return this.ClientData.CodeUri + codeUrl.Encode(), state
}

func (this *GoogleClient) GetTokenUrl() (string, string) {
	tokenUrl := url.Values{}
	tokenUrl.Set(`client_id`, this.ClientData.ClientId)
	tokenUrl.Set(`client_secret`, this.ClientData.ClientSecret)
	tokenUrl.Set(`code`, this.ClientData.Code)
	tokenUrl.Set(`redirect_uri`, this.ClientData.Referer+this.ClientData.CallbackUri)
	tokenUrl.Set(`grant_type`, `authorization_code`)
	return this.ClientData.AccessTokenUri + tokenUrl.Encode(), `POST`
}

func (this *GoogleClient) GetUserUrl() (string, string) {
	userUrl := url.Values{}
	userUrl.Set(`access_token`, this.ClientData.Token.AccessToken)
	return this.ClientData.AuthorizationUri + userUrl.Encode(), `GET`
}
