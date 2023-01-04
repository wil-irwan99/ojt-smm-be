package controller

import (
	"net/http"
	"project-ojt/model"
	"project-ojt/model/dto"
	"project-ojt/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SensorController struct {
	router      *gin.Engine
	ucAddSensor usecase.AddNewSensorUsecase
}

func (s *SensorController) AddSensor(ctx *gin.Context) {
	var input dto.DataSensorInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	convToIntBandwidth, _ := strconv.ParseInt(input.Bandwidth, 10, 16)
	convToInt16Bandwidth := int16(convToIntBandwidth)

	err := s.ucAddSensor.AddNewSensor(&model.Sensor{
		Site:      input.Site,
		Link:      input.Link,
		Id:        input.Id,
		Type:      input.Type,
		Bandwidth: convToInt16Bandwidth,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "data sensor failed to add",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": "data sensor added",
	})

}

func (s *SensorController) DeleteSensor(ctx *gin.Context) {
	var input dto.IdInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	err := s.ucAddSensor.DeleteSensor(input.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "data sensor failed to delete",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": "data sensor deleted",
	})

}

func (s *SensorController) RetrieveSensors(ctx *gin.Context) {
	page := ctx.Query("page")

	convInt64Page, _ := strconv.ParseInt(page, 10, 16)
	convIntPage := int(convInt64Page)

	offset := (convIntPage - 1) * 10

	result, err := s.ucAddSensor.RetriveDataSensorPaging(offset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "FAILED",
			"message": "can't retrieve sensors data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      "SUCCESS",
		"message":     "retrieve sensors data success",
		"dataSensors": result,
	})
}

func NewSensorController(router *gin.Engine, ucAddSensor usecase.AddNewSensorUsecase) *SensorController {
	controller := SensorController{
		router:      router,
		ucAddSensor: ucAddSensor,
	}

	rSensor := router.Group("/sensor")
	{
		rSensor.POST("/add", controller.AddSensor)
		rSensor.POST("/delete", controller.DeleteSensor)
		rSensor.GET("/get-sensors", controller.RetrieveSensors)
	}

	return &controller
}
