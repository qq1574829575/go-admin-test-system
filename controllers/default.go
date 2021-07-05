package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kirinlabs/HttpRequest"
	"github.com/tidwall/gjson"
	utils "go-admin-test-system/utils"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetGraCertUrl() {
	jsonData := string(c.Ctx.Input.RequestBody)
	year := gjson.Get(jsonData,"year").String()
	xm := gjson.Get(jsonData,"name").String()
	zsbh := gjson.Get(jsonData,"CertId").String()
	req := HttpRequest.NewRequest()
	resp, _ := req.Get("http://zzzs.jxedu.gov.cn/zzquery", nil)
	if resp.StatusCode() == 200 {
		var cookies = map[string]string{}
		for _,val:=range resp.Cookies(){
			name:=val.Name
			value:=val.Value
			cookies[name] = value
		}

		body, _ := resp.Body()
		RequestVerificationToken := utils.GetBetweenStr(string(body),"<input name=\"__RequestVerificationToken\" type=\"hidden\" value=\"","\"")

		req.SetCookies(cookies)
		postData := "year="+year+"&xm="+xm+"&zsbh="+zsbh+"&__RequestVerificationToken="+RequestVerificationToken
		resp, _ := req.Post("http://zzzs.jxedu.gov.cn/zzquery",postData)
		body_, _ := resp.Body()
		if strings.Index(string(body_),"输入信息有误") == -1 {
			url:="http://zzzs.jxedu.gov.cn"+utils.GetBetweenStr(string(body_),"<iframe id=\"ifrPic\" src=\"","\"")
			c.Data["json"] = CertRes{Code: 200, Msg: "获取成功",Url: url}
			c.ServeJSON()
		}else {
			c.Data["json"] = CertRes{Code: 0, Msg: "输入信息有误"}
			c.ServeJSON()
		}
	}
	c.Data["json"] = CertRes{Code: 0, Msg: "获取证书信息失败"}
	c.ServeJSON()
}

type CertRes struct {
	Code int	`json:"code"`
	Msg string	`json:"msg"`
	Url string	`json:"url"`
}
