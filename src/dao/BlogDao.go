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

func (b *BlogDao) GetOneById(blogId int64, publishType int32) (*model.Blog, error) {
	blog := &model.Blog{}
	_, err := b.engine.ID(blogId).Where("publish_type = ?",publishType).Get(blog)
	return blog, err
}

func (b *BlogDao) FindBlogs(offset, limit int, publishType int32) ([]*model.Blog, error) {
	var blogs [] *model.Blog
	error := b.engine.Where("publish_type = ?",publishType).OrderBy(" create_time desc").Limit(limit, offset).Find(&blogs)
	return blogs, error
}

func (b *BlogDao) Count(publishType int32) (int64, error) {
	condition := model.Blog{}
	condition.PublishType = publishType
	count, error := b.engine.Count(&condition)
	return count, error
}

func (b *BlogDao) FindClickTop10Blogs( publishType int32) ([]*model.Blog, error) {
	var blogs [] *model.Blog
	error := b.engine.Where("publish_type = ?",publishType).OrderBy("view_count desc").Limit(10, 0).Find(&blogs)
	return blogs, error
}


func (b *BlogDao) FindNewRecommendBlogs( publishType int32) ([]*model.Blog, error) {
	var blogs [] *model.Blog
	error := b.engine.Where("publish_type = ?",publishType).OrderBy("update_time desc").Limit(10, 0).Find(&blogs)
	return blogs, error
}