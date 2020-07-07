package router

import (
	"github.com/gin-gonic/gin"
	"golang-blog/src/common/error"
	"golang-blog/src/controller"
	"golang-blog/src/service"
)

var blogController controller.BlogController
var typeController controller.TypeController
var tagController controller.TagController
var htmlController controller.HtmlController

func init() {
	blogService := service.NewBlogServiceImpl()
	blogController = controller.BlogController{BlogService: blogService}

	typeService := service.NewTypeServiceImpl()
	typeController = controller.TypeController{TypeService:typeService }

	tagService := service.NewTagServiceImpl()
	tagController = controller.TagController{TagService:tagService }

	htmlController = controller.HtmlController{}
	htmlController.BlogService = blogService

}
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.NoRoute(NoRouterHandler)
	r.NoMethod(NoMethodHandler)
	r.Use(ErrorHandler())
	r.LoadHTMLGlob("src/static/*.html") // html模板
	r.GET("/index", htmlController.IndexHtml)
	r.GET("/header", htmlController.HeaderHtml)
	r.GET("/footer", htmlController.FooterHtml)

	r.GET("/about", htmlController.AboutHtml)
	r.GET("/blog", htmlController.BlogHtml)
	r.GET("/search", htmlController.SearchHtml)
	r.GET("/tags", htmlController.TagsHtml)
	r.GET("/types", htmlController.TypesHtml)
	r.GET("/friend", htmlController.FriendHtml)
	r.GET("/archives", htmlController.ArchivesHtml)


	typeGroup :=r.Group("/type")
	typeGroup.POST("/add",typeController.AddType)
	typeGroup.PUT("/update",typeController.UpdateType)
	typeGroup.DELETE("/delete", typeController.DeleteType)
	typeGroup.GET("/get", typeController.GetType)
	typeGroup.POST("/find", typeController.FindTypes)


	tagGroup :=r.Group("/tag")
	tagGroup.POST("/add",tagController.AddTag)
	tagGroup.PUT("/update",tagController.UpdateTag)
	tagGroup.DELETE("/delete", tagController.DeleteTag)
	tagGroup.GET("/get", tagController.GetTag)
	tagGroup.POST("/find", tagController.FindTags)

	blogGroup := r.Group("/blog")
	blogGroup.POST("/add", blogController.AddBlog)
	blogGroup.POST("/get", blogController.GetBlog)
	blogGroup.POST("/find", blogController.FindBlogs)
	blogGroup.POST("/find_for_time", blogController.FindBlogsForTime)
	blogGroup.GET("/find_click_top_10", blogController.FindClickTop10Blogs)
	blogGroup.GET("/find_new_recommend", blogController.FindNewRecommendBlogs)
	r.Static("/static/", "./src/static/")
	return r
}

func NoRouterHandler(ctx *gin.Context) {
	e := code.NotFound
	ctx.JSON(e.StatusCode, e)
	return
}

func NoMethodHandler(ctx *gin.Context) {
	e := code.NoMethod
	ctx.JSON(e.StatusCode, e)
	return
}

// 统一异常处理
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var exception *code.Exception
				if e, ok := err.(*code.Exception); ok {
					exception = e
				} else if e, ok := err.(error); ok {
					exception = code.OtherException(e.Error())
				} else {
					exception = code.ServerException
				}
				ctx.JSON(exception.StatusCode, exception)
				return
			}
		}()
		ctx.Next()
	}
}
