package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	global *gorm.DB
)

func open(dsn string, conf *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), conf)
}

func Init(dsn string) (err error) {
	global, err = open(dsn, &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func New(dsn string, conf *gorm.Config) (*gorm.DB, error) {
	return open(dsn, conf)
}
