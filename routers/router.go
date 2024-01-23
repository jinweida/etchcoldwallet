package routers

import (
	"fclink.cn/ethcoldwallet/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// reg router
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	eth := &controller.EthController{}
	// ?address=0x00000000
	api.GET("/eth/create", eth.Create)

	/**
	{
		"nonce":"1",
		"gasprice":""
		"gaslimit":"",
		"to":"",
		"value":"",
		"chainid":"",
		"data":"data"
		"tx":""
	}
	*/
	api.GET("/eth/send", eth.SendTransaction)
	api.GET("/eth/query", eth.QueryTransaction)
	return r
}
