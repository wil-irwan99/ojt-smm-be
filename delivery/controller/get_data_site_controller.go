package controller

import (
	"net/http"
	"project-ojt/config"
	"project-ojt/model/dto"
	"project-ojt/usecase"

	"github.com/gin-gonic/gin"
)

type GetDataSiteController struct {
	router       *gin.Engine
	ucGetData    usecase.GetInternetDataSiteUsecase
	ucGetDataCPU usecase.GetCPUDataSiteUsecase
	config       config.Config
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
	var resultIntranetArr []dto.DataOutput
	var resultCPUArr []dto.DataOutputDevice

	switch site {
	case "BIB":
		ipSiteConfig = append(ipSiteConfig, []string{"BIB", g.config.PRTGConfig.IpBIB, g.config.PRTGConfig.User, g.config.PRTGConfig.Password})
	case "KIM":
		ipSiteConfig = append(ipSiteConfig, []string{"KIM", g.config.PRTGConfig.IpKIM, g.config.PRTGConfig.User, g.config.PRTGConfig.Password})
	case "MAL":
		ipSiteConfig = append(ipSiteConfig, []string{"MAL", g.config.PRTGConfig.IpMAL, g.config.PRTGConfig.User, g.config.PRTGConfig.Password})
	case "BSL":
		ipSiteConfig = append(ipSiteConfig, []string{"BSL", g.config.PRTGConfig.IpBSL, g.config.PRTGConfig.User, g.config.PRTGConfig.Password})
	case "SML":
		ipSiteConfig = append(ipSiteConfig, []string{"SML", g.config.PRTGConfig.IpSML, g.config.PRTGConfig.User, g.config.PRTGConfig.PasswordSML})
	case "MSIG":
		ipSiteConfig = append(ipSiteConfig, []string{"MSIG", g.config.PRTGConfig.IpMSIG, g.config.PRTGConfig.User, g.config.PRTGConfig.Password})
	case "BCHO":
		ipSiteConfig = append(ipSiteConfig, []string{"BCHO", g.config.PRTGConfig.IpBCHO, g.config.PRTGConfig.UserBCHO, g.config.PRTGConfig.PasswordBCHO})
	case "all-site":
		ipSiteConfig = append(ipSiteConfig,
			[]string{"BIB", g.config.PRTGConfig.IpBIB, g.config.PRTGConfig.User, g.config.PRTGConfig.Password},
			[]string{"KIM", g.config.PRTGConfig.IpKIM, g.config.PRTGConfig.User, g.config.PRTGConfig.Password},
			[]string{"MAL", g.config.PRTGConfig.IpMAL, g.config.PRTGConfig.User, g.config.PRTGConfig.Password},
			[]string{"BSL", g.config.PRTGConfig.IpBSL, g.config.PRTGConfig.User, g.config.PRTGConfig.Password},
			[]string{"SML", g.config.PRTGConfig.IpSML, g.config.PRTGConfig.User, g.config.PRTGConfig.PasswordSML},
			[]string{"MSIG", g.config.PRTGConfig.IpMSIG, g.config.PRTGConfig.User, g.config.PRTGConfig.Password},
			[]string{"BCHO", g.config.PRTGConfig.IpBCHO, g.config.PRTGConfig.UserBCHO, g.config.PRTGConfig.PasswordBCHO})
	}

	for i := 0; i < len(ipSiteConfig); i++ {
		resultInternet, err := g.ucGetData.GetInternetDataSite(ipSiteConfig[i][0], "internet", ipSiteConfig[i][1], ipSiteConfig[i][2], ipSiteConfig[i][3], sdate, edate, stime, etime)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "FAILED",
				"message": "data not found, date time input maybe wrong",
			})
			return
		}
		resultInternetArr = append(resultInternetArr, resultInternet...)

		resultIntranet, err := g.ucGetData.GetInternetDataSite(ipSiteConfig[i][0], "intranet", ipSiteConfig[i][1], ipSiteConfig[i][2], ipSiteConfig[i][3], sdate, edate, stime, etime)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "FAILED",
				"message": "data not found, date time input maybe wrong",
			})
			return
		}
		resultIntranetArr = append(resultIntranetArr, resultIntranet...)

		resultCPU, err := g.ucGetDataCPU.GetCPUDataSite(ipSiteConfig[i][0], ipSiteConfig[i][1], ipSiteConfig[i][2], ipSiteConfig[i][3], sdate, edate, stime, etime)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "FAILED",
				"message": "data not found, date time input maybe wrong",
			})
			return
		}
		resultCPUArr = append(resultCPUArr, resultCPU...)

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
		"resultIntranet": resultIntranetArr,
		"resultCPU":      resultCPUArr,
	})

}

func NewGetDataSiteController(router *gin.Engine, ucGetData usecase.GetInternetDataSiteUsecase, ucGetDataCPU usecase.GetCPUDataSiteUsecase, config config.Config) *GetDataSiteController {
	controller := GetDataSiteController{
		router:       router,
		ucGetData:    ucGetData,
		ucGetDataCPU: ucGetDataCPU,
		config:       config,
	}

	router.GET("/get-data", controller.GetDataSite)

	return &controller
}
