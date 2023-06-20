package demo1

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	api "github.com/knightxv/go-project-drill/pkg/base_info"
	"github.com/knightxv/go-project-drill/pkg/common/config"
	"github.com/knightxv/go-project-drill/pkg/common/log"
	"github.com/knightxv/go-project-drill/pkg/grpc-etcdv3/getcdv3"
	pbDemo1 "github.com/knightxv/go-project-drill/pkg/proto/demo1"
	"github.com/knightxv/go-project-drill/pkg/utils"
)

// @Summary
// @Description	获取好友
// @Tags			好友
// @ID				CreateOneWorkMoment
// @Accept			json
// @Param			token	header	string						true	"im token"
// @Param			req		body	api.GetFriendFromDB	true	"请求"
// @Produce		json
// @Success		0	{object}	api.CreateOneWorkMomentResp
// @Failure		500	{object}	api.Swagger500Resp	"errCode为500 一般为服务器内部错误"
// @Failure		400	{object}	api.Swagger400Resp	"errCode为400 一般为参数输入错误, token未带上等"
// @Router			/demo1/get_friend_from_db [get]
func GetFriendFromDB(c *gin.Context) {
	var (
		req    api.ApiGetFriendFromDBReq
		resp   api.ApiGetFriendFromDBResp
		reqPb  pbDemo1.GetFriendFromDBReq
		respPb *pbDemo1.GetFriendFromDBResp
	)
	if err := c.BindJSON(&req); err != nil {
		log.NewError(req.OperationID, utils.GetSelfFuncName(), "bind json failed", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "bind json failed " + err.Error()})
		return
	}
	log.NewInfo(req.OperationID, utils.GetSelfFuncName(), "req: ", req)

	if err := utils.CopyStructFields(&reqPb, req); err != nil {
		log.NewDebug(req.OperationID, utils.GetSelfFuncName(), "CopyStructFields failed", err.Error())
	}
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImRelayName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

	client := pbDemo1.NewChainServiceClient(etcdConn)
	respPb, err := client.GetFriendFromDB(context.Background(), &reqPb)
	if err != nil {
		log.NewError(req.OperationID, utils.GetSelfFuncName(), "CreateOneWorkMoment rpc failed", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "CreateOneWorkMoment rpc server failed" + err.Error()})
		return
	}
	resp.ApiCommonResp = api.ApiCommonResp{
		ErrCode: respPb.CommonResp.ErrCode,
		ErrMsg:  respPb.CommonResp.ErrMsg,
	}
	log.NewInfo(req.OperationID, utils.GetSelfFuncName(), "resp: ", resp)
	c.JSON(http.StatusOK, resp)
}
