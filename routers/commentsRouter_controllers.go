package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["genesis/controllers:AdministratorController"] = append(beego.GlobalControllerRouter["genesis/controllers:AdministratorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AdministratorController"] = append(beego.GlobalControllerRouter["genesis/controllers:AdministratorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AdministratorController"] = append(beego.GlobalControllerRouter["genesis/controllers:AdministratorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AdministratorController"] = append(beego.GlobalControllerRouter["genesis/controllers:AdministratorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AdministratorController"] = append(beego.GlobalControllerRouter["genesis/controllers:AdministratorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

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
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
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
			Method: "Put",
			Router: `/:mediaId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MaterialController"] = append(beego.GlobalControllerRouter["genesis/controllers:MaterialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:mediaId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MaterialController"] = append(beego.GlobalControllerRouter["genesis/controllers:MaterialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:media_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MaterialController"] = append(beego.GlobalControllerRouter["genesis/controllers:MaterialController"],
		beego.ControllerComments{
			Method: "GetMaterialCount",
			Router: `/count`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MaterialController"] = append(beego.GlobalControllerRouter["genesis/controllers:MaterialController"],
		beego.ControllerComments{
			Method: "PostFile",
			Router: `/image`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MenuController"] = append(beego.GlobalControllerRouter["genesis/controllers:MenuController"],
		beego.ControllerComments{
			Method: "CreateMenu",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MenuController"] = append(beego.GlobalControllerRouter["genesis/controllers:MenuController"],
		beego.ControllerComments{
			Method: "GetMenu",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MenuController"] = append(beego.GlobalControllerRouter["genesis/controllers:MenuController"],
		beego.ControllerComments{
			Method: "DeleteMenu",
			Router: `/[delete]`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MenuController"] = append(beego.GlobalControllerRouter["genesis/controllers:MenuController"],
		beego.ControllerComments{
			Method: "AddConditionalMenu",
			Router: `/conditional`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MenuController"] = append(beego.GlobalControllerRouter["genesis/controllers:MenuController"],
		beego.ControllerComments{
			Method: "DeleteConditionalMenu",
			Router: `/conditional`,
			AllowHTTPMethods: []string{"delete"},
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

	beego.GlobalControllerRouter["genesis/controllers:SendMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:SendMessageController"],
		beego.ControllerComments{
			Method: "PostAllSendMessage",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:SendMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:SendMessageController"],
		beego.ControllerComments{
			Method: "DeleteAllSendMessage",
			Router: `/:msgID/:articleIDX`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:SendMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:SendMessageController"],
		beego.ControllerComments{
			Method: "CheckAllSendMessage",
			Router: `/:msgID/status`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:SendMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:SendMessageController"],
		beego.ControllerComments{
			Method: "PostAllSendImageMessage",
			Router: `/image`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:SendMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:SendMessageController"],
		beego.ControllerComments{
			Method: "UploadNewsMessageImage",
			Router: `/image/uplaod`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:SendMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:SendMessageController"],
		beego.ControllerComments{
			Method: "PostAllSendNewsMessage",
			Router: `/news`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:SendMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:SendMessageController"],
		beego.ControllerComments{
			Method: "UploadNewsMessage",
			Router: `/news/uplaod`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:SendMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:SendMessageController"],
		beego.ControllerComments{
			Method: "PostPreviewMessage",
			Router: `/preview`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:SendMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:SendMessageController"],
		beego.ControllerComments{
			Method: "PostAllSendTextMessage",
			Router: `/text`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:SendMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:SendMessageController"],
		beego.ControllerComments{
			Method: "PostAllSendVoiceMessage",
			Router: `/voice`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:SessionController"] = append(beego.GlobalControllerRouter["genesis/controllers:SessionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:SessionController"] = append(beego.GlobalControllerRouter["genesis/controllers:SessionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
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
