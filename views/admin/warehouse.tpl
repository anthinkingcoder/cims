{{ template "admin/base.tpl" .}}
{{ define "meta"}}
<title>系统管理</title>
{{end}}

{{ define "css"}}

{{end}}

{{define "content"}}
<div class="row">
    <div class="col-sm-12">
        <div class="ibox float-e-margins">
            <div class="ibox-title">
                <h5>仓库管理
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
                    <button class="btn btn-info" data-toggle="modal" data-target="#myModal5">新增仓库
                        <i class="fa fa-plus"></i></button>
                </div>
                <div style="height: 10px"></div>
                <table id="user_manager_table"
                       class="table table-striped table-bordered table-hover dataTables-example">
                    <thead>
                    <tr>
                        <th>序号</th>
                        <th>仓库编号</th>
                        <th>仓库名称</th>
                        <th>仓储量</th>
                        <th>当前库存</th>
                        <th>创建时间</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range .warehouses }}
                    <tr>
                        <td>{{.Id}}</td>
                        <td><a href="#" onclick="openEditWareHouseModal({{.}})">{{.WareHouseNum}}</a></td>
                        <td>{{.Name}}</td>
                        <td>{{.Capacity}}</td>
                        <td>{{.Stock}}</td>
                        <td>{{date .GmtCreate "Y-m-d h:i:s"}}</td>
                        <td>
                            <button class="btn btn-danger btn-sm"
                                    onclick="deleteWareHouse({{.Id }})">
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


<!--//新增仓库模态-->
<div class="modal fade" id="myModal5" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">新增仓库</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
                    <div class="form-group">
                        <label class="col-sm-2 control-label">仓库编号</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入仓库编号" class="form-control" id="wareHouseNum">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">仓库名称</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入仓库名称" class="form-control" id="name">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">联系电话</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入联系电话" class="form-control" id="tel">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">联系人</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入联系人" class="form-control" id="contact">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">仓储量</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入仓储量" class="form-control" id="capacity">
                        </div>
                    </div>


                    <div class="hr-line-dashed"></div>
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="createWareHouse();">提交</button>
            </div>
        </div>
    </div>
</div>


<!--//编辑仓库模态-->
<div class="modal fade" id="editWareHouseModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">编辑仓库</h4>
            </div>
            <div class="modal-body">
                <form class="form-horizontal">
                    <input id="edit_id" hidden>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">仓库编号</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入仓库编号" class="form-control" id="edit_wareHouseNum">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">仓库名称</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入仓库名称" class="form-control" id="edit_name">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">联系电话</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入联系电话" class="form-control" id="edit_tel">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">联系人</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入联系人" class="form-control" id="edit_contact">
                        </div>
                    </div>
                    <div class="hr-line-dashed"></div>
                    <div class="form-group">
                        <label class="col-sm-2 control-label">仓储量</label>
                        <div class="col-sm-10">
                            <input type="text" placeholder="输入仓储量" class="form-control" id="edit_capacity">
                        </div>
                    </div>
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="editWareHouse();">保存</button>
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
{{end}}


{{define "js"}}
<script>

    function openEditWareHouseModal(wareHouse) {
        $('#edit_wareHouseNum').val(wareHouse.WareHouseNum);
        $('#edit_id').val(wareHouse.Id);
        $('#edit_stock').val(wareHouse.Stock);
        $('#edit_name').val(wareHouse.Name);
        $('#edit_tel').val(wareHouse.Tel);
        $('#edit_contact').val(wareHouse.Contact);
        $('#edit_capacity').val(wareHouse.Capacity);
        $('#editWareHouseModal').modal('show');

    }

   function editWareHouse() {
       var id = $('#edit_id').val();
       var wareHouseNum = $('#edit_wareHouseNum').val();
       var name = $('#edit_name').val();
       var contact = $('#edit_contact').val();
       var tel = $('#edit_tel').val();
       var capacity = $('#edit_capacity').val();
       if (!name || !wareHouseNum || !tel || !contact || !capacity) {
           alert("填写完整");
           return
       }
       $.ajax({
           url: '/admin/ware_house/' + id +  '/update',
           type: 'post',
           data: {
               wareHouseNum: wareHouseNum,
               name: name,
               tel: tel,
               contact: contact,
               capacity: capacity
           },
           success: function (data) {
               if (data.Success) {
                   alert("更新成功");
                   window.location.href = "/admin/ware_house"
               } else {
                   alert(data.ErrorMsg);
               }
           },
           error: function (data) {
               console.info(data)
           }
       })


   }

    function deleteWareHouse(id) {
        var result = window.confirm("是否删除该仓库");
        if (result == true) {
            $.ajax({
                url: '/admin/ware_house/' + id + "/delete",
                type: 'delete',
                success: function (data) {
                    if (data.Success) {
                        alert("删除成功");
                        window.location.href = "/admin/ware_house"
                    } else {
                        alert(data.ErrorMsg);
                    }
                }, error: function (data) {
                    if (data.status == 401) {
                        alert("请重新登录!")
                    }
                }
            })
        }

    }


    function createWareHouse() {
        var wareHouseNum = $('#wareHouseNum').val();
        var name = $('#name').val();
        var tel = $('#tel').val();
        var contact = $('#contact').val();
        var capacity = $('#capacity').val();

        if (!name || !wareHouseNum || !tel || !contact) {
            alert("填写完整");
            return
        }
        $.ajax({
            url: '/admin/ware_house/create',
            type: 'post',
            data: {
                wareHouseNum: wareHouseNum,
                name: name,
                tel: tel,
                contact: contact,
                capacity: capacity
            },
            success: function (data) {
                if (data.Success) {
                    alert("创建成功");
                    window.location.href = "/admin/ware_house"
                } else {
                    alert(data.ErrorMsg);
                }
            },
            error: function (data) {
                console.info(data)
            }
        })
    }


</script>

{{end}}