package usecase

import (
	"math"
	"project-ojt/model"
	"project-ojt/model/dto"
	"project-ojt/repository"
)

type GetInternetDataSiteUsecase interface {
	GetInternetDataSite(site string, tipe string, ip string, user string, password string, sdate string, edate string, stime string, etime string) ([]dto.DataOutput, error)
}

type getInternetDataSiteUsecase struct {
	getDataRepo       repository.GetDataRepository
	getDataSensorRepo repository.GetDataSensorRepository
}

func (g *getInternetDataSiteUsecase) GetInternetDataSite(site string, tipe string, ip string, user string, password string, sdate string, edate string, stime string, etime string) ([]dto.DataOutput, error) {
	var sensors []model.Sensor
	sensors, err := g.getDataSensorRepo.RetriveSensors(site, tipe)
	if err != nil {
		return nil, err
	}

	// var bandwidth model.BandwidthCapacity
	// bandwidth, err = g.getDataSensorRepo.RetriveBandwidth(site)
	// if err != nil {
	// 	return nil, err
	// }

	var resultArr []dto.DataOutput

	for i := 0; i < len(sensors); i++ {

		result, err := g.getDataRepo.GetJson(sensors[i].Id, ip, user, password, sdate, edate, stime, etime)

		if err == nil {

			var averageMbitResultTrafficIn float64 = 0
			var averageMbitResultTrafficOut float64 = 0
			var utilizationTrafficIn float64
			var utilizationTrafficOut float64

			var averageUp float64

			for i := 0; i < len(result.HistDatas); i++ {

				rawDataTrafficIn := result.HistDatas[i]["Traffic In (speed)"]
				rawDataTrafficOut := result.HistDatas[i]["Traffic Out (speed)"]
				rawDataDowntime := result.HistDatas[i]["Downtime"]

				var convertMbitResultTrafficIn float64
				var convertMbitResultTrafficOut float64
				var convertFloatDowntime float64

				switch j := rawDataTrafficIn.(type) {
				case float64:
					convertMbitResultTrafficIn = j * 0.008
				case float32:
					convertMbitResultTrafficIn = float64(j) * 0.008
				case int64:
					convertMbitResultTrafficIn = float64(j) * 0.008
				default:
					convertMbitResultTrafficIn = 0
				}

				switch k := rawDataTrafficOut.(type) {
				case float64:
					convertMbitResultTrafficOut = k * 0.008
				case float32:
					convertMbitResultTrafficOut = float64(k) * 0.008
				case int64:
					convertMbitResultTrafficOut = float64(k) * 0.008
				default:
					convertMbitResultTrafficOut = 0
				}

				switch l := rawDataDowntime.(type) {
				case float64:
					convertFloatDowntime = l
				case float32:
					convertFloatDowntime = float64(l)
				case int64:
					convertMbitResultTrafficOut = float64(l)
				default:
					convertFloatDowntime = 100
				}

				convertMbitResultTrafficIn = math.Round(convertMbitResultTrafficIn / 1000)
				convertMbitResultTrafficOut = math.Round(convertMbitResultTrafficOut / 1000)

				averageMbitResultTrafficIn += convertMbitResultTrafficIn
				averageMbitResultTrafficOut += convertMbitResultTrafficOut
				averageUp += 100 - convertFloatDowntime
			}

			averageMbitResultTrafficIn = averageMbitResultTrafficIn / float64(len(result.HistDatas))
			averageMbitResultTrafficOut = averageMbitResultTrafficOut / float64(len(result.HistDatas))
			averageMbitResultTrafficIn = math.Round(averageMbitResultTrafficIn)
			averageMbitResultTrafficOut = math.Round(averageMbitResultTrafficOut)

			averageUp = math.Round(averageUp / float64(len(result.HistDatas)))
			// averageDown = 100 - averageUp

			utilizationTrafficIn = math.Round(((averageMbitResultTrafficIn/float64(sensors[i].Bandwidth))*100)*100) / 100
			utilizationTrafficOut = math.Round(((averageMbitResultTrafficOut/float64(sensors[i].Bandwidth))*100)*100) / 100

			resultData := dto.DataOutput{
				Id:                    sensors[i].Id,
				Site:                  site,
				Link:                  sensors[i].Link,
				AverageUp:             averageUp,
				UtilizationTrafficIn:  utilizationTrafficIn,
				UtilizationTrafficOut: utilizationTrafficOut,
				TrafficIn:             averageMbitResultTrafficIn,
				TrafficOut:            averageMbitResultTrafficOut,
				Notes:                 "",
				BandwidthCap:          float64(sensors[i].Bandwidth),
			}

			resultArr = append(resultArr, resultData)

		} else {
			return nil, err
		}

	}

	return resultArr, nil

}

func NewGetInternetDataSiteUsecase(getDataRepo repository.GetDataRepository, getDataSensorRepo repository.GetDataSensorRepository) GetInternetDataSiteUsecase {
	return &getInternetDataSiteUsecase{
		getDataRepo:       getDataRepo,
		getDataSensorRepo: getDataSensorRepo,
	}
}
