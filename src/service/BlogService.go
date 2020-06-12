package service

import (
	"golang-blog/src/common/datasource"
	"golang-blog/src/dao"
	"golang-blog/src/model"
	"time"
)

type IBlogService interface {
	Add(blog *model.Blog) error
	Update(blogId int64, blog *model.Blog) int64
	Delete(blogId int64) int64
	GetOneById(blogId int64, pushType int32) *model.Blog
}

type BlogServiceImpl struct {
	dao *dao.BlogDao
}

func (b *BlogServiceImpl) Add(blog *model.Blog) error {
	blog.CreateTime = time.Now()
	blog.UpdateTime = time.Now()
	_, err := b.dao.AddBlog(blog)
	if err != nil {
		// todo
	}
	return err
}

func NewBlogServiceImpl() *BlogServiceImpl {
	return &BlogServiceImpl{
		dao: dao.NewBlogDao(datasource.InstanceDbMaster()),
	}
}

func (b *BlogServiceImpl) Update(blogId int64, blog *model.Blog) int64 {
	affected, err := b.dao.UpdateBlogById(blogId, blog)
	if err != nil {
		// todo
	}
	return affected
}

func (b *BlogServiceImpl) Delete(blogId int64) int64 {
	affected, err := b.dao.DeleteBlogById(blogId)
	if err != nil {
		// todo
	}
	return affected
}

func (b *BlogServiceImpl) GetOneById(blogId int64, pushType int32) *model.Blog {
	blog, err := b.dao.GetOneById(blogId, pushType)
	if err != nil {
		// todo
	}
	return blog
}
