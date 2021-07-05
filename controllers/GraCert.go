package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/tidwall/gjson"
	"go-admin-test-system/models"
)

func (c *MainController) AddGraStudents() {
	var maps []orm.Params
	var o = orm.NewOrm()
	jsonData := string(c.Ctx.Input.RequestBody)
	graStudents := gjson.Parse(jsonData)
	graStudents.ForEach(func(key, value gjson.Result) bool {
		name := value.Get("name").String()
		xh := value.Get("xh").String()
		IdCard := value.Get("IdCard").String()
		GraYear := value.Get("GraYear").String()
		GraCertId := value.Get("GraCertId").String()
		num,err := o.Raw("select id from gra_students where idcard = ?", IdCard).Values(&maps)
		if num > 0 && err == nil {
			fmt.Println(name,IdCard,"记录已存在")
		}else {
			_, _ = o.Raw("INSERT INTO `gra_students` (`name`,`xh`,`idcard`,`gra_year`,`gra_cert_id`) VALUES (?,?,?,?,?)", name, xh, IdCard, GraYear,GraCertId).Exec()
		}
		return true
	})

	c.Data["json"] = models.Response{Code: 200, Msg: "success"}
	c.ServeJSON()
}

func (c *MainController) EditGraStudent() {
	var maps []orm.Params
	var o = orm.NewOrm()
	jsonData := string(c.Ctx.Input.RequestBody)
	id := gjson.Get(jsonData,"id").String()
	name := gjson.Get(jsonData,"name").String()
	xh := gjson.Get(jsonData,"xh").String()
	IdCard := gjson.Get(jsonData,"IdCard").String()
	GraYear := gjson.Get(jsonData,"GraYear").String()
	GraCertId := gjson.Get(jsonData,"GraCertId").String()
	num,err := o.Raw("select * from gra_students where id = ?", id).Values(&maps)
	if err == nil && num > 0 {
		_, errUpdate := o.Raw("UPDATE `gra_students` SET `name` = ?,`xh` = ?,`idcard` = ?,`gra_year` = ?,`gra_cert_id` = ? WHERE `id` = ?",name,xh,IdCard,GraYear,GraCertId,id).Exec()
		if errUpdate == nil {
			c.Data["json"] = models.Response{Code: 200, Msg: "修改成功"}
			c.ServeJSON()
		}else {
			c.Data["json"] = models.Response{Code: 0, Msg: "修改失败："+errUpdate.Error()}
			c.ServeJSON()
		}
	}else {
		c.Data["json"] = models.Response{Code: 0, Msg: "修改失败，信息不存在"}
		c.ServeJSON()
	}
}

func (c *MainController) DelGraStudent()  {
	var maps []orm.Params
	var o = orm.NewOrm()
	jsonData := string(c.Ctx.Input.RequestBody)
	id := gjson.Get(jsonData,"id").String()
	num,err := o.Raw("select * from gra_students where id = ?", id).Values(&maps)
	if err == nil && num > 0 {
		//存在则删除
		_, errDel := o.Raw("DELETE FROM gra_students WHERE id = ?",id).Exec()
		if errDel == nil {
			c.Data["json"] = models.Response{Code: 200,Msg: "删除成功"}
			c.ServeJSON()
		}else {
			c.Data["json"] = models.Response{Code: 0,Msg: "删除失败:"+errDel.Error()}
			c.ServeJSON()
		}
		return
	}
	c.Data["json"] = models.Response{Code: 0,Msg: "删除失败，数据不存在"}
	c.ServeJSON()
}

func (c *MainController) GetGraStudents() {
	var graStudents []graStudent
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM gra_students").QueryRows(&graStudents)
	if err == nil {
		c.Data["json"] = graStudentsRes{Code: 200,Msg: "success",Data: graStudents}
		c.ServeJSON()
	}else {
		c.Data["json"] = models.Response{Code: 0,Msg: "获取数据失败，请重试："+err.Error()}
		c.ServeJSON()
	}
}

type graStudent struct {
	Id 		   string `json:"id"   orm:"column(id)"`
	Name       string `json:"name"   orm:"column(name)"`
	Xh         string `json:"xh"  orm:"column(xh)"`
	IdCard     string `json:"IdCard"  orm:"column(idcard)"`
	GraYear    string `json:"GraYear"  orm:"column(gra_year)"`
	GraCertId  string `json:"GraCertId"  orm:"column(gra_cert_id)"`
}

type graStudentsRes struct {
	Code int			`json:"code"`
	Msg  string			`json:"msg"`
	Data []graStudent 	`json:"data"`
}