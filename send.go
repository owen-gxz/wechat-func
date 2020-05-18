package wechat

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	"github.com/owen-gxz/wechat-func/util"
)

// SendMsg 发送消息
func SendMsg(token string, v interface{}) *WxErr {
	url := WXAPIMsg + token
	body, err := util.PostJson(url, v)
	if err != nil {
		return &WxErr{-1, err.Error()}
	}
	rst := new(WxErr)
	err = json.Unmarshal(body, rst)
	if err != nil {
		return &WxErr{-1, err.Error()}
	}
	return rst
}

// SendText 发送客服text消息,过长时按500长度自动拆分
func SendText(token, to, msg string) (e *WxErr) {
	leng := utf8.RuneCountInString(msg)
	n := leng/500 + 1

	if n == 1 {
		return SendMsg(token, NewText(to, msg))
	}
	for i := 0; i < n; i++ {
		msg := fmt.Sprintf("%s\n(%v/%v)", util.Substr(msg, i*500, (i+1)*500), i+1, n)
		e = SendMsg(token, NewText(to, msg))
	}

	return
}

// SendImage 发送客服Image消息
func SendImage(token, to string, mediaId string) *WxErr {
	return SendMsg(token, NewImage(to, mediaId))
}

// SendVoice 发送客服Voice消息
func SendVoice(token, to string, mediaId string) *WxErr {
	return SendMsg(token, NewVoice(to, mediaId))
}

// SendFile 发送客服File消息
func SendFile(token, to string, mediaId string) *WxErr {
	return SendMsg(token, NewFile(to, mediaId))
}

// SendVideo 发送客服Video消息
func SendVideo(token, to string, mediaId, title, desc string) *WxErr {
	return SendMsg(token, NewVideo(to, mediaId, title, desc))
}

// SendTextcard 发送客服extcard消息
func SendTextcard(token, to string, title, desc, url string) *WxErr {
	return SendMsg(token, NewTextcard(to, title, desc, url))
}

// SendMusic 发送客服Music消息
func SendMusic(token, to string, mediaId, title, desc, musicUrl, qhMusicUrl string) *WxErr {
	return SendMsg(token, NewMusic(to, mediaId, title, desc, musicUrl, qhMusicUrl))
}

// SendNews 发送客服news消息
func SendNews(token, to string, arts ...Article) *WxErr {
	return SendMsg(token, NewNews(to, arts...))
}

// SendMpNews 发送加密新闻mpnews消息(仅企业号可用)
func SendMpNews(token, to string, arts ...MpArticle) *WxErr {
	return SendMsg(token, NewMpNews(to, arts...))
}

// SendMpNewsId 发送加密新闻mpnews消息(直接使用mediaId)
func SendMpNewsId(token, to string, mediaId string) *WxErr {
	return SendMsg(token, NewMpNewsId(to, mediaId))
}

// SendMarkDown 发送加密新闻mpnews消息(直接使用mediaId)
func SendMarkDown(token, to string, content string) *WxErr {
	return SendMsg(token, NewMarkDown(to, content))
}

// SendTaskCard 发送任务卡片taskcard消息
func SendTaskCard(token, to string, Title, Desc, Url, TaskId, Btn string) *WxErr {
	return SendMsg(token, NewTaskCard(to, Title, Desc, Url, TaskId, Btn))
}
