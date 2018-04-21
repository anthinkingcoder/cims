package admin

import (
	"github.com/astaxie/beego"
	"cims/models"
)

type AjaxController struct {
	beego.Controller
}
type Result struct {
	ErrorMsg string
	Success  bool
	Data     interface{}
}

func (this *AjaxController) Verity() bool {
	if userinfo := this.GetSession("userinfo"); userinfo != nil {
		user := userinfo.(models.User)
		if user.Flag == 1 {
			return true
		} else {
			this.ReturnUnAuth()
			return false
		}
	} else {
		this.ReturnUnAuth()
		return false
	}

}

func (this *AjaxController) ReturnJson(result Result) {
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *AjaxController) ReturnSuccess() {
	this.Data["json"] = Result{Success:true}
	this.ServeJSON()
}

func (this *AjaxController) ReturnError(errorMsg string) {
	this.Data["json"] = Result{Success:false, ErrorMsg:errorMsg}
	this.ServeJSON()
}

func (this *AjaxController) ReturnData(data interface{}) {
	this.Data["json"] = Result{Success:true, Data:data}
	this.ServeJSON()
}

func (this *AjaxController) ReturnUnAuth() {
	this.Ctx.ResponseWriter.WriteHeader(401)
	this.Data["json"] = Result{Success:false, ErrorMsg:"未授权"}
	this.ServeJSON()
}