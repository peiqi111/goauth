package provider

type OAuthConfig struct {
	ClientId     string `json:"id|key"`
	ClientSecret string `json:"secret"`
	CallbackUri  string `json:"redict"`
	Scope        string `json:"scope"`
	AuthResUri   string `json:"authres"`
}

type ClientData struct {
	Referer          string
	ClientId         string
	ClientSecret     string
	Scope            string
	CallbackUri      string
	AuthResUri       string
	CodeUri          string
	AccessTokenUri   string
	AuthorizationUri string
	Code             string
	Token            Token
}

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Time         int64
	ExpiresIn    int64  `json:"expires_in"`
	Scope        string `json:"scope"`
	Openid       string `json:"openid"`
}
