package admin

import (
	"cims/models"
	"github.com/astaxie/beego/orm"
	"cims/tool"
	"strconv"
	"time"
)

//@router /admin/user/create [post]
func (this *AjaxController)CreateUser() {
	username := this.GetString("username")
	password := this.GetString("password")
	profile := this.GetString("profile")
	name := this.GetString("name")

	user := models.User{Username:username}
	if err := user.GetByUsername(); err != nil {
		if (err == orm.ErrNoRows) {
			user.Password = tool.Md5(password)
			user.Name = name
			user.Profile = profile
			if _, err := user.Insert(); err != nil {
				this.ReturnError("system error")
				return
			} else {
				this.ReturnSuccess()
				return
			}

		} else {
			this.ReturnError("system error")
			return
		}
	} else {
		this.ReturnError("用户名已经存在")
		return
	}

}


//@router /admin/user/:id/update [post]
func (this *AjaxController)UpdateUser() {
	uid, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	username := this.GetString("username")
	password := this.GetString("password")
	profile := this.GetString("profile")
	name := this.GetString("name")

	user := models.User{Id:uid}
	count, err := orm.NewOrm().QueryTable("user").Exclude("id", uid).Filter("username", username).Count();
	if (err != nil) {
		this.ReturnError("update user fail")
		return
	}
	if count >= 1 {
		this.ReturnError("用户名已存在")
		return
	}
	user.Password = tool.Md5(password)
	user.Name = name
	user.Profile = profile
	user.Username = username
	user.GmtModifier = time.Now()
	if _, err := user.Update(); err != nil {
		this.ReturnError("system error")
		return
	}
	this.ReturnSuccess()

}

//@router /admin/user/:id/delete [delete]
func (this *AjaxController) DeleteUser() {
	if this.Verity() == false {
		return
	}

	id, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	user := models.User{Id : id}

	if err := user.Read(); err == nil {
		user.Lock = 1
		if _, err := user.Delete(); err != nil {
			this.ReturnError("delete user error")
			return
		}
	} else {
		this.ReturnError("用户不存在")
	}
	this.ReturnSuccess()
}


//@router /admin/password/update [post]
func (this *AjaxController) UpdatePassword() {
	if this.Verity() == false {
		return
	}
	newPassword := this.GetString("newPassword")
	oldPassword := this.GetString("oldPassword")
	userId := this.GetSession("userinfo").(models.User).Id

	user := models.User{Id:userId}
	if err := user.Read(); err != nil {
		this.ReturnUnAuth()
		return
	}
	if (tool.Md5(oldPassword) != user.Password) {
		this.ReturnError("旧密码错误")
		return
	}
	if result, err := orm.NewOrm().QueryTable("user").Filter("flag", 1).Filter("id", user.Id).Update(orm.Params{"password":tool.Md5(newPassword)});
		err != nil {
		this.ReturnError("update password fail")
		return
	} else {
		if (result == 0) {
			this.ReturnError("update password fail")
			return
		}
	}
	this.ReturnSuccess()
}


