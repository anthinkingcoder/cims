function deleteTip(callback) {
    swal({
        title: "您确定要删除这条信息吗",
        text: "删除后将无法恢复，请谨慎操作！",
        type: "warning",
        showCancelButton: true,
        confirmButtonColor: "#DD6B55",
        confirmButtonText: "删除",
        closeOnConfirm: true,
        cancelButtonText: "取消"
    }, function (isConfirm) {
        console.info(isConfirm)
        if(isConfirm) {
            callback()
        }
    });
}

