package controller

import (
	"github.com/gin-gonic/gin"
	"golang-blog/src/service"
)

type BlogController struct {
	BlogService service.IBlogService
}

func (c *BlogController) AddBlog(ctx *gin.Context) {

}
