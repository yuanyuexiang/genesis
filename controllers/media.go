package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"genesis/models"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
)

// MediaController for Media
type MediaController struct {
	beego.Controller
}

// URLMapping URLMapping
func (c *MediaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Media
// @Param	body		body 	models.Media	true		"body for Media content"
// @Success 201 {int} models.Media
// @Failure 403 body is empty
// @router / [post]
func (c *MediaController) Post() {
	f, h, err := c.GetFile("file")
	introduction := c.GetString("introduction")
	title := c.GetString("title")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	filename := h.Filename
	fmt.Println(h.Size)
	contentType := h.Header.Get("Content-Type")
	contentType = contentType[0:strings.Index(contentType, "/")]
	filePath := getFilePath(contentType, filename)
	fmt.Println(filePath)
	c.SaveToFile("file", filePath)
	v := models.Media{URL: filePath, Type: contentType, Title: title, Introduction: introduction}
	fmt.Println(v)
	mediaInfo, err := models.AddMedia(&v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
		c.ServeJSON()
		return
	}
	c.Data["json"] = models.GetReturnData(0, "OK", mediaInfo)
	c.ServeJSON()
}

// @Title Get
// @Description get Media by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Media
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MediaController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetMediaByID(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Media
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Media
// @Failure 403
// @router / [get]
func (c *MediaController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllMedia(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the Media
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Media	true		"body for Media content"
// @Success 200 {object} models.Media
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MediaController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Media{ID: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateMediaByID(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Media
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MediaController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteMedia(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func getFilePath(contentType, fileName string) (filePath string) {
	dirPath := "static/files/" + contentType + "/"
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+dirPath, os.ModePerm) //生成多级目录
	if err != nil {
		fmt.Println(err)
	}
	u1 := uuid.NewV4()
	postName := fileName[strings.Index(fileName, "."):len(fileName)]
	filePath = dirPath + u1.String() + postName
	return
}
