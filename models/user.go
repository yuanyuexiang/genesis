package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

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
	ID            int64  `orm:"column(id);auto" json:"id"`
	Subscribe     byte   `orm:"column(subscribe)" json:"subscribe"`
	OpenID        string `orm:"column(openid)" json:"openid"`
	NickName      string `orm:"column(nickname)" json:"nickname"`
	Sex           byte   `orm:"column(sex)" json:"sex"`
	City          string `orm:"column(city)" json:"city"`
	Country       string `orm:"column(country)" json:"country"`
	Province      string `orm:"column(province)" json:"province"`
	Language      string `orm:"column(language)" json:"language"`
	HeadImgurl    string `orm:"column(headimgurl)" json:"headimgurl"`
	SubscribeTime int64  `orm:"column(subscribe_time)" json:"subscribe_time"`
	UnionID       string `orm:"column(unionid)" json:"unionid"`
	Remark        string `orm:"column(remark)" json:"remark"`
	GroupID       byte   `orm:"column(groupid)" json:"groupid"`
	Type          string `orm:"column(type)" json:"type"` //admin
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

// GetAllUserCount retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUserCount(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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
					return 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return 0, errors.New("Error: unused 'order' fields")
		}
	}

	qs = qs.OrderBy(sortFields...)
	count, err = qs.Count()
	return
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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

	var l []User
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
