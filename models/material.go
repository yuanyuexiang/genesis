package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
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

// WechatMaterialTotalCount 获取素材总数 图片和图文消息素材（包括单图文和多图文）的总数上限为5000，其他素材的总数上限为1000
type WechatMaterialTotalCount struct {
	VoiceCount int64 `json:"voice_count"`
	VideoCount int64 `json:"video_count"`
	ImageCount int64 `json:"image_count"`
	NewsCount  int64 `json:"news_count"`
}

//WechatMaterialCount MaterialCountWechat
type WechatMaterialCount struct {
	TotalCount int32 `json:"total_count"`
	ItemCount  int32 `json:"item_count"`
}

//WechatMaterialNewsList MaterialNewsListWechat
type WechatMaterialNewsList struct {
	WechatMaterialCount
	Item []WechatMaterialNews `json:"item"`
}

//WechatMaterialNews MaterialNews
type WechatMaterialNews struct {
	MediaID    string                    `json:"media_id"`
	Content    WechatMaterialNewsContent `json:"content"`
	UpdateTime int64                     `json:"update_time"`
}

//WechatMaterialNewsContent WechatMaterialNewsContent
type WechatMaterialNewsContent struct {
	NewsItem   []WechatNewsItem `json:"news_item"`
	CreateTime int64            `json:"create_time"`
	UpdateTime int64            `json:"update_time"`
}

//WechatNewsItem NewsItemWechat
type WechatNewsItem struct {
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

// WechatMaterialArticle 图文
type WechatMaterialArticle struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     int64  `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
}

// MaterialArticle 图文
type MaterialArticle struct {
	ID               int64         `orm:"column(id)" json:"id"`
	Title            string        `orm:"column(title)" json:"title"`
	ThumbMediaID     string        `orm:"column(thumb_media_id)" json:"thumb_media_id"`
	Author           string        `orm:"column(author)" json:"author"`
	Digest           string        `orm:"column(digest)" json:"digest"`
	ShowCoverPic     int64         `orm:"column(show_cover_pic)" json:"show_cover_pic"`
	Content          string        `orm:"column(content)" json:"content"`
	ContentSourceURL string        `orm:"column(content_source_url)" json:"content_source_url"`
	ThumbURL         string        `orm:"column(thumb_url)" json:"thumb_url"`
	MaterialNews     *MaterialNews `orm:"rel(fk)" json:"-"`
}

// WechatMaterialArticles 图文 图文
type WechatMaterialArticles struct {
	Items []WechatMaterialArticle `json:"articles"`
}

// MaterialNews MaterialNews
type MaterialNews struct {
	ID         int64              `orm:"column(id)" json:"id"`
	MediaID    string             `orm:"column(media_id)" json:"media_id"`
	CreateTime time.Time          `orm:"column(create_time)" json:"create_time"`
	UpdateTime time.Time          `orm:"column(update_time)" json:"update_time"`
	Items      []*MaterialArticle `orm:"reverse(many)" json:"items"`
}

//WechatMaterialMultimediaList WechatMaterialMultimediaList
type WechatMaterialMultimediaList struct {
	WechatMaterialCount
	Item []WechatMultimedia `json:"item"`
}

//WechatMultimedia WechatMultimedia
type WechatMultimedia struct {
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

//WechatMaterialUpdate WechatMaterialUpdate
type WechatMaterialUpdate struct {
	MediaID string           `json:"media_id"`
	Index   int64            `json:"index"`
	Article WechatNewsUpdate `json:"articles"`
}

//MaterialUpdate MaterialUpdate
type MaterialUpdate struct {
	MediaID string
	Index   int64
	Article WechatNewsUpdate
}

//WechatNewsUpdate WechatNewsUpdate
type WechatNewsUpdate struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	ShowCoverPic     int64  `json:"show_cover_pic"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
}

// WechatMaterialInfoResponse 添加素材返回说明
type WechatMaterialInfoResponse struct {
	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

// MaterialMedia MaterialMedia
type MaterialMedia struct {
	ID           int64     `orm:"column(id)" json:"id"`
	Title        string    `orm:"column(title)" json:"title"`
	Introduction string    `orm:"column(introduction)" json:"introduction"`
	MediaType    string    `orm:"column(media_type)" json:"media_type"`
	MediaID      string    `orm:"column(media_id)" json:"media_id"`
	MediaURL     string    `orm:"column(media_url)" json:"media_url"`
	CreateTime   time.Time `orm:"column(create_time)" json:"create_time"`
	UpdateTime   time.Time `orm:"column(update_time)" json:"update_time"`
	Path         string    `orm:"-"`
}

func init() {
	orm.RegisterModel(new(MaterialNews), new(MaterialArticle), new(MaterialMedia))
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

// AddMaterialNews 新增其他类型永久素材
// 通过POST表单来调用接口，表单id为media，包含需要上传的素材内容，有filename、filelength、content-type等信息。请注意：图片素材将进入公众平台官网素材管理模块中的默认分组。
func AddMaterialNews(materialNews *MaterialNews) (v *MaterialNews, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}
	wechatArticles := WechatMaterialArticles{}

	for i, article := range materialNews.Items {
		wechatArticles.Items = append(wechatArticles.Items, WechatMaterialArticle{
			Title:            article.Title,
			ThumbMediaID:     article.ThumbMediaID,
			Author:           article.Author,
			Digest:           article.Digest,
			ShowCoverPic:     article.ShowCoverPic,
			Content:          article.Content,
			ContentSourceURL: article.ContentSourceURL})

		materialNews.Items[i].ID = 0
		var media *MaterialMedia
		media, err = GetMaterialMediaByMediaID(article.ThumbMediaID)
		if err != nil {
			return
		}
		var bytes []byte
		bytes, err = json.Marshal(media)
		fmt.Println("-----------------2--", string(bytes))
		fmt.Println("-----------------1--", article.ThumbMediaID)
		if err == nil {
			article.ThumbURL = media.MediaURL
		} else {
			return
		}
		article.MaterialNews = materialNews
	}
	strURL := materialAddNews + "access_token=" + accessToken

	postData, err := json.Marshal(wechatArticles)
	if err != nil {
		return
	}
	body, err := post(strURL, postData)

	if err != nil {
		fmt.Println(err)
	}
	mediaInfo := WechatMaterialInfoResponse{}
	err = json.Unmarshal(body, &mediaInfo)

	bytes, err := json.Marshal(mediaInfo)
	fmt.Println("-----------------mediaInfo--", string(bytes))
	if mediaInfo.MediaID != "" {
		materialNews.MediaID = mediaInfo.MediaID
	} else {
		err = errors.New("上传次数已经用完")
		return
	}
	materialNews.UpdateTime = time.Now()
	materialNews.CreateTime = time.Now()
	o := orm.NewOrm()
	_, err = o.Insert(materialNews)
	_, err = o.InsertMulti(len(materialNews.Items), materialNews.Items)
	v = materialNews
	return
}

// GetMaterialNewsByID GetMaterialNewsByID
func GetMaterialNewsByID(id int64) (v *MaterialNews, err error) {
	o := orm.NewOrm()
	v = &MaterialNews{ID: id}
	err = o.Read(v)
	_, err = o.LoadRelated(v, "Items")
	if err == nil {
		return v, nil
	}
	return nil, err
}

// GetMaterialNewsByMediaID GetMaterialNewsByID
func GetMaterialNewsByMediaID(mediaID string) (v *MaterialNews, err error) {
	o := orm.NewOrm()
	v = &MaterialNews{MediaID: mediaID}
	err = o.Read(v, "media_id")
	_, err = o.LoadRelated(v, "Items")
	if err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMaterialNews retrieves all Article matches certain condition. Returns empty list if
// no records exist
func GetAllMaterialNews(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MaterialNews))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []MaterialNews
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// DeleteMaterialNewsByID deletes Article by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMaterialNewsByID(id int64) (err error) {
	o := orm.NewOrm()
	v := MaterialNews{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		err = DeleteMaterialByMediaID(v.MediaID)
		if err != nil {
			return
		}
		var num int64
		if num, err = o.Delete(&MaterialNews{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// AddMaterialMedia 新增其他类型永久素材
// 通过POST表单来调用接口，表单id为media，包含需要上传的素材内容，有filename、filelength、content-type等信息。请注意：图片素材将进入公众平台官网素材管理模块中的默认分组。
func AddMaterialMedia(materialMedia *MaterialMedia) (v *MaterialMedia, err error) {
	o := orm.NewOrm()
	v = &MaterialMedia{ID: materialMedia.ID}
	if err = o.Read(v); err == nil {
		return
	}
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}
	strURL := materialAddMaterial + "access_token=" + accessToken + "&type=" + materialMedia.MediaType
	desc := map[string]string{"title": materialMedia.Title, "introduction": materialMedia.Introduction}
	description, err := json.Marshal(desc)
	if err != nil {
		return
	}
	body, err := postFile(strURL, string(description), materialMedia.Path)
	if err != nil {
		fmt.Println(err)
	}
	mediaInfo := WechatMaterialInfoResponse{}
	err = json.Unmarshal(body, &mediaInfo)
	if mediaInfo.MediaID != "" {
		materialMedia.MediaID = mediaInfo.MediaID
		materialMedia.MediaURL = mediaInfo.URL
	} else {
		err = errors.New("上传次数已经用完")
		return
	}
	materialMedia.CreateTime = time.Now()
	materialMedia.UpdateTime = time.Now()
	_, err = o.Insert(materialMedia)
	v = materialMedia
	return
}

// GetMaterialMediaByID GetMaterialMediaByID
func GetMaterialMediaByID(id int64) (v *MaterialMedia, err error) {
	o := orm.NewOrm()
	v = &MaterialMedia{ID: id}
	err = o.Read(v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

// GetMaterialMediaByMediaID retrieves Media by ID. Returns error if
// ID doesn't exist
func GetMaterialMediaByMediaID(mediaID string) (v *MaterialMedia, err error) {
	o := orm.NewOrm()
	v = &MaterialMedia{MediaID: mediaID}
	if err = o.Read(v, "media_id"); err == nil {
		return v, nil
	}
	return nil, err
}

// DeleteMaterialMediaByID deletes Article by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMaterialMediaByID(id int64) (err error) {
	o := orm.NewOrm()
	v := MaterialMedia{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		err = DeleteMaterialByMediaID(v.MediaID)
		if err != nil {
			return
		}
		var num int64
		if num, err = o.Delete(&MaterialMedia{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetAllMaterialMedia retrieves all Article matches certain condition. Returns empty list if
// no records exist
func GetAllMaterialMedia(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MaterialMedia))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []MaterialMedia
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
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

/*
// UploadImageToWechat 上传图文消息内的图片获取URL
//本接口所上传的图片不占用公众号的素材库中图片数量的5000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下。
func UploadImageToWechat(filePath string) (mediaInfo WechatMaterialInfoResponse, err error) {
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

//GetMaterialByMediaID 获取永久素材
func GetMaterialByMediaID(mediaID string) (v *WechatMaterialNewsContent, err error) {
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
	v = &WechatMaterialNewsContent{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

//GetAllMaterialNewsList 永久图文消息素材列表
func GetAllMaterialNewsList(offset int64, count int64) (v *WechatMaterialNewsList, err error) {
	body, err := getAllMaterialListFromWechat("news", offset, count)
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
	}
	return v, err
}

//GetAllMaterialMultimediaList 其他类型（图片、语音、视频）
func GetAllMaterialMultimediaList(materialType string, offset int64, count int64) (v *WechatMaterialMultimediaList, err error) {
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

// GetMaterialcount 获取素材总数
func GetMaterialcount() (v *WechatMaterialTotalCount, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	strURL := materialGetMaterialcount + "access_token=" + accessToken
	body, err := get(strURL)
	v = &WechatMaterialTotalCount{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
*/
