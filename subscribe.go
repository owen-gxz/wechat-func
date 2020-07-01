package wechat

import (
	"encoding/json"
	"github.com/owen-gxz/wechat-func/util"
)

const (
	SubscribeSendUrl = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token="
)

type SubscribeH map[string]SubscribeSendDate
type SubscribeSendRequest struct {
	Touser           string     `json:"touser"`
	TemplateID       string     `json:"template_id"`
	Page             string     `json:"page"`
	MiniprogramState string     `json:"miniprogram_state"`
	Lang             string     `json:"lang"`
	Data             SubscribeH `json:"data"`
}

type SubscribeSendDate struct {
	Value string `json:"value"`
}

// subscribeMessage.send
func SubscribeSend(token string, v SubscribeSendRequest) *WxErr {
	url := SubscribeSendUrl + token
	body, err := util.PostJson(url, v)
	if err != nil {
		return &WxErr{-1, err.Error()}
	}
	rst := new(WxErr)
	err = json.Unmarshal(body, rst)
	if err != nil {
		return &WxErr{-1, err.Error()}
	}
	return rst
}
