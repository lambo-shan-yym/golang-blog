package service

import (
	"golang-blog/src/common/datasource"
	"golang-blog/src/dao"
	"golang-blog/src/model"
)

type IBlogTagService interface {
	Add(blog *model.BlogTag) error
	Update(blogId int64, blog *model.BlogTag) int64
	Delete(blogId int64) int64
	GetOneById(blogId int64, pushType int32) *model.BlogTag
}

type BlogTagServiceImpl struct {
	dao *dao.BlogTagDao
}

func (b *BlogTagServiceImpl) Add(blog *model.BlogTag) error {
	_, err := b.dao.AddBlogTag(blog)
	if err != nil {
		// todo
	}
	return err
}

func NewBlogTagServiceImpl() *BlogTagServiceImpl {
	return &BlogTagServiceImpl{
		dao: dao.NewBlogTagDao(datasource.InstanceDbMaster()),
	}
}

func (b *BlogTagServiceImpl) Update(blogId int64, blog *model.BlogTag) int64 {
	affected, err := b.dao.UpdateBlogTagById(blogId, blog)
	if err != nil {
		// todo
	}
	return affected
}

func (b *BlogTagServiceImpl) Delete(blogId int64) int64 {
	affected, err := b.dao.DeleteBlogTagById(blogId)
	if err != nil {
		// todo
	}
	return affected
}

func (b *BlogTagServiceImpl) GetOneById(blogId int64, pushType int32) *model.BlogTag {
	blog, err := b.dao.GetOneById(blogId, pushType)
	if err != nil {
		// todo
	}
	return blog
}
