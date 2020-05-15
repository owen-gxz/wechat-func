// 目前官方未提供golang版，本SDK实现参考了php版官方库
// @woylin, since 2016-1-6

package wechat

// WXAPI 订阅号，服务号，小程序接口，相关接口常量统一以此开头
const (
	WXAPI    = "https://api.weixin.qq.com/cgi-bin/"
	WXAPIMsg = WXAPI + "message/custom/send?access_token="
)

var (
	// Debug is a flag to Println()
	Debug bool = false
)
