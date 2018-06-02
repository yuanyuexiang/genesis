package models

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
)

var (
	heavenSecret string
)

const (
	pattern = "((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\\d{8}"
)

func init() {
	heavenSecret = beego.AppConfig.String("heaven.secret")
}

// GoHeaven GoHeaven
func GoHeaven(fromUserName, content string) {
	if strings.Contains(content, heavenSecret) && heavenSecret != "" {
		u, err := GetUserByOpenID(fromUserName)
		if err != nil {
			return
		}
		u.Type = "admin"
		err = UpdateUserTypeByOpenID(u)
		if err != nil {
			return
		}
		r, _ := regexp.Compile(pattern)
		phoneNumber := r.FindString(content)
		fmt.Println(phoneNumber)
		if phoneNumber != "" {
			administrator, err := GetAdministratorByPhoneNumber(phoneNumber)
			if err == nil {
				if administrator.OpenID != u.OpenID {
					administrator.OpenID = u.OpenID
					UpdateAdministratorOpenIDByID(administrator)
				}
			} else {
				administrator := &Administrator{Name: u.NickName,
					PhoneNumber: phoneNumber,
					Password:    "e10adc3949ba59abbe56e057f20f883e",
					Status:      true,
					Role:        "editor",
					OpenID:      u.OpenID}
				AddAdministrator(administrator)
			}
		} else if strings.Contains(content, "取消发布") {
			StopAllAnnouncementTimingSendMessage()
		}
	}
}
