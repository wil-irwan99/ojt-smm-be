package usecase

import (
	"project-ojt/model"
	"project-ojt/repository"
)

type AddNewSensorUsecase interface {
	AddNewSensor(sensor *model.Sensor) error
}

type addNewSensorUsecase struct {
	getDataSensorRepo repository.GetDataSensorRepository
}

func (a *addNewSensorUsecase) AddNewSensor(sensor *model.Sensor) error {
	err := a.getDataSensorRepo.AddSensor(sensor)
	if err != nil {
		return err
	}
	return nil
}

func NewAddNewSensorUsecase(getDataSensorRepo repository.GetDataSensorRepository) AddNewSensorUsecase {
	return &addNewSensorUsecase{
		getDataSensorRepo: getDataSensorRepo,
	}
}
