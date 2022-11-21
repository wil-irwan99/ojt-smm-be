package model

type BandwidthCapacity struct {
	Site         string
	BandWidthCap int16
}

func (BandwidthCapacity) TableName() string {
	return "data_bandwidth_table"
}
