package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego/orm"
)

/**
{
   "subscribe": 1,
   "openid": "o6_bmjrPTlm6_2sgVt7hMZOPfL2M",
   "nickname": "Band",
   "sex": 1,
   "language": "zh_CN",
   "city": "广州",
   "province": "广东",
   "country": "中国",
   "headimgurl":  "http://wx.qlogo.cn/mmopen/eMsv84eavHiaiceqxibJxCfHe/0",
   "subscribe_time": 1382694957,
   "unionid": " o6_bmasdasdsad6_2sgVt7hMZOPfL"
   "remark": "",
   "groupid": 0,
   "tagid_list":[128,2]
}
**/

//User User
type User struct {
	ID            int64  `orm:"column(id);auto"`
	Subscribe     byte   `orm:"column(subscribe)"`
	Openid        string `orm:"column(openid)"`
	Sex           byte   `orm:"column(sex)"`
	Language      string `orm:"column(language)"`
	City          string `orm:"column(city)"`
	Province      string `orm:"column(province)"`
	Country       string `orm:"column(country)"`
	Headimgurl    string `orm:"column(headimgurl)"`
	SubscribeTime int64  `orm:"column(subscribe_timeid)"`
	Unionid       string `orm:"column(unionid)"`
	Remark        string `orm:"column(remark)"`
	Groupid       byte   `orm:"column(groupid)"`
	UserType      byte   `orm:"column(user_typeid)"`
}

//UserWechat UserWechat
type UserWechat struct {
	Subscribe     byte   `json:"subscribe"`
	Openid        string `json:"openid"`
	Sex           byte   `json:"sex"`
	Language      string `json:"language"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Country       string `json:"country"`
	Headimgurl    string `json:"headimgurl"`
	SubscribeTime int64  `json:"subscribe_time"`
	Unionid       string `json:"unionid"`
	Remark        string `json:"remark"`
	Groupid       byte   `json:"groupid"`
}

func init() {
	orm.RegisterModel(new(User))
}

//接口调用请求说明
//http请求方式: GET https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
const (
	//BaseUserInfoURL BaseUserInfoURL
	BaseUserInfoURL = "https://api.weixin.qq.com/cgi-bin/user/info?"
)

//GetUserWechat GetUserWechat
func GetUserWechat(openid string) (v *UserWechat, err error) {
	accessToken, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	strRequest := "access_token=" + accessToken + "&openid=" + openid + "&lang=zh_CN "
	strURL := BaseUserInfoURL + strRequest
	request, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		return
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)

		bodystr := string(body)
		fmt.Println(bodystr)
		err = json.Unmarshal(body, &v)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		body, err := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		if err != nil {
			fmt.Println(err)
		}
	}

	return v, err
}

//AddUser AddUser
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//GetUserByOpenid GetUserByOpenid
func GetUserByOpenid(openid string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Openid: openid}
	if err = o.Read(v); err == nil {
		return v, nil
	}

	userWechat, err := GetUserWechat(openid)
	if err == nil {
		v.City = userWechat.City
		v.Country = userWechat.Country
		v.Groupid = userWechat.Groupid
		v.Headimgurl = userWechat.Headimgurl
		v.Language = userWechat.Language
		v.Openid = userWechat.Openid
		v.Province = userWechat.Province
		v.Remark = userWechat.Remark
		v.Sex = userWechat.Sex
		v.Subscribe = userWechat.Subscribe
		v.SubscribeTime = userWechat.SubscribeTime
		v.Unionid = userWechat.Unionid
		AddUser(v)
		return v, nil
	}
	return nil, err
}

//UpdateUserByOpenid UpdateUserByOpenid
func UpdateUserByOpenid(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Openid: m.Openid}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//DeleteUser DeleteUser
func DeleteUser(openid string) (err error) {
	o := orm.NewOrm()
	v := User{Openid: openid}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Openid: openid}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
