package baseinfo

type CreateUserScoreRewardReq struct {
	OperationID string   `json:"operationID" binding:"required"`
	UserIDList  []string `json:"userIDList" binding:"required"`
}
type CreateUserScoreRewardResp struct {
}
