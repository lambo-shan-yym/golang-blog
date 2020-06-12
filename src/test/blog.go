package main

import (
	"golang-blog/src/model"
	"golang-blog/src/service"
)

func main() {
	blogService:=service.NewBlogServiceImpl()
	blogService.Add(&model.Blog{Title:"test"})
}