package models

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// Permission 权限
type Permission struct {
	ID          int64  `orm:"column(id);auto"`
	Role        string `orm:"column(role)"`
	Action      string `orm:"column(action)"`
	Resource    string `orm:"column(resource)"`
	Description string `orm:"column(description)"`
	CreateTime  int64  `orm:"column(create_time)"`
	UpdateTime  int64  `orm:"column(update_time)"`
}

func init() {
	orm.RegisterModel(new(Permission))
}

// CheckPermission updates Session by Id and returns error if
// the record to be updated doesn't exist
func CheckPermission(m *Permission) (err error) {
	o := orm.NewOrm()
	fmt.Println(m)
	qs := o.QueryTable(new(Permission))
	var l []Permission
	qs.Filter("Role", m.Role).Filter("Action", m.Action)
	if _, err = qs.All(&l, "Role", "Action", "Resource"); err != nil {
		err = errors.New("no permission")
	} else {
		for _, v := range l {
			m, _ := regexp.MatchString(v.Resource, m.Resource)
			if m {
				return
			}
		}
		err = errors.New("no permission")
	}
	return
}

// AddPermission insert a new Permission into database and returns
// last inserted Id on success.
func AddPermission(m *Permission) (id int64, err error) {
	o := orm.NewOrm()
	m.UpdateTime = time.Now().Unix()
	m.CreateTime = time.Now().Unix()
	id, err = o.Insert(m)
	return
}

// GetPermissionByID retrieves Permission by Id. Returns error if
// Id doesn't exist
func GetPermissionByID(id int64) (v *Permission, err error) {
	o := orm.NewOrm()
	v = &Permission{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPermission retrieves all Permission matches certain condition. Returns empty list if
// no records exist
func GetAllPermission(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Permission))
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

	var l []Permission
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

// UpdatePermissionByID updates Permission by Id and returns error if
// the record to be updated doesn't exist
func UpdatePermissionByID(m *Permission) (err error) {
	o := orm.NewOrm()
	v := Permission{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		m.UpdateTime = time.Now().Unix()
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePermission deletes Permission by Id and returns error if
// the record to be deleted doesn't exist
func DeletePermission(id int64) (err error) {
	o := orm.NewOrm()
	v := Permission{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Permission{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
