package usecase

import (
	"math"
	"project-ojt/repository"
)

type GetInternetDataUsecase interface {
	GetInternetData(id_sensor string, sdate string, edate string, stime string, etime string) (float64, float64, float64, float64, float64, float64, string, string, string, string, error)
}

type getInternetDataUsecase struct {
	getDataRepo repository.GetDataRepository
}

func (g *getInternetDataUsecase) GetInternetData(id_sensor string, sdate string, edate string, stime string, etime string) (float64, float64, float64, float64, float64, float64, string, string, string, string, error) {

	result, err := g.getDataRepo.GetJson(id_sensor, sdate, edate, stime, etime)

	if err == nil {

		var averageMbitResultTrafficIn float64 = 0
		var averageMbitResultTrafficOut float64 = 0
		var utilizationTrafficIn float64
		var utilizationTrafficOut float64

		var averageUp float64
		var averageDown float64

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
				convertFloatDowntime = 0
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
		averageDown = 100 - averageUp

		utilizationTrafficIn = math.Round(((averageMbitResultTrafficIn/85)*100)*100) / 100
		utilizationTrafficOut = math.Round(((averageMbitResultTrafficOut/85)*100)*100) / 100

		return averageMbitResultTrafficIn, utilizationTrafficIn, averageMbitResultTrafficOut, utilizationTrafficOut, averageUp, averageDown, sdate, edate, stime, etime, nil

	} else {
		return 0, 0, 0, 0, 0, 0, sdate, edate, stime, etime, err
	}

}

func NewGetInternetDataUsecase(getDataRepo repository.GetDataRepository) GetInternetDataUsecase {
	return &getInternetDataUsecase{
		getDataRepo: getDataRepo,
	}
}
