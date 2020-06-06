package service

import (
	"errors"
	"fmt"
	"github.com/owen-gxz/wechat-func"
	"github.com/owen-gxz/wechat-func/types"
	"github.com/owen-gxz/wechat-func/util"
	"log"
	"sync"
	"time"
)

type DefaultAccessToken struct {
	sync.Mutex // accessToken读取锁

	AppID     string `json:"app_id"`
	Appsecret string `json:"app_secret"`

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (s *DefaultAccessToken) getAccessToken() (*AccessToken, error) {
	url := fmt.Sprintf(wechat.WXAPIToken, s.AppID, s.Appsecret)
	at := new(AccessToken)
	if err := util.GetJson(url, at); err != nil {
		return nil, err
	}
	return at, nil
}

func (d *DefaultAccessToken) Token() string {
	at, err := d.GetToken()
	if err != nil {
		return ""
	}
	return at.Token
}

func (d *DefaultAccessToken) SetToken(t types.AccessToken) error {
	d.ExpiresIn = t.ExpiresIn
	d.AccessToken = t.Token
	return nil
}

func (d *DefaultAccessToken) GetToken() (*types.AccessToken, error) {
	d.Lock()
	defer d.Unlock()
	if d.AccessToken == "" || d.ExpiresIn < time.Now().Unix() {
		for i := 0; i < 3; i++ {
			at, err := d.getAccessToken()
			if err != nil {
				continue
			}
			if at.ErrCode > 0 {
				continue
			}
			ats := types.AccessToken{}
			ats.ExpiresIn = time.Now().Unix() + at.ExpiresIn - 5
			ats.Token = at.AccessToken
			err = d.SetToken(ats)
			if err != nil {
				continue
			}
			//Printf("***%v[%v]本地获取token:%v", util.Substr(s.Appsecret, 14, 30), s.Appsecret, s.Appsecret)
			return &ats, nil
		}
		return nil, errors.New("get token errror")
	}
	ats := types.AccessToken{}
	ats.Token = d.AccessToken
	ats.ExpiresIn = d.ExpiresIn
	return &ats, nil
}

var _ types.AccessTokenServer = (*DefaultAccessToken)(nil)

// FetchDelay 默认5分钟同步一次
var FetchDelay time.Duration = 5 * time.Minute

// AccessToken 回复体
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	wechat.WxErr
}

// GetAccessToken 读取AccessToken
func (s *Server) GetAccessToken() string {
	return s.tokenService.Token()
}

// GetUserAccessToken 获取企业微信通讯录AccessToken
func (s *Server) GetUserAccessToken() string {
	return s.GetAccessToken()
}

// Ticket JS-SDK
type Ticket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
	wechat.WxErr
}

// GetTicket 读取获取Ticket
func (s *Server) GetTicket() string {
	if s.ticket == nil || s.ticket.ExpiresIn < time.Now().Unix() {
		for i := 0; i < 3; i++ {
			err := s.getTicket()
			if err != nil {
				log.Printf("getTicket[%v] err:%v", s.AgentId, err)
				time.Sleep(time.Second)
				continue
			}
			break
		}
	}
	return s.ticket.Ticket
}

func (s *Server) getTicket() (err error) {
	url := wechat.WXAPIJsapi + s.GetAccessToken()
	at := new(Ticket)
	if err = util.GetJson(url, at); err != nil {
		return
	}
	if at.ErrCode > 0 {
		return at.Error()
	}
	Printf("[%v::%v-JsApi] >>> %+v", s.AppId, s.AgentId, *at)
	at.ExpiresIn = time.Now().Unix() + 500
	s.ticket = at
	return
}

// JsConfig Jssdk配置
type JsConfig struct {
	Beta      bool     `json:"beta"`
	Debug     bool     `json:"debug"`
	AppId     string   `json:"appId"`
	Timestamp int64    `json:"timestamp"`
	Nonsestr  string   `json:"nonceStr"`
	Signature string   `json:"signature"`
	JsApiList []string `json:"jsApiList"`
	Url       string   `json:"jsurl"`
	App       int      `json:"jsapp"`
}

// GetJsConfig 获取Jssdk配置
func (s *Server) GetJsConfig(Url string) *JsConfig {
	jc := &JsConfig{Beta: true, Debug: Debug, AppId: s.AppId}
	jc.Timestamp = time.Now().Unix()
	jc.Nonsestr = "esap"
	jc.Signature = util.SortSha1(fmt.Sprintf("jsapi_ticket=%v&noncestr=%v&timestamp=%v&url=%v", s.GetTicket(), jc.Nonsestr, jc.Timestamp, Url))
	// TODO：可加入其他apilist
	jc.JsApiList = []string{"scanQRCode"}
	jc.Url = Url
	jc.App = s.AgentId
	Println("jsconfig:", jc) // Debug
	return jc
}
