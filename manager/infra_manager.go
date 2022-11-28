package manager

import (
	"log"
	"project-ojt/config"
	"project-ojt/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	ConfigData() config.Config
	SqlDb() *gorm.DB
	Dummy() string
}

type infra struct {
	db    *gorm.DB
	cfg   config.Config
	dummy string
}

func (i *infra) ConfigData() config.Config {
	return i.cfg
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func (i *infra) Dummy() string {
	return i.dummy
}

func NewInfra(config config.Config) Infra {

	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	infra := infra{
		cfg:   config,
		db:    resource,
		dummy: "",
	}
	return &infra
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("connected")
	}

	// env := "migration"

	// if env == "dev" {
	// 	db = db.Debug()
	// } else if env == "migration" {
	db = db.Debug()
	db.AutoMigrate(
		&model.Sensor{},
		//&model.BandwidthCapacity{},
		&model.Device{},
	)
	if err != nil {
		return nil, err
	}
	// }
	return db, nil
}
