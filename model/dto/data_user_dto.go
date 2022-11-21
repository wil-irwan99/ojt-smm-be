package dto

type DataInput struct {
	Site     string `json:"site"`
	IdSensor string `json:"id_sensor"`
	SDate    string `json:"sdate"`
	EDate    string `json:"edate"`
	STime    string `json:"stime"`
	ETime    string `json:"etime"`
}

type DataOutput struct {
	Id                    string  `json:"id"`
	Site                  string  `json:"site"`
	Link                  string  `json:"link"`
	AverageUp             float64 `json:"average_up"`
	UtilizationTrafficIn  float64 `json:"uti_traffic_in"`
	UtilizationTrafficOut float64 `json:"uti_traffic_out"`
	TrafficIn             float64 `json:"traffic_in"`
	TrafficOut            float64 `json:"traffic_out"`
	Notes                 string  `json:"notes"`
	BandwidthCap          float64 `json:"bandwidth_cap"`
}
