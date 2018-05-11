package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// Media Media
type Media struct {
	ID           int64     `orm:"column(id)" json:"id"`
	URL          string    `orm:"column(url)" json:"url"`
	Type         string    `orm:"column(type)" json:"type"`
	Title        string    `orm:"column(title)" json:"title"`
	Introduction string    `orm:"column(introduction)" json:"introduction"`
	ReviewStatus bool      `orm:"column(review_status)" json:"review_status"`
	CreateTime   time.Time `orm:"column(create_time)" json:"create_time"`
	UpdateTime   time.Time `orm:"column(update_time)" json:"update_time"`
}

func init() {
	orm.RegisterModel(new(Media))
}

// AddMedia insert a new Media into database and returns
// last inserted ID on success.
func AddMedia(m *Media) (v *Media, err error) {
	o := orm.NewOrm()
	m.UpdateTime = time.Now()
	m.CreateTime = time.Now()
	_, err = o.Insert(m)
	v = m
	return
}

// GetMediaByID retrieves Media by ID. Returns error if
// ID doesn't exist
func GetMediaByID(id int64) (v *Media, err error) {
	o := orm.NewOrm()
	v = &Media{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMedia retrieves all Media matches certain condition. Returns empty list if
// no records exist
func GetAllMedia(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Media))
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

	var l []Media
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

// UpdateMediaByID updates Media by ID and returns error if
// the record to be updated doesn't exist
func UpdateMediaByID(m *Media) (err error) {
	o := orm.NewOrm()
	v := Media{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// UpdateMediaReviewStatusByID updates Media by ID and returns error if
// the record to be updated doesn't exist
func UpdateMediaReviewStatusByID(m *Media) (err error) {
	o := orm.NewOrm()
	v := Media{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		m.UpdateTime = time.Now()
		if num, err = o.Update(m, "review_status", "update_time"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMedia deletes Media by ID and returns error if
// the record to be deleted doesn't exist
func DeleteMedia(id int64) (err error) {
	o := orm.NewOrm()
	v := Media{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Media{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// ChangeMediaURL ChangeMediaURL
func ChangeMediaURL(m *Media, url string) *Media {
	if m != nil {
		m.URL = url + "/file"
	}
	return m
}

// ChangeMediaListURL ChangeMediaListURL
func ChangeMediaListURL(mediaList []interface{}, url string) []interface{} {
	newMediaList := []interface{}{}
	if mediaList != nil && len(mediaList) > 0 {
		for _, item := range mediaList {
			m := item.(Media)
			m.URL = url + "/" + strconv.FormatInt(m.ID, 10) + "/file"
			newMediaList = append(newMediaList, m)
		}
	} else {
		return nil
	}
	return newMediaList
}
