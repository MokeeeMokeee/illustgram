package handler

import (
	"shalust/api/pkg/server/model"

	"github.com/gin-gonic/gin"
)

func GetIllustratio(c *gin.Context) {

	data, _ := model.GetAllIllustratio()

	c.JSON(200, data)
}
