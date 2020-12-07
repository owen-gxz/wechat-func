package service

import (
	"fmt"
	"github.com/owen-gxz/wechat-func"
	"github.com/owen-gxz/wechat-func/util"
	"net/url"
)

// WXAPIOauth2 oauth2鉴权
const (
	WXAPIOauth2           = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%v&redirect_uri=%v&response_type=code&scope=%s&state=%s#wechat_redirect"
	WXAPIJscode2session   = "https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code"
	CorpAPIJscode2session = "https://qyapi.weixin.qq.com/cgi-bin/miniprogram/jscode2session?access_token=%v&js_code=%v&grant_type=authorization_code"
	JSAPIOauthToken       = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
)
const (
	ScopeBase     = "snsapi_base"
	ScopeUserinfo = "snsapi_userinfo"
)

// WxSession 兼容企业微信和服务号
type WxSession struct {
	wechat.WxErr
	SessionKey string `json:"session_key"`
	// corp
	CorpId string `json:"corpid"`
	UserId string `json:"userid"`
	// mp
	OpenId  string `json:"openid"`
	UnionId string `json:"unionid"`
}

type JsTokenResp struct {
	wechat.WxErr
	OpenId       string `json:"openid"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
}

// GetOauth2Url 获取鉴权页面
func GetOauth2Url(appId, host, scope, state string) string {
	if scope == "" {
		scope = ScopeBase
	}
	if state == "" {
		state = appId
	}
	return fmt.Sprintf(WXAPIOauth2, appId, url.QueryEscape(host), scope, state)
}

//Jscode code换token
func (s *Server) JSAPIOauthToken(code string) (*JsTokenResp, error) {
	url := fmt.Sprintf(JSAPIOauthToken, s.AppId, s.Secret, code)
	ws := new(JsTokenResp)
	err := util.GetJson(url, ws)

	if ws.Error() != nil {
		err = ws.Error()
		return nil, err
	}
	return ws, err
}

//Jsget userinfo
func (s *Server) JSAPIGetUserInfo(token, openID string) (*wechat.MpUserInfo, error) {
	return wechat.GetMpUserInfoBySns(token, openID)
}

// Jscode2Session code换session
func (s *Server) Jscode2Session(code string) (ws *WxSession, err error) {
	url := fmt.Sprintf(WXAPIJscode2session, s.AppId, s.Secret, code)
	ws = new(WxSession)
	err = util.GetJson(url, ws)

	if ws.Error() != nil {
		err = ws.Error()
	}
	return
}

// Jscode2SessionEnt code换session（企业微信）
func (s *Server) Jscode2SessionEnt(code string) (ws *WxSession, err error) {
	url := fmt.Sprintf(CorpAPIJscode2session, s.GetAccessToken(), code)
	ws = new(WxSession)
	err = util.GetJson(url, ws)

	if ws.Error() != nil {
		err = ws.Error()
	}
	return
}
