package controller

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type HtmlController struct {
}

func (c *HtmlController) IndexHtml(ctx *gin.Context) {
	//模板文件的拼接
	template.ParseFiles("static/index.html", "static/header.html","static/footer.html")
	ctx.HTML(http.StatusOK, "index.html", gin.H{
	})
}

func (c *HtmlController) HeaderHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "header.html", gin.H{
	})
}

func (c *HtmlController) FooterHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "footer.html", gin.H{
	})
}
