$(document).ready(function () {
    blogBody = {limit: 200, offset: 0};
    $.ajax({
        async: true,
        type: 'POST',
        url: "/blog/find_for_time",
        contentType: 'application/json',
        data: JSON.stringify(blogBody),
        success: function (data) {
            //后台用JsonResponse返回数据
            //data 就会被转成字典
            console.log(data)
            //JSON.parse(data) 把字符串类型转成字典
            var content = ""
            if (data.code == 0) {
                var blogList = data.result["list"]
                var map = new Map();
                for (j = 0, len = blogList.length; j < len; j++) {
                    var blog = blogList[j]
                    if (!map.has(blog["create_year"])){
                        map.set(blog["create_year"],[])
                    }
                    map.get(blog["create_year"]).push(blog)
                }
                var content =""
                content +="<div class=\"ui top attached padded segment m-bg animated zoomIn\">\n" +
                    "            <div class=\"ui middle aligned two column grid \">\n" +
                    "                <div class=\"column\">\n" +
                    "                    <h3 class=\"m-pink\">归档</h3>\n" +
                    "                </div>\n" +
                    "                <div class=\"right aligned  column\">\n" +
                    "                    共 <h3 class=\"ui m-pink header m-inline-block m-text-thin\">15</h3>&nbsp;篇博客\n" +
                    "                </div>\n" +
                    "            </div>\n" +
                    "        </div>"
                map.forEach(function(value,key){
                    content+=" <h3 class=\"ui center animated m-pink fadeInUp aligned header\">"+key+"</h3>\n  <div class=\"ui fluid vertical menu animated fadeInUp m-bg\">\n"
                    for(var i=0;i<value.length;i++){
                        var item = value[i];
                        var blog = item["blog"]
                        console.log(blog)
                        content+=" <a href=\"/blog?id="+blog["id"]+"\" target=\"_blank\"\n" +
                            "               class=\"item\">\n" +
                            "          <span>\n" +
                            "            <i class=\"\"></i><span>"+blog["title"]+"</span>\n" +
                            "            <div class=\"ui mini m-white  m-bg-pink left pointing  label m-padded-mini\">"+item["create_month_day"]+"</div>\n" +
                            "          </span>\n" +
                            "                <div class=\" m-pink-label m-padded-mini \">原创</div>\n" +
                            "            </a> "
                    }
                    content +="</div>"
                });
                $("#archivesContent").html(content)
            } else {

            }
        },
        error: function (data) {
            console.log("发生异常")
        }
    });
})