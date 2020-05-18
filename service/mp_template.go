package service

import (
	"github.com/owen-gxz/wechat-func"
)

// AddTemplate 获取模板
func (s *Server) AddTemplate(IdShort string) (id string, err error) {
	return wechat.AddTemplate(s.GetAccessToken(), IdShort)
}

// DelTemplate 删除模板
func (s *Server) DelTemplate(id string) (err error) {
	return wechat.DelTemplate(s.GetAccessToken(), id)
}

// GetAllTemplate 获取模板
func (s *Server) GetAllTemplate() (templist []wechat.MpTemplate, err error) {
	return wechat.GetAllTemplate(s.GetAccessToken())
}

// SendTemplate 发送模板消息，data通常是map[string]struct{value string,color string}
func (s *Server) SendTemplate(to, id, url, appid, pagepath string, data interface{}) *wechat.WxErr {
	return wechat.SendTemplate(s.GetAccessToken(), to, id, url, appid, pagepath, data)
}
