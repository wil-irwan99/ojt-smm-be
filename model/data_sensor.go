package model

type Sensor struct {
	Site      string
	Link      string
	Id        string `gorm:"primaryKey"`
	Type      string
	Bandwidth int16
}

func (Sensor) TableName() string {
	return "data_sensor_table"
}
