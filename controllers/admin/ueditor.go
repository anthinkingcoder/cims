package admin

import (
	"encoding/base64"
	"encoding/json"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
	"fmt"
	"github.com/pborman/uuid"
)

type UeditorController struct {
	beego.Controller
}

type UploadimageUE struct {
	url      string
	title    string
	original string
	state    string
	// "url": fmt.Sprintf("/static/upload/%s", filename),
	// "title": "demo.jpg",
	// "original": header.Filename,
	// "state": "SUCCESS"
}

type List struct {
	Url string `json:"url"`
	// Source string
	// State  string
}

type Listimage struct {
	State string `json:"state"` //这些第一个字母要大写，否则不出结果
	List  []List `json:"list"`
	Start int    `json:"start"`
	Total int    `json:"total"`
				    // Name        string
				    // Age         int
				    // Slices      []string //slice
				    // Mapstring   map[string]string
				    // StructArray []List            //结构体的切片型
				    // MapStruct   map[string][]List //map:key类型是string或struct，value类型是切片，切片的类型是string或struct
				    //	Desks  List
}
type ListCatch struct {
	Url    string `json:"url"`
	Source string `json:"source"`
	State  string `json:"state"`
}
type Catchimage struct {
	State string      `json:"state"` //这些第一个字母要大写，否则不出结果
	List  []ListCatch `json:"list"`
}

//@router /controller [*]
func (c *UeditorController) ControllerUE() {
	op := c.Input().Get("action")
	// key := c.Input().Get("key") //这里进行判断各个页面，如果是addtopic，如果是addcategory
	switch op {
	case "config": //这里是conf/config.json
		file, err := os.Open("conf/config.json")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		fd, _ := ioutil.ReadAll(file)
		src := string(fd)
		re, _ := regexp.Compile("\\/\\*[\\S\\s]+?\\*\\/") //参考php的$CONFIG = json_decode(preg_replace("/\/\*[\s\S]+?\*\//", "", file_get_contents("config.json")), true);
		//将php中的正则移植到go中，需要将/ \/\*[\s\S]+?\*\/  /去掉前后的/，然后将\改成2个\\
		//参考//去除所有尖括号内的HTML代码，并换成换行符
		// re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
		// src = re.ReplaceAllString(src, "\n")
		//当把<和>换成/*和*\时，斜杠/和*之间加双斜杠\\才行。
		src = re.ReplaceAllString(src, "")
		tt := []byte(src)
		var r interface{}
		json.Unmarshal(tt, &r) //这个byte要解码
		c.Data["json"] = r
		c.ServeJSON()
	case "uploadimage", "uploadfile", "uploadvideo":
		path := c.GetPathSeparator()
		err := os.MkdirAll("." + path + "attachment" + path, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
		if err != nil {
			beego.Error(err)
		}
		//保存上传的图片
		//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
		_, h, err := c.GetFile("upfile")
		if err != nil {
			beego.Error(err)
		}
		path1 := "." + path + "attachment" + path + h.Filename
		fmt.Println(path1)
		err = c.SaveToFile("upfile", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/attachment/" + h.Filename, "title": h.Filename, "original": h.Filename}
		c.ServeJSON()
	// 		{
	//     "state": "SUCCESS",
	//     "url": "upload/demo.jpg",
	//     "title": "demo.jpg",
	//     "original": "demo.jpg"
	// }
	case "uploadscrawl":
		number := c.Input().Get("number")
		name := c.Input().Get("name")
		path := c.GetPathSeparator()
		err := os.MkdirAll("." + path + "attachment" + path  + number + name, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
		if err != nil {
			beego.Error(err)
		}
		path1 := "." + path + "attachment" + path + number + name + path
		//保存上传的图片
		//upfile为base64格式文件，转成图片保存
		ww := c.Input().Get("upfile")
		ddd, _ := base64.StdEncoding.DecodeString(ww)           //成图片文件并把文件写入到buffer
		newname := strconv.FormatInt(time.Now().Unix(), 10)     // + "_" + filename
		err = ioutil.WriteFile(path1 + newname + ".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = map[string]interface{}{
			"state":    "SUCCESS",
			"url":      "/attachment/" + number + name + "/" + newname + ".jpg",
			"title":    newname + ".jpg",
			"original": newname + ".jpg",
		}
		c.ServeJSON()
	case "listimage":
		list := []List{
			{"/static/upload/1.jpg"},
			{"/static/upload/2.jpg"},
		}
		listimage := Listimage{"SUCCESS", list, 1, 21}
		c.Data["json"] = listimage
		c.ServeJSON()
	// 需要支持callback参数,返回jsonp格式
	// {
	//     "state": "SUCCESS",
	//     "list": [{
	//         "url": "upload/1.jpg"
	//     }, {
	//         "url": "upload/2.jpg"
	//     }, ],
	//     "start": 20,
	//     "total": 100
	// }
	case "catchimage":
		list := []ListCatch{
			{"/static/upload/1.jpg", "https://pic2.zhimg.com/7c4a389acaa008a6d1fe5a0083c86975_b.png", "SUCCESS"},
			{"/static/upload/2.jpg", "https://pic2.zhimg.com/7c4a389acaa008a6d1fe5a0083c86975_b.png", "SUCCESS"},
		}
		catchimage := Catchimage{"SUCCESS", list}
		c.Data["json"] = catchimage
		c.ServeJSON()

		file, header, err := c.GetFile("source") // r.FormFile("upfile")
		beego.Info(header.Filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
		err = os.MkdirAll(path.Join("static", "upload"), 0775)
		if err != nil {
			panic(err)
		}
		outFile, err := os.Create(path.Join("static", "upload", filename))
		if err != nil {
			panic(err)
		}
		defer outFile.Close()
		io.Copy(outFile, file)
	// 需要支持callback参数,返回jsonp格式
	// list项的state属性和最外面的state格式一致
	// {
	//     "state": "SUCCESS",
	//     "list": [{
	//         "url": "upload/1.jpg",
	//         "source": "http://b.com/2.jpg",
	//         "state": "SUCCESS"
	//     }, {
	//         "url": "upload/2.jpg",
	//         "source": "http://b.com/2.jpg",
	//         "state": "SUCCESS"
	//     }, ]
	// }
	}

}
func (*UeditorController) GetPathSeparator() string {
	var path string
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	return path;
}
