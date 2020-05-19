package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/owen-gxz/wechat-func/util"
)

var (
	// WXAPIMediaUpload 临时素材上传
	WXAPIMediaUpload = WXAPI + "media/upload?access_token=%s&type=%s"
	// WXAPIMediaGet 临时素材下载
	WXAPIMediaGet = WXAPI + "media/get?access_token=%s&media_id=%s"
	// WXAPIMediaGetJssdk 高清语言素材下载
	WXAPIMediaGetJssdk = WXAPI + "media/get/jssdk?access_token=%s&media_id=%s"

	// 新增永久图文消息
	WXAPIMediaNews = WXAPI + "material/add_news?access_token=%s"
	// 上传图文消息内的图片获取URL
	WXAPIMediaNewsImage = WXAPI + "media/uploadimg?access_token=%s"
	// 上传其他永久素材
	WXAPIAddMedia = WXAPI + "material/add_material?access_token=%s&type=%s"
	// 获取永久素材
	WXAPIGetMedia = WXAPI + "material/get_material?access_token=%s"
	// 删除永久素材
	WXAPIDelMedia = WXAPI + "material/del_material?access_token=%s"
	// 永久素材总数
	WXAPIMediaCount = WXAPI + "material/get_materialcount?access_token=%s"
	// 永久素材列表
	WXAPIBatchgetMaterial = WXAPI + "material/batchget_material?access_token=%s"
)

// MediaResponse 上传回复体
type MediaResponse struct {
	WxErr
	MediaID string `json:"media_id"`
	//	临时素材返回
	Type         string `json:"type"`
	ThumbMediaId string `json:"thumb_media_id"`
	CreatedAt    int64  `json:"created_at"` // 服务号是int,
	// 永久素材返回
	Url string `json:"url"` // 新增永久素材返回
	// 图文素材
	NewsItem []*NewsArticle `json:"news_item"`
	// 视频消息素材
	Title       string `json:"title"`
	Description string `json:"description"`
	DownUrl     string `json:"down_url"`
}

// MediaUpload 临时素材上传，mediaType选项如下：
//	TypeImage  = "image"
//	TypeVoice  = "voice"
//	TypeVideo  = "VideoInfo"
func MediaUpload(token, mediaType string, filename string) (media MediaResponse, err error) {
	uri := fmt.Sprintf(WXAPIMediaUpload, token, mediaType)
	return mediaUpload(uri, filename)
}

func mediaUpload(uri string, filename string) (media MediaResponse, err error) {
	b, err := util.PostFile("Media", filename, uri)
	if err != nil {
		return
	}
	if err = json.Unmarshal(b, &media); err != nil {
		return
	}
	err = media.Error()
	return
}

// MediaUpload 临时素材上传，mediaType选项如下：
//	TypeImage  = "image"
//	TypeVoice  = "voice"
//	TypeVideo  = "VideoInfo"
//	TypeThumb  = "thumb"
func MediaUploadBytes(token, mediaType string, file []byte) (media MediaResponse, err error) {
	uri := fmt.Sprintf(WXAPIMediaUpload, token, mediaType)
	return mediaUploadBytes(uri, file)
}

func mediaUploadBytes(uri string, file []byte) (media MediaResponse, err error) {
	var b []byte
	b, err = util.PostFile2Byte("Media", file, uri)
	if err != nil {
		return
	}
	if err = json.Unmarshal(b, &media); err != nil {
		return
	}
	err = media.Error()
	return
}

// GetMedia 下载临时素材
func GetMedia(token, filename, mediaId string) error {
	url := fmt.Sprintf(WXAPIMediaGet, token, mediaId)
	return util.GetFile(filename, url)
}

// GetMediaBytes 下载临时素材,返回body字节
func GetMediaBytes(token, mediaId string) ([]byte, error) {
	url := fmt.Sprintf(WXAPIMediaGet, token, mediaId)
	return util.GetBody(url)
}

// GetJsMedia 下载高清语言素材(通过JSSDK上传)
func GetJsMedia(token, filename, mediaId string) error {
	url := fmt.Sprintf(WXAPIMediaGetJssdk, token, mediaId)
	return util.GetFile(filename, url)
}

// GetJsMediaBytes 下载高清语言素材,返回body字节
func GetJsMediaBytes(token, mediaId string) ([]byte, error) {
	url := fmt.Sprintf(WXAPIMediaGetJssdk, token, mediaId)
	return util.GetBody(url)
}

//  永久素材接口
type NewsArticle struct {
	Title        string `json:"title"`
	ThumbMediaID string `json:"thumb_media_id"` //图文消息的封面图片素材id（必须是永久mediaID）
	Author       string `json:"author"`         // 作者
	//图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前64个字。
	Digest string `json:"digest"`
	//是否显示封面，0为false，即不显示，1为true，即显示
	ShowCoverPic int `json:"show_cover_pic"`
	// 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片url将被过滤。
	Content string `json:"Content"`
	// 图文消息的原文地址，即点击“阅读原文”后的URL
	ContentSourceUrl string `json:"content_source_url"`
	// Uint32 是否打开评论，0不打开，1打开
	NeedOpenComment int `json:"need_open_comment"`
	//Uint32 是否粉丝才可评论，0所有人可评论，1粉丝才可评论
	OnlyFansCanComment int `json:"only_fans_can_comment"`
}

// 新增永久图文接口
func MediaNews(token string, articles ...NewsArticle) (media MediaResponse, err error) {
	uri := fmt.Sprintf(WXAPIMediaNews, token)
	b, err := util.PostJson(uri, articles)
	if err != nil {
		return
	}
	if err = json.Unmarshal(b, &media); err != nil {
		return
	}
	err = media.Error()
	return
}

// 上传图文消息内的图片获取URL
/*
本接口所上传的图片不占用公众号的素材库中图片数量的100000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下。
*/
func MediaNewsImage(token string, fileName string) (media MediaResponse, err error) {
	uri := fmt.Sprintf(WXAPIMediaNewsImage, token)
	return mediaUpload(uri, fileName)
}
func MediaNewsImageBytes(token string, file []byte) (media MediaResponse, err error) {
	uri := fmt.Sprintf(WXAPIMediaNewsImage, token)
	return mediaUploadBytes(uri, file)
}

// 新增其他类型永久素材
func AddMedia(token, mediaType string, fileName string) (media MediaResponse, err error) {
	uri := fmt.Sprintf(WXAPIAddMedia, token, mediaType)
	return mediaUpload(uri, fileName)
}
func AddMediaBytes(token, mediaType string, file []byte) (media MediaResponse, err error) {
	uri := fmt.Sprintf(WXAPIAddMedia, token, mediaType)
	return mediaUploadBytes(uri, file)
}

// 获取永久素材
func GetForeverMedia(token, mediaID string) (media MediaResponse, err error) {
	uri := fmt.Sprintf(WXAPIGetMedia, token)
	form := H{"media_id": mediaID}

	b, err := util.PostJson(uri, form)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &media)
	return
}

// 删除永久素材
func RemoveForeverMedia(token, mediaID string) (media MediaResponse, err error) {
	uri := fmt.Sprintf(WXAPIDelMedia, token)
	form := H{"media_id": mediaID}

	b, err := util.PostJson(uri, form)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &media)
	return
}

type MediaCountResponse struct {
	WxErr
	VoiceCount int64 `json:"voice_count"`
	ImageCount int64 `json:"image_count"`
	VideoCount int64 `json:"video_count"`
	NewsCount  int64 `json:"news_count"`
}

//获取永久素材总数
func MediaCount(token string) (media MediaCountResponse, err error) {
	uri := fmt.Sprintf(WXAPIMediaCount, token)

	b, err := util.GetBody(uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &media)
	return
}

type BatchgetMediaResponse struct {
	WxErr
	TotalCount int64           `json:"total_count"`
	ItemCount  int64           `json:"item_count"`
	Item       []MediaResponse `json:"item"`
}

type BatchgetMediaResponseItem struct {
	MediaID string `json:"media_id"`
	// 图文
	Content BatchgetMediaResponseItemContent `json:"Content"`
	// 其他类型
	Name       string      `json:"name"`
	Url        string      `json:"url"`
	UpdateTime interface{} `json:"update_time"` //todo 未只类型

}
type BatchgetMediaResponseItemContent struct {
	NewsItem []NewsArticle `json:"news_item"`
}

//获取永久素材列表
func BatchgetMedia(token, mediaType string, offset, count int64) (media BatchgetMediaResponse, err error) {
	uri := fmt.Sprintf(WXAPIBatchgetMaterial, token)
	req := H{
		"type":   mediaType,
		"offset": offset,
		"count":  count,
	}

	b, err := util.PostJson(uri, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &media)
	return
}
