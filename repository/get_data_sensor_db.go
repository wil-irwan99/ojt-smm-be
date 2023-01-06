package repository

import (
	"errors"
	"project-ojt/model"

	"gorm.io/gorm"
)

type GetDataSensorRepository interface {
	RetriveSensors(site string, tipe string) ([]model.Sensor, error)
	RetriveSensorsPaging(offset int) ([]model.Sensor, error)
	AddSensor(sensor *model.Sensor) error
	DeleteSensor(id string) error
	CountSensors() int16
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

func (g *getDataSensorRepository) RetriveSensorsPaging(offset int) ([]model.Sensor, error) {
	var sensors []model.Sensor
	res := g.db.Offset(offset).Limit(10).Find(&sensors)

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

func (g *getDataSensorRepository) CountSensors() int16 {
	var count int64
	g.db.Model(&model.Sensor{}).Count(&count)
	return int16(count)
}

func NewGetDataSensorRepository(db *gorm.DB) GetDataSensorRepository {
	repo := new(getDataSensorRepository)
	repo.db = db
	return repo
}
