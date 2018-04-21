package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
	"log"
	"database/sql"
	"fmt"
	"net/url"
	"github.com/astaxie/beego"
	"cims/tool"
)

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Color))
	orm.RegisterModel(new(Size))
	orm.RegisterModel(new(ArtNo))
	orm.RegisterModel(new(Lining))
	orm.RegisterModel(new(Fabric))
	orm.RegisterModel(new(WareHouse))
	orm.RegisterModel(new(Sku))
	orm.RegisterModel(new(InStock))
	orm.RegisterModel(new(InStockDetail))
	orm.RegisterModel(new(OutStock))
	orm.RegisterModel(new(OutStockDetail))
}

func createDb() {
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")

	var dns string
	var sqlstring string
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&loc=%s", db_user, db_pass, db_host, db_port, url.QueryEscape("Asia/Shanghai"))
	sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8mb4 COLLATE utf8mb4_general_ci", db_name)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()
}

func Connect() {
	var dns string
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s", db_user, db_pass, db_host, db_port, db_name, url.QueryEscape("Asia/Shanghai"))
	orm.RegisterDataBase("default", "mysql", dns)
}

var o orm.Ormer
func Syncbd(force bool) {
	createDb()
	Connect()
	o = orm.NewOrm()
	// 数据库别名
	name := "default"
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	if force {
		//添加系统用户
		AddSystemAdmin()
		//添加Size数据表
		AddSizes()
		//添加Color数据表
		AddColors()
		//添加里料表
		AddFabrics()
		//添加面料表
		AddLinings()
	}
	fmt.Println("database init is complete.\nPlease restart the application")
}

func AddSizes() {
	for i:= 150; i < 195; i += 5 {
		var size Size;
		size.GmtCreate = time.Now()
		size.GmtModifier = time.Now()
		size.Size = i
		if _, err := size.Insert(); err != nil{
			fmt.Println("insert size error")
		} else {
			fmt.Println(size.Size)
		}
	}
}

func AddColors() {
	AddColor("大红色")
	AddColor("浅红色")
	AddColor("紫红色")
	AddColor("纯白色")
	AddColor("米白色")
	AddColor("深蓝色")
	AddColor("淡蓝色")
	AddColor("黑色")
	AddColor("棕色")

}

func AddLinings() {
	AddLining("永俊长丝里布")
	AddLining("永俊创新的里布")
	AddLining("永俊条子里布")
}

func AddFabrics() {
	AddFabric("棉布")
	AddFabric("麻布")
	AddFabric("丝绸")
	AddFabric("呢绒")
	AddFabric("皮革")
	AddFabric("化纤")
	AddFabric("混纺")
	AddFabric("莫代尔")
}
func AddFabric(fabric string) {
	var fa Fabric
	fa.GmtCreate = time.Now()
	fa.GmtModifier = time.Now()
	fa.Fabric = fabric
	if _, err := fa.Insert(); err != nil {
		fmt.Println(err.Error())
		fmt.Println("insert fabric error")
	} else {
		fmt.Println(fa)
	}
}

func AddLining(lining string) {
	var li Lining
	li.GmtCreate = time.Now()
	li.GmtModifier = time.Now()
	li.Lining = lining
	if _, err := li.Insert(); err != nil {
		fmt.Println("insert lining error")
	} else {
		fmt.Println(lining)
	}
}

func AddColor(colorName string) {
	var color Color
	color.GmtCreate = time.Now()
	color.GmtModifier = time.Now()
	color.Color = colorName
	if _, err := color.Insert(); err != nil {
		fmt.Println("insert color error")
	} else {
		fmt.Println(colorName)
	}

}

func AddSystemAdmin() {
	var user User
	fmt.Println("please input username for system administrator")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("please input password for system administrator")
	var password string
	fmt.Scanf("%s", &password)

	user.Name = "SystemUser"
	user.Username = name
	user.Flag = 1
	user.Password = tool.Md5(password)
	if _, err := user.Insert(); err != nil {
		fmt.Println("admin create error,please run this application again")
	} else {
		fmt.Println("admin create finished")
		fmt.Println("管理员账户：" + user.Username)
		fmt.Println("管理员密码：" + user.Password)

	}
}

