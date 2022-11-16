package controller

import (
	"net/http"
	"project-ojt/model/dto"
	"project-ojt/usecase"

	"github.com/gin-gonic/gin"
)

type GetDataSiteAController struct {
	router    *gin.Engine
	ucGetData usecase.GetInternetDataUsecase
}

func (g *GetDataSiteAController) GetDataSiteA(ctx *gin.Context) {
	var input dto.DataInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	resultTraffIn, resUtilizationTraffIn, resultTraffOut, resUtilizationTraffOut, averageUp, _, sdate, edate, stime, etime, err := g.ucGetData.GetInternetData(input.IdSensor, input.SDate, input.EDate, input.STime, input.ETime)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "FAILED",
			"message": "data not found, date time input maybe wrong",
		})
		return
	}

	resultData := dto.DataOutput{
		TrafficIn:             resultTraffIn,
		UtilizationTrafficIn:  resUtilizationTraffIn,
		TrafficOut:            resultTraffOut,
		UtilizationTrafficOut: resUtilizationTraffOut,
		AverageUp:             averageUp,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":    "SUCCESS",
		"date time": sdate + " " + stime + " - " + edate + " " + etime,
		"result":    resultData,
	})

}

func NewGetDataSiteAController(router *gin.Engine, ucGetData usecase.GetInternetDataUsecase) *GetDataSiteAController {
	controller := GetDataSiteAController{
		router:    router,
		ucGetData: ucGetData,
	}

	router.POST("/get-data-site-a", controller.GetDataSiteA)

	return &controller
}
