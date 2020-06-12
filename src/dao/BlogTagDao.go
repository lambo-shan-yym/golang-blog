package dao

import (
	"github.com/xormplus/xorm"
	"golang-blog/src/model"
)

type BlogTagDao struct {
	engine *xorm.Engine
}

func NewBlogTagDao(engine *xorm.Engine) *BlogTagDao {
	return &BlogTagDao{
		engine: engine,
	}
}

func (b *BlogTagDao) AddBlogTag(blog *model.BlogTag) (int64, error) {
	affected, err := b.engine.Insert(blog)
	return affected, err
}

func (b *BlogTagDao) UpdateBlogTagById(blogId int64, blog *model.BlogTag) (int64, error) {
	affected, err := b.engine.ID(blogId).Update(blog)
	return affected, err
}

func (b *BlogTagDao) DeleteBlogTagById(blogId int64) (int64, error) {
	affected, err := b.engine.ID(blogId).Delete(nil)
	return affected, err
}

func (b *BlogTagDao) GetOneById(blogId int64, pushType int32) (*model.BlogTag, error) {
	blog := &model.BlogTag{}
	_, err := b.engine.ID(blogId).Get(blog)
	return blog, err
}
