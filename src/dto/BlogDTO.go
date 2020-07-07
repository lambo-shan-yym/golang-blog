package dto

import "golang-blog/src/model"


type GetBlogReq struct {
	Id int64 `json:"id"`
}

type FindBlogsReq struct {
	Offset int
	Limit int
}

type FindBlogsResp struct {
	Count int64  `json:"count"`
	List []*model.Blog `json:"list"`
}


type FindClickTop10BlogsResp struct {
	List []*model.Blog `json:"list"`
}


type FindNewRecommendBlogsResp struct {
	List []*model.Blog `json:"list"`
}


type FindBlogsReqForTime struct {
	Offset int
	Limit int
}

type FindBlogsRespForTime struct {
	Count int64  `json:"count"`
	List []*FindBlogsRespForTimeItem `json:"list"`
}

type FindBlogsRespForTimeItem struct {
	Blog *model.Blog `json:"blog"`
	CreateYear int	`json:"create_year"`
	CreateMonthDay string `json:"create_month_day"`
}