<!DOCTYPE html>
<html>

<head>
    <script src="/web-static/admin/js/jquery.min.js"></script>

    {{template "meta" .}}
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="renderer" content="webkit">
    <meta name="keywords" content="">
    <meta name="description" content="">
    <link rel="shortcut icon" href="favicon.ico">
    <link href="/web-static/admin/css/bootstrap.min.css" rel="stylesheet">
    <link href="/web-static/admin/css/plugins/dataTables/dataTables.bootstrap.css" rel="stylesheet">

    <link href="/web-static/admin/css/font-awesome.min.css?v=4.4.0" rel="stylesheet">
    <link href="/web-static/admin/css/animate.css" rel="stylesheet">
    <link href="/web-static/admin/css/style.css?v=4.1.0" rel="stylesheet">

    <link href="/web-static/admin/css/plugins/steps/jquery.steps.css" rel="stylesheet">
    <link href="/web-static/admin/css/plugins/iCheck/custom.css" rel="stylesheet">
    <!-- Data Tables -->
    {{template "css" .}}
    <style>
        .modal-backdrop {
            z-index: 0 !important;
            pointer-events: none;
        }

        .modal-dialog {
            margin-top: 13.5% !important;
        }

    </style>
</head>

<body class="fixed-sidebar full-height-layout gray-bg" style="overflow:hidden">
<div id="wrapper">
    <!--左侧导航开始-->
    <nav class="navbar-default navbar-static-side" role="navigation">
        <div class="nav-close"><i class="fa fa-times-circle"></i>
        </div>
        <div class="sidebar-collapse" style="">
            <ul class="nav" id="side-menu">
                <li class="nav-header">
                    <div class="dropdown profile-element">
                        <a data-toggle="dropdown" class="dropdown-toggle" href="#">
                                <span class="clear">
                                    <span class="block m-t-xs" style="font-size:20px;">
                                        <i class="fa fa-shopping-cart"> </i>
                                        <strong class="font-bold">Cims</strong>
                                    </span>
                                </span>
                        </a>
                    </div>
                    <div class="logo-element">Cims
                    </div>
                </li>
                <li>
                    <a class="J_menuItem" href="/admin">
                        <i class="fa fa-home"></i>
                        <span class="nav-label">主页</span>
                    </a>
                </li>

                <li>
                    <a href="#" id="systemManageMenu"><i class="fa fa-shopping-cart"></i> <span
                            class="nav-label">系统管理</span><span class="fa arrow"></span></a>
                    <ul class="nav nav-second-level">
                        <li><a href="/admin/user" class="J_menuItem">用户管理 <span
                                class="label label-info pull-right"></span></a>
                        </li>
                        <li><a href="/admin/art_no" class="J_menuItem">货号管理 <span
                                class="label label-info pull-right"></span></a>
                        </li>
                        <li><a href="/admin/ware_house" class="J_menuItem">仓库管理 <span
                                class="label label-info pull-right"></span></a>
                        </li>
                    </ul>
                </li>
                <li>
                    <a href="#" id="stockManageMenu">
                        <i class="fa fa-flask"></i>
                        <span class="nav-label">库存管理</span><span class="fa arrow"></span>
                        <span class="label label-warning pull-right" id="orderInfo"></span>
                    </a>
                    <ul class="nav nav-second-level">
                        <li>
                            <a class="J_menuItem" href="/admin/in_stock">入库管理</a>
                        </li>
                        <li><a href="/admin/out_stock" class="J_menuItem">出库管理 <span
                                class="label label-info pull-right"></span></a>
                        </li>
                    </ul>

                </li>

                <!--<li>-->
                    <!--<a href="#" id="articleMenu">-->
                        <!--<i class="fa fa-flask"></i>-->
                        <!--<span class="nav-label">文章管理</span><span class="fa arrow"></span>-->
                        <!--<span class="label label-warning pull-right" id="article"></span>-->
                    <!--</a>-->
                    <!--<ul class="nav nav-second-level">-->
                        <!--<li>-->
                            <!--<a class="J_menuItem" href="/admin/new">文章列表</a>-->
                        <!--</li>-->
                    <!--</ul>-->

                <!--</li>-->

            </ul>
        </div>
    </nav>
    <!--左侧导航结束-->
    <!--右侧部分开始-->
    <div id="page-wrapper" class="gray-bg dashbard-1">
        <div class="row border-bottom">
            <nav class="navbar navbar-static-top" role="navigation" style="margin-bottom: 0">
                <ul class="nav navbar-top-links navbar-right">
                    <li><a href="#">{{.user.Username}}</a></li>
                    <li><a href="#" data-toggle="modal" data-target="#editPasswordModal">修改密码</a></li>
                    <li><a href="/admin/logout"><i class="fa fa-sign-out"></i>注销</a></li>

                </ul>
            </nav>
        </div>
        <div class="row J_mainContent" id="content-main">
            <!--代码部分-->
            <div class="gray-bg">
                <div class="wrapper wrapper-content">
                    {{template "content" .}}

                </div>
            </div>


        </div>
    </div>
    <!--右侧部分结束-->
</div>

<!--//密码修改模态-->
<div class="modal fade" id="editPasswordModal" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="modal-dialog modal-md">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span
                        aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">密码修改</h4>
            </div>
            <div class="modal-body">
                <form role="form">
                    <div class="form-group">
                        <input type="hidden" class="form-control" name="userId" id="edit_user_id">
                    </div>

                    <div class="form-group">
                        <label>旧密码</label>
                        <input type="password" name="password" placeholder="输入旧密码"
                               class="form-control"
                               id="edit_old_password">
                    </div>

                    <div class="form-group">
                        <label>新密码</label>
                        <input type="password" name="password" placeholder="输入新密码"
                               class="form-control"
                               id="edit_new_password">
                    </div>

                    <div class="form-group">
                        <label>确认密码</label>
                        <input type="password" name="password" placeholder="输入新密码"
                               class="form-control"
                               id="edit_re_new_password">
                    </div>
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="changePassword();">提交
                </button>
            </div>
        </div>
    </div>
</div>


<!-- 全局js -->
<script src="/web-static/admin/js/bootstrap.min.js?v=3.3.6"></script>
<script src="/web-static/admin/js/plugins/metisMenu/jquery.metisMenu.js"></script>
<script src="/web-static/admin/js/plugins/slimscroll/jquery.slimscroll.min.js"></script>
<script src="/web-static/admin/js/plugins/layer/layer.min.js"></script>


<!-- 自定义js -->
<script src="/web-static/admin/js/hAdmin.js?v=4.1.0"></script>
<script type="text/javascript" src="/web-static/admin/js/index.js"></script>
<script src="/web-static/admin/js/validate.js"></script>

<!-- 第三方插件 -->
<script src="/web-static/admin/js/plugins/pace/pace.min.js"></script>

<!-- Data Tables -->
<script src="/web-static/admin/js/plugins/jeditable/jquery.jeditable.js"></script>
<script src="/web-static/admin/js/plugins/dataTables/jquery.dataTables.js"></script>
<script src="/web-static/admin/js/plugins/dataTables/dataTables.bootstrap.js"></script>
<script src="/web-static/admin/js/content.js?v=1.0.0"></script>

<!--引入yifanToast-->
<link rel="stylesheet" href="/web-static/admin/css/yifanToast.css">
<script src="/web-static/admin/js/yifanToast.js" type="text/javascript"></script>


<!-- alert -->
<link rel="stylesheet" href="/web-static/admin/css/plugins/sweetalert/sweetalert.css">
<script src="/web-static/admin/js/plugins/sweetalert/sweetalert.min.js"></script>
<script src="/web-static/admin/js/sweetalertUtil.js"></script>

<!-- Steps -->
<script src="/web-static/admin/js/plugins/staps/jquery.steps.min.js"></script>

<!-- Jquery Validate -->
<script src="/web-static/admin/js/plugins/validate/jquery.validate.min.js"></script>
<script src="/web-static/admin/js/plugins/validate/messages_zh.min.js"></script>

{{template "js" .}}
<script>
    function changePassword() {
        var oldPassword = $('#edit_old_password').val();
        var newPassword = $('#edit_new_password').val();
        var newRePassword = $('#edit_re_new_password').val();
        if (!oldPassword || !newPassword || !newRePassword) {
            alert("密码不能为空");
            return;
        }
        if (newPassword != newRePassword) {
            alert("两次输入的密码不一致");
            return;
        }
        $.ajax({
            url: '/admin/password/update',
            type: 'post',
            data: {
                oldPassword: oldPassword,
                newPassword: newPassword
            },
            success: function (data) {
                if (data.Success) {
                    alert("密码修改成功");
                    window.location.href = "/admin/login"
                } else {
                    alert(data.ErrorMsg)
                }
            }
        })

    }
</script>
</body>

<script>
    window.onload = function () {
        {{template "onload" .}}
    };

</script>

</html>
