// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"genesis/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/genesis/v1",

		beego.NSNamespace("/weblog",
			beego.NSInclude(
				&controllers.WeblogController{},
			),
		),
		beego.NSNamespace("/receiveMessage",
			beego.NSInclude(
				&controllers.ReceiveMessageController{},
			),
		),
		beego.NSNamespace("/sendMessage",
			beego.NSInclude(
				&controllers.SendMessageController{},
			),
		),
		beego.NSNamespace("/file",
			beego.NSInclude(
				&controllers.FileController{},
			),
		),
		beego.NSNamespace("/material",
			beego.NSInclude(
				&controllers.MaterialController{},
			),
		),
		beego.NSNamespace("/menu",
			beego.NSInclude(
				&controllers.MenuController{},
			),
		),
		beego.NSNamespace("/session",
			beego.NSInclude(
				&controllers.SessionController{},
			),
		),
		beego.NSNamespace("/administrator",
			beego.NSInclude(
				&controllers.AdministratorController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
