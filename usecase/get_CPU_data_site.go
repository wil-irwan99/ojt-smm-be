package usecase

import (
	"math"
	"project-ojt/model"
	"project-ojt/model/dto"
	"project-ojt/repository"
)

type GetCPUDataSiteUsecase interface {
	GetCPUDataSite(site string, ip string, user string, password string, sdate string, edate string, stime string, etime string) ([]dto.DataOutputDevice, error)
}

type getCPUDataSiteUsecase struct {
	getDataRepo       repository.GetDataRepository
	getDataSensorRepo repository.GetDataSensorRepository
}

func (g *getCPUDataSiteUsecase) GetCPUDataSite(site string, ip string, user string, password string, sdate string, edate string, stime string, etime string) ([]dto.DataOutputDevice, error) {
	var devices []model.Device
	devices, err := g.getDataSensorRepo.RetriveDevices(site)
	if err != nil {
		return nil, err
	}

	var resultArr []dto.DataOutputDevice

	for i := 0; i < len(devices); i++ {

		result, err := g.getDataRepo.GetJson(devices[i].Id, ip, user, password, sdate, edate, stime, etime)

		var key string

		for k := range result.HistDatas[i] {
			if k != "datetime" {
				if k != "Downtime" {
					if k != "coverage" {
						key = k
					}
				}
			}
		}

		if err == nil {

			var averageValue float64

			for i := 0; i < len(result.HistDatas); i++ {

				rawDataValue := result.HistDatas[i][key]

				var convertFloatValue float64

				switch j := rawDataValue.(type) {
				case float64:
					convertFloatValue = j
				case float32:
					convertFloatValue = float64(j)
				case int64:
					convertFloatValue = float64(j)
				default:
					convertFloatValue = 0
				}

				averageValue += convertFloatValue
			}

			averageValue = math.Round((averageValue/float64(len(result.HistDatas)))*100) / 100

			var condition string

			if averageValue <= 70 {
				condition = "Normal"
			} else if averageValue <= 80 {
				condition = "Warning"
			} else {
				condition = "Critical"
			}

			resultData := dto.DataOutputDevice{
				Id:        devices[i].Id,
				Location:  devices[i].Location,
				Type:      devices[i].Type,
				Category:  devices[i].Category,
				Usage:     averageValue,
				Condition: condition,
				Notes:     "",
			}

			resultArr = append(resultArr, resultData)
		} else {
			return nil, err
		}
	}

	return resultArr, nil
}

func NewGetCPUDataSiteUsecase(getDataRepo repository.GetDataRepository, getDataSensorRepo repository.GetDataSensorRepository) GetCPUDataSiteUsecase {
	return &getCPUDataSiteUsecase{
		getDataRepo:       getDataRepo,
		getDataSensorRepo: getDataSensorRepo,
	}
}
