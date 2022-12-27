package usecase

import (
	"project-ojt/model"
	"project-ojt/repository"
)

type AddNewSensorUsecase interface {
	AddNewSensor(sensor *model.Sensor) error
	AddNewSensorDevice(device *model.Device) error
	DeleteSensor(id string) error
	DeleteSensorDevice(id string) error
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

func (a *addNewSensorUsecase) AddNewSensorDevice(device *model.Device) error {
	err := a.getDataSensorRepo.AddSensorDevice(device)
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

func (a *addNewSensorUsecase) DeleteSensorDevice(id string) error {
	err := a.getDataSensorRepo.DeleteSensorDevice(id)
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
