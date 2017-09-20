package controllers

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("/files", "static")
}

// oprations for File
type FileController struct {
	beego.Controller
}

func (c *FileController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create File
// @Param	body		body 	models.File	true		"body for File content"
// @Success 201 {int} models.File
// @Failure 403 body is empty
// @router / [post]
func (c *FileController) Post() {
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	filePath := "static/files/" + h.Filename
	fmt.Println(filePath)
	c.SaveToFile("uploadname", filePath) // 保存位置在 static/upload, 没有文件夹要先创建
	returnData := map[string]string{"filePath": filePath}
	c.Data["json"] = returnData
	c.ServeJSON()
}

// @Title Delete
// @Description delete the File
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FileController) Delete() {
	c.Data["json"] = "OK"

	c.ServeJSON()
}
