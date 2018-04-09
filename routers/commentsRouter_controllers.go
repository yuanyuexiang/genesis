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

	beego.GlobalControllerRouter["genesis/controllers:AdministratorController"] = append(beego.GlobalControllerRouter["genesis/controllers:AdministratorController"],
		beego.ControllerComments{
			Method: "PutName",
			Router: `/:id/name`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AdministratorController"] = append(beego.GlobalControllerRouter["genesis/controllers:AdministratorController"],
		beego.ControllerComments{
			Method: "PutPassword",
			Router: `/:id/password`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AdministratorController"] = append(beego.GlobalControllerRouter["genesis/controllers:AdministratorController"],
		beego.ControllerComments{
			Method: "PutPhoneNumber",
			Router: `/:id/phoneNumber`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AdministratorController"] = append(beego.GlobalControllerRouter["genesis/controllers:AdministratorController"],
		beego.ControllerComments{
			Method: "PutRole",
			Router: `/:id/role`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AdministratorController"] = append(beego.GlobalControllerRouter["genesis/controllers:AdministratorController"],
		beego.ControllerComments{
			Method: "PutPhoneStatus",
			Router: `/:id/status`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"],
		beego.ControllerComments{
			Method: "PostAllAnnouncement",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"],
		beego.ControllerComments{
			Method: "DeleteAllAnnouncement",
			Router: `/:msgID/:articleIDX`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"],
		beego.ControllerComments{
			Method: "CheckAllAnnouncement",
			Router: `/:msgID/status`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"],
		beego.ControllerComments{
			Method: "PostAllSendImageMessage",
			Router: `/image`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"],
		beego.ControllerComments{
			Method: "PostAllSendNewsMessage",
			Router: `/news`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"],
		beego.ControllerComments{
			Method: "PostPreviewMessage",
			Router: `/preview`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"],
		beego.ControllerComments{
			Method: "PostAllSendTextMessage",
			Router: `/text`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["genesis/controllers:AnnouncementController"],
		beego.ControllerComments{
			Method: "PostAllSendVoiceMessage",
			Router: `/voice`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ArticleController"] = append(beego.GlobalControllerRouter["genesis/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ArticleController"] = append(beego.GlobalControllerRouter["genesis/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ArticleController"] = append(beego.GlobalControllerRouter["genesis/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ArticleController"] = append(beego.GlobalControllerRouter["genesis/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ArticleController"] = append(beego.GlobalControllerRouter["genesis/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ArticleController"] = append(beego.GlobalControllerRouter["genesis/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "PutReviewStatus",
			Router: `/:id/reviewStatus`,
			AllowHTTPMethods: []string{"put"},
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
			Method: "GetAllMaterialNewsList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
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

	beego.GlobalControllerRouter["genesis/controllers:MaterialController"] = append(beego.GlobalControllerRouter["genesis/controllers:MaterialController"],
		beego.ControllerComments{
			Method: "GetMaterialCount",
			Router: `/count`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MaterialController"] = append(beego.GlobalControllerRouter["genesis/controllers:MaterialController"],
		beego.ControllerComments{
			Method: "PostNews",
			Router: `/news`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MaterialController"] = append(beego.GlobalControllerRouter["genesis/controllers:MaterialController"],
		beego.ControllerComments{
			Method: "GetOneNews",
			Router: `/news/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MediaController"] = append(beego.GlobalControllerRouter["genesis/controllers:MediaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MediaController"] = append(beego.GlobalControllerRouter["genesis/controllers:MediaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MediaController"] = append(beego.GlobalControllerRouter["genesis/controllers:MediaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MediaController"] = append(beego.GlobalControllerRouter["genesis/controllers:MediaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MediaController"] = append(beego.GlobalControllerRouter["genesis/controllers:MediaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MediaController"] = append(beego.GlobalControllerRouter["genesis/controllers:MediaController"],
		beego.ControllerComments{
			Method: "GetOneFile",
			Router: `/:id/file`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:MediaController"] = append(beego.GlobalControllerRouter["genesis/controllers:MediaController"],
		beego.ControllerComments{
			Method: "PutReviewStatus",
			Router: `/:id/reviewStatus`,
			AllowHTTPMethods: []string{"put"},
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

	beego.GlobalControllerRouter["genesis/controllers:PermissionController"] = append(beego.GlobalControllerRouter["genesis/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:PermissionController"] = append(beego.GlobalControllerRouter["genesis/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:PermissionController"] = append(beego.GlobalControllerRouter["genesis/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:PermissionController"] = append(beego.GlobalControllerRouter["genesis/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:PermissionController"] = append(beego.GlobalControllerRouter["genesis/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ReplyController"] = append(beego.GlobalControllerRouter["genesis/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "PostDefult",
			Router: `/defult`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ReplyController"] = append(beego.GlobalControllerRouter["genesis/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "DeleteDefult",
			Router: `/defult/:type`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ReplyController"] = append(beego.GlobalControllerRouter["genesis/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "GetOneDefult",
			Router: `/defult/:type`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ReplyController"] = append(beego.GlobalControllerRouter["genesis/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "PutDefult",
			Router: `/defult/:type`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ReplyController"] = append(beego.GlobalControllerRouter["genesis/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "GetAllKey",
			Router: `/key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ReplyController"] = append(beego.GlobalControllerRouter["genesis/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "PostKey",
			Router: `/key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ReplyController"] = append(beego.GlobalControllerRouter["genesis/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "DeleteKey",
			Router: `/key/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ReplyController"] = append(beego.GlobalControllerRouter["genesis/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "GetOneKey",
			Router: `/key/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:ReplyController"] = append(beego.GlobalControllerRouter["genesis/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "PutKey",
			Router: `/key/:key`,
			AllowHTTPMethods: []string{"put"},
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

	beego.GlobalControllerRouter["genesis/controllers:WechatRequestMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:WechatRequestMessageController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["genesis/controllers:WechatRequestMessageController"] = append(beego.GlobalControllerRouter["genesis/controllers:WechatRequestMessageController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
