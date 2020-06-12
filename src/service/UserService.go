package service

import (
	"golang-blog/src/common/datasource"
	"golang-blog/src/dao"
	"golang-blog/src/model"
	"time"
)

type IUserService interface {
	Add(blog *model.User) error
	Update(blogId int64, blog *model.User) int64
	Delete(blogId int64) int64
	GetOneById(blogId int64, pushType int32) *model.User
}

type UserServiceImpl struct {
	dao *dao.UserDao
}

func (b *UserServiceImpl) Add(blog *model.User) error {
	blog.CreateTime = time.Now()
	blog.UpdateTime = time.Now()
	_, err := b.dao.AddUser(blog)
	if err != nil {
		// todo
	}
	return err
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		dao: dao.NewUserDao(datasource.InstanceDbMaster()),
	}
}

func (b *UserServiceImpl) Update(blogId int64, blog *model.User) int64 {
	affected, err := b.dao.UpdateUserById(blogId, blog)
	if err != nil {
		// todo
	}
	return affected
}

func (b *UserServiceImpl) Delete(blogId int64) int64 {
	affected, err := b.dao.DeleteUserById(blogId)
	if err != nil {
		// todo
	}
	return affected
}

func (b *UserServiceImpl) GetOneById(blogId int64, pushType int32) *model.User {
	blog, err := b.dao.GetOneById(blogId, pushType)
	if err != nil {
		// todo
	}
	return blog
}
