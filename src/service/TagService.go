package service

import (
	"golang-blog/src/common/datasource"
	"golang-blog/src/common/error"
	"golang-blog/src/dao"
	"golang-blog/src/model"
	"time"
)

type ITagService interface {
	Add(blog *model.Tag) error
	Update(blogId int64, blog *model.Tag) int64
	Delete(blogId int64) int64
	GetOneById(blogId int64) *model.Tag
	CheckTagIsExist(tagId int64) *model.Tag
	FindTags(offset int, limit int) []*model.Tag
	CountTags() int64
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

func (b *TagServiceImpl) GetOneById(blogId int64) *model.Tag {
	blog, err := b.dao.GetOneById(blogId)
	if err != nil {
		// todo
	}
	return blog
}

func (b *TagServiceImpl) CheckTagIsExist(tagId int64) *model.Tag {
	tagRecord := b.GetOneById(tagId)
	if tagRecord == nil {
		panic(code.DataNotExistInDb)
	}
	return tagRecord
}

func (b *TagServiceImpl) FindTags(offset int, limit int) []*model.Tag {
	tags, err := b.dao.FindTypes(offset, limit)
	if err != nil {
		// todo
	}
	if tags == nil {
		return make([]*model.Tag, 0)
	}
	return tags
}

func (b *TagServiceImpl) CountTags() int64 {
	count, err := b.dao.Count()
	if err != nil {
		// todo
	}
	return count
}
