package admin
import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
	"fmt"
)

type AttachController struct {
	beego.Controller
}
//@router /attachment/*
func (c *AttachController) Get() {
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:])
	fmt.Println("文件路径:" + filePath);
	_, err = url.QueryUnescape(c.Ctx.Request.RequestURI)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		c.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()

	_, err = io.Copy(c.Ctx.ResponseWriter, f)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
}
