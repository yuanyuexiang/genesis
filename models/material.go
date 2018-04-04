package models

import (
	"encoding/json"
	"fmt"
)

/*
https://api.weixin.qq.com/cgi-bin/material/batchget_material?accessToken=accessToken
{
  "total_count": TOTAL_COUNT,
  "item_count": ITEM_COUNT,
  "item": [{
      "media_id": MEDIA_ID,
      "content": {
          "news_item": [{
              "title": TITLE,
              "thumb_media_id": THUMB_MEDIA_ID,
              "show_cover_pic": SHOW_COVER_PIC(0 / 1),
              "author": AUTHOR,
              "digest": DIGEST,
              "content": CONTENT,
              "url": URL,
              "content_source_url": CONTETN_SOURCE_URL
          },
          //多图文消息会在此处有多篇文章
          ]
       },
       "update_time": UPDATE_TIME
   },
   //可能有多个图文消息item结构
 ]
}

{
  "total_count": TOTAL_COUNT,
  "item_count": ITEM_COUNT,
  "item": [{
      "media_id": MEDIA_ID,
      "name": NAME,
      "update_time": UPDATE_TIME,
      "url":URL
  },
  //可能会有多个素材
  ]
}
*/

// MaterialTotalCount 获取素材总数 图片和图文消息素材（包括单图文和多图文）的总数上限为5000，其他素材的总数上限为1000
type MaterialTotalCount struct {
	VoiceCount int64 `json:"voice_count"`
	VideoCount int64 `json:"video_count"`
	ImageCount int64 `json:"image_count"`
	NewsCount  int64 `json:"news_count"`
}

//MaterialCount MaterialCount
type MaterialCount struct {
	TotalCount int32 `json:"total_count"`
	ItemCount  int32 `json:"item_count"`
}

//MaterialNewsList MaterialNewsList
type MaterialNewsList struct {
	MaterialCount
	Item []MaterialNews `json:"item"`
}

//MaterialNews MaterialNews
type MaterialNews struct {
	MediaID    string              `json:"media_id"`
	Content    MaterialNewsContent `json:"content"`
	UpdateTime int64               `json:"update_time"`
}

//MaterialNewsContent MaterialNewsContent
type MaterialNewsContent struct {
	NewsItem   []NewsItem `json:"news_item"`
	CreateTime int64      `json:"create_time"`
	UpdateTime int64      `json:"update_time"`
}

//NewsItem NewsItem
type NewsItem struct {
	Title              string `json:"title"`
	ThumbMediaID       string `json:"thumb_media_id"`
	ShowCoverPic       int64  `json:"show_cover_pic"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	Content            string `json:"content"`
	URL                string `json:"url"`
	ContentSourceURL   string `json:"content_source_url"`
	ThumbURL           string `json:"thumb_url"`
	NeedOpenComment    int64  `json:"need_open_comment"`
	OnlyFansCanComment int64  `json:"only_fans_can_comment"`
}

// MaterialArticleItem 图文
type MaterialArticleItem struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     int64  `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
}

// MaterialArticles 图文 图文
type MaterialArticles struct {
	Items []MaterialArticleItem `json:"articles"`
}

//MaterialMultimediaList MaterialMultimediaList
type MaterialMultimediaList struct {
	MaterialCount
	Item []Multimedia `json:"item"`
}

//Multimedia Multimedia
type Multimedia struct {
	MediaID    string `json:"media_id"`
	Name       string `json:"name"`
	UpdateTime int64  `json:"update_time"`
	URL        string `json:"url"`
}

/*
{
 "media_id":MEDIA_ID,
 "index":INDEX,
 "articles": {
      "title": TITLE,
      "thumb_media_id": THUMB_MEDIA_ID,
      "author": AUTHOR,
      "digest": DIGEST,
      "show_cover_pic": SHOW_COVER_PIC(0 / 1),
      "content": CONTENT,
      "content_source_url": CONTENT_SOURCE_URL
   }
}
*/

//MaterialUpdate MaterialUpdate
type MaterialUpdate struct {
	MediaID string     `json:"media_id"`
	Index   int64      `json:"index"`
	Article NewsUpdate `json:"articles"`
}

//NewsUpdate NewsUpdate
type NewsUpdate struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	ShowCoverPic     int64  `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
}

// MaterialInfoResponse 添加素材返回说明
type MaterialInfoResponse struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

const (
	mediaUploadimg           = "https://api.weixin.qq.com/cgi-bin/media/uploadimg?"
	materialAddMaterial      = "https://api.weixin.qq.com/cgi-bin/material/add_material?"
	materialAddNews          = "https://api.weixin.qq.com/cgi-bin/material/add_news?"
	materialBatchgetMaterial = "https://api.weixin.qq.com/cgi-bin/material/batchget_material?"
	materialGetMaterial      = "https://api.weixin.qq.com/cgi-bin/material/get_material?"
	materialDelMaterial      = "https://api.weixin.qq.com/cgi-bin/material/del_material?"
	materialUpdateNews       = "https://api.weixin.qq.com/cgi-bin/material/update_news?"
	materialGetMaterialcount = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?"
)

// AddNews 新增其他类型永久素材
// 通过POST表单来调用接口，表单id为media，包含需要上传的素材内容，有filename、filelength、content-type等信息。请注意：图片素材将进入公众平台官网素材管理模块中的默认分组。
func AddNews(articles *MaterialArticles) (mediaInfo MaterialInfoResponse, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	strURL := materialAddNews + "access_token=" + accessToken

	postData, err := json.Marshal(articles)
	if err != nil {
		return
	}
	body, err := post(strURL, postData)

	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &mediaInfo)
	return
}

// AddMaterialImage 上传图文消息内的图片获取URL
//本接口所上传的图片不占用公众号的素材库中图片数量的5000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下。
func AddMaterialImage(filePath string) (mediaInfo MaterialInfoResponse, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	strURL := mediaUploadimg + "access_token=" + accessToken

	if err != nil {
		return
	}
	body, err := postFile(strURL, "", filePath)

	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &mediaInfo)
	return
}

// AddMaterial 新增其他类型永久素材
// 通过POST表单来调用接口，表单id为media，包含需要上传的素材内容，有filename、filelength、content-type等信息。请注意：图片素材将进入公众平台官网素材管理模块中的默认分组。
func AddMaterial(filePath, materialType string) (mediaInfo MaterialInfoResponse, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	strURL := materialPictureAddMaterial + "access_token=" + accessToken + "&type=" + materialType

	desc := map[string]string{"title": "title", "introduction": "introduction"}

	description, err := json.Marshal(desc)

	if err != nil {
		return
	}
	body, err := postFile(strURL, string(description), filePath)

	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &mediaInfo)
	return
}

//GetMaterialByMediaID 获取永久素材
func GetMaterialByMediaID(mediaID string) (v *MaterialNewsContent, err error) {
	postData := map[string]interface{}{"media_id": mediaID}
	postByte, err := json.Marshal(postData)

	if err != nil {
		return
	}

	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	strURL := materialGetMaterial + "access_token=" + accessToken
	body, err := post(strURL, postByte)
	v = &MaterialNewsContent{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

//GetAllMaterialNewsList 永久图文消息素材列表
func GetAllMaterialNewsList(offset int64, count int64) (v *MaterialNewsList, err error) {

	body, err := getAllMaterialListFromWechat("news", offset, count)
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

//GetAllMaterialMultimediaList 其他类型（图片、语音、视频）
func GetAllMaterialMultimediaList(materialType string, offset int64, count int64) (v *MaterialMultimediaList, err error) {
	body, err := getAllMaterialListFromWechat(materialType, offset, count)
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func getAllMaterialListFromWechat(materialType string, offset int64, count int64) (v []byte, err error) {
	postDataQuery := map[string]interface{}{"type": materialType, "offset": offset, "count": count}
	if err != nil {
		return
	}

	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	strURL := materialBatchgetMaterial + "access_token=" + accessToken
	postData, err := json.Marshal(postDataQuery)

	if err != nil {
		return
	}
	body, err := post(strURL, postData)
	v = body
	return
}

//UpdateMaterialByID 修改永久图文素材
func UpdateMaterialByID(m *MaterialUpdate) (err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	strURL := materialUpdateNews + "access_token=" + accessToken
	postData, err := json.Marshal(m)

	if err != nil {
		return
	}
	body, err := post(strURL, postData)
	bodystr := string(body)
	fmt.Println(bodystr)
	if err != nil {
		fmt.Println(err)
	}
	return
}

//DeleteMaterialByMediaID  删除永久素材
func DeleteMaterialByMediaID(mediaID string) (err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	strURL := materialDelMaterial + "access_token=" + accessToken
	requestData := map[string]interface{}{"media_id": mediaID}
	postData, err := json.Marshal(requestData)
	if err != nil {
		return
	}
	body, err := post(strURL, postData)
	bodystr := string(body)
	fmt.Println(bodystr)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// GetMaterialcount 获取素材总数
func GetMaterialcount() (v *MaterialTotalCount, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	strURL := materialGetMaterialcount + "access_token=" + accessToken
	body, err := get(strURL)
	v = &MaterialTotalCount{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
