package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["genesis/controllers:FileController"] = append(beego.GlobalControllerRouter["genesis/controllers:FileController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:FileController"] = append(beego.GlobalControllerRouter["genesis/controllers:FileController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MaterialController"] = append(beego.GlobalControllerRouter["genesis/controllers:MaterialController"],
		beego.ControllerComments{
			Method: "GetAllMaterialNewsList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MaterialController"] = append(beego.GlobalControllerRouter["genesis/controllers:MaterialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MaterialController"] = append(beego.GlobalControllerRouter["genesis/controllers:MaterialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:mediaId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MaterialController"] = append(beego.GlobalControllerRouter["genesis/controllers:MaterialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:media_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ReceiveMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:ReceiveMessageController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ReceiveMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:ReceiveMessageController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:WeblogController"] = append(beego.GlobalControllerRouter["genesis/controllers:WeblogController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:WeblogController"] = append(beego.GlobalControllerRouter["genesis/controllers:WeblogController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:WeblogController"] = append(beego.GlobalControllerRouter["genesis/controllers:WeblogController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:WeblogController"] = append(beego.GlobalControllerRouter["genesis/controllers:WeblogController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:WeblogController"] = append(beego.GlobalControllerRouter["genesis/controllers:WeblogController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:WeblogController"] = append(beego.GlobalControllerRouter["genesis/controllers:WeblogController"],
		beego.ControllerComments{
			Method: "PutReviewed",
			Router: `/:id/reviewed`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

}
