package usecase

import (
	"project-ojt/model"
	"project-ojt/repository"
)

type AddNewSensorUsecase interface {
	AddNewSensor(sensor *model.Sensor) error
	DeleteSensor(id string) error
	RetriveDataSensorPaging(offset int) ([]model.Sensor, error)
	CountDataSensor() int16
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

func (a *addNewSensorUsecase) DeleteSensor(id string) error {
	err := a.getDataSensorRepo.DeleteSensor(id)
	if err != nil {
		return err
	}
	return nil
}

func (a *addNewSensorUsecase) RetriveDataSensorPaging(offset int) ([]model.Sensor, error) {
	result, err := a.getDataSensorRepo.RetriveSensorsPaging(offset)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *addNewSensorUsecase) CountDataSensor() int16 {
	result := a.getDataSensorRepo.CountSensors()
	return result
}

func NewAddNewSensorUsecase(getDataSensorRepo repository.GetDataSensorRepository) AddNewSensorUsecase {
	return &addNewSensorUsecase{
		getDataSensorRepo: getDataSensorRepo,
	}
}
