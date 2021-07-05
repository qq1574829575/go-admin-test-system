package models

type AdminUser struct {
	UserName             string `json:"username"   orm:"column(username)"`
	PermissionRoutesId 	 string `json:"permission_routes_id"  orm:"column(permission_routes_id)"`
	AssPerRoutesId		 string `json:"ass_per_routes_id"  orm:"column(ass_per_routes_id)"`
	Remarks              string `json:"remarks"  orm:"column(remarks)"`
	CreateName           string `json:"create_name"  orm:"column(create_name)"`
}

type AdminUserInfoRes struct {
	Code int   							`json:"code"`
	Msg string 							`json:"msg"`
	AdminUserInfo AdminUser 			`json:"admin_userinfo"`
	PermissionRoutes PermissionRoutes 	`json:"permission_routes"`
}

type AdminUsersRes struct {
	Code int			`json:"code"`
	Msg  string			`json:"msg"`
	Data []AdminUser 	`json:"data"`
}
