package main

import (
	"flag"
	"io"
	"os"
	"strconv"

	_ "github.com/knightxv/go-project-drill/api/docs"
	"github.com/knightxv/go-project-drill/pkg/common/config"
	"github.com/knightxv/go-project-drill/pkg/common/log"
	"github.com/knightxv/go-project-drill/pkg/utils"

	"github.com/knightxv/go-project-drill/internal/api/demo1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	//"syscall"
	"github.com/knightxv/go-project-drill/pkg/common/constant"
)

// @title			open-IM-Server API
// @version		1.0
// @description	open-IM-Server 的API服务器文档, 文档中所有请求都有一个operationID字段用于链路追踪
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath		/
func main() {
	log.NewPrivateLog(constant.LogFileName)
	gin.SetMode(gin.ReleaseMode)
	f, _ := os.Create("../logs/api.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(utils.CorsHandler())
	log.Info("load config: ", config.Config)
	if !config.Config.IsPublicEnv {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	demo1GroupRouter := r.Group("demo1")
	{
		demo1GroupRouter.GET("/demo1/get_friend_from_db", demo1.GetFriendFromDB)
		demo1GroupRouter.GET("/demo1/get_friend_from_redis", demo1.GetFriendFromRedis)
		demo1GroupRouter.GET("/demo1/get_friend_from_local_cache", demo1.GetFriendFromLocalCache)
	}

	defaultPorts := config.Config.Api.GinPort
	ginPort := flag.Int("port", defaultPorts[0], "get ginServerPort from cmd,default 10002 as port")
	flag.Parse()
	address := "0.0.0.0:" + strconv.Itoa(*ginPort)
	if config.Config.Api.ListenIP != "" {
		address = config.Config.Api.ListenIP + ":" + strconv.Itoa(*ginPort)
	}
	log.NewInfo("", "start api server, address: ", address, "OpenIM version: ", constant.CurrentVersion)
	err := r.Run(address)
	if err != nil {
		log.Error("", "api run failed ", address, err.Error())
		panic("api start failed " + err.Error())
	}
}
