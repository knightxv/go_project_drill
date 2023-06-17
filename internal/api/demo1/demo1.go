package demo1

// @Summary		创建一条工作圈
// @Description	用户创建一条工作圈
// @Tags			工作圈
// @ID				CreateOneWorkMoment
// @Accept			json
// @Param			token	header	string						true	"im token"
// @Param			req		body	api.CreateOneWorkMomentReq	true	"请求 atUserList likeUserList permissionGroupList permissionUserList 字段中userName可以不填"
// @Produce		json
// @Success		0	{object}	api.CreateOneWorkMomentResp
// @Failure		500	{object}	api.Swagger500Resp	"errCode为500 一般为服务器内部错误"
// @Failure		400	{object}	api.Swagger400Resp	"errCode为400 一般为参数输入错误, token未带上等"
// @Router			/office/create_one_work_moment [post]
func CreateOneWorkMoment(c *gin.Context) {
	var (
		req    api.CreateOneWorkMomentReq
		resp   api.CreateOneWorkMomentResp
		reqPb  pbOffice.CreateOneWorkMomentReq
		respPb *pbOffice.CreateOneWorkMomentResp
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
	reqPb.WorkMoment.UserID = userID
	etcdConn := getcdv3.GetDefaultConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImOfficeName, req.OperationID)
	if etcdConn == nil {
		errMsg := req.OperationID + "getcdv3.GetDefaultConn == nil"
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

	client := pbOffice.NewOfficeServiceClient(etcdConn)
	respPb, err := client.CreateOneWorkMoment(context.Background(), &reqPb)
	if err != nil {
		log.NewError(req.OperationID, utils.GetSelfFuncName(), "CreateOneWorkMoment rpc failed", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"errCode": 500, "errMsg": "CreateOneWorkMoment rpc server failed" + err.Error()})
		return
	}
	resp.CommResp = api.CommResp{
		ErrCode: respPb.CommonResp.ErrCode,
		ErrMsg:  respPb.CommonResp.ErrMsg,
	}
	log.NewInfo(req.OperationID, utils.GetSelfFuncName(), "resp: ", resp)
	c.JSON(http.StatusOK,  gin.H{"errCode": 0, "errMsg": "success"})
}

