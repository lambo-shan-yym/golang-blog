package service

import (
	"golang-blog/src/common/datasource"
	"golang-blog/src/common/error"
	"golang-blog/src/dao"
	"golang-blog/src/model"
	"time"
)

type IBlogService interface {
	Add(blog *model.Blog) error
	Update(blogId int64, blog *model.Blog) int64
	Delete(blogId int64) int64
	GetOneById(blogId int64, publishType int32) *model.Blog

	FindBlogs(offset int, limit int, publishType int32) []*model.Blog

	CountBlogs(publishType int32) int64

	FindClickTop10Blogs(publishType int32) []*model.Blog

	FindNewRecommendBlogs(publishType int32) []*model.Blog

	CheckBlogIsExist(id int64, publishType int32) *model.Blog
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



func (b *BlogServiceImpl) FindBlogs(offset int, limit int, publishType int32) []*model.Blog {
	blogs, err := b.dao.FindBlogs(offset, limit, publishType)
	if err != nil {
		// todo
	}
	if blogs == nil {
		return make([]*model.Blog, 0)
	}
	return blogs
}

func (b *BlogServiceImpl) CountBlogs(publishType int32) int64 {
	count, err := b.dao.Count(publishType)
	if err != nil {
		// todo
	}
	return count
}


func (b *BlogServiceImpl)FindClickTop10Blogs(publishType int32) []*model.Blog{
	blogs, err := b.dao.FindClickTop10Blogs(publishType)
	if err != nil {
		// todo
	}
	if blogs == nil {
		return make([]*model.Blog, 0)
	}
	return blogs
}



func (b *BlogServiceImpl)FindNewRecommendBlogs(publishType int32) []*model.Blog{
	blogs, err := b.dao.FindNewRecommendBlogs(publishType)
	if err != nil {
		// todo
	}
	if blogs == nil {
		return make([]*model.Blog, 0)
	}
	return blogs
}

func (b *BlogServiceImpl) CheckBlogIsExist(id int64, publishType int32) *model.Blog  {
	blogRecord := b.GetOneById(id, publishType)
	if blogRecord == nil {
		panic(code.DataNotExistInDb)
	}
	return blogRecord
}


func (b *BlogServiceImpl) GetOneById(blogId int64, publishType int32) *model.Blog {
	blog, err := b.dao.GetOneById(blogId,publishType)
	if err != nil {
		// todo
	}
	return blog
}

