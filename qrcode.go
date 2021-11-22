package wechat

import (
	"fmt"
	"github.com/owen-gxz/wechat-func/util"
)

var (
	GetUnlimitedQrcodeUrl = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s"
)

func MiniprogramGetUnlimitedQrcode(token, scene string) ([]byte, error) {
	uri := fmt.Sprintf(GetUnlimitedQrcodeUrl, token)
	form := H{"scene": scene}
	b, err := util.PostJson(uri, form)
	if err != nil {
		return nil, err
	}
	//err = json.Unmarshal(b, &m)
	return b, err
}
