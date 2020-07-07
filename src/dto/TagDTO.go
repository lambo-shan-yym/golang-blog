package dto

import "golang-blog/src/model"

type AddTagReq struct {
	Name string `json:"name" `
	Color string `json:"color"`
}

type UpdateTagReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Color string `json:"color"`
}

type DeleteTagReq struct {
	Id int64 `json:"id"`
}

type FindTagsReq struct {
	Offset int
	Limit int
}

type FindTagsResp struct {
	Count int64  `json:"count"`
	List []*model.Tag `json:"list"`
}



