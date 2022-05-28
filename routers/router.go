package routers

import (
	"training_project/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// beego.Router("/v1/user", &controllers.UserController{}, "get:GetAll")
	beego.Router("/v1/partner/:uid([0-9]+)", &controllers.PartnerController{}, "get:Get")
	beego.Router("/v1/partner/", &controllers.PartnerController{}, "post:Post")

}
