package repository

import (
	"errors"
	"project-ojt/model"

	"gorm.io/gorm"
)

type GetDataSensorRepository interface {
	RetriveSensors(site string, tipe string) ([]model.Sensor, error)
	AddSensor(sensor *model.Sensor) error
	DeleteSensor(id string) error
}

type getDataSensorRepository struct {
	db *gorm.DB
}

func (g *getDataSensorRepository) RetriveSensors(site string, tipe string) ([]model.Sensor, error) {
	var sensors []model.Sensor
	res := g.db.Where("site = ? AND type = ?", site, tipe).Find(&sensors)

	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return sensors, nil
}

func (g *getDataSensorRepository) AddSensor(sensor *model.Sensor) error {
	result := g.db.Create(sensor)
	return result.Error
}

func (g *getDataSensorRepository) DeleteSensor(id string) error {
	result := g.db.Unscoped().Where("id = ?", id).Delete(&model.Sensor{})
	return result.Error
}

func NewGetDataSensorRepository(db *gorm.DB) GetDataSensorRepository {
	repo := new(getDataSensorRepository)
	repo.db = db
	return repo
}
