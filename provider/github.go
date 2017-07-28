package provider

import (
	utils "github.com/peiqi/goauth/utils"
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
	UserCompany  string `json:"company"`
	UserLink     string `json:"html_url"`
	UserLogin    string `json:"login"`
	UserBio      string `json:"bio"`
	UserLocation string `json:"location"`
}

func (u *GitHubUser) Id() string       { return u.UserLogin }
func (u *GitHubUser) Provider() string { return "github.com" }
func (u *GitHubUser) Name() string     { return u.UserName }
func (u *GitHubUser) Email() string    { return u.UserEmail }
func (u *GitHubUser) Org() string      { return u.UserCompany }
func (u *GitHubUser) Picture() string  { return "https://secure.gravatar.com/avatar/" + u.UserGravatar }
func (u *GitHubUser) Link() string     { return u.UserLink }
func (u *GitHubUser) Bio() string      { return u.UserBio }
func (u *GitHubUser) Location() string { return u.UserLocation }

type GitHubError struct {
	ErrorCode string `json:"error"`
	ErrorDes  string `json:"error_description"`
	ErrorUri  string `json:"error_uri"`
	ErrorJson string
}

func (e *GitHubError) GetErrorDes() string { return e.ErrorUri }

func (e *GitHubError) GetErrorJson() string { return e.ErrorJson }

func (e *GitHubError) SetErrorJson(err string) { e.ErrorJson = err }

func init() {
	client := &GitHubClient{
		ClientData: ClientData{
			CodeUri:          "https://github.com/login/oauth/authorize",
			AccessTokenUri:   "https://github.com/login/oauth/access_token",
			AuthorizationUri: "https://api.github.com/user",
			Token:            Token{},
		},
		User:  &GitHubUser{},
		Error: &GitHubError{},
	}
	Register("github", client)
	Register("Github", client)
	Register("GitHub", client)
}

func (this *GitHubClient) GetData() *ClientData { return &this.ClientData }

func (this *GitHubClient) GetUser() *User { return &this.User }

func (this *GitHubClient) GetError() *Error { return &this.Error }

func (this *GitHubClient) GetCodeUrl() (string, string) {
	state := utils.GetState()
	url := this.ClientData.CodeUri +
		"?client_id=" + this.ClientData.ClientId +
		"&redirect_uri=" + this.ClientData.Referer + this.ClientData.CallbackUri +
		"&scope=" + this.ClientData.Scope +
		"&state=" + state
	return url, state
}

func (this *GitHubClient) GetTokenUrl() (string, string) {
	url := this.ClientData.AccessTokenUri +
		"?client_id=" + this.ClientData.ClientId +
		"&client_secret=" + this.ClientData.ClientSecret +
		"&code=" + this.ClientData.Code +
		"&redirect_uri=" + this.ClientData.Referer + this.ClientData.CallbackUri
	return url, "POST"
}

func (this *GitHubClient) GetUserUrl() (string, string) {
	url := this.ClientData.AuthorizationUri +
		"?access_token=" + this.ClientData.Token.AccessToken
	return url, "GET"
}
