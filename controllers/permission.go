package controllers

import (
	"github.com/astaxie/beego/orm"
	"go-admin-test-system/models"
)

func (c *MainController) AddPermissionRoutes() {
	var o = orm.NewOrm()
	desc := c.GetString("desc")
	routes := c.GetString("routes")
	createName := c.GetString("create_name")
	_, errInsertRole := o.Raw("INSERT INTO `permission_routes` (`desc`,`routes`,`create_name`) VALUES (?,?,?)", desc,routes,createName).Exec()
	if errInsertRole == nil {
		c.Data["json"] = models.Response{Code: 200, Msg: "创建路由权限成功"}
		c.ServeJSON()
		return
	}
	c.Data["json"] = models.Response{Code: 0, Msg: "创建路由权限失败："+errInsertRole.Error()}
	c.ServeJSON()
}

func (c *MainController) EditPermissionRoutes() {
	var maps []orm.Params
	var o = orm.NewOrm()
	id := c.GetString("id")
	desc := c.GetString("desc")
	routes := c.GetString("routes")
	createName := c.GetString("create_name")
	num, errSearch := o.Raw("select * from `permission_routes` where id = ?", id).Values(&maps)
	if errSearch == nil && num > 0 {
		_, errUpdate := o.Raw("UPDATE `permission_routes` SET `desc` = ?,`routes` = ?,`create_name` = ? WHERE `id` = ?",desc,routes,createName,id).Exec()
		if errUpdate == nil {
			c.Data["json"] = models.Response{Code: 200, Msg: "修改路由权限成功"}
			c.ServeJSON()
		}else {
			c.Data["json"] = models.Response{Code: 0, Msg: "修改路由权限失败："+errUpdate.Error()}
			c.ServeJSON()
		}
	}else {
		c.Data["json"] = models.Response{Code: 0, Msg: "路由权限不存在！"}
		c.ServeJSON()
	}
}

func (c *MainController) DelPermissionRoutes()  {
	var maps []orm.Params
	var o = orm.NewOrm()
	id := c.GetString("id")
	num,err := o.Raw("select id from permission_routes where id = ?", id).Values(&maps)
	if err == nil && num > 0 {
		//存在则删除
		_, errDel := o.Raw("DELETE FROM permission_routes WHERE id = ?",id).Exec()
		if errDel == nil {
			c.Data["json"] = models.Response{Code: 200,Msg: "删除成功"}
			c.ServeJSON()
		}else {
			c.Data["json"] = models.Response{Code: 0,Msg: "删除失败:"+errDel.Error()}
			c.ServeJSON()
		}
		return
	}
	c.Data["json"] = models.Response{Code: 0,Msg: "删除失败，路由权限不存在"}
	c.ServeJSON()
}

func (c *MainController) GetAllPermissionRoutes() {
	var allPermissionRoutes []models.PermissionRoutes
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM permission_routes").QueryRows(&allPermissionRoutes)
	if err == nil {
		c.Data["json"] = models.AllPermissionRoutesRes{Code: 200,Msg: "success",Data: allPermissionRoutes}
		c.ServeJSON()
	}else {
		c.Data["json"] = models.Response{Code: 0,Msg: "获取管理权限用户失败："+err.Error()}
		c.ServeJSON()
	}
}

func (c *MainController) GetPermissionRoutes() {
	o := orm.NewOrm()
	id := c.GetString("id")
	var permissionRoutes models.PermissionRoutes
	_ = o.Raw("select * from permission_routes where id = ?", id).QueryRow(&permissionRoutes)
	c.Data["json"] = models.PermissionRoutesRes{Code: 200, Msg: "success",Data: permissionRoutes}
	c.ServeJSON()
}