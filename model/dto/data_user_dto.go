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
	TrafficIn             float64 `json:"traffic_in"`
	UtilizationTrafficIn  float64 `json:"uti_traffic_in"`
	TrafficOut            float64 `json:"traffic_out"`
	UtilizationTrafficOut float64 `json:"uti_traffic_out"`
	AverageUp             float64 `json:"average_up"`
}
