package controller

import (
	"net/http"
	"project-ojt/model"
	"project-ojt/model/dto"
	"project-ojt/usecase"

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

	err := s.ucAddSensor.AddNewSensor(&model.Sensor{
		Site:      input.Site,
		Link:      input.Link,
		Id:        input.Id,
		Type:      input.Type,
		Bandwidth: input.Bandwidth,
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

func (s *SensorController) AddSensorDevice(ctx *gin.Context) {
	var input dto.DataDeviceInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	err := s.ucAddSensor.AddNewSensorDevice(&model.Device{
		Location: input.Location,
		Type:     input.Type,
		Category: input.Category,
		Id:       input.Id,
		Site:     input.Site,
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

func (s *SensorController) DeleteSensorDevice(ctx *gin.Context) {
	var input dto.IdInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	err := s.ucAddSensor.DeleteSensorDevice(input.Id)

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

func NewSensorController(router *gin.Engine, ucAddSensor usecase.AddNewSensorUsecase) *SensorController {
	controller := SensorController{
		router:      router,
		ucAddSensor: ucAddSensor,
	}

	rSensor := router.Group("/sensor")
	{
		rSensor.POST("/add", controller.AddSensor)
		rSensor.POST("/add-device", controller.AddSensorDevice)
		rSensor.POST("/delete", controller.DeleteSensor)
		rSensor.POST("/delete-device", controller.DeleteSensorDevice)
	}

	return &controller
}
