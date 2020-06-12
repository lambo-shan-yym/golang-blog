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

func (c *HtmlController) BlogHtml(ctx *gin.Context) {
	//模板文件的拼接
	template.ParseFiles("static/blog.html", "static/header.html","static/footer.html")
	ctx.HTML(http.StatusOK, "blog.html", gin.H{
	})
}

func (c *HtmlController) SearchHtml(ctx *gin.Context) {
	//模板文件的拼接
	template.ParseFiles("static/search.html", "static/header.html","static/footer.html")
	ctx.HTML(http.StatusOK, "search.html", gin.H{
	})
}
func (c *HtmlController) TagsHtml(ctx *gin.Context) {
	//模板文件的拼接
	template.ParseFiles("static/tags.html", "static/header.html","static/footer.html")
	ctx.HTML(http.StatusOK, "tags.html", gin.H{
	})
}

func (c *HtmlController) TypesHtml(ctx *gin.Context) {
	//模板文件的拼接
	template.ParseFiles("static/types.html", "static/header.html","static/footer.html")
	ctx.HTML(http.StatusOK, "types.html", gin.H{
	})
}

func (c *HtmlController) FriendHtml(ctx *gin.Context) {
	//模板文件的拼接
	template.ParseFiles("static/friend.html", "static/header.html","static/footer.html")
	ctx.HTML(http.StatusOK, "friend.html", gin.H{
	})
}

func (c *HtmlController) AboutHtml(ctx *gin.Context) {
	//模板文件的拼接
	template.ParseFiles("static/about.html", "static/header.html","static/footer.html")
	ctx.HTML(http.StatusOK, "about.html", gin.H{
	})
}
func (c *HtmlController) ArchivesHtml(ctx *gin.Context) {
	//模板文件的拼接
	template.ParseFiles("static/archives.html", "static/header.html","static/footer.html")
	ctx.HTML(http.StatusOK, "archives.html", gin.H{
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
