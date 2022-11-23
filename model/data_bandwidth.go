package model

type BandwidthCapacity struct {
	Site         string
	BandwidthCap int16
}

func (BandwidthCapacity) TableName() string {
	return "data_bandwidth_table"
}
