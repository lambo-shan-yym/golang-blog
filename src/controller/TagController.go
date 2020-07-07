package controller

import (
	"github.com/gin-gonic/gin"
	"golang-blog/src/dto"
	"golang-blog/src/model"
	"golang-blog/src/service"
	"golang-blog/src/utils"
	"time"
)

type TagController struct {
	TagService service.ITagService
}

func (c *TagController) AddTag(ctx *gin.Context) {
	var tagDTO dto.AddTagReq
	ctx.BindJSON(&tagDTO)
	//todo 参数校验
	tools.CheckParamNotBlank(tagDTO.Name, "name")
	tools.CheckParamNotBlank(tagDTO.Color, "color")
	newTag := &model.Tag{
		Name:       tagDTO.Name,
		Color:      tagDTO.Color,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	c.TagService.Add(newTag)
	tools.SuccessResponse(ctx, newTag)
}

func (c *TagController) UpdateTag(ctx *gin.Context) {
	var tagDTO dto.UpdateTagReq
	ctx.BindJSON(&tagDTO)
	tools.CheckParamGreaterThanZero(tagDTO.Id, "id")
	tools.CheckParamNotBlank(tagDTO.Name, "name")

	// 校验id是否存在
	c.TagService.CheckTagIsExist(tagDTO.Id)
	newTag := &model.Tag{
		Name:       tagDTO.Name,
		UpdateTime: time.Now(),
	}
	c.TagService.Update(tagDTO.Id, newTag)
	newTag.Id = tagDTO.Id
	tools.SuccessResponse(ctx, newTag)
}

func (c *TagController) DeleteTag(ctx *gin.Context) {
	var tagDTO dto.DeleteTagReq
	ctx.BindJSON(&tagDTO)
	tools.CheckParamGreaterThanZero(tagDTO.Id, "id")

	// 校验id是否存在
	tagRecord := c.TagService.CheckTagIsExist(tagDTO.Id)
	// 删除
	c.TagService.Delete(tagDTO.Id)
	tools.SuccessResponse(ctx, tagRecord)

}

func (c *TagController) GetTag(ctx *gin.Context) {
	var tagDTO dto.DeleteTagReq
	ctx.BindJSON(&tagDTO)
	tools.CheckParamGreaterThanZero(tagDTO.Id, "id")

	// 校验id是否存在
	tagRecord := c.TagService.CheckTagIsExist(tagDTO.Id)
	tools.SuccessResponse(ctx, tagRecord)

}

func (c *TagController) FindTags(ctx *gin.Context) {
	var tagDTO dto.FindTagsReq
	ctx.BindJSON(&tagDTO)
	tools.CheckParams4Paginator(&tagDTO.Limit, tagDTO.Offset, "limit", "offset")

	countTags := c.TagService.CountTags()
	tags := c.TagService.FindTags(tagDTO.Offset, tagDTO.Limit)
	result := &dto.FindTagsResp{
		Count: countTags,
		List:  tags,
	}
	tools.SuccessResponse(ctx, result)

}
