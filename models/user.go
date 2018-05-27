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
	OpenID        string `orm:"column(openid)"`
	NickName      string `orm:"column(nickname)"`
	Sex           byte   `orm:"column(sex)"`
	City          string `orm:"column(city)"`
	Country       string `orm:"column(country)"`
	Province      string `orm:"column(province)"`
	Language      string `orm:"column(language)"`
	HeadImgurl    string `orm:"column(headimgurl)"`
	SubscribeTime int64  `orm:"column(subscribe_time)"`
	UnionID       string `orm:"column(unionid)"`
	Remark        string `orm:"column(remark)"`
	GroupID       byte   `orm:"column(groupid)"`
	Type          string `orm:"column(type)"` //admin
}

//UserWechat UserWechat
type UserWechat struct {
	Subscribe     byte   `json:"subscribe"`
	OpenID        string `json:"openid"`
	NickName      string `json:"nickname"`
	Sex           byte   `json:"sex"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Province      string `json:"province"`
	Language      string `json:"language"`
	HeadImgurl    string `json:"headimgurl"`
	SubscribeTime int64  `json:"subscribe_time"`
	UnionID       string `json:"unionid"`
	Remark        string `json:"remark"`
	GroupID       byte   `json:"groupid"`
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

//GetUserByOpenID GetUserByOpenID
func GetUserByOpenID(openid string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{OpenID: openid}
	if err = o.Read(v, "openid"); err == nil {
		return v, nil
	}

	userWechat, err := GetUserWechat(openid)
	if err == nil {
		v.Subscribe = userWechat.Subscribe
		v.OpenID = userWechat.OpenID
		v.NickName = userWechat.NickName
		v.Sex = userWechat.Sex
		v.City = userWechat.City
		v.Country = userWechat.Country
		v.Province = userWechat.Province
		v.Language = userWechat.Language
		v.HeadImgurl = userWechat.HeadImgurl
		v.SubscribeTime = userWechat.SubscribeTime
		v.UnionID = userWechat.UnionID
		v.Remark = userWechat.Remark
		v.GroupID = userWechat.GroupID
		AddUser(v)
		return v, nil
	}
	return nil, err
}

//UpdateUserByOpenID UpdateUserByOpenID
func UpdateUserByOpenID(m *User) (err error) {
	o := orm.NewOrm()
	v := User{OpenID: m.OpenID}
	// ascertain id exists in the database
	if err = o.Read(&v, "openid"); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//UpdateUserTypeByOpenID UpdateUserTypeByOpenID
func UpdateUserTypeByOpenID(m *User) (err error) {
	o := orm.NewOrm()
	v := User{OpenID: m.OpenID}
	// ascertain id exists in the database
	if err = o.Read(&v, "openid"); err == nil {
		var num int64
		if num, err = o.Update(m, "type"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//DeleteUser DeleteUser
func DeleteUser(openid string) (err error) {
	o := orm.NewOrm()
	v := User{OpenID: openid}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{OpenID: openid}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
