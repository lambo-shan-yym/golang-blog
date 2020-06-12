package dao

import (
	"github.com/xormplus/xorm"
	"golang-blog/src/model"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{
		engine: engine,
	}
}

func (b *UserDao) AddUser(blog *model.User) (int64, error) {
	affected, err := b.engine.Insert(blog)
	return affected, err
}

func (b *UserDao) UpdateUserById(blogId int64, blog *model.User) (int64, error) {
	affected, err := b.engine.ID(blogId).Update(blog)
	return affected, err
}

func (b *UserDao) DeleteUserById(blogId int64) (int64, error) {
	affected, err := b.engine.ID(blogId).Delete(nil)
	return affected, err
}

func (b *UserDao) GetOneById(blogId int64, pushType int32) (*model.User, error) {
	blog := &model.User{}
	_, err := b.engine.ID(blogId).Get(blog)
	return blog, err
}
