package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	uuid "github.com/satori/go.uuid"
)

// Session 用户会话
type Session struct {
	ID         int64  `orm:"column(id);auto"`
	Token      string `orm:"column(token)"`
	CreateTime int64  `orm:"column(create_time)"`
	UpdateTime int64  `orm:"column(update_time)"`
}

// AuthInfo 用户名及密码
type AuthInfo struct {
	PhoneNumber string
	Password    string
}

func init() {
	orm.RegisterModel(new(Session))
}

// CreateSession insert a new Session into database and returns
// last inserted Id on success.
func CreateSession(m *AuthInfo) (session *Session, err error) {
	administrator, err := GetAdministratorByPhoneNumber(m)
	if err != nil {
		return nil, errors.New("NO User")
	}
	if m.Password != administrator.Password {
		return nil, errors.New("Password Error")
	}
	u1 := uuid.NewV4()
	session = &Session{Token: u1.String(), CreateTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	o := orm.NewOrm()
	_, err = o.Insert(session)
	fmt.Println(session)
	return session, nil
}

// CheckSessionByToken updates Session by Id and returns error if
// the record to be updated doesn't exist
func CheckSessionByToken(token string) (err error) {
	o := orm.NewOrm()
	v := Session{Token: token}
	// ascertain id exists in the database
	if err = o.Read(&v, "Token"); err == nil {
		x := time.Now().Unix() - v.UpdateTime
		fmt.Println(x)
		if x > 600 {
			err = errors.New("Time Out")
		} else {
			var num int64
			v.UpdateTime = time.Now().Unix()
			if num, err = o.Update(&v); err == nil {
				fmt.Println("Number of records updated in database:", num)
			}
		}
	}
	return
}

// DeleteSessionByToken deletes Session by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSessionByToken(token string) (err error) {
	o := orm.NewOrm()
	v := Session{Token: token}
	// ascertain id exists in the database
	if err = o.Read(&v, "Token"); err == nil {
		var num int64
		if num, err = o.Delete(&Session{Token: token}, "Token"); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
