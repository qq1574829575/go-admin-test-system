package models

type PermissionRoutes struct {
	Id       string `json:"id"   orm:"column(id)"`
	Desc 	 string `json:"desc"  orm:"column(desc)"`
	Routes   string `json:"routes"  orm:"column(routes)"`
	CreateName string `json:"create_name"  orm:"column(create_name)"`
}

type PermissionRoutesRes struct {
	Code int   		`json:"code"`
	Msg string 		`json:"msg"`
	Data PermissionRoutes 	`json:"data"`
}

type AllPermissionRoutesRes struct {
	Code int					`json:"code"`
	Msg  string					`json:"msg"`
	Data []PermissionRoutes 	`json:"data"`
}
