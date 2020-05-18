package service

import (
	"github.com/owen-gxz/wechat-func"
)

// GetMenu 获取应用菜单
func (s *Server) GetMenu() (m *wechat.Menu, err error) {
	return wechat.GetMenu(s.GetAccessToken())

}

// AddMenu 创建应用菜单
func (s *Server) AddMenu(m *wechat.Menu) (err error) {
	return wechat.AddMenu(s.GetAccessToken(), m)

}

// DelMenu 删除应用菜单
func (s *Server) DelMenu() (err error) {
	return wechat.DelMenu(s.GetAccessToken())

}
