package service

import (
	"encoding/xml"
	"errors"
	"github.com/owen-gxz/wechat-func"
	"net/http"
	"strings"
	"time"
)

// Context 消息上下文
type Context struct {
	*Server
	Timestamp string
	Nonce     string
	Msg       *WxMsg
	Resp      interface{}
	Writer    http.ResponseWriter
	Request   *http.Request
	hasReply  bool
}

// Reply 被动回复消息
func (c *Context) Reply() (err error) {
	if c.hasReply {
		return errors.New("重复调用错误")
	}

	c.hasReply = true

	if c.Resp == nil {
		return nil
	}

	Printf("Wechat <== %+v", c.Resp)
	if c.SafeMode {
		b, err := xml.MarshalIndent(c.Resp, "", "  ")
		if err != nil {
			return err
		}
		c.Resp, err = c.EncryptMsg(b, c.Timestamp, c.Nonce)
		if err != nil {
			return err
		}
	}
	c.Writer.Header().Set("Content-Type", "application/xml;charset=UTF-8")
	return xml.NewEncoder(c.Writer).Encode(c.Resp)
}

// Send 主动发送消息(客服)
func (c *Context) Send() *Context {
	c.AddMsg(c.Resp)
	return c
}

func (c *Context) newResp(msgType string) wechat.WxResp {
	return wechat.WxResp{
		FromUserName: wechat.CDATA(c.Msg.ToUserName),
		ToUserName:   wechat.CDATA(c.Msg.FromUserName),
		MsgType:      wechat.CDATA(msgType),
		CreateTime:   time.Now().Unix(),
		AgentId:      c.Msg.AgentID,
		//Safe:         c.Safe,
	}
}

// NewText Text消息
func (c *Context) NewText(text ...string) *Context {
	c.Resp = &wechat.Text{
		WxResp:  c.newResp(wechat.TypeText),
		Content: wechat.Content{wechat.CDATA(strings.Join(text, ""))}}
	return c
}

// NewImage Image消息
func (c *Context) NewImage(mediaId string) *Context {
	c.Resp = &wechat.ImageInfo{
		WxResp: c.newResp(wechat.TypeImage),
		Image:  wechat.Media{wechat.CDATA(mediaId)}}
	return c
}

// NewVoice Voice消息
func (c *Context) NewVoice(mediaId string) *Context {
	c.Resp = &wechat.Voice{
		WxResp: c.newResp(wechat.TypeVoice),
		Voice:  wechat.Media{wechat.CDATA(mediaId)}}
	return c
}

// NewFile File消息
func (c *Context) NewFile(mediaId string) *Context {
	c.Resp = &wechat.File{
		WxResp: c.newResp(wechat.TypeFile),
		File:   wechat.Media{wechat.CDATA(mediaId)}}
	return c
}

// NewVideo Video消息
func (c *Context) NewVideo(mediaId, title, desc string) *Context {
	c.Resp = &wechat.Video{
		WxResp: c.newResp(wechat.TypeVideo),
		Video:  wechat.VideoInfo{wechat.CDATA(mediaId), wechat.CDATA(title), wechat.CDATA(desc)}}
	return c
}

// NewTextcard Textcard消息
func (c *Context) NewTextcard(title, description, url string) *Context {
	c.Resp = &wechat.Textcard{
		WxResp:   c.newResp(wechat.TypeTextcard),
		Textcard: wechat.TextcardInfo{wechat.CDATA(title), wechat.CDATA(description), wechat.CDATA(url)}}
	return c
}

// NewNews News消息
func (c *Context) NewNews(arts ...wechat.Article) *Context {
	news := wechat.News{
		WxResp:       c.newResp(wechat.TypeNews),
		ArticleCount: len(arts),
	}
	news.Articles.Item = arts
	c.Resp = &news
	return c
}

// NewMpNews News消息
func (c *Context) NewMpNews(mediaId string) *Context {
	news := wechat.MpNewsId{
		WxResp: c.newResp(wechat.TypeMpNews),
	}
	news.MpNews.MediaId = wechat.CDATA(mediaId)
	c.Resp = &news
	return c
}

// NewMusic Music消息
func (c *Context) NewMusic(mediaId, title, desc, musicUrl, hqMusicUrl string) *Context {
	c.Resp = &wechat.Music{
		WxResp: c.newResp(wechat.TypeMusic),
		Music: wechat.MusicInfo{wechat.CDATA(mediaId), wechat.CDATA(title),
			wechat.CDATA(desc), wechat.CDATA(musicUrl), wechat.CDATA(hqMusicUrl)}}
	return c
}

// Id 返回消息的来源与去向，可作为多应用管理时的用户组Id
func (c *Context) Id() string {
	return c.Msg.FromUserName + "|" + c.Msg.ToUserName
}
