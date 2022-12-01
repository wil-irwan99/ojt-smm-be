package model

type Device struct {
	Location string
	Type     string
	Category string
	Id       string `gorm:"primaryKey"`
	Site     string
}

func (Device) TableName() string {
	return "data_sensor_network_device"
}
