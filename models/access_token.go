package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//Token Token
type Token struct {
	ID          int64  `orm:"column(id);auto"`
	AccessToken string `orm:"column(access_token)"`
	ExpiresTime int64  `orm:"column(expires_time)"`
	UpdateTime  int64  `orm:"column(update_time)"`
}

//TokenResponse TokenResponse
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
func init() {
	orm.RegisterModel(new(Token))
}

const (
	baseURL = "https://api.weixin.qq.com/cgi-bin/token?"
)

//GetTokenResponseFromWeChat GetTokenResponseFromWeChat
func GetTokenResponseFromWeChat() (v *TokenResponse, errReturn error) {
	client := &http.Client{}
	strRequest := "grant_type=client_credential&appid=" + beego.AppConfig.String("appid") + "&secret=" + beego.AppConfig.String("secret")
	strURL := baseURL + strRequest
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

		err = json.Unmarshal(body, &v)
		if err != nil {
			fmt.Println(err)
			errReturn = err
		}
		if v.AccessToken == "" {
			err = errors.New(string(body))
			errReturn = err
		}
	} else {
		body, err := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		if err != nil {
			errReturn = err
			fmt.Println(err)
		}
	}
	return
}

//GetToken GetToken
func GetToken() (v string, err error) {
	o := orm.NewOrm()
	token := &Token{ID: 1}
	if err = o.Read(token); err == nil {
		if token.ExpiresTime-time.Now().Unix() < 0 {
			tokenResponse, err := GetTokenResponseFromWeChat()
			if err != nil {
				return v, err
			}
			token.AccessToken = tokenResponse.AccessToken
			token.UpdateTime = time.Now().Unix()
			token.ExpiresTime = time.Now().Unix() + tokenResponse.ExpiresIn
			UpdateTokenByID(token)
		}
		v = token.AccessToken
		return v, nil
	}
	tokenResponse, err := GetTokenResponseFromWeChat()
	if err != nil {
		return v, err
	}
	token.AccessToken = tokenResponse.AccessToken
	token.UpdateTime = time.Now().Unix()
	token.ExpiresTime = time.Now().Unix() + tokenResponse.ExpiresIn
	SaveToken(token)
	v = token.AccessToken
	return v, nil
}

//UpdateTokenByID UpdateTokenByID
func UpdateTokenByID(m *Token) (err error) {
	o := orm.NewOrm()
	v := Token{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//SaveToken SaveToken
func SaveToken(m *Token) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//DeleteToken DeleteToken
func DeleteToken(id int64) (err error) {
	o := orm.NewOrm()
	v := Token{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Token{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
