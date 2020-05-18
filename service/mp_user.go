package service

import (
	"github.com/owen-gxz/wechat-func"
)

// BatchGetAll 获取所有公众号用户
func (s *Server) BatchGetAll() (ui []wechat.MpUserInfo, err error) {
	return wechat.BatchGetAll(s.GetAccessToken())
}

// BatchGet 批量获取公众号用户信息
func (s *Server) BatchGet(ul []string) (ui []wechat.MpUserInfo, err error) {
	return wechat.BatchGet(s.GetAccessToken(), ul)
}

// GetAllMpUserList 获取所有用户ID
func (s *Server) GetAllMpUserList() (ul []string, err error) {
	return wechat.GetAllMpUserList(s.GetAccessToken())
}

// GetMpUserList 获取用户信息，根据openid
func (s *Server) GetMpUserList(openid ...string) (ul *wechat.MpUser, err error) {
	return s.GetMpUserList(openid...)
}

// GetMpUserInfo 获取用户详情
func (s *Server) GetMpUserInfo(openid string, lang ...string) (user *wechat.MpUserInfo, err error) {
	return wechat.GetMpUserInfo(s.GetAccessToken(), openid, lang...)
}
