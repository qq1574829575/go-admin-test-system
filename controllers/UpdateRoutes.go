package controllers

import (
	"github.com/astaxie/beego/orm"
	"go-admin-test-system/models"
)

func (c *MainController) UpdateRoutes() {
	var maps []orm.Params
	var o = orm.NewOrm()
	id := c.GetString("id")
	routes := c.GetString("routes")
	num, errSearch := o.Raw("select * from `routes` where id = ?", id).Values(&maps)
	if errSearch == nil && num > 0 {
		_, errUpdate := o.Raw("UPDATE `routes` SET `routes` = ? WHERE `id` = ?",routes,id).Exec()
		if errUpdate == nil {
			c.Data["json"] = models.Response{Code: 200, Msg: "操作成功"}
			c.ServeJSON()
		}else {
			c.Data["json"] = models.Response{Code: 0, Msg: "操作失败："+errUpdate.Error()}
			c.ServeJSON()
		}
	}else {
		c.Data["json"] = models.Response{Code: 0, Msg: "路由不存在"}
		c.ServeJSON()
	}
}

func (c *MainController) GetRoutes() {
	var maps []orm.Params
	o := orm.NewOrm()
	id := c.GetString("id")
	num, errSearch := o.Raw("select * from `routes` where id = ?", id).Values(&maps)
	if errSearch == nil && num > 0 {
		routes := maps[0]["routes"]
		c.Data["json"] = RoutesRes{Code: 200,Msg: "success", Routes: routes.(string)}
		c.ServeJSON()
	}else {
		c.Data["json"] = models.Response{Code: 0, Msg: "路由不存在"}
		c.ServeJSON()
	}
}

type RoutesRes struct {
	Code int   		`json:"code"`
	Msg string 		`json:"msg"`
	Routes string 	`json:"routes"`
}