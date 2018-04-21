{{ template "admin/base.tpl" .}}
{{ define "meta"}}
<title>文章管理</title>
{{end}}

{{ define "css"}}

{{end}}

{{define "content"}}
<body>
<div id="content">

</div>
<button class="btn btn-info" onclick="submit()">提交</button>
<script id="container" name="content" type="text/plain" style="height:300px">
        这里写你的初始化内容


</script>


</body>
{{end}}
{{define "onload"}}


{{end}}
<script>
    var ue = UE.getEditor('container');
    function submit() {
        console.info(ue.getContent());
        $('#content').append(ue.getContent());
    }
</script>

{{define "js"}}
<!-- 加载编辑器的容器 -->

<!-- 配置文件 -->
<script type="text/javascript" src="/web-static/admin/ueditor/ueditor.config.js"></script>
<!-- 编辑器源码文件 -->
<script type="text/javascript" src="/web-static/admin/ueditor/ueditor.all.js"></script>
<!-- 实例化编辑器 -->

{{end}}