package controller

import (
	"github.com/gin-gonic/gin"
	"golang-blog/src/dto"
	"golang-blog/src/model"
	"golang-blog/src/service"
	"golang-blog/src/utils"
	"time"
)

type TypeController struct {
	TypeService service.ITypeService
}

func (c *TypeController) AddType(ctx *gin.Context) {
	var typeDTO dto.AddTypeReq
	ctx.BindJSON(&typeDTO)
	//todo 参数校验
	tools.CheckParamNotBlank(typeDTO.Name, "name")

	newType := &model.Type{
		Name:       typeDTO.Name,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	c.TypeService.Add(newType)
	tools.SuccessResponse(ctx, newType)
}

func (c *TypeController) UpdateType(ctx *gin.Context) {
	var typeDTO dto.UpdateTypeReq
	ctx.BindJSON(&typeDTO)
	tools.CheckParamGreaterThanZero(typeDTO.Id, "id")
	tools.CheckParamNotBlank(typeDTO.Name, "name")

	// 校验id是否存在
	c.TypeService.CheckTypeIsExist(typeDTO.Id)
	newType := &model.Type{
		Name:       typeDTO.Name,
		UpdateTime: time.Now(),
	}
	c.TypeService.Update(typeDTO.Id, newType)
	newType.Id = typeDTO.Id
	tools.SuccessResponse(ctx, newType)
}

func (c *TypeController) DeleteType(ctx *gin.Context) {
	var typeDTO dto.DeleteTypeReq
	ctx.BindJSON(&typeDTO)
	tools.CheckParamGreaterThanZero(typeDTO.Id, "id")

	// 校验id是否存在
	typeRecord := c.TypeService.CheckTypeIsExist(typeDTO.Id)
	// 删除
	c.TypeService.Delete(typeDTO.Id)
	tools.SuccessResponse(ctx, typeRecord)

}

func (c *TypeController) GetType(ctx *gin.Context) {
	var typeDTO dto.DeleteTypeReq
	ctx.BindJSON(&typeDTO)
	tools.CheckParamGreaterThanZero(typeDTO.Id, "id")

	// 校验id是否存在
	typeRecord := c.TypeService.CheckTypeIsExist(typeDTO.Id)
	tools.SuccessResponse(ctx, typeRecord)

}

func (c *TypeController) FindTypes(ctx *gin.Context) {
	var typeDTO dto.FindTypesReq
	ctx.BindJSON(&typeDTO)
	tools.CheckParams4Paginator(&typeDTO.Limit,typeDTO.Offset,"limit","offset")

	countTypes := c.TypeService.CountTypes()
	types := c.TypeService.FindTypes(typeDTO.Offset, typeDTO.Limit)
	result := &dto.FindTypesResp{
		Count: countTypes,
		List:  types,
	}
	tools.SuccessResponse(ctx, result)

}
