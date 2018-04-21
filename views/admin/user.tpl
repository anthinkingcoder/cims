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
                <h5>用户管理
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
                    <button class="btn btn-info" data-toggle="modal" data-target="#myModal5">新增用户
                        <i class="fa fa-plus"></i></button>
                </div>
                <div style="height: 10px"></div>
                <table id="user_manager_table"
                       class="table table-striped table-bordered table-hover dataTables-example">
                    <thead>
                    <tr>
                        <th>序号</th>
                        <th>用户名</th>
                        <th>用户姓名</th>
                        <th>注册时间</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range $index, $user := .users }}
                    <tr>
                        <td>{{$user.Id}}</td>
                        <td><a href="#" onclick="openEditUserModal({{.}})">{{.Username}}</a></td>
                        <td>{{$user.Name}}</td>
                        <td>{{date $user.GmtCreate "Y-m-d h:i:s"}}</td>
                        <td>
                            <button class="btn btn-danger btn-sm"
                                    onclick="deleteNormalUser({{ $user.Id }})">
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


<!--//新增用户模态-->
<div class="modal fade" id="myModal5" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">新增用户</h4>
            </div>
            <div class="modal-body">
                <form role="form">
                    <div class="form-group">
                        <label>登录名</label>
                        <input type="text" placeholder="输入用户名" class="form-control" id="username">
                    </div>
                    <div class="form-group">
                        <label>密码</label>
                        <input type="password" placeholder="输入登录密码" class="form-control"
                               id="password">
                    </div>
                    <div class="form-group">
                        <label>姓名</label>
                        <input type="text" placeholder="姓名" class="form-control" name="name" id="name">
                    </div>

                    <div class="form-group">
                        <label>简介</label>
                        <textarea class="form-control" rows="6" name="profile" id="profile"></textarea>
                    </div>


                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="createSubmit();">提交</button>
            </div>
        </div>
    </div>
</div>

<!--//更新用户模态-->
<div class="modal fade" id="editUserModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">用户信息</h4>
            </div>
            <div class="modal-body">
                <form role="form">
                    <input id="uid" hidden/>
                    <div class="form-group">
                        <label>登录名</label>
                        <input type="text" placeholder="输入用户名" class="form-control" id="edit_username">
                    </div>
                    <div class="form-group">
                        <label>密码</label>
                        <input type="password" placeholder="输入登录密码" class="form-control"
                               id="edit_password">
                    </div>
                    <div class="form-group">
                        <label>姓名</label>
                        <input type="text" placeholder="姓名" class="form-control" name="name" id="edit_name">
                    </div>

                    <div class="form-group">
                        <label>简介</label>
                        <textarea class="form-control" rows="6" name="profile" id="edit_profile"></textarea>
                    </div>


                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="editUser();">提交</button>
            </div>
        </div>
    </div>
</div>


{{end}}


{{define "onload"}}
$("#systemManageMenu").click()
$('#user_manager_table').dataTable({
"order": [0,'desc']
});
{{end}}


{{define "js"}}
<script>


    function deleteNormalUser(id) {
        var result = window.confirm("是否删除该用户");
        if (result == true) {
            $.ajax({
                url: '/admin/user/' + id + "/delete",
                type: 'delete',
                success: function (data) {
                    if (data.Success) {
                        window.location.href = "/admin/user"
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

    function createSubmit() {
        var username = $('#username').val();
        var password = $('#password').val();
        var name = $('#name').val();
        var profile = $('#profile').val();
        if (!username || !password || !name) {
            alert("填写完整");
            return;
        }
        $.ajax({
            url: '/admin/user/create',
            type: 'post',
            data: {
                username: username,
                password: password,
                name: name,
                profile: profile
            },
            success: function (data) {
                if (data.Success) {
                    alert("创建成功");
                    window.location.href = "/admin/user"
                } else {
                    alert(data.ErrorMsg);
                }
            },
            error: function (data) {
                console.info(data)
            }
        })
    }

    function openEditUserModal(user) {
       $('#uid').val(user.Id);
        $('#edit_username').val(user.Username);
        $('#edit_password').val(user.Password);
        $('#edit_name').val(user.Name);
        $('#edit_profile').val(user.Profile);
        $('#editUserModal').modal("show");
    }
    function editUser() {
        var uid = $('#uid').val();
        var username = $('#edit_username').val();
        var password = $('#edit_password').val();
        var name = $('#edit_name').val();
        var profile = $('#edit_profile').val();
        if (!username || !password || !name) {
            alert("填写完整");
            return;
        }
        $.ajax({
            url: '/admin/user/' + uid + '/update',
            type: 'post',
            data: {
                username: username,
                password: password,
                name: name,
                profile: profile
            },
            success: function (data) {
                if (data.Success) {
                    alert("更新成功");
                    window.location.href = "/admin/user"
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