package service

import (
	"github.com/owen-gxz/wechat-func"
)

// MediaUpload 临时素材上传，mediaType选项如下：
//	TypeImage  = "image"
//	TypeVoice  = "voice"
//	TypeVideo  = "video"
//	TypeFile   = "file" // 仅企业微信可用
func (s *Server) MediaUpload(mediaType string, filename string) (media wechat.MediaResponse, err error) {
	return wechat.MediaUpload(s.GetAccessToken(), mediaType, filename)
}

func (s *Server) MediaUploadBytes(mediaType string, file []byte) (media wechat.MediaResponse, err error) {
	return wechat.MediaUploadBytes(s.GetAccessToken(), mediaType, file)
}

// GetMedia 下载临时素材
func (s *Server) GetMedia(filename, mediaId string) error {
	return wechat.GetMedia(s.GetAccessToken(), filename, mediaId)

}

// GetMediaBytes 下载临时素材,返回body字节
func (s *Server) GetMediaBytes(mediaId string) ([]byte, error) {
	return wechat.GetMediaBytes(s.GetAccessToken(), mediaId)
}

// GetJsMedia 下载高清语言素材(通过JSSDK上传)
func (s *Server) GetJsMedia(filename, mediaId string) error {
	return wechat.GetJsMedia(s.GetAccessToken(), filename, mediaId)
}

// GetJsMediaBytes 下载高清语言素材,返回body字节
func (s *Server) GetJsMediaBytes(mediaId string) ([]byte, error) {
	return wechat.GetJsMediaBytes(s.GetAccessToken(), mediaId)
}

// 新增永久图文接口
func (s *Server) MediaNews(articles ...wechat.NewsArticle) (media wechat.MediaResponse, err error) {
	return wechat.MediaNews(s.GetAccessToken(), articles...)
}

// 上传图文消息内的图片获取URL
/*
本接口所上传的图片不占用公众号的素材库中图片数量的100000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下。
*/
func (s *Server) MediaNewsImage(fileName string) (media wechat.MediaResponse, err error) {
	return wechat.MediaNewsImage(s.GetAccessToken(), fileName)
}
func (s *Server) MediaNewsImageBytes(file []byte) (media wechat.MediaResponse, err error) {
	return wechat.MediaNewsImageBytes(s.GetAccessToken(), file)
}

// 新增其他类型永久素材
func (s *Server) AddMedia(mediaType string, fileName string) (media wechat.MediaResponse, err error) {
	return wechat.AddMedia(s.GetAccessToken(), mediaType, fileName)
}
func (s *Server) AddMediaBytes(mediaType string, file []byte) (media wechat.MediaResponse, err error) {
	return wechat.AddMediaBytes(s.GetAccessToken(), mediaType, file)
}

// 获取永久素材
func (s *Server) GetForeverMedia(mediaID string) (media wechat.MediaResponse, err error) {
	return wechat.GetForeverMedia(s.GetAccessToken(), mediaID)
}
//删除永久素材
func (s *Server) RemoveForeverMedia(mediaID string) (media wechat.MediaResponse, err error) {
	return wechat.RemoveForeverMedia(s.GetAccessToken(), mediaID)
}

//获取永久素材总数
func (s *Server)MediaCount() (media wechat.MediaCountResponse, err error) {
	return wechat.MediaCount(s.GetAccessToken())

}

////获取永久素材列表
func (s *Server)BatchgetMedia(mediaType string, offset, count int64) (media wechat.BatchgetMediaResponse, err error) {
	return wechat.BatchgetMedia(s.GetAccessToken(),mediaType,offset,count)

}
