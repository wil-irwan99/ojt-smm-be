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

func NewSensorController(router *gin.Engine, ucAddSensor usecase.AddNewSensorUsecase) *SensorController {
	controller := SensorController{
		router:      router,
		ucAddSensor: ucAddSensor,
	}

	rSensor := router.Group("/sensor")
	{
		rSensor.POST("/add", controller.AddSensor)
	}

	return &controller
}
