package controller

import (
	"fclink.cn/ethcoldwallet/entity"
	"fclink.cn/ethcoldwallet/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EthController struct {
}

func (this *EthController) SendTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Res.Ok().Result(nil))
}

func (this *EthController) QueryTransaction(c *gin.Context) {
	address := c.DefaultQuery("address", "")
	contractAddr := c.DefaultQuery("contractaddr", "")
	op := c.DefaultQuery("op", "0")

	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))

	if err != nil {
		c.JSON(http.StatusOK, entity.Res.Error())
	} else {
		if size > 50 {
			size = 10
		}
		if len(address) > 0 {
			list := models.FindByAddress(op, address, contractAddr, page, size)
			for i, m := range list {
				if len(m.ContractAddr) > 0 {
					list[i].TxValue = m.TokenCount
				}
			}
			c.JSON(http.StatusOK, entity.Res.Ok().Result(list))
		} else {
			c.JSON(http.StatusOK, entity.Res.Error().SetMsg("address invaild"))
		}
	}
}

func (this *EthController) Create(c *gin.Context) {
	address := c.DefaultQuery("address", "")
	if len(address) > 0 {
		if err := models.CreateFund(address); err != nil {
			c.JSON(http.StatusOK, entity.Res.Ok())
		} else {
			c.JSON(http.StatusOK, entity.Res.Error().SetMsg(fmt.Sprintf("address %s exist", address)))
		}
	} else {
		c.JSON(http.StatusOK, entity.Res.Error().SetMsg("address invaild"))
	}
}
