package service

import (
	"github.com/owen-gxz/wechat-func"
)

// subscribeMessage.send
func (s *Server) SubscribeSend( v wechat.SubscribeSendRequest) *wechat.WxErr {
	return wechat.SubscribeSend(s.GetAccessToken(), v)
}
