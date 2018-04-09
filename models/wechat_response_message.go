package models

import (
	"encoding/xml"
	"fmt"
	"time"
)

/*
回复文本消息

<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>12345678</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[你好]]></Content>
</xml>

回复图片消息

<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>12345678</CreateTime>
<MsgType><![CDATA[image]]></MsgType>
<Image>
	<MediaId><![CDATA[media_id]]></MediaId>
</Image>
</xml>

回复语音消息

<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>12345678</CreateTime>
<MsgType><![CDATA[voice]]></MsgType>
<Voice>
	<MediaId><![CDATA[media_id]]></MediaId>
</Voice>
</xml>

回复视频消息

<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>12345678</CreateTime>
<MsgType><![CDATA[video]]></MsgType>
<Video>
	<MediaId><![CDATA[media_id]]></MediaId>
	<Title><![CDATA[title]]></Title>
	<Description><![CDATA[description]]></Description>
</Video>
</xml>

回复音乐消息

<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>12345678</CreateTime>
<MsgType><![CDATA[music]]></MsgType>
<Music>
	<Title><![CDATA[TITLE]]></Title>
	<Description><![CDATA[DESCRIPTION]]></Description>
	<MusicUrl><![CDATA[MUSIC_Url]]></MusicUrl>
	<HQMusicUrl><![CDATA[HQ_MUSIC_Url]]></HQMusicUrl>
	<ThumbMediaId><![CDATA[media_id]]></ThumbMediaId>
</Music>
</xml>

回复图文消息

<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>12345678</CreateTime>
<MsgType><![CDATA[news]]></MsgType>
<ArticleCount>2</ArticleCount>
<Articles>
	<item>
		<Title><![CDATA[title1]]></Title>
		<Description><![CDATA[description1]]></Description>
		<PicUrl><![CDATA[picurl]]></PicUrl>
		<Url><![CDATA[url]]></Url>
	</item>
	<item>
		<Title><![CDATA[title]]></Title>
		<Description><![CDATA[description]]></Description>
		<PicUrl><![CDATA[picurl]]></PicUrl>
		<Url><![CDATA[url]]></Url>
	</item>
</Articles>
</xml>
*/

//ReplayMessage ReplayMessage
type ReplayMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int64
	MsgType      CDATA
}

//ReplayMessageText ReplayMessageText
type ReplayMessageText struct {
	ReplayMessage
	Content CDATA
}

//ReplayMessageImage ReplayMessageImage
type ReplayMessageImage struct {
	ReplayMessage
	Image struct {
		MediaId CDATA
	}
}

//ReplayMessageVoice ReplayMessageVoice
type ReplayMessageVoice struct {
	ReplayMessage
	Voice struct {
		MediaId CDATA
	}
}

//ReplayMessageVideo ReplayMessageVideo
type ReplayMessageVideo struct {
	ReplayMessage
	Video struct {
		MediaId     CDATA
		Title       CDATA
		Description CDATA
	}
}

//ReplayMessageMusic ReplayMessageMusic
type ReplayMessageMusic struct {
	ReplayMessage
	Music struct {
		Title        CDATA
		Description  CDATA
		MusicUrl     CDATA
		HQMusicUrl   CDATA
		ThumbMediaId CDATA
	}
}

//ReplayMessageArticles ReplayMessageArticles
type ReplayMessageArticles struct {
	ReplayMessage
	ArticleCount int
	Articles     struct {
		Articles []ArticleItem
	}
}

//ArticleItem ArticleItem
type ArticleItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       CDATA
	Description CDATA
	PicUrl      CDATA
	Url         CDATA
}

// GetReplayMessage GetReplayMessage
func GetReplayMessage() (rm *ReplayMessageArticles, err error) {
	rm = &ReplayMessageArticles{}
	rm.CreateTime = 123123
	rm.FromUserName.Text = "test"

	a := ArticleItem{}
	a.Description.Text = "dasdsadsd"
	a.PicUrl.Text = "saasdasdasd"
	a.Title.Text = "asdasdasdasd"
	a.Url.Text = "dasdasdsad"
	rm.ArticleCount = 1
	rm.Articles.Articles = append(rm.Articles.Articles, a)
	output, err := xml.MarshalIndent(rm, "  ", "    ")

	fmt.Println(string(output))
	return
}

// GetReplayMessageUserInput GetReplayMessageUserInput
func GetReplayMessageUserInput(fromUserName, toUserName, key string) (returnData interface{}, err error) {
	var v *ReplyKey
	v, err = GetReplyKeyByKey(key)
	if err != nil {
		var r *ReplyDefult
		r, err = GetReplyDefultByType("receiveMessage")
		if err != nil {
			replayMessageText := ReplayMessageText{}
			replayMessageText.ToUserName.Text = fromUserName
			replayMessageText.FromUserName.Text = toUserName
			replayMessageText.CreateTime = time.Now().Unix()
			replayMessageText.MsgType.Text = "text"
			replayMessageText.Content.Text = "佑恩堂欢迎你！"
			returnData = replayMessageText
		} else {
			returnData, err = getReplyByContentTypeAndContent(fromUserName, toUserName, r.ContentType, r.Content)
		}
	} else {
		returnData, err = getReplyByContentTypeAndContent(fromUserName, toUserName, v.ContentType, v.Content)
	}
	return
}

// GetReplayMessageReceiveMessage GetReplayMessageReceiveMessage
func GetReplayMessageReceiveMessage(fromUserName, toUserName string) (returnData interface{}, err error) {

	var r *ReplyDefult
	r, err = GetReplyDefultByType("receiveMessage")
	if err != nil {
		replayMessageText := ReplayMessageText{}
		replayMessageText.ToUserName.Text = fromUserName
		replayMessageText.FromUserName.Text = toUserName
		replayMessageText.CreateTime = time.Now().Unix()
		replayMessageText.MsgType.Text = "text"
		replayMessageText.Content.Text = "佑恩堂欢迎你！"
		returnData = replayMessageText
	} else {
		returnData, err = getReplyByContentTypeAndContent(fromUserName, toUserName, r.ContentType, r.Content)
	}
	return
}

// GetReplayMessageSubscribe GetReplayMessageSubscribe
func GetReplayMessageSubscribe(fromUserName, toUserName string) (returnData interface{}, err error) {

	var r *ReplyDefult
	r, err = GetReplyDefultByType("subscribe")
	if err != nil {
		replayMessageText := ReplayMessageText{}
		replayMessageText.ToUserName.Text = fromUserName
		replayMessageText.FromUserName.Text = toUserName
		replayMessageText.CreateTime = time.Now().Unix()
		replayMessageText.MsgType.Text = "text"
		replayMessageText.Content.Text = "佑恩堂欢迎你！"
		returnData = replayMessageText
	} else {
		returnData, err = getReplyByContentTypeAndContent(fromUserName, toUserName, r.ContentType, r.Content)
	}
	return
}

func getReplyByContentTypeAndContent(fromUserName, toUserName, contentType, content string) (returnData interface{}, err error) {
	if contentType == "text" {
		replayMessageText := ReplayMessageText{}
		replayMessageText.ToUserName.Text = fromUserName
		replayMessageText.FromUserName.Text = toUserName
		replayMessageText.CreateTime = time.Now().Unix()
		replayMessageText.MsgType.Text = "text"
		replayMessageText.Content.Text = content
		returnData = replayMessageText
	} else if contentType == "image" {
		replayMessageImage := ReplayMessageImage{}
		replayMessageImage.ToUserName.Text = fromUserName
		replayMessageImage.FromUserName.Text = toUserName
		replayMessageImage.CreateTime = time.Now().Unix()
		replayMessageImage.MsgType.Text = "image"
		replayMessageImage.Image.MediaId.Text = content
		returnData = replayMessageImage
	} else if contentType == "voice" {
		replayMessageVoice := ReplayMessageVoice{}
		replayMessageVoice.ToUserName.Text = fromUserName
		replayMessageVoice.FromUserName.Text = toUserName
		replayMessageVoice.CreateTime = time.Now().Unix()
		replayMessageVoice.MsgType.Text = "voice"
		replayMessageVoice.Voice.MediaId.Text = content
		returnData = replayMessageVoice
	} else if contentType == "video" {
		replayMessageVideo := ReplayMessageVideo{}
		replayMessageVideo.ToUserName.Text = fromUserName
		replayMessageVideo.FromUserName.Text = toUserName
		replayMessageVideo.CreateTime = time.Now().Unix()
		replayMessageVideo.MsgType.Text = "video"
		var m *MaterialMedia
		m, err = GetMaterialMediaByMediaID(content)
		replayMessageVideo.Video.MediaId.Text = m.MediaID
		replayMessageVideo.Video.Title.Text = m.Title
		replayMessageVideo.Video.Description.Text = m.Introduction
		returnData = replayMessageVideo
	} else if contentType == "music" {
		replayMessageText := ReplayMessageText{}
		replayMessageText.ToUserName.Text = fromUserName
		replayMessageText.FromUserName.Text = toUserName
		replayMessageText.CreateTime = time.Now().Unix()
		replayMessageText.MsgType.Text = "text"
		replayMessageText.Content.Text = "佑恩堂欢迎你！"
		returnData = replayMessageText
	} else if contentType == "news" {
		replayMessageText := ReplayMessageText{}
		replayMessageText.ToUserName.Text = fromUserName
		replayMessageText.FromUserName.Text = toUserName
		replayMessageText.CreateTime = time.Now().Unix()
		replayMessageText.MsgType.Text = "text"
		replayMessageText.Content.Text = "佑恩堂欢迎你！"
		returnData = replayMessageText
	} else {
		replayMessageText := ReplayMessageText{}
		replayMessageText.ToUserName.Text = fromUserName
		replayMessageText.FromUserName.Text = toUserName
		replayMessageText.CreateTime = time.Now().Unix()
		replayMessageText.MsgType.Text = "text"
		replayMessageText.Content.Text = "佑恩堂欢迎你！"
		returnData = replayMessageText
	}
	return
}
