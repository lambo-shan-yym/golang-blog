package dao

import (
	"github.com/xormplus/xorm"
	"golang-blog/src/model"
)

type TypeDao struct {
	engine *xorm.Engine
}

func NewTypeDao(engine *xorm.Engine) *TypeDao {
	return &TypeDao{
		engine: engine,
	}
}

func (b *TypeDao) AddType(blog *model.Type) (int64, error) {
	affected, err := b.engine.Insert(blog)
	return affected, err
}

func (b *TypeDao) UpdateTypeById(blogId int64, blog *model.Type) (int64, error) {
	affected, err := b.engine.ID(blogId).Update(blog)
	return affected, err
}

func (b *TypeDao) DeleteTypeById(blogId int64) (int64, error) {
	affected, err := b.engine.ID(blogId).Delete(nil)
	return affected, err
}

func (b *TypeDao) GetOneById(blogId int64, pushType int32) (*model.Type, error) {
	blog := &model.Type{}
	_, err := b.engine.ID(blogId).Get(blog)
	return blog, err
}
