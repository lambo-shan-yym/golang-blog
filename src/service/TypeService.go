package service

import (
	"golang-blog/src/common/datasource"
	"golang-blog/src/dao"
	"golang-blog/src/model"
	"time"
)

type ITypeService interface {
	Add(blog *model.Type) error
	Update(blogId int64, blog *model.Type) int64
	Delete(blogId int64) int64
	GetOneById(blogId int64, pushType int32) *model.Type
}

type TypeServiceImpl struct {
	dao *dao.TypeDao
}

func (b *TypeServiceImpl) Add(blog *model.Type) error {
	blog.CreateTime = time.Now()
	blog.UpdateTime = time.Now()
	_, err := b.dao.AddType(blog)
	if err != nil {
		// todo
	}
	return err
}

func NewTypeServiceImpl() *TypeServiceImpl {
	return &TypeServiceImpl{
		dao: dao.NewTypeDao(datasource.InstanceDbMaster()),
	}
}

func (b *TypeServiceImpl) Update(blogId int64, blog *model.Type) int64 {
	affected, err := b.dao.UpdateTypeById(blogId, blog)
	if err != nil {
		// todo
	}
	return affected
}

func (b *TypeServiceImpl) Delete(blogId int64) int64 {
	affected, err := b.dao.DeleteTypeById(blogId)
	if err != nil {
		// todo
	}
	return affected
}

func (b *TypeServiceImpl) GetOneById(blogId int64, pushType int32) *model.Type {
	blog, err := b.dao.GetOneById(blogId, pushType)
	if err != nil {
		// todo
	}
	return blog
}
