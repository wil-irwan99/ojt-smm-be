package controller

import (
	"net/http"
	"project-ojt/config"
	"project-ojt/model/dto"
	"project-ojt/usecase"

	"github.com/gin-gonic/gin"
)

type GetDataSiteController struct {
	router    *gin.Engine
	ucGetData usecase.GetInternetDataSiteUsecase
	config    config.Config
}

func (g *GetDataSiteController) GetDataSite(ctx *gin.Context) {

	site := ctx.Query("site")
	sdate := ctx.Query("sdate")
	edate := ctx.Query("edate")
	stime := ctx.Query("stime")
	etime := ctx.Query("etime")

	// var input dto.DataInput
	// if err := ctx.BindJSON(&input); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "can't bind struct",
	// 	})
	// 	return
	// }

	var ipSiteConfig [][]string
	var resultInternetArr []dto.DataOutput

	switch site {
	case "BIB":
		ipSiteConfig = append(ipSiteConfig, []string{g.config.PRTGConfig.Ip, g.config.PRTGConfig.User, g.config.PRTGConfig.Password})
	case "Berau":

	default:

	}

	for i := 0; i < len(ipSiteConfig); i++ {
		resultInternet, err := g.ucGetData.GetInternetDataSite(site, "internet", ipSiteConfig[i][0], ipSiteConfig[i][1], ipSiteConfig[i][2], sdate, edate, stime, etime)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "FAILED",
				"message": "data not found, date time input maybe wrong",
			})
			return
		}
		resultInternetArr = append(resultInternetArr, resultInternet...)
	}

	// resultTraffIn, resUtilizationTraffIn, resultTraffOut, resUtilizationTraffOut, averageUp, _, sdate, edate, stime, etime, err := g.ucGetData.GetInternetData(input.IdSensor, input.SDate, input.EDate, input.STime, input.ETime)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
	// 		"status":  "FAILED",
	// 		"message": "data not found, date time input maybe wrong",
	// 	})
	// 	return
	// }

	//var resultArr []dto.DataOutput

	// resultData := dto.DataOutput{
	// 	Id:                    input.IdSensor,
	// 	Site:                  input.Site,
	// 	Link:                  "XL",
	// 	AverageUp:             averageUp,
	// 	UtilizationTrafficIn:  resUtilizationTraffIn,
	// 	UtilizationTrafficOut: resUtilizationTraffOut,
	// 	TrafficIn:             resultTraffIn,
	// 	TrafficOut:            resultTraffOut,
	// 	Notes:                 "notes",
	// 	BandwidthCap:          85,
	// }

	//resultArr = append(resultArr, result...)

	ctx.JSON(http.StatusOK, gin.H{
		"status":         "SUCCESS",
		"datetime":       sdate + " " + stime + " - " + edate + " " + etime,
		"resultInternet": resultInternetArr,
	})

}

func NewGetDataSiteController(router *gin.Engine, ucGetData usecase.GetInternetDataSiteUsecase, config config.Config) *GetDataSiteController {
	controller := GetDataSiteController{
		router:    router,
		ucGetData: ucGetData,
		config:    config,
	}

	router.GET("/get-data", controller.GetDataSite)

	return &controller
}
