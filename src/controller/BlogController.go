package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-blog/src/common/constant"
	"golang-blog/src/dto"
	"golang-blog/src/service"
	"golang-blog/src/utils"
)

type BlogController struct {
	BlogService service.IBlogService
}

func (c *BlogController) AddBlog(ctx *gin.Context) {

}

func (c *BlogController) FindBlogs(ctx *gin.Context) {
	var blogDTO dto.FindBlogsReq
	ctx.BindJSON(&blogDTO)
	tools.CheckParams4Paginator(&blogDTO.Limit, blogDTO.Offset, "limit", "offset")

	countTags := c.BlogService.CountBlogs(constant.PUBLISH_TYPE_ON)
	blogs := c.BlogService.FindBlogs(blogDTO.Offset, blogDTO.Limit, constant.PUBLISH_TYPE_ON)
	result := &dto.FindBlogsResp{
		Count: countTags,
		List:  blogs,
	}
	tools.SuccessResponse(ctx, result)

}

func (c *BlogController) GetBlog(ctx *gin.Context) {
	var blogDTO dto.GetBlogReq
	ctx.BindJSON(&blogDTO)
	tools.CheckParamGreaterThanZero(blogDTO.Id, "id")

	// 校验id是否存在
	blogRecord := c.BlogService.CheckBlogIsExist(blogDTO.Id, constant.PUBLISH_TYPE_ON)

	tools.SuccessResponse(ctx, blogRecord)

}

func (c *BlogController) FindClickTop10Blogs(ctx *gin.Context) {

	blogs := c.BlogService.FindClickTop10Blogs(constant.PUBLISH_TYPE_ON)
	result := &dto.FindClickTop10BlogsResp{
		List: blogs,
	}
	tools.SuccessResponse(ctx, result)

}

func (c *BlogController) FindNewRecommendBlogs(ctx *gin.Context) {

	blogs := c.BlogService.FindNewRecommendBlogs(constant.PUBLISH_TYPE_ON)
	result := &dto.FindNewRecommendBlogsResp{
		List: blogs,
	}
	tools.SuccessResponse(ctx, result)

}

func (c *BlogController) FindBlogsForTime(ctx *gin.Context) {
	var blogDTO dto.FindBlogsReqForTime
	ctx.BindJSON(&blogDTO)
	tools.CheckParams4Paginator(&blogDTO.Limit, blogDTO.Offset, "limit", "offset")

	countTags := c.BlogService.CountBlogs(constant.PUBLISH_TYPE_ON)
	blogs := c.BlogService.FindBlogs(blogDTO.Offset, blogDTO.Limit, constant.PUBLISH_TYPE_ON)

	list := make([] *dto.FindBlogsRespForTimeItem, 0)
	for index := range blogs {
		blog := blogs[index]
		year := blog.CreateTime.Year()
		month := int(blog.CreateTime.Month())
		day := blog.CreateTime.Day()

		monthDay := fmt.Sprintf("%d月%d日", month, day)

		item := &dto.FindBlogsRespForTimeItem{
			Blog:           blogs[index],
			CreateYear:     year,
			CreateMonthDay: monthDay,
		}
		list = append(list, item)
	}
	result := &dto.FindBlogsRespForTime{
		Count: countTags,
		List:  list,
	}
	tools.SuccessResponse(ctx, result)

}
