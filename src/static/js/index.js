$(document).ready(function () {
    body = {limit: 10, offset: 0};
    $.ajax({
        async: true,
        type: 'POST',
        url: "/tag/find",
        contentType: 'application/json',
        data: JSON.stringify(body),
        success: function (data) {
            //后台用JsonResponse返回数据
            //data 就会被转成字典
            console.log(data)
            //JSON.parse(data) 把字符串类型转成字典
            var content = "<ul><li>"
            if (data.code == 0) {
                var tagList = data.result["list"]
                for (j = 0, len = tagList.length; j < len; j++) {
                    var tag = tagList[j]
                    content+="<a href=\"/tags/"+tag["id"]+"\" style=\"background: "+tag["color"]+"\"> <span >"+tag["name"]+"</span>\n" +
                        "                                    <div class=\"m-inline-block\">3</div>\n" +
                        "                                </a>";
                }
                content +="</li></ul>"
                $("#tag_content").html(content)
            } else {

            }
        },
        error: function (data) {
            console.log("发生异常")
        }
    });








    typeBody = {limit: 5, offset: 0};
    $.ajax({
        async: true,
        type: 'POST',
        url: "/type/find",
        contentType: 'application/json',
        data: JSON.stringify(typeBody),
        success: function (data) {
            //后台用JsonResponse返回数据
            //data 就会被转成字典
            console.log(data)
            //JSON.parse(data) 把字符串类型转成字典
            var content = ""
            if (data.code == 0) {
                var typeList = data.result["list"]
                for (j = 0, len = typeList.length; j < len; j++) {
                    var type = typeList[j]
                    content+=" <a href=\"/types/"+type["id"]+"\" class=\"item\">\n" +
                        "                                <span>"+type["name"]+"</span>\n" +
                        "                                <div class=\"m-pink-label\">11</div>\n" +
                        "                            </a>"
                }
                $("#type_content").html(content)
            } else {

            }
        },
        error: function (data) {
            console.log("发生异常")
        }
    });





    blogBody = {limit: 10, offset: 0};
    $.ajax({
        async: true,
        type: 'POST',
        url: "/blog/find",
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
                for (j = 0, len = blogList.length; j < len; j++) {
                    var blog = blogList[j]
                    content+=" <div class=\"ui padded vertical segment  m-padding-tb-large m-mobile-clear\">\n" +
                        "                        <div class=\"ui mobile reversed stackable grid\">\n" +
                        "                            <div class=\"ui eleven wide column\">\n" +
                        "                                <h3 class=\"ui header\"><a href=\""+"/blog?id="+blog["id"]+"\" target=\"_blank\"\n" +
                        "                                                         class=\"m-black\">"+blog["title"]+"</a></h3>\n" +
                        "                                <p class=\"m-text\">"+ blog["description"]+
                        "                                </p>\n" +
                        "                                <div class=\"ui stackable grid\">\n" +
                        "                                    <div class=\"eleven wide column\">\n" +
                        "                                        <div class=\"ui mini horizontal link list\">\n" +
                        "                                            <div class=\"item\">\n" +
                        "                                                <img src=\"/static/images/admin_header.jpg\" alt=\"\"\n" +
                        "                                                     class=\"ui avatar image\">\n" +
                        "                                                <div class=\"content\"><a href=\"/about\" style=\"color: #2E64FE!important;\"\n" +
                        "                                                                        class=\"header m-pink\">管理员</a></div>\n" +
                        "                                            </div>\n" +
                        "                                            <div class=\"item\">\n" +
                        "                                                <a href=\"/archives\"><i\n" +
                        "                                                            class=\"calendar icon m-pink\"></i><span>2020-06-07</span></a>\n" +
                        "\n" +
                        "                                                <!--                                          <i class=\"calendar icon m-pink\"></i><a href=\"javascript:void(0);\">2019-11-20</a>-->\n" +
                        "                                            </div>\n" +
                        "                                            <div class=\"item\">\n" +
                        "                                                <i class=\"eye icon\"></i> <span>"+blog["view_count"]+"</span>\n" +
                        "                                            </div>\n" +
                        "                                            <!--                                              <div class=\"item\">-->\n" +
                        "                                            <!--                                                  <a th:href=\"@{/blog/{id}#comment-form(id=${blog.id})}\"><i class=\"comment icon m-pink\"></i> <span th:text=\"${blog.views}\">2342</span></a>-->\n" +
                        "                                            <!--                                              </div>-->\n" +
                        "\n" +
                        "                                            <!--<div class=\"item\">\n" +
                        "                                                <i class=\"tag icon m-pink\"></i>\n" +
                        "                                                <a th:text=\"${blog.typeName}\"  href=\"#\">html</a>\n" +
                        "                                            </div>-->\n" +
                        "                                        </div>\n" +
                        "                                    </div>\n" +
                        "                                    <div class=\"right aligned five wide column\">\n" +
                        "                                        <a href=\"/types/8\" target=\"_blank\"\n" +
                        "                                           class=\"m-pink-label m-padding-tiny m-text-thin\">Java框架</a>\n" +
                        "                                    </div>\n" +
                        "                                </div>\n" +
                        "\n" +
                        "                            </div>\n" +
                        "\n" +
                        "\n" +
                        "                            <div class=\"ui five wide column\"><!--图片-->\n" +
                        "                                <a href=\"/blog/1591456383956\" target=\"_blank\">\n" +
                        "                                    <img src=\""+blog["first_picture_url"]+"\"\n" +
                        "                                         alt=\"\" class=\"ui m-image rounded image\">\n" +
                        "                                </a>\n" +
                        "                            </div>\n" +
                        "                        </div>\n" +
                        "                    </div> "
                }
                $("#blog_content").html(content)
            } else {

            }
        },
        error: function (data) {
            console.log("发生异常")
        }
    });




    $.ajax({
        async: true,
        type: 'GET',
        url: "/blog/find_click_top_10",
        contentType: 'application/json',
        success: function (data) {
            //后台用JsonResponse返回数据
            //data 就会被转成字典
            console.log(data)
            //JSON.parse(data) 把字符串类型转成字典
            var content = " <div class=\"ui segment m-bg m-top-color\">\n" +
                "                        <i class=\"bookmark icon m-pink m-bg\"></i>点击排行\n" +
                "                    </div>"
            if (data.code == 0) {
                var blogList = data.result["list"]
                for (j = 0, len = blogList.length; j < len; j++) {
                    var blog = blogList[j]
                    content+="  <div class=\"ui m-bg segment\">\n" +
                        "                        <a href=\"/blog/"+blog["id"]+"\" class=\"m-black  m-text-thin\">"+blog["title"]+"</a>\n" +
                        "                    </div> "
                }
                 $("#click_top10_content").html(content)
            } else {

            }
        },
        error: function (data) {
            console.log("发生异常")
        }
    });





    $.ajax({
        async: true,
        type: 'GET',
        url: "/blog/find_new_recommend",
        contentType: 'application/json',
        success: function (data) {
            //后台用JsonResponse返回数据
            //data 就会被转成字典
            console.log(data)
            //JSON.parse(data) 把字符串类型转成字典
            var content = " <div class=\"ui segment m-bg m-top-color\">\n" +
                "                        <i class=\"bookmark icon m-pink m-bg\"></i>最新推荐\n" +
                "                    </div> "
            if (data.code == 0) {
                var blogList = data.result["list"]
                for (j = 0, len = blogList.length; j < len; j++) {
                    var blog = blogList[j]
                    content+="  <div class=\"ui m-bg segment\">\n" +
                        "                        <a href=\"/blog/"+blog["id"]+"\" class=\"m-black  m-text-thin\">"+blog["title"]+"</a>\n" +
                        "                    </div>"
                }
                $("#new_recommend_content").html(content)
            } else {

            }
        },
        error: function (data) {
            console.log("发生异常")
        }
    });

})