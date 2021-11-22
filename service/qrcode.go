package service

import "github.com/owen-gxz/wechat-func"

func (s *Server) MiniprogramGetUnlimitedQrcode(scene string) ([]byte, error) {
	return wechat.MiniprogramGetUnlimitedQrcode(s.GetAccessToken(), scene)
}
