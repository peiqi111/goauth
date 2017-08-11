package provider

import (
	"net/http"
)

var OAuthProvider map[string]Client

type Result struct {
	User
	Error
}

type User interface {
	Id() string
	Provider() string
	Name() string
	Email() string
	Picture() string
	Link() string
	Bio() string
	Location() string
}

type Error interface {
	GetError() string
	SetError(err string)
	HttpErrToJson(r *http.Request)
}

type Client interface {
	GetData() *ClientData
	GetUser() *User
	GetError() *Error
	GetCodeUrl() (string, string)
	GetTokenUrl() (string, string)
	GetUserUrl() (string, string)
}

func init() {
	OAuthProvider = make(map[string]Client)
}

func Register(name string, client Client) {
	OAuthProvider[name] = client
}
