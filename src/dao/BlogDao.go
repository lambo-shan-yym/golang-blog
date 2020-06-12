package dao

import (
	"github.com/xormplus/xorm"
	"golang-blog/src/model"
)

type BlogDao struct {
	engine *xorm.Engine
}

func NewBlogDao(engine *xorm.Engine) *BlogDao {
	return &BlogDao{
		engine: engine,
	}
}

func (b *BlogDao) AddBlog(blog *model.Blog) (int64, error) {
	affected, err := b.engine.Insert(blog)
	return affected, err
}

func (b *BlogDao) UpdateBlogById(blogId int64, blog *model.Blog) (int64, error) {
	affected, err := b.engine.ID(blogId).Update(blog)
	return affected, err
}

func (b *BlogDao) DeleteBlogById(blogId int64) (int64, error) {
	affected, err := b.engine.ID(blogId).Delete(nil)
	return affected, err
}

func (b *BlogDao) GetOneById(blogId int64, pushType int32) (*model.Blog, error) {
	blog := &model.Blog{}
	_, err := b.engine.ID(blogId).Get(blog)
	return blog, err
}
