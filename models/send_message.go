package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
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

type ArticleItem struct {
	ThumbMediaId     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Title            string `json:"title"`
	ContentSourceUrl string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
	ShowCoverPic     string `json:"show_cover_pic"`
}

type Articles struct {
	Articles []ArticleItem `json:"articles"`
}

const (
	WechatBaseUrl = "https://api.weixin.qq.com/cgi-bin/"
)

//上传图文消息内的图片获取URL【订阅号与服务号认证后均可用】
/*
http请求方式: POST
https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
调用示例（使用curl命令，用FORM表单方式上传一个图片）：
curl -F media=@test.jpg "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN"
*/
func UploadNewsMessagePicture(file string) (picUrl string, err error) {
	access_token, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	str_request := "media/uploadimg?access_token=" + access_token
	str_url := BaseUserInfoUrl + str_request

	body, err := postFile(str_url, "media=@test.jpg", file)
	var data map[string]string
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}
	picUrl = data["url"]
	return
}

//上传图文消息素材【订阅号与服务号认证后均可用】
/*
http请求方式: POST
https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
*/
func UploadNewsMessage(articles *Articles) (data map[string]interface{}, err error) {
	access_token, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	str_request := "media/uploadimg?access_token=" + access_token
	str_url := BaseUserInfoUrl + str_request
	postData, err := json.Marshal(articles)
	if err != nil {
		return
	}
	body, err := post(str_url, postData)
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

type AllSendNewsMessage struct {
	Filter struct {
		IsToAll bool  `json:"is_to_all"`
		TagId   int64 `json:"tag_id"`
	} `json:"filter"`
	mpnews struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
	SendIgnoreReprint int64 `json:"send_ignore_reprint"`
}

type AllSendTextMessage struct {
	Filter struct {
		IsToAll bool  `json:"is_to_all"`
		TagId   int64 `json:"tag_id"`
	} `json:"filter"`
	Text struct {
		MediaId string `json:"media_id"`
	} `json:"text"`
	MsgType string `json:"msgtype"`
}

type AllSendVoiceMessage struct {
	Filter struct {
		IsToAll bool  `json:"is_to_all"`
		TagId   int64 `json:"tag_id"`
	} `json:"filter"`
	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
	MsgType string `json:"msgtype"`
}

type AllSendImageMessage struct {
	Filter struct {
		IsToAll bool  `json:"is_to_all"`
		TagId   int64 `json:"tag_id"`
	} `json:"filter"`
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
	MsgType string `json:"msgtype"`
}

//根据标签进行群发【订阅号与服务号认证后均可用】
func PostAllSendMessage(requestData interface{}) (data map[string]interface{}, err error) {
	access_token, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	str_request := "message/mass/sendall?access_token=" + access_token
	str_url := BaseUserInfoUrl + str_request
	postData, err := json.Marshal(requestData)
	if err != nil {
		return
	}
	body, err := post(str_url, postData)
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

//删除群发【订阅号与服务号认证后均可用】
func DeleteAllSendMessage(msg_id, article_idx int64) (data map[string]interface{}, err error) {
	access_token, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	str_request := "message/mass/delete?access_token=" + access_token
	str_url := BaseUserInfoUrl + str_request
	requestData := map[string]int64{"msg_id": msg_id, "article_idx": article_idx}
	postData, err := json.Marshal(requestData)
	if err != nil {
		return
	}
	body, err := post(str_url, postData)
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

//查询群发消息发送状态【订阅号与服务号认证后均可用】
func CheckAllSendMessage(msg_id int64) (data map[string]interface{}, err error) {
	access_token, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	str_request := "message/mass/get?access_token=" + access_token
	str_url := BaseUserInfoUrl + str_request
	requestData := map[string]int64{"msg_id": msg_id}
	postData, err := json.Marshal(requestData)
	if err != nil {
		return
	}
	body, err := post(str_url, postData)
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func post(url string, postData []byte) (data []byte, err error) {
	request, err := http.NewRequest("POST", url, strings.NewReader(string(postData)))
	if err != nil {
		return
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		err = errors.New("wechat server error")
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	bodystr := string(body)
	fmt.Println(bodystr)
	data = body
	return
}

func postFile(url, formDataTag, filePath string) (data []byte, err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()
	fw, err := w.CreateFormFile(formDataTag, filePath)
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	w.Close()
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	bodystr := string(body)
	fmt.Println(bodystr)
	data = body
	return
}
