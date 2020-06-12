package model

type BlogTag struct {
	Id     int64 `xorm:"pk autoincr BIGINT(20)"`
	BlogId int64 `xorm:"not null default 0 index(idx_blog_id) unique(uni_blog_type) BIGINT(20)"`
	TagId  int64 `xorm:"not null default 0 index(idx_tag_id) unique(uni_blog_type) BIGINT(20)"`
}
