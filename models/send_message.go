package models

import (
	"encoding/json"
	"fmt"
)

/*
{
   "articles": [
		 {
            "thumb_media_id":"qI6_Ze_6PtV7svjolgs-rN6stStuHIjs9_DidOHaj0Q-mwvBelOXCFZiq2OsIU-p",
            "author":"xxx",
			"title":"Happy Day",
			"content_source_url":"www.qq.com",
			"content":"content",
			"digest":"digest",
            "show_cover_pic":1
		 },
		 {
            "thumb_media_id":"qI6_Ze_6PtV7svjolgs-rN6stStuHIjs9_DidOHaj0Q-mwvBelOXCFZiq2OsIU-p",
            "author":"xxx",
			"title":"Happy Day",
			"content_source_url":"www.qq.com",
			"content":"content",
			"digest":"digest",
           	"show_cover_pic":0
		 }
   ]
}
*/

//ArticleItem ArticleItem
type ArticleItem struct {
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Title            string `json:"title"`
	ContentSourceURL string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
	ShowCoverPic     string `json:"show_cover_pic"`
}

//Articles Articles
type Articles struct {
	Articles []ArticleItem `json:"articles"`
}

const (
	wechatBaseURL = "https://api.weixin.qq.com/cgi-bin/"
)

//上传图文消息内的图片获取URL【订阅号与服务号认证后均可用】
/*
http请求方式: POST
https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
调用示例（使用curl命令，用FORM表单方式上传一个图片）：
curl -F media=@test.jpg "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN"
*/
func UploadNewsMessagePicture(file string) (picURL string, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	strRequest := "media/uploadimg?access_token=" + accessToken
	strURL := BaseUserInfoURL + strRequest

	body, err := postFile(strURL, "media=@test.jpg", file)
	var data map[string]string
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}
	picURL = data["url"]
	return
}

//上传图文消息素材【订阅号与服务号认证后均可用】
/*
http请求方式: POST
https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
*/
func UploadNewsMessage(articles *Articles) (data map[string]interface{}, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	strRequest := "media/uploadimg?access_token=" + accessToken
	strURL := BaseUserInfoURL + strRequest
	postData, err := json.Marshal(articles)
	if err != nil {
		return
	}
	body, err := post(strURL, postData)
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

/*
{
   "filter":{
      "is_to_all":false,
      "tag_id":2
   },
   "mpnews":{
      "media_id":"123dsdajkasd231jhksad"
   },
    "msgtype":"mpnews",
    "send_ignore_reprint":0
}
*/

//AllSendNewsMessage AllSendNewsMessage
type AllSendNewsMessage struct {
	Filter struct {
		IsToAll bool  `json:"is_to_all"`
		TagID   int64 `json:"tag_id"`
	} `json:"filter"`
	Mpnews struct {
		MediaID string `json:"media_id"`
	} `json:"mpnews"`
	SendIgnoreReprint int64 `json:"send_ignore_reprint"`
}

//AllSendTextMessage AllSendTextMessage
type AllSendTextMessage struct {
	Filter struct {
		IsToAll bool  `json:"is_to_all"`
		TagID   int64 `json:"tag_id"`
	} `json:"filter"`
	Text struct {
		MediaID string `json:"media_id"`
	} `json:"text"`
	MsgType string `json:"msgtype"`
}

//AllSendVoiceMessage AllSendVoiceMessage
type AllSendVoiceMessage struct {
	Filter struct {
		IsToAll bool  `json:"is_to_all"`
		TagID   int64 `json:"tag_id"`
	} `json:"filter"`
	Voice struct {
		MediaID string `json:"media_id"`
	} `json:"voice"`
	MsgType string `json:"msgtype"`
}

//AllSendImageMessage AllSendImageMessage
type AllSendImageMessage struct {
	Filter struct {
		IsToAll bool  `json:"is_to_all"`
		TagID   int64 `json:"tag_id"`
	} `json:"filter"`
	Image struct {
		MediaID string `json:"media_id"`
	} `json:"image"`
	MsgType string `json:"msgtype"`
}

//PostAllSendMessage 根据标签进行群发【订阅号与服务号认证后均可用】
func PostAllSendMessage(requestData interface{}) (data map[string]interface{}, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	strRequest := "message/mass/sendall?access_token=" + accessToken
	strURL := BaseUserInfoURL + strRequest
	postData, err := json.Marshal(requestData)
	if err != nil {
		return
	}
	body, err := post(strURL, postData)
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

//DeleteAllSendMessage 删除群发【订阅号与服务号认证后均可用】
func DeleteAllSendMessage(msg_id, article_idx int64) (data map[string]interface{}, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	strRequest := "message/mass/delete?access_token=" + accessToken
	strURL := BaseUserInfoURL + strRequest
	requestData := map[string]int64{"msg_id": msg_id, "article_idx": article_idx}
	postData, err := json.Marshal(requestData)
	if err != nil {
		return
	}
	body, err := post(strURL, postData)
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

//CheckAllSendMessage 查询群发消息发送状态【订阅号与服务号认证后均可用】
func CheckAllSendMessage(msgID int64) (data map[string]interface{}, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	strRequest := "message/mass/get?access_token=" + accessToken
	strURL := BaseUserInfoURL + strRequest
	requestData := map[string]int64{"msg_id": msgID}
	postData, err := json.Marshal(requestData)
	if err != nil {
		return
	}
	body, err := post(strURL, postData)
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}
	return
}
