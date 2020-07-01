package wechat

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

// Type io类型汇总
const (
	TypeText     = "text"
	TypeImage    = "image"
	TypeVoice    = "voice"
	TypeMusic    = "MusicInfo"
	TypeVideo    = "VideoInfo"
	TypeTextcard = "TextcardInfo" // 仅企业微信可用
	TypeWxCard   = "wxcard"       // 仅服务号可用
	TypeMarkDown = "markdown"     // 仅企业微信可用
	TypeTaskCard = "taskcard"     // 仅企业微信可用
	TypeFile     = "file"         // 仅企业微信可用
	TypeNews     = "news"
	TypeMpNews   = "mpnews" // 仅企业微信可用
)

//WxErr 通用错误
type WxErr struct {
	ErrCode int
	ErrMsg  string
}

func (w *WxErr) Error() error {
	if w.ErrCode != 0 {
		return fmt.Errorf("err: errcode=%v , errmsg=%v", w.ErrCode, w.ErrMsg)
	}
	return nil
}

// CDATA 标准规范，XML编码成 `<![CDATA[消息内容]]>`
type CDATA string

// MarshalXML 自定义xml编码接口，实现讨论: http://stackoverflow.com/q/41951345/7493327
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// WxResp 响应消息共用字段
// 响应消息被动回复为XML结构，文本类型采用CDATA编码规范
// 响应消息主动发送为json结构，即客服消息
type WxResp struct {
	XMLName      xml.Name `xml:"xml" json:"-"`
	ToUserName   CDATA    `json:"touser"`
	ToParty      CDATA    `xml:"-" json:"toparty"` // 企业号专用
	ToTag        CDATA    `xml:"-" json:"totag"`   // 企业号专用
	FromUserName CDATA    `json:"-"`
	CreateTime   int64    `json:"-"`
	MsgType      CDATA    `json:"msgtype"`
	AgentId      int      `xml:"-" json:"agentid"`
	Safe         int      `xml:"-" json:"safe"`
}

// to字段格式："userid1|userid2 deptid1|deptid2 tagid1|tagid2"
func newWxResp(msgType, to string) (r WxResp) {
	toArr := strings.Split(to, " ")
	r = WxResp{
		ToUserName: CDATA(toArr[0]),
		MsgType:    CDATA(msgType),
	}
	if len(toArr) > 1 {
		r.ToParty = CDATA(toArr[1])
	}
	if len(toArr) > 2 {
		r.ToTag = CDATA(toArr[2])
	}
	return
}

// Text 文本消息
type (
	Text struct {
		WxResp
		Content `xml:"Content" json:"text"`
	}

	Content struct {
		Content CDATA `json:"Content"`
	}
)

// NewText Text 文本消息
func NewText(to string, msg ...string) Text {
	return Text{
		newWxResp(TypeText, to),
		Content{CDATA(strings.Join(msg, ""))},
	}
}

// ImageInfo 图片消息
type (
	ImageInfo struct {
		WxResp
		Image Media `json:"image"`
	}

	Media struct {
		MediaId CDATA `json:"media_id"`
	}
)

// NewImage ImageInfo 消息
func NewImage(to, mediaId string) ImageInfo {
	return ImageInfo{
		newWxResp(TypeImage, to),
		Media{CDATA(mediaId)},
	}
}

// Voice 语音消息
type Voice struct {
	WxResp
	Voice Media `json:"voice"`
}

// NewVoice Voice消息
func NewVoice(to, mediaId string) Voice {
	return Voice{
		newWxResp(TypeVoice, to),
		Media{CDATA(mediaId)},
	}
}

// File 文件消息，仅企业号支持
type File struct {
	WxResp
	File Media `json:"file"`
}

// NewFile File消息
func NewFile(to, mediaId string) File {
	return File{
		newWxResp(TypeFile, to),
		Media{CDATA(mediaId)},
	}
}

// Video 视频消息
type (
	Video struct {
		WxResp
		Video VideoInfo `json:"VideoInfo"`
	}

	VideoInfo struct {
		MediaId     CDATA `json:"media_id"`
		Title       CDATA `json:"title"`
		Description CDATA `json:"description"`
	}
)

// NewVideo Video消息
func NewVideo(to, mediaId, title, desc string) Video {
	return Video{
		newWxResp(TypeVideo, to),
		VideoInfo{CDATA(mediaId), CDATA(title), CDATA(desc)},
	}
}

// Textcard 卡片消息，仅企业微信客户端有效
type (
	Textcard struct {
		WxResp
		Textcard TextcardInfo `json:"TextcardInfo"`
	}

	TextcardInfo struct {
		Title       CDATA `json:"title"`
		Description CDATA `json:"description"`
		Url         CDATA `json:"url"`
	}
)

// NewTextcard Textcard消息
func NewTextcard(to, title, description, url string) Textcard {
	return Textcard{
		newWxResp(TypeTextcard, to),
		TextcardInfo{CDATA(title), CDATA(description), CDATA(url)},
	}
}

// Music 音乐消息，企业微信不支持
type (
	Music struct {
		WxResp
		Music MusicInfo `json:"MusicInfo"`
	}

	MusicInfo struct {
		Title        CDATA `json:"title"`
		Description  CDATA `json:"description"`
		MusicUrl     CDATA `json:"musicurl"`
		HQMusicUrl   CDATA `json:"hqmusicurl"`
		ThumbMediaId CDATA `json:"thumb_media_id"`
	}
)

// NewMusic Music消息
func NewMusic(to, mediaId, title, desc, musicUrl, qhMusicUrl string) Music {
	return Music{
		newWxResp(TypeMusic, to),
		MusicInfo{CDATA(title), CDATA(desc), CDATA(musicUrl), CDATA(qhMusicUrl), CDATA(mediaId)},
	}
}

// News 新闻消息
type News struct {
	WxResp
	ArticleCount int
	Articles     struct {
		Item []Article `xml:"item" json:"articles"`
	} `json:"news"`
}

// NewNews news消息
func NewNews(to string, arts ...Article) (news News) {
	news.WxResp = newWxResp(TypeNews, to)
	news.ArticleCount = len(arts)
	news.Articles.Item = arts
	return
}

// Article 文章
type Article struct {
	Title       CDATA `json:"title"`
	Description CDATA `json:"description"`
	PicUrl      CDATA `json:"picurl"`
	Url         CDATA `json:"url"`
}

// NewArticle 先创建文章，再传给NewNews()
func NewArticle(title, desc, picUrl, url string) Article {
	return Article{CDATA(title), CDATA(desc), CDATA(picUrl), CDATA(url)}
}

type (
	// MpNews 加密新闻消息，仅企业微信支持
	MpNews struct {
		WxResp
		MpNews struct {
			Articles []MpArticle `json:"articles"`
		} `json:"mpnews"`
	}

	// MpNewsId 加密新闻消息(通过mediaId直接发)
	MpNewsId struct {
		WxResp
		MpNews struct {
			MediaId CDATA `json:"media_id"`
		} `json:"mpnews"`
	}
)

// NewMpNews 加密新闻mpnews消息(仅企业微信可用)
func NewMpNews(to string, arts ...MpArticle) (news MpNews) {
	news.WxResp = newWxResp(TypeMpNews, to)
	news.MpNews.Articles = arts
	return
}

// NewMpNewsId 加密新闻mpnews消息(仅企业微信可用)
func NewMpNewsId(to string, mediaId string) (news MpNewsId) {
	news.WxResp = newWxResp(TypeMpNews, to)
	news.MpNews.MediaId = CDATA(mediaId)
	return
}

// MpArticle 加密文章
type MpArticle struct {
	Title        string `json:"title"`
	ThumbMediaId string `json:"thumb_media_id"`
	Author       string `json:"author"`
	Url          string `json:"content_source_url"`
	Content      string `json:"Content"`
	Digest       string `json:"digest"`
}

// NewMpArticle 先创建加密文章，再传给NewMpNews()
func NewMpArticle(title, mediaId, author, url, content, digest string) MpArticle {
	return MpArticle{title, mediaId, author, url, content, digest}
}

// WxCard 卡券
type WxCard struct {
	WxResp
	WxCard struct {
		CardId string `json:"card_id"`
	} `json:"wxcard"`
}

// NewWxCard 卡券消息，服务号可用
func NewWxCard(to, cardId string) (c WxCard) {
	c.WxResp = newWxResp(TypeWxCard, to)
	c.WxCard.CardId = cardId
	return
}

// MarkDown markdown消息，仅企业微信支持，上限2048字节，utf-8编码
type MarkDown struct {
	WxResp
	MarkDown struct {
		Content string `json:"Content"`
	} `json:"markdown"`
}

// NewMarkDown markdown消息，企业微信可用
func NewMarkDown(to, content string) (md MarkDown) {
	md.WxResp = newWxResp(TypeMarkDown, to)
	md.MarkDown.Content = content
	return
}

// TaskCard 任务卡片消息，仅企业微信支持，支持一到两个按钮设置
type TaskCard struct {
	WxResp
	TaskCard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		TaskId      string `json:"task_id"`
		Btn         []H    `json:"btn"`
	} `json:"taskcard"`
}

// NewTaskCard 任务卡片消息，企业微信可用
func NewTaskCard(to, Title, Desc, Url, TaskId, Btn string) (tc TaskCard) {
	tc.WxResp = newWxResp(TypeTaskCard, to)
	tc.TaskCard.Title = Title
	tc.TaskCard.Description = Desc
	tc.TaskCard.Url = Url
	tc.TaskCard.TaskId = TaskId
	mp := make([]H, 0)
	if Btn != "" {
		if err := json.Unmarshal([]byte(Btn), &mp); err != nil {
			fmt.Println("create taskcard btn err:", err)
		} else {
			tc.TaskCard.Btn = mp
		}
	}
	return
}

type Userinfo struct {
	OpenID    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	UnionID   string `json:"unionId"`
	Watermark struct {
		Appid     string `json:"appid"`
		Timestamp int32  `json:"timestamp"`
	} `json:"watermark"`
}
