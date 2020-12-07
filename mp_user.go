package wechat

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/owen-gxz/wechat-func/util"
)

// MPUserGetList 公众号用户接口
var (
	MPUserGetList  = WXAPI + "user/get?access_token=%s&next_openid=%s"
	MPUserBatchGet = WXAPI + "user/info/batchget?access_token="
	MPUserInfo     = WXAPI + "user/info?access_token=%s&openid=%v&lang=%v"

	SnsUserInfo = " https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
)

type (
	// MpUserInfoList 公众号用户信息列表
	MpUserInfoList struct {
		WxErr
		MpUserInfoList []MpUserInfo `json:"user_info_list"`
	}

	// MpUserInfo 公众号用户信息
	MpUserInfo struct {
		Subscribe     int
		OpenId        string
		NickName      string
		Sex           int
		Language      string
		City          string
		Province      string
		Country       string
		HeadImgUrl    string
		SubscribeTime int `json:"subscribe_time"`
		UnionId       string
		Remark        string
		GroupId       int
		TagIdList     []int `json:"tagid_list"`
	}

	// MpUser 服务号用户
	MpUser struct {
		WxErr
		Total int
		Count int
		Data  struct {
			OpenId []string
		}
		NextOpenId string
	}

	// MpUserListReq 公众号用户请求
	MpUserListReq struct {
		UserList interface{} `json:"user_list"`
	}
)

// BatchGetAll 获取所有公众号用户
func BatchGetAll(token string) (ui []MpUserInfo, err error) {
	var ul []string
	ul, err = GetAllMpUserList(token)
	if err != nil {
		return
	}
	leng := len(ul)
	if leng <= 100 {
		return BatchGet(token, ul)
	}
	for i := 0; i < leng/100+1; i++ {
		end := (i + 1) * 100
		if end > leng {
			end = leng
		}
		ui2 := make([]MpUserInfo, 0)
		ui2, err = BatchGet(token, ul[i*100:end])
		if err != nil {
			return
		}
		ui = append(ui, ui2...)
	}
	return
}

// BatchGet 批量获取公众号用户信息
func BatchGet(token string, ul []string) (ui []MpUserInfo, err error) {
	m := make([]H, len(ul))

	for k, v := range ul {
		m[k] = H{}
		m[k]["openid"] = v
	}
	ml := new(MpUserInfoList)
	err = util.PostJsonPtr(MPUserBatchGet+token, MpUserListReq{m}, ml)
	return ml.MpUserInfoList, ml.Error()
}

// GetAllMpUserList 获取所有用户ID
func GetAllMpUserList(token string) (ul []string, err error) {
	ul = make([]string, 0)
	mul, err := GetMpUserList(token)
	if err != nil {
		return
	}
	if mul.Error() == nil {
		ul = append(ul, mul.Data.OpenId...)
	}
	for mul.Count == 10000 {
		mul, err = GetMpUserList(token, mul.NextOpenId)
		if err != nil {
			return
		}
		if mul.Error() == nil {
			ul = append(ul, mul.Data.OpenId...)
		}
	}
	return
}

// GetMpUserList 获取用户信息，根据openid
func GetMpUserList(token string, openid ...string) (ul *MpUser, err error) {
	if len(openid) == 0 {
		openid = append(openid, "")
	}
	mpuser := new(MpUser)
	url := fmt.Sprintf(MPUserGetList, token, openid[0])
	if err = util.GetJson(url, &mpuser); err != nil {
		return
	}
	return mpuser, mpuser.Error()
}

// GetMpUserInfo 获取用户详情
func GetMpUserInfo(token, openid string, lang ...string) (user *MpUserInfo, err error) {
	if len(lang) == 0 {
		lang = append(lang, "zh_CN")
	}
	user = new(MpUserInfo)
	url := fmt.Sprintf(MPUserInfo, token, openid, lang[0])
	if err = util.GetJson(url, &user); err != nil {
		return
	}
	return
}

// GetMpUserInfo 获取用户详情
func GetMpUserInfoBySns(token, openid string, lang ...string) (user *MpUserInfo, err error) {
	if len(lang) == 0 {
		lang = append(lang, "zh_CN")
	}
	user = new(MpUserInfo)
	url := fmt.Sprintf(MPUserInfo, token, openid, lang[0])
	if err = util.GetJson(url, &user); err != nil {
		return
	}
	return
}

// 获取微信小程序userInfo
func GetUserInfo(data, key, iv string) (*Userinfo, error) {
	k, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	ivDecode, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	dataDecode, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	ud, err := util.AesDecrypt2(dataDecode, k, ivDecode)
	if err != nil {
		return nil, err
	}
	u := Userinfo{}
	err = json.Unmarshal(ud, &u)
	if err != nil {
		return nil, err

	}
	return &u, nil
}

