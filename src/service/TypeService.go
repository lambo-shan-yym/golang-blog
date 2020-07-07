package service

import (
	"golang-blog/src/common/datasource"
	"golang-blog/src/common/error"
	"golang-blog/src/dao"
	"golang-blog/src/model"
)

type ITypeService interface {
	Add(blog *model.Type) error
	Update(blogId int64, blog *model.Type) int64
	Delete(blogId int64) int64
	GetOneById(blogId int64) *model.Type
	CheckTypeIsExist(typeId int64) *model.Type
	FindTypes(offset int, limit int) []*model.Type
	CountTypes() int64
}

type TypeServiceImpl struct {
	dao *dao.TypeDao
}

func (b *TypeServiceImpl) Add(blog *model.Type) error {
	_, err := b.dao.AddType(blog)
	if err != nil {
		// todo

		panic(code.DataHadExistInDb)
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

func (b *TypeServiceImpl) GetOneById(blogId int64) *model.Type {
	blog, err := b.dao.GetOneById(blogId)
	if err != nil {
		// todo
	}
	if blog.Id == 0 {
		return nil
	}
	return blog
}

func (b *TypeServiceImpl) CheckTypeIsExist(typeId int64) *model.Type {
	typeRecord := b.GetOneById(typeId)
	if typeRecord == nil {
		panic(code.DataNotExistInDb)
	}
	return typeRecord
}

func (b *TypeServiceImpl) FindTypes(offset int, limit int) []*model.Type {
	types, err := b.dao.FindTypes(offset, limit)
	if err != nil {
		// todo
	}
	if types == nil {
		return make([]*model.Type, 0)
	}
	return types
}

func (b *TypeServiceImpl) CountTypes() int64 {
	count, err := b.dao.Count()
	if err != nil {
		// todo
	}
	return count
}
