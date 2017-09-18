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
	NewsItem []News `json:"news_item"`
}

//News News
type News struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	ShowCoverPic     string `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	URL              string `json:"url"`
	ContentSourceURL string `json:"content_source_url"`
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
	ShowCoverPic     string `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
}

const (
	materialAddMaterial      = "https://api.weixin.qq.com/cgi-bin/material/add_material?"
	materialAddNews          = "https://api.weixin.qq.com/cgi-bin/material/add_news?"
	materialBatchgetMaterial = "https://api.weixin.qq.com/cgi-bin/material/batchget_material?"
	materialGetMaterial      = "https://api.weixin.qq.com/cgi-bin/material/get_material?"
	materialDelMaterial      = "https://api.weixin.qq.com/cgi-bin/material/del_material?"
	materialUpdateNews       = "https://api.weixin.qq.com/cgi-bin/material/update_news?"
)

//GetMaterialByMediaID 获取永久素材
func GetMaterialByMediaID(mediaID string) (v *MaterialNews, err error) {
	postData := map[string]interface{}{"media_id": mediaID}

	postByte, err := json.Marshal(postData)

	if err != nil {
		return
	}

	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	strURL := materialGetMaterial + "accessToken=" + accessToken
	body, err := post(strURL, postByte)
	materialNewsContent := MaterialNewsContent{}
	err = json.Unmarshal(body, &materialNewsContent)
	if err != nil {
		fmt.Println(err)
	}
	v.Content = materialNewsContent

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

	strURL := materialBatchgetMaterial + "accessToken=" + accessToken
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

	strURL := materialUpdateNews + "accessToken=" + accessToken
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

	strURL := materialDelMaterial + "accessToken=" + accessToken
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
