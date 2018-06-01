package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Article 福音文章
type Article struct {
	ID               int64     `orm:"column(id);auto" json:"id"`
	Title            string    `orm:"column(title)" json:"title"`
	ThumbMediaID     string    `orm:"column(thumb_media_id)" json:"thumb_media_id"`
	ShowCoverPic     int64     `orm:"column(show_cover_pic)" json:"show_cover_pic"`
	Author           string    `orm:"column(author)" json:"author"`
	Digest           string    `orm:"column(digest)" json:"digest"`
	Content          string    `orm:"column(content)" json:"content"`
	ContentSourceURL string    `orm:"column(content_source_url)" json:"content_source_url"`
	ThumbID          int64     `orm:"column(thumb_id)" json:"thumb_id"`
	ThumbURL         string    `orm:"column(thumb_url)" json:"thumb_url"`
	ReviewStatus     int64     `orm:"column(review_status)" json:"review_status"` //1通过0未审查-1未通过
	CreateTime       time.Time `orm:"column(create_time)" json:"create_time"`
	UpdateTime       time.Time `orm:"column(update_time)" json:"update_time"`
}

func init() {
	orm.RegisterModel(new(Article))
}

// AddArticle insert a new Article into database and returns
// last inserted Id on success.
func AddArticle(m *Article) (id int64, err error) {
	o := orm.NewOrm()
	m.UpdateTime = time.Now()
	m.CreateTime = time.Now()
	id, err = o.Insert(m)
	return
}

// GetArticleByID retrieves Article by Id. Returns error if
// Id doesn't exist
func GetArticleByID(id int64) (v *Article, err error) {
	o := orm.NewOrm()
	v = &Article{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllArticle retrieves all Article matches certain condition. Returns empty list if
// no records exist
func GetAllArticle(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
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

	var l []Article
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

// UpdateArticleByID updates Article by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticleByID(m *Article) (err error) {
	o := orm.NewOrm()
	v := Article{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		m.CreateTime = time.Now()
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// UpdateArticleReviewedByID UpdateArticleReviewedByID
func UpdateArticleReviewedByID(m *Article) (err error) {
	o := orm.NewOrm()
	v := Article{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var materialMedia *MaterialMedia
		materialMedia, err = GetMaterialMediaByID(v.ThumbID)
		if err != nil && m.ReviewStatus == 1 {
			var media *Media
			media, err = GetMediaByID(v.ThumbID)
			if err != nil {
				return
			}
			materialMedia = &MaterialMedia{ID: media.ID, Path: media.URL, Title: media.Title, Introduction: media.Introduction, MediaType: "thumb"}
			materialMedia, err = AddMaterialMedia(materialMedia)
			if err != nil {
				return
			}
		}
		m.ThumbMediaID = materialMedia.MediaID
		m.CreateTime = time.Now()
		var num int64
		if num, err = o.Update(m, "review_status", "thumb_media_id", "update_time"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArticle deletes Article by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArticle(id int64) (err error) {
	o := orm.NewOrm()
	v := Article{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Article{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
