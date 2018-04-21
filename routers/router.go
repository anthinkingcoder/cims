package routers

import (
	"github.com/astaxie/beego"
	"cims/controllers/admin"
)

func init() {
	beego.SetStaticPath("/web-static", "web-static")
	beego.Include(  &admin.AdminController{},&admin.AjaxController{},&admin.UeditorController{},&admin.AttachController{})

}
