package models

import (
	"encoding/xml"
	"fmt"
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
		Articles []Article
	}
}

//Article Article
type Article struct {
	XMLName     xml.Name `xml:"item"`
	Title       CDATA
	Description CDATA
	PicUrl      CDATA
	Url         CDATA
}

func GetReplayMessage() (rm *ReplayMessageArticles, err error) {
	rm = &ReplayMessageArticles{}
	rm.CreateTime = 123123
	rm.FromUserName.Text = "test"

	a := Article{}
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
