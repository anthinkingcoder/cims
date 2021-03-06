<!DOCTYPE html>
<html>
<head>
    <title>登录</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="/web-static/admin/css/bootstrap.min.css">
    <style>
        html,body{
            height: 100%;
        }
        section.registerSection{
            text-align: center;
            margin:30px
        }
        .registerSection input{
            max-width: 400px;
            height: 40px;
            margin: 0 auto;
        }
        .margin-top-15{
            margin-top: 15px;
        }
        section.registerSection h3{
            color:#999;
            margin-bottom: 20px;
        }
        #registerButton{
            margin-top: 15px;
            margin-bottom: 5px;
            width: 150px;
        }
        #loginButton{
            margin-top: 15px;
        }
        .registerSection a{
            margin-top: 15px;
        }
        .background{
            background-color: #ddd;
            height: 100%;
        }
        .colorLogin{
            background-color: rgba(68,181,73,0.6);
            border-color:rgba(68,181,73,0.5);
        }
        .colorLogin:hover{
            background-color: rgba(68,181,73,0.8);
            border-color:rgba(68,181,73,0.8);
        }

        #username,#password{
            color:white;
            background-color: rgba(68,181,73,0.1);

        }
        input:-webkit-autofill {
            -webkit-box-shadow: 0 0 0px 1000px white inset;
            -webkit-text-fill-color: #333;
        }
        #username:focus,#password:focus{
            border:1px solid rgba(68,181,73,0.5);
        }
    </style>
</head>
<body onkeydown="keyLogin();">
<div class="background">
    <header><div style="height: 180px"></div></header>
    <section class="registerSection">
        <h3>服装库存管理系统</h3>
        <form autocomplete="off">
            <input id="username" type="text" class="form-control" value="" placeholder="username">
            <div class="margin-top-15"></div>
            <input id="password" type="password" class="form-control" value="" placeholder="password">

        </form>

        <button id="registerButton" class="btn btn-info colorLogin" onclick="login()"> 登录 </button><br>
    </section>
</div>
</body>

<script src="/web-static/admin/js/jquery.min.js"></script>
<script src="/web-static/admin/js/bootstrap.min.js"></script>
<!--引入yifanToast-->
<link rel="stylesheet" href="/web-static/admin/css/yifanToast.css">
<script src="/web-static/admin/js/yifanToast.js" type="text/javascript"></script>
<script>
    function keyLogin(){
        if (event.keyCode==13)  //回车键的键值为13
            login();
    }

    function login() {
        var username = $("#username").val();
        var password = $("#password").val();
        if(username==""||password==""){
            yifanToast('请填写完整信息！','toast-md');
        }else{
            $.ajax({
                url: '/admin/authorize',
                type: 'post',
                dataType: 'json',
                data:{
                    'username':username,
                    'password':password
                },
                contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
                cache: false,
                success: function(data) {
                    console.info(data)
                    if(data.Success) {
                        window.location.href = "/admin/index"
                    }else {
                        yifanToast(data.ErrorMsg,'toast-md');
                    }

                },
                error : function() {
                    yifanToast('服务器失去响应，请检查您的网络！','toast-md');
                }
            });
        }


    }


</script>
</html>
