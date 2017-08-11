package provider

import (
	utils "github.com/peiqi/goauth/utils"

	"net/http"
	"net/url"
)

type GitHubClient struct {
	ClientData
	User
	Error
}

type GitHubUser struct {
	UserEmail    string `json:"email"`
	UserName     string `json:"name"`
	UserGravatar string `json:"gravatar_id"`
	UserLink     string `json:"html_url"`
	UserLogin    string `json:"login"`
	UserBio      string `json:"bio"`
	UserLocation string `json:"location"`
}

func (u *GitHubUser) Id() string       { return u.UserLogin }
func (u *GitHubUser) Provider() string { return `github.com` }
func (u *GitHubUser) Name() string     { return u.UserName }
func (u *GitHubUser) Email() string    { return u.UserEmail }
func (u *GitHubUser) Picture() string  { return `https://secure.gravatar.com/avatar/` + u.UserGravatar }
func (u *GitHubUser) Link() string     { return u.UserLink }
func (u *GitHubUser) Bio() string      { return u.UserBio }
func (u *GitHubUser) Location() string { return u.UserLocation }

type GitHubError struct {
	ErrorStr string `json:"error_description"`
}

func (e *GitHubError) GetError() string    { return e.ErrorStr }
func (e *GitHubError) SetError(err string) { e.ErrorStr = err }
func (e *GitHubError) HttpErrToJson(r *http.Request) {
	e.ErrorStr = r.URL.Query().Get(`error_description`)
}

func init() {
	client := &GitHubClient{
		ClientData: ClientData{
			CodeUri:          `https://github.com/login/oauth/authorize?`,
			AccessTokenUri:   `https://github.com/login/oauth/access_token?`,
			AuthorizationUri: `https://api.github.com/user?`,
			Token:            Token{},
		},
		User:  &GitHubUser{},
		Error: &GitHubError{},
	}
	Register(`github`, client)
	Register(`Github`, client)
	Register(`GitHub`, client)
}

func (this *GitHubClient) GetData() *ClientData { return &this.ClientData }

func (this *GitHubClient) GetUser() *User { return &this.User }

func (this *GitHubClient) GetError() *Error { return &this.Error }

func (this *GitHubClient) GetCodeUrl() (string, string) {
	state := utils.GetState(16)
	codeUrl := url.Values{}
	codeUrl.Set(`client_id`, this.ClientData.ClientId)
	codeUrl.Set(`redirect_uri`, this.ClientData.Referer+this.ClientData.CallbackUri)
	codeUrl.Set(`scope`, this.ClientData.Scope)
	codeUrl.Set(`state`, state)
	return this.ClientData.CodeUri + codeUrl.Encode(), state
}

func (this *GitHubClient) GetTokenUrl() (string, string) {
	tokenUrl := url.Values{}
	tokenUrl.Set(`client_id`, this.ClientData.ClientId)
	tokenUrl.Set(`client_secret`, this.ClientData.ClientSecret)
	tokenUrl.Set(`code`, this.ClientData.Code)
	tokenUrl.Set(`redirect_uri`, this.ClientData.Referer+this.ClientData.CallbackUri)
	return this.ClientData.AccessTokenUri + tokenUrl.Encode(), `POST`
}

func (this *GitHubClient) GetUserUrl() (string, string) {
	userUrl := url.Values{}
	userUrl.Set(`access_token`, this.ClientData.Token.AccessToken)
	return this.ClientData.AuthorizationUri + userUrl.Encode(), `GET`
}
