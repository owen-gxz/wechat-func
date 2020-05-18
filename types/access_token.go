package types

type AccessTokenServer interface {
	Token() string
	SetToken(t AccessToken) error
	GetToken() (*AccessToken, error)
}

//
type AccessToken struct {
	Token     string `json:"token"` // token
	ExpiresIn int64  `json:"expires_in"`
	// 第三方平台使用
	AppID        string `json:"app_id"`        //
	RefreshToken string `json:"refresh_token"` //
}
