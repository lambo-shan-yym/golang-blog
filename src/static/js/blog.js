$(document).ready(function () {
    body = {id: parseInt($("#blog_id_hidden").text())};
    console.log(body)
    $.ajax({
        async: true,
        type: 'POST',
        url: "/blog/get",
        contentType: 'application/json',
        data: JSON.stringify(body),
        success: function (data) {
            //后台用JsonResponse返回数据
            //data 就会被转成字典
            console.log(data)
            //JSON.parse(data) 把字符串类型转成字典
            if (data.code == 0) {
                var blog = data.result
                var content = blog["content"]
                $("#blogContent").html(content)
            } else {

            }
        },
        error: function (data) {
            console.log("发生异常")
        }
    });
})