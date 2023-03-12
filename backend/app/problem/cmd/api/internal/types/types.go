// Code generated by goctl. DO NOT EDIT.
package types

type Problem struct {
	ProblemId   string  `json:"problemId"`
	Title       string  `json:"title"`
	PassingRate float64 `json:"passingRate"`
	Difficulty  int64   `json:"difficulty"`
	IsAccepted  bool    `json:"isAccepted"`
}

type GetProblemListInfoReq struct {
	CurrentPage int `form:"currentPage" validate:"required,gte=1"`
}

type GetProblemListInfoResp struct {
	Current  int       `json:"current"`
	Total    int64     `json:"total"`
	Problems []Problem `json:"problems"`
}

type GetProblemInfoReq struct {
	ProblemId string `path:"problemId" validate:"required"`
}

type GetProblemInfoResp struct {
	ProblemInfo Problem `json:"problemInfo"`
}
