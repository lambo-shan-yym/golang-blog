package router

import (
	"github.com/gin-gonic/gin"
	"golang-blog/src/common/error"
	"golang-blog/src/controller"
	"golang-blog/src/service"
)

var blogController controller.BlogController
var htmlController controller.HtmlController

func init() {
	blogService := service.NewBlogServiceImpl()
	blogController = controller.BlogController{BlogService: blogService}
	htmlController = controller.HtmlController{}
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
	blogGroup := r.Group("/blog")
	blogGroup.POST("/add", blogController.AddBlog)
	r.Static("/static", "./src/static")
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
