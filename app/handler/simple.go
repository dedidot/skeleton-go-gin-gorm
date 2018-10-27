package handler

import(
	"../helper"
	"../model"
	"github.com/gin-gonic/gin"
)


func GetData(c *gin.Context) {
	var scheme []model.Scheme
	err := model.GetAll(&scheme)
	if err != nil {
		helper.RespondJSON(c, 404, scheme)
	} else {
		helper.RespondJSON(c, 200, scheme)
	}
}