package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

// Administrator Administrator
type Administrator struct {
	ID          int64  `orm:"column(id);auto"`
	Name        string `orm:"column(name)"`
	PhoneNumber string `orm:"column(phone_number)"`
	Password    string `orm:"column(password)" json:"Password,omitempty"`
	Status      bool   `orm:"column(status)"`
	Role        string `orm:"column(role)"`
}

func init() {
	orm.RegisterModel(new(Administrator))
}

// AddAdministrator insert a new Administrator into database and returns
// last inserted Id on success.
func AddAdministrator(m *Administrator) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAdministratorByPhoneNumber retrieves Administrator by Id. Returns error if
// Id doesn't exist
func GetAdministratorByPhoneNumber(m *AuthInfo) (v *Administrator, err error) {
	o := orm.NewOrm()
	v = &Administrator{PhoneNumber: m.PhoneNumber}
	fmt.Println(m)
	if err = o.Read(v, "PhoneNumber"); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAdministratorByID retrieves Administrator by Id. Returns error if
// Id doesn't exist
func GetAdministratorByID(id int64) (v *Administrator, err error) {
	o := orm.NewOrm()
	v = &Administrator{ID: id}
	if err = o.Read(v); err == nil {
		v.Password = ""
		return v, nil
	}
	return nil, err
}

// GetAllAdministrator retrieves all Administrator matches certain condition. Returns empty list if
// no records exist
func GetAllAdministrator(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (total int64, ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Administrator))
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
					return 0, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return 0, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return 0, nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return 0, nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Administrator
	qs = qs.OrderBy(sortFields...)

	total, err = qs.Count()
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				v.Password = ""
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				v.Password = ""
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return total, ml, nil
	}
	return 0, nil, err
}

// UpdateAdministratorByID updates Administrator by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdministratorByID(m *Administrator) (err error) {
	o := orm.NewOrm()
	v := Administrator{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Name", "PhoneNumber", "Role", "Status"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// UpdateAdministratorRoleByID updates Administrator by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdministratorRoleByID(m *Administrator) (err error) {
	o := orm.NewOrm()
	v := Administrator{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Role"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// UpdateAdministratorPasswordByID updates Administrator by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdministratorPasswordByID(m *Administrator) (err error) {
	o := orm.NewOrm()
	v := Administrator{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Password"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//UpdateAdministratorStatusByID UpdateAdministratorStatusByID
func UpdateAdministratorStatusByID(m *Administrator) (err error) {
	o := orm.NewOrm()
	v := Administrator{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Status"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//UpdateAdministratorNameByID UpdateAdministratorNameByID
func UpdateAdministratorNameByID(m *Administrator) (err error) {
	o := orm.NewOrm()
	v := Administrator{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "Name"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//UpdateAdministratorPhoneNumberByID UpdateAdministratorPhoneNumberByID
func UpdateAdministratorPhoneNumberByID(m *Administrator) (err error) {
	o := orm.NewOrm()
	v := Administrator{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "PhoneNumber"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdministrator deletes Administrator by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdministrator(id int64) (err error) {
	o := orm.NewOrm()
	v := Administrator{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Administrator{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
