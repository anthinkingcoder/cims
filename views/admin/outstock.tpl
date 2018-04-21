{{ template "admin/base.tpl" .}}
{{ define "meta"}}
<title>库存管理</title>
{{end}}

{{ define "css"}}

{{end}}

{{define "content"}}
<div class="row">
    <div class="col-sm-12">
        <div class="ibox float-e-margins">
            <div class="ibox-title">
                <h5>出库管理
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
                    <button class="btn btn-info" data-toggle="modal" onclick="openNewOutStockModal()">新增出库单
                        <i class="fa fa-plus"></i></button>
                </div>
                <div style="height: 10px"></div>
                <table id="user_manager_table"
                       class="table table-striped table-bordered table-hover dataTables-example">
                    <thead>
                    <tr>
                        <th>序号</th>
                        <th>单据编号</th>
                        <th>所出仓库</th>
                        <th>出库日期</th>
                        <th>接受人</th>
                        <th>发往地址</th>
                        <th>单据详情</th>
                        <th>状态</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range .outstocks}}
                    <tr>
                        <td>{{.Id}}</td>
                        <td><a href="#"
                               onclick='openEditOutStockModal({{.Id}},{{.OutStockNum}},{{.WareHouse.Id}},{{date .DeliveryTime "Y-m-d h:i:s"}},{{.Remark}},{{.ReceiverTel}},{{.ReceiverName}},{{.Address}},{{.Status}})'>{{.OutStockNum}}</a>
                        </td>
                        <td>{{.WareHouse.Name}}</td>
                        <td>{{date .DeliveryTime "Y-m-d h:i:s"}}</td>
                        <td>{{.ReceiverName}}</td>
                        <td>{{.Address}}</td>
                        <td>
                            <button class="btn btn-info btn-sm"
                                    onclick="openOutStockDetailModal({{.Status}},{{.Id}})">
                                <i
                                        class="fa fa-edit"></i> 查看
                            </button>
                        </td>
                        <td>
                            {{if eq .Status 0}}
                            <button class="btn btn-primary btn-sm"
                                    onclick="submitOutStock({{.Id }})">
                                <i
                                        class="fa fa-edit"></i> 提交
                            </button>
                            {{ end }}


                            {{if eq .Status 1}}
                            <button class="btn btn-success btn-circle"
                                    disabled>
                                <i
                                        class="fa fa-check"></i>
                            </button>
                            {{ end }}
                        </td>
                        <td>

                            <button class="btn btn-danger btn-sm"
                                    onclick="deleteOutStock({{.Id }})">
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


<!--//新增出库单模态-->
<div class="modal fade" id="myModal5" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">新增出库单</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
                    <div class="form-group">
                        <label class="col-sm-2 control-label">单据编号</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="" disabled class="form-control" id="outStockNum">
                        </div>
                    </div>

                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">所出仓库</label>

                        <div class="col-sm-10">
                            <select class="form-control m-b" name="" id="warehouse">
                                <option value="-1">请选择</option>
                                {{ range .warehouses}}
                                <option value="{{.Id}}">{{.Name}}</option>
                                {{end}}
                            </select>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">接受人</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入接受人" class="form-control" id="receiverName">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">接受人电话</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入接受人电话" class="form-control" id="receiverTel">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">发往地址</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入发往地址" class="form-control" id="address">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">备注</label>
                        <div class="col-sm-10">
                            <textarea class="form-control" rows="6" name="profile" id="remark"></textarea>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">出库时间：</label>
                        <div class="col-sm-10">
                            <input class="laydate-icon form-control layer-date" id="create_time">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div style="height: 60px;"></div>
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="createOutStock();">提交</button>
            </div>
        </div>
    </div>
</div>

<!--//编辑出库单模态-->
<div class="modal fade" id="editOutStockModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">出库单信息</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
                    <input id="edit_outStockId" hidden>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">单据编号</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="" disabled class="form-control" id="edit_outStockNum">
                        </div>
                    </div>

                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">所出仓库</label>

                        <div class="col-sm-10">
                            <select class="form-control m-b" name="" id="edit_warehouse">
                                <option value="-1">请选择</option>
                                {{ range .warehouses}}
                                <option value="{{.Id}}">{{.Name}}</option>
                                {{end}}
                            </select>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">接受人</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入接受人" class="form-control" id="edit_receiverName">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">接受人电话</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入接受人电话" class="form-control" id="edit_receiverTel">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">发往地址</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入发往地址" class="form-control" id="edit_address">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">备注</label>
                        <div class="col-sm-10">
                            <textarea class="form-control" rows="6" name="profile" id="edit_remark"></textarea>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">出库时间：</label>
                        <div class="col-sm-10">
                            <input class="laydate-icon form-control layer-date" id="edit_create_time">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div style="height: 60px;"></div>
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" id="updateOutStockBtn" onclick="updateOutStock();">保存
                </button>
            </div>
        </div>
    </div>
</div>


<!-- outstock detail详情模态框-->
<div class="modal fade" id="outStockDetailModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">单据详情</h4>
            </div>
            <div class="modal-body">
                <div style="text-align: center;">
                    <button class="btn btn-info" id="createOutStockBtn" onclick="openNewOutStockDetail()">新增出库单详情
                        <i class="fa fa-plus"></i></button>
                </div>
                <input id="outStockId" hidden>
                <div style="height: 30px;"></div>
                <span class="help-block m-b-none" id="outStockDetailTableTip"><i class="fa fa-info-circle"></i> 失效表示该商品已经被删除或者该商品的颜色尺寸被修改,需删除才能提交订单</span>
                <table id=""
                       class="table table-striped table-bordered table-hover">
                    <thead>
                    <tr>
                        <th>序号</th>
                        <th>货号</th>
                        <th>品名</th>
                        <th>颜色</th>
                        <th>尺寸</th>
                        <th>数量</th>
                        <th>状态</th>
                        <th id="outStockDetailOper">操作</th>
                    </tr>
                    </thead>
                    <tbody id="outStockDetailTBody">
                    </tbody>
                </table>
            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
            </div>
        </div>
    </div>
</div>

<!--//新增出库单详情模态-->
<div class="modal fade" id="newOutStockDetailModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button disabled type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">新增出库单详情</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
                    <div class="hr-line-dashed"></div>

                    <div class="form-group">
                        <label class="col-sm-2 control-label">货号</label>
                        <div class="col-sm-10">
                            <select class="form-control m-b"
                                    onchange="selectedArtNo({{.artnos}},'#artNo','#skuTable',0)"
                                    id="artNo">
                                <option value="-1">请选择</option>
                                {{range .artnos}}
                                <option value="{{.Id}}">{{.ArtNo}}</option>
                                {{end}}
                            </select>

                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">出库商品</label>
                        <div class="col-sm-10">

                            <span class="help-block m-b-none"><i class="fa fa-info-circle"></i> 选择一个需要出库的商品</span>
                            <table id="skuTable" style="table-layout:fixed"
                                   class="table table-striped table-bordered table-hover">
                                <thead>
                                <tr>
                                    <th>颜色</th>
                                    <th>尺寸</th>
                                    <th>库存</th>
                                    <th>选中</th>

                                </tr>
                                </thead>
                                <tbody id="skuTableTBody">
                                </tbody>
                            </table>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">数量</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="请输入数量" class="form-control" id="count">
                        </div>
                    </div>
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="createOutStockDetail();">提交</button>
            </div>
        </div>
    </div>
</div>

<!--//编辑出库单详情模态-->
<div class="modal fade" id="editOutStockDetailModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button disabled type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">编辑出库单详情</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
                    <div class="hr-line-dashed"></div>
                    <input id="outStockDetailId" hidden>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">货号</label>
                        <div class="col-sm-10">
                            <select class="form-control m-b"
                                    onchange="selectedArtNo({{.artnos}},'#edit_artNo','#edit_skuTable',1)"
                                    id="edit_artNo">
                                {{range .artnos}}
                                <option value="{{.Id}}">{{.ArtNo}}</option>
                                {{end}}
                            </select>

                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">出库商品</label>
                        <div class="col-sm-10">
                            <table id="edit_skuTable" style="table-layout:fixed"
                                   class="table table-striped table-bordered table-hover">
                                <thead>
                                <tr>
                                    <th>颜色</th>
                                    <th>尺寸</th>
                                    <th>库存</th>
                                    <th>选中</th>
                                </tr>
                                </thead>
                                <tbody id="edit_skuTableTBody">
                                </tbody>
                            </table>
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div id="editOutStockDetailCountFG" class="form-group">
                        <label class="col-sm-2 control-label">数量</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="请输入数量" class="form-control" id="edit_count">
                        </div>
                    </div>
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="editOutStockDetail();">提交</button>
            </div>
        </div>
    </div>
</div>
{{end}}


{{define "onload"}}
$("#stockManageMenu").click()
$('.dataTables-example').dataTable({
"order": [0,'desc']

});
{{end}}


{{define "js"}}
<!-- layerDate plugin javascript -->
<script src="/web-static/admin/js/plugins/layer/laydate/laydate.js"></script>
<script>

    function createOutStock() {
        var outStockNum = $('#outStockNum').val();
        var create_time = $("#create_time").val();
        var remark = $('#remark').val();
        var receiverTel = $('#receiverTel').val();
        var receiverName = $('#receiverName').val();
        var address = $('#address').val();
        var warehouseId = $('#warehouse').val();

        if (!outStockNum) {
            alert("出库单据号未生成,请重新打开");
            return
        }
        if (!create_time) {
            alert("请填写出库时间");
            return;
        }
        if (!receiverTel) {
            alert("请填写接受人电话");
            return;
        }
        if (!receiverName) {
            alert("请填写接受人");
            return;
        }
        if (!address) {
            alert("请填写地址");
            return;
        }
        if (warehouseId == -1) {
            alert("请填写入库仓库");
            return;
        }
        if (!warehouseId) {
            alert("请填写入库仓库");
            return;
        }
        $.ajax({
            url: '/admin/out_stock/create',
            type: 'post',
            data: {
                deliveryTime: create_time,
                remark: remark,
                warehouseId: warehouseId,
                outStockNum: outStockNum,
                receiverName: receiverName,
                receiverTel: receiverTel,
                address: address
            },
            success: function (data) {
                if (data.Success) {
                    alert("新增成功");
                    window.location.href = "/admin/out_stock"
                } else {
                    alert(data.ErrorMsg);
                }
            }
        })

    }


    function updateOutStock() {
        var outStockId = $('#edit_outStockId').val();
        var outStockNum = $('#edit_outStockNum').val();
        var create_time = $("#edit_create_time").val();
        var remark = $('#edit_remark').val();
        var receiverTel = $('#edit_receiverTel').val();
        var receiverName = $('#edit_receiverName').val();
        var address = $('#edit_address').val();
        var warehouseId = $('#edit_warehouse').val();

        if (!create_time) {
            alert("请填写出库时间");
            return;
        }
        if (!receiverTel) {
            alert("请填写接受人电话");
            return;
        }
        if (!receiverName) {
            alert("请填写接受人");
            return;
        }
        if (!address) {
            alert("请填写地址");
            return;
        }
        if (warehouseId == -1) {
            alert("请填写入库仓库");
            return;
        }
        if (!warehouseId) {
            alert("请填写入库仓库");
            return;
        }
        $.ajax({
            url: '/admin/out_stock/' + outStockId + '/update',
            type: 'post',
            data: {
                deliveryTime: create_time,
                remark: remark,
                warehouseId: warehouseId,
                outStockNum: outStockNum,
                receiverName: receiverName,
                receiverTel: receiverTel,
                address: address
            },
            success: function (data) {
                if (data.Success) {
                    alert("更新成功");
                    window.location.href = "/admin/out_stock"
                } else {
                    alert(data.ErrorMsg);
                }
            }
        })

    }

    function openEditOutStockModal(id, outStockNum, wareHouseId, deliveryTime, remark, receiverTel, receiverName, address, status) {

        if (status == 1) {
            $('#updateOutStockBtn').attr('disabled', true)

        } else {
            $('#updateOutStockBtn').removeAttr('disabled')
        }
        $('#edit_outStockId').val(id);
        $('#edit_outStockNum').val(outStockNum);
        $('#edit_warehouse').val(parseInt(wareHouseId));
        $('#edit_create_time').val(deliveryTime);
        $('#edit_remark').val(remark);
        $('#edit_receiverName').val(receiverName);
        $('#edit_receiverTel').val(receiverTel);
        $('#edit_address').val(address);

        $('#editOutStockModal').modal('show');
    }


    function openNewOutStockModal() {
        $.ajax({
            url: '/admin/out_stock/getOutStockNum',
            type: 'get',
            success: function (data) {
                if (data.Success) {
                    $('#outStockNum').val(data.Data);
                    $('#myModal5').modal("show");
                }
            }
        });
    }

    function openNewOutStockDetail() {
        checkedSku = null;
        $('#artNo').val(-1);
        $('#skuTableTBody').empty();
        $('#outStockNum').val(null);
        $('#newOutStockDetailModal').modal("show")
    }

    var editColorId;
    var editSizeId;

    function openEditOutStockDetailModal(id, artNoId, colorId, sizeId, count, can) {
        if (can == -2) {
            $('#editOutStockDetailCountFG').addClass("has-error")
        } else {
            $('#editOutStockDetailCountFG').removeClass("has-error")
        }
        checkedSku = null;
        console.info(id);
        $('#outStockDetailId').val(id);
        console.info($('#outStockDetailId').val());
        $('#editOutStockDetailModal').modal("show");
        $('#editOutStockDetailModal').on('shown.bs.modal', function () {
            $("#edit_artNo").val(parseInt(artNoId));
            editColorId = colorId;
            editSizeId = sizeId;
            $('#edit_count').val(count);
            $("#edit_artNo").change();
            var flag = 1;
            selectSkus.forEach(function (value) {
                if (artNoId == value.ArtNo.Id && value.Color.Id == colorId
                        && value.Size.Id == sizeId) {
                    selectSku(value.Id, 1);
                    flag = 0;
                }
            });
            if (flag == 1) {
                alert("该商品可能已经不存在或者被修改");

            }
        })


    }

    function editOutStockDetail() {
        if (checkedSku == null) {
            alert("请选择入库商品");
            return;
        }
        var artNoId = $('#edit_artNo').val();
        var colorId = checkedSku.Color.Id;
        var sizeId = checkedSku.Size.Id;
        var count = $('#edit_count').val();
        var outStockId = $('#outStockId').val();
        var outStockDetailId = $('#outStockDetailId').val();
        if (!count) {
            alert("请填写数量");
            return;
        }

        if (!isPositiveInteger(count)) {
            alert("请输入正确的数量");
            return;
        }
        if (count > checkedSku.Stock) {
            alert("不能大于库存量");
            return
        }

        $.ajax({
            url: '/admin/out_stock_detail/' + outStockDetailId + '/update',
            type: 'post',
            data: {
                artNoId: artNoId,
                colorId: colorId,
                sizeId: sizeId,
                count: count
            },
            success: function (data) {
                if (data.Success) {
                    alert("更新成功");
                    $('#editOutStockDetailModal').modal('hide');
                    $('#edit_count').val(null);
                    openOutStockDetailModal(0, outStockId);
                } else {
                    alert(data.ErrorMsg);
                }
            }
        });
    }

    function createOutStockDetail() {
        if (checkedSku == null) {
            alert("请选择入库商品");
            return;
        }

        var artNoId = $('#artNo').val();
        var colorId = checkedSku.Color.Id;
        var sizeId = checkedSku.Size.Id;
        var count = $('#count').val();
        var outStockId = $('#outStockId').val();

        if (artNoId == -1 || !colorId || !sizeId || !outStockId) {
            alert("请选择货号");
            return;
        }
        if (!count) {
            alert("请填写数量");
            return;
        }
        if (!isPositiveInteger(count)) {
            alert("请输入正确的数量");
            return;
        }
        if (count > checkedSku.Stock) {
            alert("不能大于库存量");
            return
        }
        $.ajax({
            url: '/admin/out_stock_detail/create',
            type: 'post',
            data: {
                artNoId: artNoId,
                colorId: colorId,
                sizeId: sizeId,
                count: count,
                outStockId: outStockId
            },
            success: function (data) {
                if (data.Success) {
                    alert("新增成功");
                    $('#newOutStockDetailModal').modal('hide');
                    $('#artNo').val(-1);
                    $('#count').val(null);
                    openOutStockDetailModal(0, outStockId);
                } else {
                    alert(data.ErrorMsg);
                }
            }
        })
    }

    function deleteOutStockDetail(id) {
        var result = window.confirm("是否删除?");
        if (result == true) {
            $.ajax({
                url: "/admin/out_stock_detail/" + id + "/delete",
                type: "delete",
                success: function (data) {
                    if (data.Success) {
                        alert('删除成功');
                        openOutStockDetailModal(0, $('#outStockId').val());

                    } else {
                        alert(data.ErrorMsg)
                    }
                },
                error: function (data) {
                    if (data.status == 401) {
                        alert("请重新登录!")
                    }
                }
            });

        }

    }

    function deleteOutStock(id) {
        var result = window.confirm("是否删除?");
        if (result == true) {
            $.ajax({
                url: "/admin/out_stock/" + id + "/delete",
                type: "delete",
                success: function (data) {
                    if (data.Success) {
                        alert('删除成功');
                        window.location.href = "/admin/out_stock";
                    } else {
                        alert(data.ErrorMsg)
                    }
                },
                error: function (data) {
                    if (data.status == 401) {
                        alert("请重新登录!")
                    }
                }
            });

        }

    }

    function submitOutStock(id) {
        var result = window.confirm("出库单提交之后不能在修改，确定是否提交?");
        if (result) {
            $.ajax({
                url: "/admin/out_stock/" + id + "/finish",
                type: "post",
                success: function (data) {
                    if (data.Success) {
                        alert('提交成功');
                        window.location.href = "/admin/out_stock";
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


    outStocks = [];
    function openOutStockDetailModal(status, outStockId) {
        if (status == 1) {
            $('#createOutStockBtn').attr('disabled', true)
            $('#createOutStockBtn').hide();
            $('#outStockDetailTableTip').hide();
        } else {
            $('#createOutStockBtn').removeAttr('disabled')
            $('#createOutStockBtn').show();
            $('#outStockDetailTableTip').show();
        }

        $.ajax({
            url: '/admin/out_stock/' + outStockId + "/out_stock_detail/list",
            type: 'get',
            success: function (data) {
                $('#outStockId').val(outStockId);
                $('#outStockDetailModal').modal("show");
                $('#outStockDetailTBody').empty();
                var tbody = null;
                if (data.Success) {
                    outStocks = data.Data;
                    data.Data.forEach(function (value) {
                        tbody += "<tr>";
                        tbody += "<td>" + value.Id + "</td>";
                        tbody += "<td>" + value.ArtNo.ArtNo + "</td>";
                        tbody += "<td>" + value.ArtNo.ProductName + "</td>";
                        tbody += "<td>" + value.Color.Color + "</td>";
                        tbody += "<td>" + value.Size.Size + "</td>";
                        tbody += "<td>" + value.Count + "</td>";
                        if (status == 0) {

                            if (value.Can == -1) {
                                tbody += "<td>失效</td>";
                                tbody += "<td><button class='btn btn-danger btn-sm' onclick='deleteOutStockDetail(" + value.Id + ")'> <i class='fa fa-trash'></i> 删除</button>" + "</td>";
                            } else if (value.Can == -2) {
                                tbody += "<td>库存不足</td>";
                                tbody += "<td>" + "<button class='btn btn-info btn-sm' onclick='openEditOutStockDetailModal(" + value.Id + "," + " &quot;" + value.ArtNo.Id + " &quot;" + "," + value.Color.Id + "," + value.Size.Id + "," + value.Count + "," + value.Can + ")'> <i class='fa fa-paste'></i> 编辑</button> <button class='btn btn-danger btn-sm' onclick='deleteOutStockDetail(" + value.Id + ")'> <i class='fa fa-trash'></i> 删除</button>" + "</td>";
                            } else {
                                tbody += "<td><i class='fa fa-check'></i></td>";
                                tbody += "<td>" + "<button class='btn btn-info btn-sm' onclick='openEditOutStockDetailModal(" + value.Id + "," + " &quot;" + value.ArtNo.Id + " &quot;" + "," + value.Color.Id + "," + value.Size.Id + "," + value.Count +  "," + value.Can + ")'> <i class='fa fa-paste'></i>编辑</button> <button class='btn btn-danger btn-sm' onclick='deleteOutStockDetail(" + value.Id + ")'> <i class='fa fa-trash'></i> 删除</button>" + "</td>";
                            }
                            $('#outStockDetailOper').removeAttr('hidden')
                        } else {
                            tbody += "<td><i class='fa fa-check'></i></td>";
                            $('#outStockDetailOper').attr('hidden', true)
                        }
                        tbody += "</tr>";
                    });
                    $('#outStockDetailTBody').append(tbody);
                }
            }

        })
    }


    var selectSkus = [];
    var checkedSku;
    /**
     * 当新增或者编辑出库详情的时候,货号选择框触发时调用
     * @param artNos 货号数组
     * @param artNo 当前选择的货号
     * @param table sku表格id
     * @param isEdit 区分是新增还是编辑 1是编辑 0是新增
     */
    function selectedArtNo(artNos, artNo, table, isEdit) {
        var artNoId = $(artNo).val();
        //未选择
        if (artNoId == -1) {
            $(table + "TBody").empty();
            return
        }
        $(table).dataTable({
            "destroy": true,
            "scrollY": "150px",
            "scrollCollapse": false,
            "searching": false,
            "ordering": false,
            "info": false,
            "paging": false
        });
        selectSkus = null;
        artNos.forEach(function (value) {
            if (value.Id == artNoId) {
                selectSkus = value.Skus;
                $(table + "TBody").empty();
                var tbody = null;
                selectSkus.forEach(function (value) {
                    tbody += "<tr>";
                    tbody += "<td>" + value.Color.Color + "</td>";
                    tbody += "<td>" + value.Size.Size + "</td>";
                    tbody += "<td>" + value.Stock + "</td>";
                    var flag = 1;
                    for (var i = 0; i < outStocks.length; i++) {

                        //如果是编辑入库详情，择需排除掉当前入库商品
                        if (value.ArtNo.Id == outStocks[i].ArtNo.Id && outStocks[i].Color.Id == value.Color.Id && outStocks[i].Size.Id == value.Size.Id) {
                            if (isEdit == 1) {
                                if (outStocks[i].Color.Id == editColorId
                                        && outStocks[i].Size.Id == editSizeId) {
                                    break;
                                }
                            }
                            tbody += "<td>" + "<span>已在该入库单中</span></td>";
                            flag = 0;
                            break;
                        }
                    }
                    if (flag == 1) {
                        if (isEdit == 0) {

                            tbody += "<td>" + "<input type='checkbox'   onchange='selectSku(" + value.Id + ",0)' value = '" + value.Id + "' id='" + "sku_check_" + value.Id + "'></td>";
                        } else {
                            tbody += "<td>" + "<input type='checkbox'   onchange='selectSku(" + value.Id + ",1)' value = '" + value.Id + "' id='" + "edit_sku_check_" + value.Id + "'></td>";

                        }

                    }

//                    tbody += "<td><i class='fa fa-check text-navy'></i></td>";
                    tbody += "</tr>";
                });
                $(table + "TBody").append(tbody);
            }
        });
    }
    /**
     * 当新增或者编辑出库详情的时候,出库商品选择时调用
     * @param skuId
     * @param isEdit 区分是新增还是编辑 1是编辑 0是新增
     */
    function selectSku(skuId, isEdit) {
        var checkId;
        if (isEdit == 1) {
            checkId = "#edit_sku_check_";
        } else {
            checkId = "#sku_check_";
        }
        checkedSku = null;
        selectSkus.forEach(function (value) {
            $(checkId + value.Id).prop("checked", false);
            if (value.Id == skuId) {
                $(checkId + value.Id).prop("checked", true);
                checkedSku = value;
            }
        });
    }


</script>

<script>
    //    //外部js调用
    //    laydate({
    //        elem: '#hello', //目标元素。由于laydate.js封装了一个轻量级的选择器引擎，因此elem还允许你传入class、tag但必须按照这种方式 '#id .class'
    //        event: 'focus' //响应事件。如果没有传入event，则按照默认的click
    //    });

    //日期范围限制
    var start = {
        elem: '#create_time',
        format: 'YYYY-MM-DD hh:mm:ss',
        event: 'focus',
        location: 'top',
        max: '2099-06-16 23:59:59', //最大日期
        istime: true,
        fixed: false,
        istoday: false
    };
    var edit_start = {
        elem: '#edit_create_time',
        format: 'YYYY-MM-DD hh:mm:ss',
        event: 'focus',
        location: 'top',
        max: '2099-06-16 23:59:59', //最大日期
        istime: true,
        fixed: false,
        istoday: false
    };
    laydate(start);
    laydate(edit_start);
</script>

{{end}}