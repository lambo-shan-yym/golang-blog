package dto

import "golang-blog/src/model"

type AddTypeReq struct {
	Name string `json:"name" `
}

type UpdateTypeReq struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

type DeleteTypeReq struct {
	Id int64 `json:"id"`
}

type FindTypesReq struct {
	Offset int
	Limit int
}

type FindTypesResp struct {
	Count int64  `json:"count"`
	List []*model.Type `json:"list"`
}