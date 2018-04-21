package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["cims/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Home1",
			Router: `/admin/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "ArtNo",
			Router: `/admin/art_no`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "InStock",
			Router: `/admin/in_stock`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Home",
			Router: `/admin/index`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Article",
			Router: `/admin/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "OutStock",
			Router: `/admin/out_stock`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "UserManage",
			Router: `/admin/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "WareHouse",
			Router: `/admin/ware_house`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "UpdateSku",
			Router: `/admin/art_no/:art_no_id/sku/:id/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteArtNo",
			Router: `/admin/art_no/:id/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "GetColorsAndSizes",
			Router: `/admin/art_no/:id/getColorsAndSizes`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "CreateSku",
			Router: `/admin/art_no/:id/sku/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "ListSku",
			Router: `/admin/art_no/:id/sku/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "UpdateArtNo",
			Router: `/admin/art_no/:id/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "CreateArtNo",
			Router: `/admin/art_no/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteSku",
			Router: `/admin/art_no/sku/:id/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "Authorize",
			Router: `/admin/authorize`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteInStock",
			Router: `/admin/in_stock/:id/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "FinishInStock",
			Router: `/admin/in_stock/:id/finish`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "UpdateInStock",
			Router: `/admin/in_stock/:id/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "CreateInStock",
			Router: `/admin/in_stock/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "GetInStockNum",
			Router: `/admin/in_stock/getInStockNum`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteInStockDetail",
			Router: `/admin/in_stock_detail/:id/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "UpdateInStockDetail",
			Router: `/admin/in_stock_detail/:id/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "CreateInStockDetail",
			Router: `/admin/in_stock_detail/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "ListIntStockDetails",
			Router: `/admin/int_stock/:id/int_stock_detail/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/admin/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/admin/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteOutStock",
			Router: `/admin/out_stock/:id/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "FinishOutStock",
			Router: `/admin/out_stock/:id/finish`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "ListOutStockDetails",
			Router: `/admin/out_stock/:id/out_stock_detail/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "UpdateOutStock",
			Router: `/admin/out_stock/:id/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "CreateOutStock",
			Router: `/admin/out_stock/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "GetOutStockNum",
			Router: `/admin/out_stock/getOutStockNum`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteOutStockDetail",
			Router: `/admin/out_stock_detail/:id/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "UpdateOutStockDetail",
			Router: `/admin/out_stock_detail/:id/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "CreateOutStockDetail",
			Router: `/admin/out_stock_detail/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "UpdatePassword",
			Router: `/admin/password/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteUser",
			Router: `/admin/user/:id/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "UpdateUser",
			Router: `/admin/user/:id/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/admin/user/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "DeleteWareHouse",
			Router: `/admin/ware_house/:id/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "UpdateWareHouse",
			Router: `/admin/ware_house/:id/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AjaxController"],
		beego.ControllerComments{
			Method: "CreateWareHouse",
			Router: `/admin/ware_house/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:AttachController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:AttachController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/attachment/*`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["cims/controllers/admin:UeditorController"] = append(beego.GlobalControllerRouter["cims/controllers/admin:UeditorController"],
		beego.ControllerComments{
			Method: "ControllerUE",
			Router: `/controller`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

}
