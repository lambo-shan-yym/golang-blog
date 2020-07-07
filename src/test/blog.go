package main

import (
	"golang-blog/src/common/datasource"
	"golang-blog/src/model"
)

func main() {
	dbMaster := datasource.InstanceDbMaster()
	dbMaster.Sync2(new(model.Type),new(model.Tag),new(model.BlogTag),new(model.User))
	//blogService:=service.NewBlogServiceImpl()
	//blogService.Add(&model.Blog{Title:"test"})
}