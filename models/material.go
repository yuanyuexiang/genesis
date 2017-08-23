package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=ACCESS_TOKEN
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
type MaterialCount struct {
	total_count int32 `json:"total_count"`
	item_count  int32 `json:"item_count"`
}

type MaterialNewsList struct {
	MaterialCount
	Item []MaterialNews `json:"item"`
}

type MaterialNews struct {
	MediaId    string              `json:"media_id"`
	Content    MaterialNewsContent `json:"content"`
	UpdateTime int64               `json:"update_time"`
}

type MaterialNewsContent struct {
	NewsItem []News `json:"news_item"`
}

type News struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	ShowCoverPic     string `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	Url              string `json:"url"`
	ContentSourceUrl string `json:"content_source_url"`
}

type MaterialMultimediaList struct {
	MaterialCount
	Item []Multimedia `json:"item"`
}

type Multimedia struct {
	MediaId    string `json:"media_id"`
	Name       string `json:"name"`
	UpdateTime int64  `json:"update_time"`
	Url        string `json:"url"`
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

type MaterialUpdate struct {
	MediaId string     `json:"media_id"`
	Index   int64      `json:"index"`
	Article NewsUpdate `json:"articles"`
}

type NewsUpdate struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	ShowCoverPic     string `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	ContentSourceUrl string `json:"content_source_url"`
}

const (
	material_batchget_material = "https://api.weixin.qq.com/cgi-bin/material/batchget_material?"
	material_get_material      = "https://api.weixin.qq.com/cgi-bin/material/get_material?"
	material_del_material      = "https://api.weixin.qq.com/cgi-bin/material/del_material?"
	material_update_news       = "https://api.weixin.qq.com/cgi-bin/material/update_news?"
)

//获取永久素材
func GetMaterialByMediaId(media_id string) (v *MaterialNews, err error) {
	post_data := map[string]interface{}{"media_id": media_id}

	post_string, err := json.Marshal(post_data)

	if err != nil {
		return
	}

	access_token, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	str_url := material_get_material + "access_token=" + access_token
	request, err := http.NewRequest("POST", str_url, strings.NewReader(string(post_string)))
	if err != nil {
		return
	} else {
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		if response.StatusCode == 200 {
			body, err := ioutil.ReadAll(response.Body)
			bodystr := string(body)
			fmt.Println(bodystr)
			materialNewsContent := MaterialNewsContent{}
			err = json.Unmarshal(body, &materialNewsContent)
			if err != nil {
				fmt.Println(err)
			}
			v.Content = materialNewsContent
		} else {
			body, err := ioutil.ReadAll(response.Body)
			bodystr := string(body)
			fmt.Println(bodystr)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return v, err
}

//永久图文消息素材列表
func GetAllMaterialNewsList(offset int64, count int64) (v *MaterialNewsList, err error) {

	body, err := getAllMaterialListFromWechat("news", offset, count)
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

//其他类型（图片、语音、视频）
func GetAllMaterialMultimediaList(materialType string, offset int64, count int64) (v *MaterialMultimediaList, err error) {

	body, err := getAllMaterialListFromWechat(materialType, offset, count)
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

func getAllMaterialListFromWechat(materialType string, offset int64, count int64) (v []byte, err error) {

	post_data := map[string]interface{}{"type": materialType, "offset": offset, "count": count}

	post_string, err := json.Marshal(post_data)

	if err != nil {
		return
	}

	access_token, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	str_url := material_batchget_material + "access_token=" + access_token
	request, err := http.NewRequest("POST", str_url, strings.NewReader(string(post_string)))
	if err != nil {
		return
	} else {
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		if response.StatusCode == 200 {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}
			bodystr := string(body)
			fmt.Println(bodystr)
			v = body
		} else {
			body, err := ioutil.ReadAll(response.Body)
			bodystr := string(body)
			fmt.Println(bodystr)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return v, err
}

//修改永久图文素材
func UpdateMaterialById(m *MaterialUpdate) (err error) {
	post_string, err := json.Marshal(m)

	if err != nil {
		return
	}

	access_token, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	str_url := material_del_material + "access_token=" + access_token
	request, err := http.NewRequest("POST", str_url, strings.NewReader(string(post_string)))
	if err != nil {
		return
	} else {
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		if response.StatusCode == 200 {
			body, err := ioutil.ReadAll(response.Body)
			bodystr := string(body)
			fmt.Println(bodystr)
			errorResponse := ErrorResponse{}
			err = json.Unmarshal(body, &errorResponse)
			if err != nil {
				fmt.Println(err)
			}
			if errorResponse.ErrorCode != 0 {
				err = errors.New("delete fail")
			}
		} else {
			body, err := ioutil.ReadAll(response.Body)
			bodystr := string(body)
			fmt.Println(bodystr)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return err
}

//删除永久素材
func DeleteMaterialByMediaId(media_id string) (err error) {
	post_data := map[string]interface{}{"media_id": media_id}

	post_string, err := json.Marshal(post_data)

	if err != nil {
		return
	}

	access_token, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	str_url := material_del_material + "access_token=" + access_token
	request, err := http.NewRequest("POST", str_url, strings.NewReader(string(post_string)))
	if err != nil {
		return
	} else {
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		if response.StatusCode == 200 {
			body, err := ioutil.ReadAll(response.Body)
			bodystr := string(body)
			fmt.Println(bodystr)
			errorResponse := ErrorResponse{}
			err = json.Unmarshal(body, &errorResponse)
			if err != nil {
				fmt.Println(err)
			}
			if errorResponse.ErrorCode != 0 {
				err = errors.New("delete fail")
			}
		} else {
			body, err := ioutil.ReadAll(response.Body)
			bodystr := string(body)
			fmt.Println(bodystr)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return err
}
