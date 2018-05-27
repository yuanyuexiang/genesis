package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"genesis/utils"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	wechatBaseURL = "https://api.weixin.qq.com/cgi-bin/"
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
/*
//AnnouncementItem AnnouncementItem
type AnnouncementItem struct {
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Title            string `json:"title"`
	ContentSourceURL string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
	ShowCoverPic     string `json:"show_cover_pic"`
}

//Announcements Announcements
type Announcements struct {
	Announcements []AnnouncementItem `json:"articles"`
}


//UploadNewsMessageImage 上传图文消息内的图片获取URL【订阅号与服务号认证后均可用】
/*
http请求方式: POST
https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
调用示例（使用curl命令，用FORM表单方式上传一个图片）：
curl -F media=@test.jpg "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN"
*/ /*
func UploadNewsMessageImage(filePath string) (data map[string]string, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}

	strRequest := "media/uploadimg?access_token=" + accessToken
	strURL := wechatBaseURL + strRequest

	if err != nil {
		return
	}
	body, err := postFile(strURL, "", filePath)

	if err != nil {
		fmt.Println(err)
	}
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}
	return
}

//UploadNewsMessage 上传图文消息素材【订阅号与服务号认证后均可用】
/*
http请求方式: POST
https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
*/ /*
func UploadNewsMessage(articles *Announcements) (data map[string]interface{}, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	strRequest := "media/uploadnews?access_token=" + accessToken
	strURL := wechatBaseURL + strRequest
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
*/
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

// Filter Filter
type Filter struct {
	IsToAll bool  `json:"is_to_all"`
	TagID   int64 `json:"tag_id"`
}

//MpNews MpNews
type MpNews struct {
	MediaID string `json:"media_id"`
}

//MpVideo MpVideo
type MpVideo struct {
	MediaID string `json:"media_id"`
}

//Music Music
type Music struct {
	MediaID string `json:"media_id"`
}

//Text Text
type Text struct {
	Content string `json:"content"`
}

//Voice Voice
type Voice struct {
	MediaID string `json:"media_id"`
}

//Image Image
type Image struct {
	MediaID string `json:"media_id"`
}

//WXCard WXCard
type WXCard struct {
	CardID string `json:"card_id"`
}

//AnnouncementResponseMessage AnnouncementResponseMessage
type AnnouncementResponseMessage struct {
	ErrCode   int64  `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	MsgID     int64  `json:"msg_id"`
	MsgDataID int64  `json:"msg_data_id"`
}

//AnnouncementMessageStatus AnnouncementMessageStatus
type AnnouncementMessageStatus struct {
	MsgID     int64  `json:"msg_id"`
	MsgStatus string `json:"msg_status"`
}

//AllSendNewsMessage AllSendNewsMessage
type AllSendNewsMessage struct {
	Filter            Filter `json:"filter"`
	MpNews            MpNews `json:"mpnews"`
	MsgType           string `json:"msgtype"`
	SendIgnoreReprint int64  `json:"send_ignore_reprint"`
}

//AllSendTextMessage AllSendTextMessage
type AllSendTextMessage struct {
	Filter  Filter `json:"filter"`
	Text    Text   `json:"text"`
	MsgType string `json:"msgtype"`
}

//AllSendVoiceMessage AllSendVoiceMessage
type AllSendVoiceMessage struct {
	Filter  Filter `json:"filter"`
	Voice   Voice  `json:"voice"`
	MsgType string `json:"msgtype"`
}

//AllSendImageMessage AllSendImageMessage
type AllSendImageMessage struct {
	Filter  Filter `json:"filter"`
	Image   Image  `json:"image"`
	MsgType string `json:"msgtype"`
}

//AllSendMusicMessage AllSendNewsMessage
type AllSendMusicMessage struct {
	Filter  Filter `json:"filter"`
	Music   Music  `json:"music"`
	MsgType string `json:"msgtype"`
}

//AllSendMpVideoMessage AllSendNewsMessage
type AllSendMpVideoMessage struct {
	Filter  Filter  `json:"filter"`
	MpVideo MpVideo `json:"mpvideo"`
	MsgType string  `json:"msgtype"`
}

//AllSendWXCardMessage AllSendNewsMessage
type AllSendWXCardMessage struct {
	Filter  Filter `json:"filter"`
	WXCard  WXCard `json:"wxcard"`
	MsgType string `json:"msgtype"`
}

//PreviewNewsMessage PreviewNewsMessage
type PreviewNewsMessage struct {
	ToUser            string `json:"touser"`
	MpNews            MpNews `json:"mpnews"`
	MsgType           string `json:"msgtype"`
	SendIgnoreReprint int64  `json:"send_ignore_reprint"`
}

//PreviewTextMessage PreviewTextMessage
type PreviewTextMessage struct {
	ToUser  string `json:"touser"`
	Text    Text   `json:"text"`
	MsgType string `json:"msgtype"`
}

//PreviewVoiceMessage PreviewVoiceMessage
type PreviewVoiceMessage struct {
	ToUser  string `json:"touser"`
	Voice   Voice  `json:"voice"`
	MsgType string `json:"msgtype"`
}

//PreviewImageMessage PreviewImageMessage
type PreviewImageMessage struct {
	ToUser  string `json:"touser"`
	Image   Image  `json:"image"`
	MsgType string `json:"msgtype"`
}

//PreviewMusicMessage PreviewMusicMessage
type PreviewMusicMessage struct {
	ToUser  string `json:"touser"`
	Music   Music  `json:"music"`
	MsgType string `json:"msgtype"`
}

//PreviewMpVideoMessage PreviewMpVideoMessage
type PreviewMpVideoMessage struct {
	ToUser  string  `json:"touser"`
	MpVideo MpVideo `json:"mpvideo"`
	MsgType string  `json:"msgtype"`
}

//PreviewWXCardMessage PreviewWXCardMessage
type PreviewWXCardMessage struct {
	ToUser  string `json:"touser"`
	WXCard  WXCard `json:"wxcard"`
	MsgType string `json:"msgtype"`
}

// Announcement Announcement
type Announcement struct {
	ID          int64     `orm:"column(id);auto" json:"id"`
	IsToAll     bool      `orm:"column(is_to_all)" json:"is_to_all"`
	TagID       int64     `orm:"column(tag_id)" json:"tag_id"`
	MsgType     string    `orm:"column(msgtype)" json:"msgtype"`
	Content     string    `orm:"column(content)" json:"content"`
	ErrCode     int64     `orm:"column(errcode)" json:"errcode"`
	ErrMsg      string    `orm:"column(errmsg)" json:"errmsg"`
	MsgID       int64     `orm:"column(msg_id)" json:"msg_id"`
	MsgDataID   int64     `orm:"column(msg_data_id)" json:"msg_data_id"`
	Status      int64     `orm:"column(status)" json:"status"` // 0 准备 1 成功 -1 取消
	CreateTime  time.Time `orm:"column(create_time)" json:"create_time"`
	PublishTime time.Time `orm:"column(publish_time)" json:"publish_time"`
}

var (
	//AnnouncementTimer AnnouncementTimer
	AnnouncementTimer map[int64]*time.Timer
)

func init() {
	orm.RegisterModel(new(Announcement))
	AnnouncementTimer = make(map[int64]*time.Timer)
	go restoreTimingTask()
}

func restoreTimingTask() {
	time.Sleep(100 * time.Millisecond)
	o := orm.NewOrm()
	qs := o.QueryTable(new(Announcement))
	var l []Announcement
	if _, err := qs.Filter("status", 0).All(&l); err == nil {
		for _, m := range l {
			now := time.Now()
			next := m.PublishTime
			duration := next.Sub(now)
			if duration < 0 {
				m.Status = -1
				UpdateAnnouncementByID(&m)
			} else {
				timingSendMessage(&m)
			}
		}
	}
}

func previewSendMessage(a *Announcement) (err error) {
	l, err := ListAllAdministrator()
	if err == nil {
		for _, v := range l {
			if v.OpenID != "" {
				requestData, err := getPrevieMessageFromAnnouncement(v.OpenID, a)
				data, err := PostPreviewMessage(requestData)
				utils.Println(data)
				fmt.Println(err)
			}
		}
	}
	return
}

func timingSendMessage(a *Announcement) (err error) {
	now := time.Now()
	next := a.PublishTime
	duration := next.Sub(now)
	if duration < 0 {
		err = errors.New("TIME ERROR")
		return
	}
	var requestData interface{}
	requestData, err = getAllSendMessageFromAnnouncement(a)
	go func(m Announcement, r interface{}) {
		now := time.Now()
		next := m.PublishTime
		duration := next.Sub(now)
		fmt.Println(m)
		if duration >= 0 {
			utils.Println(m)
			t := time.NewTimer(duration)
			AnnouncementTimer[m.ID] = t
			<-t.C
			delete(AnnouncementTimer, m.ID)
			data, err := PostAllSendMessage(requestData)
			fmt.Println(err)
			m.ErrCode = data.ErrCode
			m.ErrMsg = data.ErrMsg
			m.MsgID = data.MsgID
			m.MsgDataID = data.MsgDataID
			m.Status = 1
			err = UpdateAnnouncementByID(&m)
			fmt.Println(err)
			utils.Println(r)
		}
	}(*a, requestData)
	return
}

//stopAnnouncementTimingSendMessage stopAnnouncementTimingSendMessage
func stopAnnouncementTimingSendMessage(id int64) (err error) {
	if t, ok := AnnouncementTimer[id]; ok {
		t.Stop()
		delete(AnnouncementTimer, id)
	} else {
		err = errors.New("NO TIMER")
	}
	return
}

// AddAnnouncement insert a new Announcement into database and returns
// last inserted Id on success.
func AddAnnouncement(m *Announcement) (id int64, err error) {
	m.IsToAll = true
	m.TagID = 0
	m.Status = 0
	m.CreateTime = time.Now()
	o := orm.NewOrm()
	id, err = o.Insert(m)
	err = previewSendMessage(m)
	if err == nil {
		err = timingSendMessage(m)
	}
	return
}

// GetAnnouncementByID retrieves Announcement by Id. Returns error if
// Id doesn't exist
func GetAnnouncementByID(id int64) (v *Announcement, err error) {
	o := orm.NewOrm()
	v = &Announcement{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAnnouncement retrieves all Announcement matches certain condition. Returns empty list if
// no records exist
func GetAllAnnouncement(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Announcement))
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

	var l []Announcement
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

// UpdateAnnouncementByID updates Announcement by Id and returns error if
// the record to be updated doesn't exist
func UpdateAnnouncementByID(m *Announcement) (err error) {
	o := orm.NewOrm()
	v := Announcement{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		m.CreateTime = time.Now()
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// UpdateAnnouncementStatusByID updates Announcement by Id and returns error if
// the record to be updated doesn't exist
func UpdateAnnouncementStatusByID(m *Announcement) (err error) {
	o := orm.NewOrm()
	v := Announcement{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if m.Status == -1 && v.Status == 0 {
			if err = stopAnnouncementTimingSendMessage(m.ID); err == nil {
				v.Status = -1
				if num, err = o.Update(&v); err == nil {
					fmt.Println("Number of records updated in database:", num)
				}
			}
		} else if m.Status == 1 && v.Status == 0 {
			if err = stopAnnouncementTimingSendMessage(m.ID); err == nil {
				requestData, err := getAllSendMessageFromAnnouncement(m)
				data, err := PostAllSendMessage(requestData)
				fmt.Println(err)
				v.ErrCode = data.ErrCode
				v.ErrMsg = data.ErrMsg
				v.MsgID = data.MsgID
				v.MsgDataID = data.MsgDataID
				v.Status = 1
				if num, err = o.Update(&v); err == nil {
					fmt.Println("Number of records updated in database:", num)
				}
			}
		} else {
			err = errors.New("STATUS NO CHANGE")
		}
	}
	return
}

// DeleteAnnouncement deletes Announcement by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAnnouncement(id int64) (err error) {
	o := orm.NewOrm()
	v := Announcement{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Announcement{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//PostAllSendMessage 根据标签进行群发【订阅号与服务号认证后均可用】
func PostAllSendMessage(requestData interface{}) (data AnnouncementResponseMessage, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	strRequest := "message/mass/sendall?access_token=" + accessToken
	strURL := wechatBaseURL + strRequest
	fmt.Println(strURL)
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
func DeleteAllSendMessage(msgID, articleIDX int64) (data AnnouncementResponseMessage, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	strRequest := "message/mass/delete?access_token=" + accessToken
	strURL := wechatBaseURL + strRequest
	requestData := map[string]int64{"msg_id": msgID, "article_idx": articleIDX}
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

//PostPreviewMessage 预览接口【订阅号与服务号认证后均可用】
func PostPreviewMessage(requestData interface{}) (data AnnouncementResponseMessage, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	strRequest := "message/mass/preview?access_token=" + accessToken
	strURL := wechatBaseURL + strRequest
	fmt.Println(strURL)
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
func CheckAllSendMessage(msgID int64) (data AnnouncementMessageStatus, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	strRequest := "message/mass/get?access_token=" + accessToken
	strURL := wechatBaseURL + strRequest
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

func getAllSendMessageFromAnnouncement(m *Announcement) (requestData interface{}, err error) {
	if m.MsgType == "mpnews" {
		requestData = AllSendNewsMessage{
			Filter:            Filter{IsToAll: m.IsToAll, TagID: m.MsgID},
			MpNews:            MpNews{MediaID: m.Content},
			MsgType:           m.MsgType,
			SendIgnoreReprint: 1}
	} else if m.MsgType == "text" {
		requestData = AllSendTextMessage{
			Filter:  Filter{IsToAll: m.IsToAll, TagID: m.MsgID},
			Text:    Text{Content: m.Content},
			MsgType: m.MsgType}
	} else if m.MsgType == "mpvideo" {
		requestData = AllSendVoiceMessage{
			Filter:  Filter{IsToAll: m.IsToAll, TagID: m.MsgID},
			Voice:   Voice{MediaID: m.Content},
			MsgType: m.MsgType}
	} else if m.MsgType == "music" {
		requestData = AllSendMusicMessage{
			Filter:  Filter{IsToAll: m.IsToAll, TagID: m.MsgID},
			Music:   Music{MediaID: m.Content},
			MsgType: m.MsgType}
	} else if m.MsgType == "image" {
		requestData = AllSendImageMessage{
			Filter:  Filter{IsToAll: m.IsToAll, TagID: m.MsgID},
			Image:   Image{MediaID: m.Content},
			MsgType: m.MsgType}
	} else if m.MsgType == "video" {
		requestData = AllSendMpVideoMessage{
			Filter:  Filter{IsToAll: m.IsToAll, TagID: m.MsgID},
			MpVideo: MpVideo{MediaID: m.Content},
			MsgType: m.MsgType}
	} else if m.MsgType == "wxcard" {
		requestData = AllSendWXCardMessage{
			Filter:  Filter{IsToAll: m.IsToAll, TagID: m.MsgID},
			WXCard:  WXCard{CardID: m.Content},
			MsgType: m.MsgType}
	} else {
		err = errors.New("DATA ERROR")
	}
	return
}

func getPrevieMessageFromAnnouncement(touser string, m *Announcement) (requestData interface{}, err error) {
	if m.MsgType == "mpnews" {
		requestData = PreviewNewsMessage{
			ToUser:            touser,
			MpNews:            MpNews{MediaID: m.Content},
			MsgType:           m.MsgType,
			SendIgnoreReprint: 1}
	} else if m.MsgType == "text" {
		requestData = PreviewTextMessage{
			ToUser:  touser,
			Text:    Text{Content: m.Content},
			MsgType: m.MsgType}
	} else if m.MsgType == "voice" {
		requestData = PreviewVoiceMessage{
			ToUser:  touser,
			Voice:   Voice{MediaID: m.Content},
			MsgType: m.MsgType}
	} else if m.MsgType == "music" {
		requestData = PreviewMusicMessage{
			ToUser:  touser,
			Music:   Music{MediaID: m.Content},
			MsgType: m.MsgType}
	} else if m.MsgType == "image" {
		requestData = PreviewImageMessage{
			ToUser:  touser,
			Image:   Image{MediaID: m.Content},
			MsgType: m.MsgType}
	} else if m.MsgType == "mpvideo" {
		requestData = PreviewMpVideoMessage{
			ToUser:  touser,
			MpVideo: MpVideo{MediaID: m.Content},
			MsgType: m.MsgType}
	} else if m.MsgType == "wxcard" {
		requestData = PreviewWXCardMessage{
			ToUser:  touser,
			WXCard:  WXCard{CardID: m.Content},
			MsgType: m.MsgType}
	} else {
		err = errors.New("DATA ERROR")
	}
	return
}
