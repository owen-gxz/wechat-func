package wechat

import (
	"fmt"
	"github.com/owen-gxz/wechat-func/util"
)

// WXAPIMenuGet 微信菜单接口，兼容企业微信和服务号

var (
	CorpAgentid = ""
	WXAPIMenuGet = WXAPI + `menu/get?access_token=%s` + CorpAgentid
	WXAPIMenuAdd = WXAPI + `menu/create?access_token=%s` + CorpAgentid
	WXAPIMenuDel = WXAPI + `menu/delete?access_token=%s` + CorpAgentid
)

type (
	// Button 按钮
	Button struct {
		Name      string `json:"name"`
		Type      string `json:"type"`
		Key       string `json:"key"`
		Url       string `json:"url"`
		AppId     string `json:"appid"`
		PagePath  string `json:"pagepath"`
		SubButton []struct {
			Name     string `json:"name"`
			Type     string `json:"type"`
			Key      string `json:"key"`
			Url      string `json:"url"`
			AppId    string `json:"appid"`
			PagePath string `json:"pagepath"`
		} `json:"sub_button"`
	}
	// Menu 菜单
	Menu struct {
		WxErr
		Button []Button `json:"button"`

		Menu struct {
			Button []Button `json:"button"`
		} `json:"menu,omitempty"`
	}
)

// GetMenu 获取应用菜单
func GetMenu(token string) (m *Menu, err error) {
	m = new(Menu)
	url := fmt.Sprintf(WXAPIMenuGet, token)
	if err = util.GetJson(url, m); err != nil {
		return
	}
	if len(m.Menu.Button) == 0 && len(m.Button) > 0 {
		m.Menu.Button = m.Button
	}
	err = m.Error()
	return
}

// AddMenu 创建应用菜单
func AddMenu(token string, m *Menu) (err error) {
	e := new(WxErr)
	url := fmt.Sprintf(WXAPIMenuAdd, token)
	if err = util.PostJsonPtr(url, m, e); err != nil {
		return
	}
	return e.Error()
}

// DelMenu 删除应用菜单
func DelMenu(token string) (err error) {
	e := new(WxErr)
	url := fmt.Sprintf(WXAPIMenuDel, token)
	if err = util.GetJson(url, e); err != nil {
		return
	}
	return e.Error()
}
