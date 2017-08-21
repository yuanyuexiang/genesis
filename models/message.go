package models

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	_ "reflect"
	"sort"
	"strings"
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

type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int64
	MsgType      CDATA
}

type MessageText struct {
	Message
	Content CDATA
	MsgId   int64
}

type MessageVoice struct {
	Message
	MediaId CDATA
	Format  CDATA
	MsgId   int64
}

type MessageVideo struct {
	Message
	MediaId      CDATA
	ThumbMediaId CDATA
	MsgId        int64
}

type MessageShortVideo struct {
	Message
	MediaId      CDATA
	ThumbMediaId CDATA
	MsgId        int64
}

type MessageLocation struct {
	Message
	Location_X float64
	Location_Y float64
	Scale      int64
	Label      CDATA
	MsgId      int64
}

type MessageLink struct {
	Message
	Title       CDATA
	Description CDATA
	Url         CDATA
	MsgId       int64
}

type MessageEvent struct {
	Message
	Event CDATA
}

type MessageEventQR struct {
	MessageEvent
	Ticket CDATA
}

type MessageEventLOCATION struct {
	MessageEvent
	Latitude  float64
	Longitude float64
	Precision float64
}

type MessageEventMenu struct {
	MessageEvent
	EventKey CDATA
}

func init() {
}

func HandleMessage(msg []byte) (id int64, err error) {
	var message Message
	err = xml.Unmarshal(msg, &message)
	if err != nil {
		return
	}
	if message.MsgType.Text == "text" {
		//
	} else if message.MsgType.Text == "image" {
		//
	} else if message.MsgType.Text == "voice" {
		//
	} else if message.MsgType.Text == "video" {
		//
	} else if message.MsgType.Text == "shortvideo" {
		//
	} else if message.MsgType.Text == "location" {
		//
	} else if message.MsgType.Text == "link" {
		//
	} else if message.MsgType.Text == "event" {

		//
	} else {

	}
	return
}

func HandleMessageEvent(msg []byte, m *MessageEvent) (returnData interface{}, err error) {

	if m.Event.Text == "subscribe" {

	} else if m.Event.Text == "unsubscribe" {

	} else if m.Event.Text == "LOCATION" {

	} else if m.Event.Text == "CLICK" {

	} else if m.Event.Text == "VIEW" {

	} else {
	}
	return
}

//
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
