package routers

import (
	"github.com/astaxie/beego"
	"go-admin-test-system/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //管理用户接口
	beego.Router("/AddAdminUser",&controllers.MainController{},"post:AddAdminUser")
	beego.Router("/EditAdminUser",&controllers.MainController{},"post:EditAdminUser")
	beego.Router("/DelAdminUser",&controllers.MainController{},"post:DelAdminUser")
	beego.Router("/GetAdminUsers",&controllers.MainController{},"get:GetAdminUsers")
	beego.Router("/GetAdminUserInfo",&controllers.MainController{},"post:GetAdminUserInfo")
    //路由权限接口
	beego.Router("/AddPermissionRoutes",&controllers.MainController{},"post:AddPermissionRoutes")
	beego.Router("/EditPermissionRoutes",&controllers.MainController{},"post:EditPermissionRoutes")
	beego.Router("/DelPermissionRoutes",&controllers.MainController{},"post:DelPermissionRoutes")
	beego.Router("/GetAllPermissionRoutes",&controllers.MainController{},"get:GetAllPermissionRoutes")
	beego.Router("/GetPermissionRoutes",&controllers.MainController{},"post:GetPermissionRoutes")
    //路由接口
	beego.Router("/UpdateRoutes",&controllers.MainController{},"post:UpdateRoutes")
	beego.Router("/GetRoutes",&controllers.MainController{},"post:GetRoutes")
}
