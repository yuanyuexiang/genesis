package models

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

/**

文本消息

<xml>
 <ToUserName><![CDATA[toUser]]></ToUserName>
 <FromUserName><![CDATA[fromUser]]></FromUserName>
 <CreateTime>1348831860</CreateTime>
 <MsgType><![CDATA[text]]></MsgType>
 <Content><![CDATA[this is a test]]></Content>
 <MsgId>1234567890123456</MsgId>
</xml>

图片消息

<xml>
 <ToUserName><![CDATA[toUser]]></ToUserName>
 <FromUserName><![CDATA[fromUser]]></FromUserName>
 <CreateTime>1348831860</CreateTime>
 <MsgType><![CDATA[image]]></MsgType>
 <PicUrl><![CDATA[this is a url]]></PicUrl>
 <MediaId><![CDATA[media_id]]></MediaId>
 <MsgId>1234567890123456</MsgId>
</xml>

语音消息

<xml>
 <ToUserName><![CDATA[toUser]]></ToUserName>
 <FromUserName><![CDATA[fromUser]]></FromUserName>
 <CreateTime>1357290913</CreateTime>
 <MsgType><![CDATA[voice]]></MsgType>
 <MediaId><![CDATA[media_id]]></MediaId>
 <Format><![CDATA[Format]]></Format>
 <MsgId>1234567890123456</MsgId>
</xml>

视频消息

<xml>
 <ToUserName><![CDATA[toUser]]></ToUserName>
 <FromUserName><![CDATA[fromUser]]></FromUserName>
 <CreateTime>1357290913</CreateTime>
 <MsgType><![CDATA[video]]></MsgType>
 <MediaId><![CDATA[media_id]]></MediaId>
 <ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId>
 <MsgId>1234567890123456</MsgId>
</xml>

小视频消息

<xml>
 <ToUserName><![CDATA[toUser]]></ToUserName>
 <FromUserName><![CDATA[fromUser]]></FromUserName>
 <CreateTime>1357290913</CreateTime>
 <MsgType><![CDATA[shortvideo]]></MsgType>
 <MediaId><![CDATA[media_id]]></MediaId>
 <ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId>
 <MsgId>1234567890123456</MsgId>
</xml>

地理位置消息

<xml>
 <ToUserName><![CDATA[toUser]]></ToUserName>
 <FromUserName><![CDATA[fromUser]]></FromUserName>
 <CreateTime>1351776360</CreateTime>
 <MsgType><![CDATA[location]]></MsgType>
 <Location_X>23.134521</Location_X>
 <Location_Y>113.358803</Location_Y>
 <Scale>20</Scale>
 <Label><![CDATA[位置信息]]></Label>
 <MsgId>1234567890123456</MsgId>
</xml>

链接消息

<xml>
 <ToUserName><![CDATA[toUser]]></ToUserName>
 <FromUserName><![CDATA[fromUser]]></FromUserName>
 <CreateTime>1351776360</CreateTime>
 <MsgType><![CDATA[link]]></MsgType>
 <Title><![CDATA[公众平台官网链接]]></Title>
 <Description><![CDATA[公众平台官网链接]]></Description>
 <Url><![CDATA[url]]></Url>
 <MsgId>1234567890123456</MsgId>
</xml>


关注/取消关注事件

推送XML数据包示例：

<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>123456789</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[subscribe]]></Event>
</xml>

扫描带参数二维码事件

用户未关注时，进行关注后的事件推送

推送XML数据包示例：

<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>123456789</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[subscribe]]></Event>
<EventKey><![CDATA[qrscene_123123]]></EventKey>
<Ticket><![CDATA[TICKET]]></Ticket>
</xml>

用户已关注时的事件推送

推送XML数据包示例：
<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>123456789</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[SCAN]]></Event>
<EventKey><![CDATA[SCENE_VALUE]]></EventKey>
<Ticket><![CDATA[TICKET]]></Ticket>
</xml>

自定义菜单事件

点击菜单拉取消息时的事件推送

推送XML数据包示例：

<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>123456789</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[CLICK]]></Event>
<EventKey><![CDATA[EVENTKEY]]></EventKey>
</xml>


点击菜单跳转链接时的事件推送

推送XML数据包示例：

<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>123456789</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[VIEW]]></Event>
<EventKey><![CDATA[www.qq.com]]></EventKey>
</xml>

**/

//Message Message
type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int64
	MsgType      CDATA
}

//MessageText MessageText
type MessageText struct {
	Message
	Content CDATA
	MsgId   int64
}

//MessageImage MessageImage
type MessageImage struct {
	Message
	PicUrl  CDATA
	MediaId CDATA
	MsgId   int64
}

//MessageVoice MessageVoice
type MessageVoice struct {
	Message
	MediaId CDATA
	Format  CDATA
	MsgId   int64
}

//MessageVideo MessageVideo
type MessageVideo struct {
	Message
	MediaId      CDATA
	ThumbMediaId CDATA
	MsgId        int64
}

//MessageShortVideo MessageShortVideo
type MessageShortVideo struct {
	Message
	MediaId      CDATA
	ThumbMediaId CDATA
	MsgId        int64
}

//MessageLocation MessageLocation
type MessageLocation struct {
	Message
	Location_X float64
	Location_Y float64
	Scale      int64
	Label      CDATA
	MsgId      int64
}

//MessageLink MessageLink
type MessageLink struct {
	Message
	Title       CDATA
	Description CDATA
	Url         CDATA
	MsgId       int64
}

//MessageEvent MessageEvent
type MessageEvent struct {
	Message
	Event CDATA
}

//MessageEventQR MessageEventQR
type MessageEventQR struct {
	MessageEvent
	Ticket CDATA
}

//MessageEventLOCATION MessageEventLOCATION
type MessageEventLOCATION struct {
	MessageEvent
	Latitude  float64
	Longitude float64
	Precision float64
}

//MessageEventMenu MessageEventMenu
type MessageEventMenu struct {
	MessageEvent
	EventKey CDATA
}

func init() {
}

//HandleMessage HandleMessage
func HandleMessage(msg []byte) (returnData interface{}, err error) {
	var message Message
	err = xml.Unmarshal(msg, &message)
	if err != nil {
		return
	}
	if message.MsgType.Text == "text" {
		returnData, err = HandleMessageText(msg)
	} else if message.MsgType.Text == "image" {
		returnData, err = HandleMessageImage(msg)
	} else if message.MsgType.Text == "voice" {
		returnData, err = HandleMessageVoice(msg)
	} else if message.MsgType.Text == "video" {
		returnData, err = HandleMessageVideo(msg)
	} else if message.MsgType.Text == "shortvideo" {
		returnData, err = HandleMessageShortvideo(msg)
	} else if message.MsgType.Text == "location" {
		returnData, err = HandleMessageLocation(msg)
	} else if message.MsgType.Text == "link" {
		returnData, err = HandleMessageLink(msg)
	} else if message.MsgType.Text == "event" {
		returnData, err = HandleMessageEvent(msg)
	} else {
		err = errors.New("wechat message error!")
	}
	return
}

//HandleMessageText HandleMessageText
func HandleMessageText(msg []byte) (returnData interface{}, err error) {
	var messageText MessageText
	err = xml.Unmarshal(msg, &messageText)
	if err != nil {
		return
	}
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageText.FromUserName.Text
	replayMessageText.FromUserName.Text = messageText.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，已经收到你发送的文字，功能开发中"
	returnData = replayMessageText
	return
}

//HandleMessageImage HandleMessageImage
func HandleMessageImage(msg []byte) (returnData interface{}, err error) {
	var messageImage MessageImage
	err = xml.Unmarshal(msg, &messageImage)
	if err != nil {
		return
	}
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageImage.FromUserName.Text
	replayMessageText.FromUserName.Text = messageImage.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，已经收到你发送的图片，功能开发中"
	returnData = replayMessageText
	return
}

//HandleMessageVoice HandleMessageVoice
func HandleMessageVoice(msg []byte) (returnData interface{}, err error) {
	var messageVoice MessageVoice
	err = xml.Unmarshal(msg, &messageVoice)
	if err != nil {
		return
	}
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageVoice.FromUserName.Text
	replayMessageText.FromUserName.Text = messageVoice.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，已经收到你发送的音频，功能开发中"
	returnData = replayMessageText
	return
}

// HandleMessageVideo HandleMessageVideo
func HandleMessageVideo(msg []byte) (returnData interface{}, err error) {
	var messageVideo MessageVideo
	err = xml.Unmarshal(msg, &messageVideo)
	if err != nil {
		return
	}
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageVideo.FromUserName.Text
	replayMessageText.FromUserName.Text = messageVideo.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，已经收到你发送的视频，功能开发中"
	returnData = replayMessageText
	return
}

//HandleMessageShortvideo HandleMessageShortvideo
func HandleMessageShortvideo(msg []byte) (returnData interface{}, err error) {
	var messageShortVideo MessageShortVideo
	err = xml.Unmarshal(msg, &messageShortVideo)
	if err != nil {
		return
	}
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageShortVideo.FromUserName.Text
	replayMessageText.FromUserName.Text = messageShortVideo.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，已经收到你发送的小视频，功能开发中"
	returnData = replayMessageText
	return
}

//HandleMessageLocation HandleMessageLocation
func HandleMessageLocation(msg []byte) (returnData interface{}, err error) {
	var messageLocation MessageLocation
	err = xml.Unmarshal(msg, &messageLocation)
	if err != nil {
		return
	}
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageLocation.FromUserName.Text
	replayMessageText.FromUserName.Text = messageLocation.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，已经收到你发送的位置，功能开发中"
	returnData = replayMessageText
	return
}

//HandleMessageLink HandleMessageLink
func HandleMessageLink(msg []byte) (returnData interface{}, err error) {
	var messageLink MessageLink
	err = xml.Unmarshal(msg, &messageLink)
	if err != nil {
		return
	}
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageLink.FromUserName.Text
	replayMessageText.FromUserName.Text = messageLink.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，已经收到你发送的链接，功能开发中"
	returnData = replayMessageText
	return
}

//HandleMessageEvent HandleMessageEvent
func HandleMessageEvent(msg []byte) (returnData interface{}, err error) {
	var messageEvent MessageEvent
	err = xml.Unmarshal(msg, &messageEvent)
	if err != nil {
		return
	}
	if messageEvent.Event.Text == "subscribe" {
		returnData, err = HandleMessageEventSubscribe(&messageEvent)
	} else if messageEvent.Event.Text == "unsubscribe" {
		returnData, err = HandleMessageEventUnsubscribe(&messageEvent)
	} else if messageEvent.Event.Text == "LOCATION" {
		returnData, err = HandleMessageEventLOCATION(msg)
	} else if messageEvent.Event.Text == "CLICK" {
		returnData, err = HandleMessageEventCLICK(msg)
	} else if messageEvent.Event.Text == "VIEW" {
		returnData, err = HandleMessageEventVIEW(msg)
	} else {
		err = errors.New("wechat event error!")
	}
	return
}

//HandleMessageEventSubscribe HandleMessageEventSubscribe
func HandleMessageEventSubscribe(messageEvent *MessageEvent) (returnData interface{}, err error) {
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageEvent.FromUserName.Text
	replayMessageText.FromUserName.Text = messageEvent.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，欢迎关注XX教会，功能开发中"
	returnData = replayMessageText
	return
}

//HandleMessageEventUnsubscribe HandleMessageEventUnsubscribe
func HandleMessageEventUnsubscribe(messageEvent *MessageEvent) (returnData interface{}, err error) {
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageEvent.FromUserName.Text
	replayMessageText.FromUserName.Text = messageEvent.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，欢迎再次关注XX教会，功能开发中"
	returnData = replayMessageText
	return
}

//HandleMessageEventLOCATION HandleMessageEventLOCATION
func HandleMessageEventLOCATION(msg []byte) (returnData interface{}, err error) {
	var messageEventLOCATION MessageEventLOCATION
	err = xml.Unmarshal(msg, &messageEventLOCATION)
	if err != nil {
		return
	}
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageEventLOCATION.FromUserName.Text
	replayMessageText.FromUserName.Text = messageEventLOCATION.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，上报地理位置事件，功能开发中"
	returnData = replayMessageText
	return
}

//HandleMessageEventCLICK HandleMessageEventCLICK
func HandleMessageEventCLICK(msg []byte) (returnData interface{}, err error) {
	var messageEventMenu MessageEventMenu
	err = xml.Unmarshal(msg, &messageEventMenu)
	if err != nil {
		return
	}
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageEventMenu.FromUserName.Text
	replayMessageText.FromUserName.Text = messageEventMenu.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，菜单被你点击了，功能开发中"
	returnData = replayMessageText
	return
}

//HandleMessageEventVIEW HandleMessageEventVIEW
func HandleMessageEventVIEW(msg []byte) (returnData interface{}, err error) {
	var messageEventMenu MessageEventMenu
	err = xml.Unmarshal(msg, &messageEventMenu)
	if err != nil {
		return
	}
	replayMessageText := ReplayMessageText{}
	replayMessageText.ToUserName.Text = messageEventMenu.FromUserName.Text
	replayMessageText.FromUserName.Text = messageEventMenu.ToUserName.Text
	replayMessageText.CreateTime = time.Now().Unix()
	replayMessageText.MsgType.Text = "text"
	replayMessageText.Content.Text = "你好，点击菜单跳转链接时的事件推送，功能开发中"
	returnData = replayMessageText
	return
}

//CheckMessageInterface CheckMessageInterface
func CheckMessageInterface(signature, timestamp, nonce, echostr string) (err error) {
	if signature != makeSignature(timestamp, nonce) {
		err = errors.New("CHECK FAILED")
	}
	return err
}

//生成签名
func makeSignature(timestamp, nonce string) string {
	token := beego.AppConfig.String("token")
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
