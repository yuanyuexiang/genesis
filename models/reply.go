package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

// ReplyDefult ReplyDefult
type ReplyDefult struct {
	Type        string `orm:"column(type);pk" json:"type"`
	ContentType string `orm:"column(content_type)" json:"content_type"`
	Content     string `orm:"column(content)" json:"content"`
}

// ReplyKey ReplyKey
type ReplyKey struct {
	Key         string `orm:"column(key);pk" json:"key"`
	ContentType string `orm:"column(content_type)" json:"content_type"`
	Content     string `orm:"column(content)" json:"content"`
}

func init() {
	orm.RegisterModel(new(ReplyDefult), new(ReplyKey))
}

// AddReplyDefult insert a new Reply into database and returns
// last inserted Id on success.
func AddReplyDefult(m *ReplyDefult) (id int64, err error) {
	o := orm.NewOrm()
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	if created, id, err := o.ReadOrCreate(m, "type"); err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", id)
		} else {
			fmt.Println("Get an object. Id:", id)
		}
	} else {
		return id, err
	}
	return id, nil
}

// GetReplyDefultByType retrieves Reply by Id. Returns error if
// Id doesn't exist
func GetReplyDefultByType(_type string) (v *ReplyDefult, err error) {
	o := orm.NewOrm()
	v = &ReplyDefult{Type: _type}
	if err = o.Read(v, "type"); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateReplyDefultByType updates Reply by ID and returns error if
// the record to be updated doesn't exist
func UpdateReplyDefultByType(m *ReplyDefult) (err error) {
	o := orm.NewOrm()
	v := ReplyDefult{Type: m.Type}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "type"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReplyDefultByType deletes Reply by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReplyDefultByType(_type string) (err error) {
	o := orm.NewOrm()
	v := ReplyDefult{Type: _type}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ReplyDefult{Type: _type}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// AddReplyKey insert a new Reply into database and returns
// last inserted Id on success.
func AddReplyKey(m *ReplyKey) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReplyKeyByKey retrieves Reply by Id. Returns error if
// Id doesn't exist
func GetReplyKeyByKey(key string) (v *ReplyKey, err error) {
	o := orm.NewOrm()
	v = &ReplyKey{Key: key}
	if err = o.Read(v, "key"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReplyKey retrieves all Reply matches certain condition. Returns empty list if
// no records exist
func GetAllReplyKey(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ReplyKey))
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

	var l []ReplyKey
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

// UpdateReplyKeyByKey updates Reply by ID and returns error if
// the record to be updated doesn't exist
func UpdateReplyKeyByKey(m *ReplyKey) (err error) {
	o := orm.NewOrm()
	v := ReplyKey{Key: m.Key}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReplyKeyByKey deletes Reply by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReplyKeyByKey(key string) (err error) {
	o := orm.NewOrm()
	v := ReplyKey{Key: key}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ReplyKey{Key: key}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
