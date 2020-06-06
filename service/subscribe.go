package service

import (
	"github.com/owen-gxz/wechat-func"
)

// subscribeMessage.send
func (s *Server) SubscribeSend(token string, v wechat.SubscribeSendRequest) *wechat.WxErr {
	return wechat.SubscribeSend(token, v)
}
