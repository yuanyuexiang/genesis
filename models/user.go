package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
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

type User struct {
	Id            int64
	Subscribe     byte
	Openid        string
	Sex           byte
	Language      string
	City          string
	Province      string
	Country       string
	Headimgurl    string
	SubscribeTime int64
	Unionid       string
	Remark        string
	Groupid       byte
	UserType      byte
}

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
	BaseUserInfoUrl = "https://api.weixin.qq.com/cgi-bin/user/info?"
)

func GetUserWechat(openid string) (v *UserWechat, err error) {

	access_token, err := GetToken()
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	str_request := "access_token=" + access_token + "&openid=" + openid + "&lang=zh_CN "
	str_url := BaseUserInfoUrl + str_request
	request, err := http.NewRequest("GET", str_url, nil)
	if err != nil {
		return
	} else {
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
	}
	return v, err
}

func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

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
