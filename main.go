package main

import (
	"fclink.cn/ethcoldwallet/conf"
	"fclink.cn/ethcoldwallet/models"
	"fclink.cn/ethcoldwallet/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func init() {
	_ = conf.ParseConf("config.json")
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})
	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.WarnLevel)

	models.Setup()
}

func main() {
	router := routers.InitRouter()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	router.Run(fmt.Sprintf(":%d", conf.Context.Server.Port))
}

func line() {
	fmt.Println("ok")
	client := &http.Client{
		Timeout: conf.Context.Node.Interval * time.Second,
	}
	res, err := client.Get("https://openapi.wbf.live/open/api/get_ticker?symbol=gcsusdt")
	if err != nil {
		log.Errorf("get maxHeight err: %s", err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
