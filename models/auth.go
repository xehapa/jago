package models

const ClientId string = "dfuha7yjaakeroeydmikgi6hky"
const ClientSecret string = "ah6hsseus7suhiduokjy2sxfn4zvhkg7o5bzge5hbf2tuntfhgnq"
const AuthURL string = "https://id.jobadder.com/connect"

type AuthRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ApiUrl       string `json:"api"`
	RefreshToken string `json:"refresh_token"`
}
