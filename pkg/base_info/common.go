package baseinfo

type RequestPagination struct {
	PageNumber int `json:"pageNumber"  binding:"required"`
	ShowNumber int `json:"showNumber"  binding:"required"`
}

type ResponsePagination struct {
	CurrentPage int `json:"currentPage"`
	ShowNumber  int `json:"showNumber"`
}

type Swagger400Resp struct {
	ErrCode int32  `json:"errCode" example:"400"`
	ErrMsg  string `json:"errMsg" example:"err msg"`
}

type Swagger500Resp struct {
	ErrCode int32  `json:"errCode" example:"500"`
	ErrMsg  string `json:"errMsg" example:"err msg"`
}
