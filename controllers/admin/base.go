package admin

import (
	"github.com/astaxie/beego"
	"cims/models"
)

type Base struct {
	beego.Controller
	IsLogin bool
	User     models.User
}

func (this *Base) Prepare() {
	this.IsLogin = false

	if userInfo := this.GetSession("userinfo"); userInfo != nil {
		user := userInfo.(models.User)
		if err := user.Read(); err != nil {
			this.Ctx.Redirect(302, "/admin/login")
			return
		} else {
			if user.Flag == 1 {
				this.IsLogin = true
				this.User = user
				this.Data["user"] = user
			} else {
				this.Ctx.Redirect(302, "/admin/login")
				return
			}
		}
	} else {
		this.Ctx.Redirect(302, "/admin/login")
	}
}