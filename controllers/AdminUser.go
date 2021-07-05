package controllers

import (
	"github.com/astaxie/beego/orm"
	"go-admin-test-system/models"
)

func (c *MainController) AddAdminUser() {
	var maps []orm.Params
	var o = orm.NewOrm()
	username := c.GetString("username")
	permissionRoutesId := c.GetString("permission_routes_id")
	remarks := c.GetString("remarks")
	createName := c.GetString("create_name")
	numSearchUser,errSearchUser := o.Raw("select id from admin_user where username = ?", username).Values(&maps)
	if errSearchUser == nil && numSearchUser > 0 {
		//如果该管理用户已存在
		c.Data["json"] = models.Response{Code: 0, Msg: "该用户已有管理权限，不可重复创建"}
		c.ServeJSON()
		return
	}
	numSearchRoutes,errSearchRoutes := o.Raw("select id from permission_routes where id = ?", permissionRoutesId).Values(&maps)
	if errSearchRoutes == nil && numSearchRoutes > 0 {
		_, errInsertRole := o.Raw("INSERT INTO `admin_user` (`username`,`permission_routes_id`,`remarks`,`create_name`) VALUES (?,?,?,?)", username,permissionRoutesId,remarks, createName).Exec()
		if errInsertRole == nil {
			c.Data["json"] = models.Response{Code: 200, Msg: "创建管理用户成功"}
			c.ServeJSON()
			return
		}
		c.Data["json"] = models.Response{Code: 0, Msg: "创建管理用户失败："+errInsertRole.Error()}
		c.ServeJSON()
	}else {
		c.Data["json"] = models.Response{Code: 0, Msg: "创建管理用户失败，权限路由id不存在"}
		c.ServeJSON()
	}
}

func (c *MainController) EditAdminUser() {
	var maps []orm.Params
	var o = orm.NewOrm()
	username := c.GetString("username")
	permissionRoutesId := c.GetString("permission_routes_id")
	remarks := c.GetString("remarks")
	createName := c.GetString("create_name")
	numSearchUser,errSearchUser := o.Raw("select id from admin_user where username = ?", username).Values(&maps)
	if errSearchUser == nil && numSearchUser > 0 {
		//如果该管理用户已存在
		numSearchRoutes,errSearchRoutes := o.Raw("select id from permission_routes where id = ?", permissionRoutesId).Values(&maps)
		if errSearchRoutes == nil && numSearchRoutes > 0 {
			_, errUpdate := o.Raw("UPDATE `admin_user` SET `permission_routes_id` = ?,`remarks` = ?,`create_name` = ? WHERE `username` = ?",permissionRoutesId,remarks,createName,username).Exec()
			if errUpdate == nil {
				c.Data["json"] = models.Response{Code: 200, Msg: "修改成功"}
				c.ServeJSON()
			}else {
				c.Data["json"] = models.Response{Code: 0, Msg: "修改失败："+errUpdate.Error()}
				c.ServeJSON()
			}
		}else {
			c.Data["json"] = models.Response{Code: 0, Msg: "创建管理用户失败权限路由id不存在"}
			c.ServeJSON()
		}
	}else {
		c.Data["json"] = models.Response{Code: 0, Msg: "用户不存在"}
		c.ServeJSON()
	}
}

func (c *MainController) DelAdminUser()  {
	var maps []orm.Params
	var o = orm.NewOrm()
	username := c.GetString("username")
	num,err := o.Raw("select id from admin_user where username = ?", username).Values(&maps)
	if err == nil && num > 0 {
		//存在则删除
		_, errDel := o.Raw("DELETE FROM admin_user WHERE username = ?",username).Exec()
		if errDel == nil {
			c.Data["json"] = models.Response{Code: 200,Msg: "删除成功"}
			c.ServeJSON()
		}else {
			c.Data["json"] = models.Response{Code: 0,Msg: "删除失败:"+errDel.Error()}
			c.ServeJSON()
		}
		return
	}
	c.Data["json"] = models.Response{Code: 0,Msg: "该管理用户不存在"}
	c.ServeJSON()
}

func (c *MainController) GetAdminUsers() {
	var adminUsers []models.AdminUser
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM admin_user").QueryRows(&adminUsers)
	if err == nil {
		c.Data["json"] = models.AdminUsersRes{Code: 200,Msg: "success",Data: adminUsers}
		c.ServeJSON()
	}else {
		c.Data["json"] = models.Response{Code: 0,Msg: "获取管理权限用户失败："+err.Error()}
		c.ServeJSON()
	}
}

func (c *MainController) GetAdminUserInfo() {
	o := orm.NewOrm()
	username := c.GetString("username")
	var adminUser models.AdminUser
	err := o.Raw("select * from admin_user where username=?", username).QueryRow(&adminUser)
	if err == nil {
		var permissionRoutes models.PermissionRoutes
		errPerRoutes := o.Raw("select * from permission_routes where id=?", adminUser.PermissionRoutesId).QueryRow(&permissionRoutes)
		if errPerRoutes == nil {
			c.Data["json"] = models.AdminUserInfoRes{Code: 200, Msg: "success",AdminUserInfo: adminUser,PermissionRoutes: permissionRoutes}
			c.ServeJSON()
			return
		}
		c.Data["json"] = models.Response{Code: 200,Msg: "NoPermissionRoutes"}
		c.ServeJSON()
		return
	}
	c.Data["json"] = models.Response{Code: 200,Msg: "NoAdminUser"}
	c.ServeJSON()
}
