package dao

import (
	"github.com/xormplus/xorm"
	"golang-blog/src/model"
)

type TagDao struct {
	engine *xorm.Engine
}

func NewTagDao(engine *xorm.Engine) *TagDao {
	return &TagDao{
		engine: engine,
	}
}

func (b *TagDao) AddTag(blog *model.Tag) (int64, error) {
	affected, err := b.engine.Insert(blog)
	return affected, err
}

func (b *TagDao) UpdateTagById(blogId int64, blog *model.Tag) (int64, error) {
	affected, err := b.engine.ID(blogId).Update(blog)
	return affected, err
}

func (b *TagDao) DeleteTagById(blogId int64) (int64, error) {
	affected, err := b.engine.ID(blogId).Delete(nil)
	return affected, err
}

func (b *TagDao) GetOneById(blogId int64, pushType int32) (*model.Tag, error) {
	blog := &model.Tag{}
	_, err := b.engine.ID(blogId).Get(blog)
	return blog, err
}
