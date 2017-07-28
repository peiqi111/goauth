package provider

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
	Org() string
	Picture() string
	Link() string
	Bio() string
	Location() string
}

type Error interface {
	GetErrorDes() string
	GetErrorJson() string
	SetErrorJson(err string)
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
