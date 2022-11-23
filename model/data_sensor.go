package model

type Sensor struct {
	Site     string
	Link     string
	Id       string `gorm:"primaryKey"`
	Type     string
	DataType string
}

func (Sensor) TableName() string {
	return "data_sensor_table"
}
