package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
)

type AccessToken struct {
	Id          int64  `orm:"auto"`
	AccessToken string `orm:"size(128)"`
	ExpiresIn   int64  `orm:"size(128)"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
func init() {
	orm.RegisterModel(new(AccessToken))
}

const (
	Base_url = "https://api.weixin.qq.com/cgi-bin/token?"
)

func GetAccessTokenFromWeChat() (v *AccessTokenResponse, err error) {
	client := &http.Client{}
	str_request := "grant_type=client_credential&appid=" + beego.AppConfig.String("appid") + "&secret=" + beego.AppConfig.String("secret")
	str_url := Base_url + str_request
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
			err = json.Unmarshal(body, v)
			if err != nil {
				fmt.Println(bodystr)
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

func SaveAccessToken(m *AccessToken) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetAccessTokenById(id int64) (v *AccessToken, err error) {
	o := orm.NewOrm()
	v = &AccessToken{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func UpdateAccessTokenById(m *AccessToken) (err error) {
	o := orm.NewOrm()
	v := AccessToken{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteAccessToken(id int64) (err error) {
	o := orm.NewOrm()
	v := AccessToken{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AccessToken{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
