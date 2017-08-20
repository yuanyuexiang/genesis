package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
	"time"
)

type Token struct {
	Id          int64
	AccessToken string
	ExpiresTime int64
	UpdateTime  int64
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
func init() {
	orm.RegisterModel(new(Token))
}

const (
	Base_url = "https://api.weixin.qq.com/cgi-bin/token?"
)

func GetTokenResponseFromWeChat() (v *TokenResponse, err error) {
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

func GetToken() (v string, err error) {
	o := orm.NewOrm()
	token := &Token{Id: 1}
	if err = o.Read(token); err == nil {
		if token.ExpiresTime-time.Now().Unix() < 0 {
			tokenResponse, err := GetTokenResponseFromWeChat()
			if err != nil {
				return v, err
			}
			token.AccessToken = tokenResponse.AccessToken
			token.UpdateTime = time.Now().Unix()
			token.ExpiresTime = time.Now().Unix() + tokenResponse.ExpiresIn
			UpdateTokenById(token)
		}
		v = token.AccessToken
		return v, nil
	}
	tokenResponse, err := GetTokenResponseFromWeChat()
	if err != nil {
		return v, err
	}
	fmt.Println("token")
	fmt.Println(token)
	fmt.Println("tokenResponse")
	fmt.Println(tokenResponse)
	token.AccessToken = tokenResponse.AccessToken
	token.UpdateTime = time.Now().Unix()
	token.ExpiresTime = time.Now().Unix() + tokenResponse.ExpiresIn
	SaveToken(token)
	v = token.AccessToken
	return v, nil
}

func UpdateTokenById(m *Token) (err error) {
	o := orm.NewOrm()
	v := Token{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func SaveToken(m *Token) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func DeleteToken(id int64) (err error) {
	o := orm.NewOrm()
	v := Token{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Token{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
