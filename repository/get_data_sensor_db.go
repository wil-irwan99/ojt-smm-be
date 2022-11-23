package repository

import (
	"errors"
	"project-ojt/model"

	"gorm.io/gorm"
)

type GetDataSensorRepository interface {
	RetriveSensors(site string, tipe string, tipedata string) ([]model.Sensor, error)
	RetriveBandwidth(site string) (model.BandwidthCapacity, error)
}

type getDataSensorRepository struct {
	db *gorm.DB
}

func (g *getDataSensorRepository) RetriveSensors(site string, tipe string, tipedata string) ([]model.Sensor, error) {
	var sensors []model.Sensor
	res := g.db.Where("site = ? AND type = ? AND data_type = ?", site, tipe, tipedata).Find(&sensors)

	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return sensors, nil
}

func (g *getDataSensorRepository) RetriveBandwidth(site string) (model.BandwidthCapacity, error) {
	var bandwidth model.BandwidthCapacity
	res := g.db.Where("site = ?", site).Last(&bandwidth)

	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.BandwidthCapacity{}, nil
		} else {
			return model.BandwidthCapacity{}, err
		}
	}
	return bandwidth, nil
}

func NewGetDataSensorRepository(db *gorm.DB) GetDataSensorRepository {
	repo := new(getDataSensorRepository)
	repo.db = db
	return repo
}
