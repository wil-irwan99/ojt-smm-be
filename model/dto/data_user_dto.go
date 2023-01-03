package dto

type DataSensorInput struct {
	Site      string `json:"site"`
	Link      string `json:"link"`
	Id        string `json:"id"`
	Type      string `json:"type"`
	Bandwidth string `json:"bandwidth"`
}

type IdInput struct {
	Id string `json:"id"`
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

type DataOutputDevice struct {
	Id        string  `json:"id"`
	Site      string  `json:"site"`
	Device    string  `json:"device"`
	Type      string  `json:"type"`
	Usage     float64 `json:"usage"`
	Condition string  `json:"condition"`
	Notes     string  `json:"notes"`
}
