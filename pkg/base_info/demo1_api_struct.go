package baseinfo

type ApiCreateFriendReq struct {
	ApiCommonReq
}
type ApiCreateFriendResp struct {
	ApiCommonResp
}

type ApiGetFriendFromDBReq struct {
	ApiCommonReq
}
type ApiGetFriendFromDBResp struct {
	ApiCommonResp
}

type ApiGetFriendFromRedisReq struct {
	ApiCommonReq
}
type ApiGetFriendFromRedisResp struct {
	ApiCommonResp
}

type ApiGetFriendFromLocalCacheReq struct {
	ApiCommonReq
}
type ApiGetFriendFromLocalCacheResp struct {
	ApiCommonResp
}
