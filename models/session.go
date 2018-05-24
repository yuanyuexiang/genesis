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
	ID              int64  `orm:"column(id);auto"`
	Token           string `orm:"column(token)"`
	AdministratorID int64  `orm:"column(administrator_id)"`
	CreateTime      int64  `orm:"column(create_time)"`
	UpdateTime      int64  `orm:"column(update_time)"`
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
	if !administrator.Status {
		return nil, errors.New("frozen-in")
	}
	u1 := uuid.NewV4()
	session = &Session{Token: u1.String(), AdministratorID: administrator.ID, CreateTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	o := orm.NewOrm()
	_, err = o.Insert(session)
	if err != nil {
		return nil, errors.New("Insert fial")
	}
	fmt.Println(session)
	return session, nil
}

// CheckSessionByToken updates Session by Id and returns error if
// the record to be updated doesn't exist
func CheckSessionByToken(token string) (err error) {
	if token == "" {
		err = errors.New("bad request no token")
		return
	}
	o := orm.NewOrm()
	v := Session{Token: token}
	// ascertain id exists in the database
	if err = o.Read(&v, "Token"); err == nil {
		x := time.Now().Unix() - v.UpdateTime
		if x > 6000 {
			err = errors.New("Token Timeout")
		} else {
			var num int64
			v.UpdateTime = time.Now().Unix()
			if num, err = o.Update(&v); err == nil {
				fmt.Println("Number of records updated in database:", num)
			}
		}
	} else {
		err = errors.New("NO Token")
	}
	return
}

// GetSessionByToken updates Session by Id and returns error if
// the record to be updated doesn't exist
func GetSessionByToken(token string) (v *Session, err error) {
	if token == "" {
		err = errors.New("bad request")
		return
	}
	o := orm.NewOrm()
	v = &Session{Token: token}
	// ascertain id exists in the database
	if err = o.Read(v, "Token"); err == nil {
		x := time.Now().Unix() - v.UpdateTime
		if x > 600 {
			err = errors.New("Token Timeout")
		} else {
			var num int64
			v.UpdateTime = time.Now().Unix()
			if num, err = o.Update(v); err == nil {
				fmt.Println("Number of records updated in database:", num)
			}
		}
	} else {
		err = errors.New("NO Token")
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
