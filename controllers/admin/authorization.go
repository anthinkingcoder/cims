package admin

import (
	"cims/models"
	"cims/tool"
	"time"
	"fmt"
)


//@router /admin/login
func (this *AjaxController) Login() {
	this.TplName = "admin/login.tpl"
}



//@router /admin/logout
func (this *AjaxController) Logout() {
	this.DelSession("userinfo")
	this.TplName = "admin/login.tpl"
}
//@router /admin/authorize [post]
func (this *AjaxController) Authorize() {
	var result Result
	username := this.GetString("username")
	password := this.GetString("password")
	user := models.User{Username:username}

	if err := user.GetByUsername(); err != nil {
		result = Result{Success:false, ErrorMsg:"username or password error"}
		this.ReturnJson(result)
		return
	}

	if user.Flag != 1 {
		fmt.Println("Flag")
		result = Result{Success:false, ErrorMsg:"username or password error"}
		this.ReturnJson(result)
		return
	}

	if tool.Md5(password) != user.Password {
		result = Result{Success:false, ErrorMsg:"username or password error"}
		this.ReturnJson(result)
		return
	}

	user.GmtModifier = time.Now()
	user.Update()
	user.Password = ""
	this.SetSession("userinfo", user)
	this.ReturnSuccess()

}


