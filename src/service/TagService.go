package service

import (
	"golang-blog/src/common/datasource"
	"golang-blog/src/dao"
	"golang-blog/src/model"
	"time"
)

type ITagService interface {
	Add(blog *model.Tag) error
	Update(blogId int64, blog *model.Tag) int64
	Delete(blogId int64) int64
	GetOneById(blogId int64, pushType int32) *model.Tag
}

type TagServiceImpl struct {
	dao *dao.TagDao
}

func (b *TagServiceImpl) Add(blog *model.Tag) error {
	blog.CreateTime = time.Now()
	blog.UpdateTime = time.Now()
	_, err := b.dao.AddTag(blog)
	if err != nil {
		// todo
	}
	return err
}

func NewTagServiceImpl() *TagServiceImpl {
	return &TagServiceImpl{
		dao: dao.NewTagDao(datasource.InstanceDbMaster()),
	}
}

func (b *TagServiceImpl) Update(blogId int64, blog *model.Tag) int64 {
	affected, err := b.dao.UpdateTagById(blogId, blog)
	if err != nil {
		// todo
	}
	return affected
}

func (b *TagServiceImpl) Delete(blogId int64) int64 {
	affected, err := b.dao.DeleteTagById(blogId)
	if err != nil {
		// todo
	}
	return affected
}

func (b *TagServiceImpl) GetOneById(blogId int64, pushType int32) *model.Tag {
	blog, err := b.dao.GetOneById(blogId, pushType)
	if err != nil {
		// todo
	}
	return blog
}
