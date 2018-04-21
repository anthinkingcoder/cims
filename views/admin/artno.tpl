{{ template "admin/base.tpl" .}}
{{ define "meta"}}
<title>货号管理</title>
{{end}}

{{ define "css"}}
{{end}}

{{define "content"}}
<div class="row">
    <div class="col-sm-12">
        <div class="ibox float-e-margins">
            <div class="ibox-title">
                <h5>货号管理
                    <small>分类，查找</small>
                </h5>
                <div class="ibox-tools">
                    <a class="collapse-link">
                        <i class="fa fa-chevron-up"></i>
                    </a>
                    <a class="dropdown-toggle" data-toggle="dropdown"
                       href="#">
                        <i class="fa fa-wrench"></i>
                    </a>
                    <ul class="dropdown-menu dropdown-user">
                        <li><a href="#">选项1</a>
                        </li>
                        <li><a href="#">选项2</a>
                        </li>
                    </ul>
                    <a class="close-link">
                        <i class="fa fa-times"></i>
                    </a>
                </div>
            </div>
            <div class="ibox-content">
                <div style="text-align: center;">
                    <button class="btn btn-info" data-toggle="modal" data-target="#newArtNoModal">新增货号
                        <i class="fa fa-plus"></i></button>
                </div>
                <div style="height: 10px"></div>
                <table id="user_manager_table"
                       class="table table-striped table-bordered table-hover dataTables-example">
                    <thead>
                    <tr>
                        <th>序号</th>
                        <th>货号</th>
                        <th>品名</th>
                        <th>Sku</th>
                        <th>创建时间</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range .artnos }}
                    <tr>
                        <td>{{.Id}}</td>
                        <td><a href="#" data-toggle="modal" onclick="openEditArtNoModal({{.}})">{{.ArtNo}}</a>
                        </td>
                        <td>{{.ProductName}}</td>
                        <th>
                            <button class="btn btn-info btn-sm" type="button" id="openSkuDetailModalBtn_{{.Id}}"

                                    onclick="openSkuDetailModal({{.Colors}},{{.Sizes}},{{.Id}},{{.RetailPrice}},{{.FactoryPrice}})"><i
                                    class="fa fa-paste"></i>查看
                            </button>
                        </th>
                        <td>{{date .GmtCreate "Y-m-d h:i:s"}}</td>
                        <td>
                            <button class="btn btn-danger btn-sm"
                                    onclick="deleteArtNo({{ .Id }})">
                                <i
                                        class="fa fa-trash"></i> 删除
                            </button>
                        </td>
                    </tr>
                    {{end }}

                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>


<!--新增货号-->
<div class="modal fade" id="newArtNoModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">新增货号</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
                    <div class="form-group">
                        <label class="col-sm-2 control-label">货号</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入货号" class="form-control" id="artNo">

                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">品名</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入品名" class="form-control"
                                   id="productName">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">面料</label>

                        <div class="col-sm-10">
                            <select class="form-control m-b" name="" id="lining">
                                {{ range .linings}}
                                <option value="{{.Id}}">{{.Lining}}</option>
                                {{end}}
                            </select>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">里料</label>

                        <div class="col-sm-10">
                            <select class="form-control m-b" name="" id="fabric">
                                {{ range .fabrics}}
                                <option value="{{.Id}}">{{.Fabric}}</option>
                                {{end}}
                            </select>
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-sm-2 control-label">出厂价</label>

                        <div class="col-sm-10">
                            <input type="text" placeholder="默认出厂价" class="form-control"
                                   id="factoryPrice">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">零售价</label>

                        <div class="col-sm-10">
                            <input type="text" placeholder="默认零售价" class="form-control"
                                   id="retailPrice" >
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">颜色</label>
                        <div class="col-sm-10" id="color">
                            {{range .colors}}
                            <label class="checkbox-inline">
                                <input type="checkbox" value="{{.Id}}">{{.Color}}</label>
                            {{end}}

                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">尺寸</label>
                        <div class="col-sm-10" id="size">
                            {{range .sizes}}
                            <label class="checkbox-inline">
                                <input type="checkbox" value="{{.Id}}">{{ .Size }}</label>
                            {{end}}

                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                </form>
            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="newArtNo();">提交</button>
            </div>
        </div>
    </div>
</div>


<!--编辑货号-->
<div class="modal fade" id="editArtNoModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">编辑货号</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
                    <input type="text" hidden="hidden" id="edit_id">

                    <div class="form-group">
                        <label class="col-sm-2 control-label">货号</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入货号" class="form-control" id="edit_artNo">

                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">品名</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入品名" class="form-control"
                                   id="edit_productName">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">面料</label>

                        <div class="col-sm-10">
                            <select class="form-control m-b" name="" id="edit_lining">
                                {{ range .linings}}
                                <option value="{{.Id}}">{{.Lining}}</option>
                                {{end}}
                            </select>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">里料</label>

                        <div class="col-sm-10">
                            <select class="form-control m-b" name="" id="edit_fabric">
                                {{ range .fabrics}}
                                <option value="{{.Id}}">{{.Fabric}}</option>
                                {{end}}
                            </select>
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-sm-2 control-label">出厂价</label>

                        <div class="col-sm-10">
                            <input type="text" placeholder="默认出厂价" class="form-control"
                                   id="edit_factoryPrice">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">零售价</label>

                        <div class="col-sm-10">
                            <input type="text" placeholder="默认零售价" class="form-control"
                                   id="edit_retailPrice">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">颜色</label>
                        <div class="col-sm-10" id="edit_color">
                            {{range .colors}}
                            <label class="checkbox-inline">
                                <input type="checkbox" value="{{.Id}}">{{.Color}}</label>
                            {{end}}

                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">尺寸</label>
                        <div class="col-sm-10" id="edit_size">
                            {{range .sizes}}
                            <label class="checkbox-inline">
                                <input type="checkbox" value="{{.Id}}">{{ .Size }}</label>
                            {{end}}

                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                </form>
            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="editArtNo();">保存</button>
            </div>
        </div>
    </div>
</div>


<!-- sku详情模态框-->
<div class="modal fade" id="skuDetailModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">颜色尺寸</h4>
            </div>
            <p id="sku_art_no_id" hidden></p>
            <div class="modal-body">
                <div style="text-align: center;">
                    <button class="btn btn-info" data-toggle="modal" data-target="#newSkuModal">新增Sku
                        <i class="fa fa-plus"></i></button>
                </div>
                <div style="height: 30px"></div>
                <table id="skuDetailTable"
                       class="table table-striped table-bordered table-hover">
                    <thead>
                    <tr>
                        <th>序号</th>
                        <th>颜色</th>
                        <th>尺寸</th>
                        <th>出厂价</th>
                        <th>零售价</th>
                        <th>库存</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody id="skuDetailTBody">
                    </tbody>
                </table>
            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
            </div>
        </div>
    </div>
</div>


<!-- 新增sku模态框-->
<div class="modal fade" id="newSkuModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">新增Sku</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">颜色</label>
                        <div class="col-sm-10">
                            <select class="form-control m-b" name="" id="sku_color">

                            </select>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">尺寸</label>
                        <div class="col-sm-10">
                            <select class="form-control m-b" name="" id="sku_size">
                            </select>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">出厂价</label>

                        <div class="col-sm-10">
                            <input type="text" placeholder="默认出厂价" class="form-control"
                                   id="sku_factoryPrice"  value="">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">零售价</label>

                        <div class="col-sm-10">
                            <input type="text" placeholder="默认零售价" class="form-control"
                                   id="sku_retailPrice"  value="">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                </form>
            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="newSku();">保存</button>
            </div>
        </div>
    </div>
</div>


<!-- sku编辑模态框-->
<div class="modal fade" id="editSkuModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">单品详情</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
                    <input type="text" hidden="hidden" id="sku_id">
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">颜色</label>
                        <div class="col-sm-10">
                            <select class="form-control m-b" name="" id="edit_sku_color">

                            </select>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">尺寸</label>
                        <div class="col-sm-10">
                            <select class="form-control m-b" name="" id="edit_sku_size">
                            </select>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">出厂价</label>

                        <div class="col-sm-10">
                            <input type="text" placeholder="默认出厂价" class="form-control"
                                   id="edit_sku_factoryPrice">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">零售价</label>

                        <div class="col-sm-10">
                            <input type="text" placeholder="默认零售价" class="form-control"
                                   id="edit_sku_retailPrice">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">库存</label>
                        <div class="col-sm-10">
                            <input type="text" disabled class="form-control"
                                   id="edit_sku_stock">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                </form>
            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="editSku();">保存</button>
            </div>
        </div>
    </div>
</div>

{{end}}
{{define "onload"}}
$("#systemManageMenu").click()
$('.dataTables-example').dataTable({
"order": [0,'desc']

});

$("#wizard").steps();
$("#form").steps({
bodyTag: "fieldset",
onStepChanging: function (event, currentIndex, newIndex) {
// Always allow going backward even if the current step contains invalid fields!
if (currentIndex > newIndex) {
return true;
}

// Forbid suppressing "Warning" step if the user is to young
if (newIndex === 3 && Number($("#age").val()) < 18) {
return false;
}

var form = $(this);

// Clean up if user went backward before
if (currentIndex < newIndex) {
// To remove error styles
$(".body:eq(" + newIndex + ") label.error", form).remove();
$(".body:eq(" + newIndex + ") .error", form).removeClass("error");
}

// Disable validation on fields that are disabled or hidden.
form.validate().settings.ignore = ":disabled,:hidden";

// Start validation; Prevent going forward if false
return form.valid();
},
onStepChanged: function (event, currentIndex, priorIndex) {
// Suppress (skip) "Warning" step if the user is old enough.
if (currentIndex === 2 && Number($("#age").val()) >= 18) {
$(this).steps("next");
}

// Suppress (skip) "Warning" step if the user is old enough and wants to the previous step.
if (currentIndex === 2 && priorIndex === 3) {
$(this).steps("previous");
}
},
onFinishing: function (event, currentIndex) {
var form = $(this);

// Disable validation on fields that are disabled.
// At this point it's recommended to do an overall check (mean ignoring only disabled fields)
form.validate().settings.ignore = ":disabled";

// Start validation; Prevent form submission if false
return form.valid();
},
onFinished: function (event, currentIndex) {
var form = $(this);

// Submit form input
form.submit();
}
}).validate({
errorPlacement: function (error, element) {
element.before(error);
},
rules: {
confirm: {
equalTo: "#password"
}
}
});

{{end}}

{{define "js"}}
<script>

    $.fn.dataTable.ext.errMode = 'none';

    function openEditArtNoModal(artNo) {

        $('#edit_artNo').val(artNo.ArtNo);
        $('#edit_id').val(artNo.Id);
        $('#edit_productName').val(artNo.ProductName);
        $('#edit_factoryPrice').val(artNo.FactoryPrice);
        $('#edit_retailPrice').val(artNo.RetailPrice);
        $('#edit_lining').val(artNo.Lining.Id);
        $('#edit_fabric').val(artNo.Fabric.Id);
        var colors = artNo.Colors;
        var sizes = artNo.Sizes;
        if (sizes.length != 0) {
            $("#edit_size").children("label").children("input").each(function () {
                var input = $(this);
                var id = $(this).val();
                sizes.forEach(function (value) {
                    if (value.Id == id) {
                        input.prop("checked", true);
                    }
                })
            });
        }
        if (colors.length != 0) {
            $("#edit_color").children("label").children("input").each(function () {
                var input = $(this);
                var id = $(this).val();
                colors.forEach(function (value) {
                    if (value.Id == id) {
                        input.prop("checked", true);
                    }
                })
            });
        }
        $('#editArtNoModal').modal("show");
    }



    function openSkuDetailModal(colors, sizes, artNoId,retailPrice,factoryPrice) {

        $('#sku_color').empty();
        $('#edit_sku_color').empty();
        $('#sku_factoryPrice').val(factoryPrice)
        $('#sku_retailPrice').val(retailPrice)
        colors.forEach(function (value) {
            var option = "<option value='" + value.Id + "'>" + value.Color + "</option>";
            $('#sku_color').append(option);
            $('#edit_sku_color').append(option);
        });


        $('#sku_size').empty();
        $('#edit_sku_size').empty();
        sizes.forEach(function (value) {
            var option = "<option value='" + value.Id + "'>" + value.Size + "</option>";
            $('#sku_size').append(option);
            $('#edit_sku_size').append(option);
        });

        $('#sku_art_no_id').text(artNoId);
        $("#skuDetailModal").modal("show");
        $.ajax({
            url: "/admin/art_no/" + artNoId + "/sku/list",
            type: "get",
            success: function (data) {
                if (data.Success) {
                    //将sku渲染到skuDetailTable中
                    var datas = data.Data;
                    var tbody = "";
                    datas.forEach(function (value) {
                        tbody += "<tr>";
                        tbody += "<td>" + value.Id + "</td>";
                        tbody += "<td>" + value.Color.Color + "</td>";
                        tbody += "<td>" + value.Size.Size + "</td>";
                        tbody += "<td>" + value.FactoryPrice + "</td>";
                        tbody += "<td>" + value.RetailPrice + "</td>";
                        tbody += "<td>" + value.Stock + "</td>";
                        tbody += "<td>" + "<button class='btn btn-info btn-sm' onclick='openEditSkuModal(" + value.Id + "," + " &quot;" + value.Color.Id + " &quot;" + "," + "&quot;" + value.Size.Id + "&quot;" + "," + value.FactoryPrice + "," + value.RetailPrice + "," + value.Stock + ")'> <i class='fa fa-paste'></i> 编辑</button>  <button class='btn btn-danger btn-sm' onclick='deleteSku(" + value.Id + "," + artNoId + ")'> <i class='fa fa-trash'></i> 删除</button>" + "</td>";
                        tbody += "</tr>";
                    });
                    $('#skuDetailTBody').empty();
                    $('#skuDetailTBody').append(tbody);
                    $('#skuDetailTable').dataTable();
                }
            },
            error: function (data) {

            }
        })
    }

    function newSku() {
        var factoryPrice = $('#sku_factoryPrice').val();
        var retailPrice = $('#sku_retailPrice').val();
        var colorId = $('#sku_color').val();
        var sizeId = $('#sku_size').val();
        var artNoId = $('#sku_art_no_id').text();
        if (!factoryPrice || !retailPrice || !colorId || !sizeId) {
            alert('信息填写不完整');
            return
        }

        $.ajax({
            url: "/admin/art_no/" + artNoId + "/sku/create",
            type: "post",
            data: {
                factoryPrice: factoryPrice,
                retailPrice: retailPrice,
                colorId: colorId,
                sizeId: sizeId,
            },
            success: function (data) {
                if (data.Success) {
                    alert("新增成功");
                    $('#sku_factoryPrice').val(null);
                    $('#sku_retailPrice').val(null);
                    $('#sku_color').val(null);
                    $('#sku_size').val(null);
                    $('#newSkuModal').modal("hide");
                    $('#openSkuDetailModalBtn_' + artNoId).click();
                } else {
                    alert(data.ErrorMsg)
                }
            },
            error: function (data) {
                if (data.status == 401) {
                    alert("请重新登录!")
                }
            }
        })
    }

    function openEditSkuModal(skuId, color, size, factoryPrice, retailPrice, stock) {
        $('#sku_id').val(skuId);
        console.info(color)
        $('#edit_sku_color').val(parseInt(color));
        $('#edit_sku_size').val(parseInt(size));
        $('#edit_sku_factoryPrice').val(factoryPrice);
        $('#edit_sku_retailPrice').val(retailPrice);
        $('#edit_sku_stock').val(stock);
        $('#editSkuModal').modal("show");
    }
    function editSku() {
        var skuId = $('#sku_id').val();
        var factoryPrice = $('#edit_sku_factoryPrice').val();
        var retailPrice = $('#edit_sku_retailPrice').val();
        var colorId = $('#edit_sku_color').val();
        var sizeId = $('#edit_sku_size').val();
        var artNoId = $('#sku_art_no_id').text();
        if (!factoryPrice || !retailPrice || !colorId || !sizeId) {
            alert('信息填写不完整');
            return
        }
        $.ajax({
            url: "/admin/art_no/" + artNoId + "/sku/" + skuId + "/update",
            type: "post",
            data: {
                factoryPrice: factoryPrice,
                retailPrice: retailPrice,
                colorId: colorId,
                sizeId: sizeId
            },
            success: function (data) {
                if (data.Success) {
                    alert("更新成功")
                    $('#editSkuModal').modal("hide");
                    $('#openSkuDetailModalBtn_' + artNoId).click();
                } else {
                    alert(data.ErrorMsg)
                }
            },
            error: function (data) {
                if (data.status == 401) {
                    alert("请重新登录!")
                }
            }
        })


    }

    function deleteSku(skuId, artNoId) {
        console.info(artNoId);
        var result = window.confirm("是否删除?");
        if (result == true) {
            $.ajax({
                url: "/admin/art_no/sku/" + skuId + "/delete",
                type: "delete",
                success: function (data) {
                    if (data.Success) {
                        $('#openSkuDetailModalBtn_' + artNoId).click();
                    } else {
                        alert(data.ErrorMsg)
                    }
                },
                error: function (data) {
                    if (data.status == 401) {
                        alert("请重新登录!")
                    }
                }
            })

        }
    }

    function editArtNo() {
        var artId = $('#edit_id').val();
        var artNo = $('#edit_artNo').val();
        var productName = $('#edit_productName').val();
        var lining = $('#edit_lining').val();
        var fabric = $('#edit_fabric').val();
        var retailPrice = $('#edit_retailPrice').val();
        var factoryPrice = $('#edit_factoryPrice').val();
        var colors = [];
        var sizes = [];
        var i = 0;
        if (!artNo || !productName || !lining || !fabric || !retailPrice || !factoryPrice) {
            alert("信息请填写完整");
            return
        }
        //获取选中的颜色
        $("#edit_color").children("label").children("input").each(function () {
            if ($(this).is(":checked")) {
                colors[i++] = $(this).val();
            }
        });
        if (colors.length == 0) {
            alert("请选择颜色");
            return
        }
        //获取选中的尺寸
        i = 0;
        $("#edit_size").children("label").children("input").each(function () {
            if ($(this).is(":checked")) {
                sizes[i++] = $(this).val();
            }
        });
        if (sizes.length == 0) {
            alert("请选择尺寸");
            return
        }

        $.ajax({
            url: "/admin/art_no/" + artId + "/update",
            type: "post",
            data: {
                artNo: artNo,
                productName: productName,
                liningId: lining,
                fabricId: fabric,
                retailPrice: retailPrice,
                factoryPrice: factoryPrice,
                colors: colors,
                sizes: sizes
            },
            success: function (data) {
                if (data.Success) {
                    window.location.href = "/admin/art_no"
                } else {
                    alert(data.ErrorMsg)
                }
            },
            error: function (data) {
                if (data.status == 401) {
                    alert("请重新登录!")
                }
            }
        })


    }

    function deleteArtNo(artNoId) {

        var result = window.confirm("是否删除?");
        if (result == true) {
            $.ajax({
                url: "/admin/art_no/" + artNoId + "/delete",
                type: "delete",
                success: function (data) {
                    if (data.Success) {
                        window.location.href = "/admin/art_no"
                    } else {
                        alert(data.ErrorMsg)
                    }
                },
                error: function (data) {
                    if (data.status == 401) {
                        alert("请重新登录!")
                    }
                }
            })

        }

    }

    function newArtNo() {
        var artNo = $('#artNo').val();
        var productName = $('#productName').val();
        var lining = $('#lining').val();
        var fabric = $('#fabric').val();
        var retailPrice = $('#retailPrice').val();
        var factoryPrice = $('#factoryPrice').val();
        var colors = [];
        var sizes = [];
        var i = 0;
        if (!artNo || !productName || !lining || !fabric || !retailPrice || !factoryPrice) {
            alert("信息请填写完整");
            return
        }
        //获取选中的颜色
        $("#color").children("label").children("input").each(function () {
            if ($(this).is(":checked")) {
                colors[i++] = $(this).val();
            }
        });
        if (colors.length == 0) {
            alert("请选择颜色");
            return
        }
        //获取选中的尺寸
        i = 0;
        $("#size").children("label").children("input").each(function () {
            if ($(this).is(":checked")) {
                sizes[i++] = $(this).val();
            }
        });
        if (sizes.length == 0) {
            alert("请选择尺寸");
            return
        }

        $.ajax({
            url: "/admin/art_no/create",
            type: "post",
            data: {
                artNo: artNo,
                productName: productName,
                liningId: lining,
                fabricId: fabric,
                retailPrice: retailPrice,
                factoryPrice: factoryPrice,
                colors: colors,
                sizes: sizes
            },
            success: function (data) {
                if (data.Success) {
                    window.location.href = "/admin/art_no"
                } else {
                    alert(data.ErrorMsg)
                }
            },
            error: function (data) {
                if (data.status == 401) {
                    alert("请重新登录!")
                }
            }
        })
    }

</script>


{{end}}