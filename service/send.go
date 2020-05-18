package service

import (
	"github.com/owen-gxz/wechat-func"
)

// AddMsg 添加队列消息
func (s *Server) AddMsg(v interface{}) {
	s.MsgQueue <- v
}

// SendMsg 发送消息
func (s *Server) SendMsg(v interface{}) *wechat.WxErr {
	return wechat.SendMsg(s.GetAccessToken(), v)
}

// SendText 发送客服text消息,过长时按500长度自动拆分
func (s *Server) SendText(to, msg string) (e *wechat.WxErr) {
	return wechat.SendText(s.GetAccessToken(), to, msg)
}

// SendImage 发送客服Image消息
func (s *Server) SendImage(to string, mediaId string) *wechat.WxErr {
	return s.SendMsg(wechat.NewImage(to, mediaId))
}

// SendVoice 发送客服Voice消息
func (s *Server) SendVoice(to string, mediaId string) *wechat.WxErr {
	return s.SendMsg(wechat.NewVoice(to, mediaId))
}

// SendFile 发送客服File消息
func (s *Server) SendFile(to string, mediaId string) *wechat.WxErr {
	return s.SendMsg(wechat.NewFile(to, mediaId))
}

// SendVideo 发送客服Video消息
func (s *Server) SendVideo(to string, mediaId, title, desc string) *wechat.WxErr {
	return s.SendMsg(wechat.NewVideo(to, mediaId, title, desc))
}

// SendTextcard 发送客服extcard消息
func (s *Server) SendTextcard(to string, title, desc, url string) *wechat.WxErr {
	return s.SendMsg(wechat.NewTextcard(to, title, desc, url))
}

// SendMusic 发送客服Music消息
func (s *Server) SendMusic(to string, mediaId, title, desc, musicUrl, qhMusicUrl string) *wechat.WxErr {
	return s.SendMsg(wechat.NewMusic(to, mediaId, title, desc, musicUrl, qhMusicUrl))
}

// SendNews 发送客服news消息
func (s *Server) SendNews(to string, arts ...wechat.Article) *wechat.WxErr {
	return s.SendMsg(wechat.NewNews(to, arts...))
}

// SendMpNews 发送加密新闻mpnews消息(仅企业号可用)
func (s *Server) SendMpNews(to string, arts ...wechat.MpArticle) *wechat.WxErr {
	return s.SendMsg(wechat.NewMpNews(to, arts...))
}

// SendMpNewsId 发送加密新闻mpnews消息(直接使用mediaId)
func (s *Server) SendMpNewsId(to string, mediaId string) *wechat.WxErr {
	return s.SendMsg(wechat.NewMpNewsId(to, mediaId))
}

// SendMarkDown 发送加密新闻mpnews消息(直接使用mediaId)
func (s *Server) SendMarkDown(to string, content string) *wechat.WxErr {
	return s.SendMsg(wechat.NewMarkDown(to, content))
}

// SendTaskCard 发送任务卡片taskcard消息
func (s *Server) SendTaskCard(to string, Title, Desc, Url, TaskId, Btn string) *wechat.WxErr {
	return s.SendMsg(wechat.NewTaskCard(to, Title, Desc, Url, TaskId, Btn))
}
