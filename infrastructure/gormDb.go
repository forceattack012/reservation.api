package infrastructure

import (
	"fmt"

	"github.com/forceattack012/reservationroom/config"
	"github.com/forceattack012/reservationroom/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	*gorm.DB
}

func NewGormDB(dbConfig *config.Database) *GormDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Dbname, dbConfig.Dbport, dbConfig.Sslmode, dbConfig.Tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: true,
	})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(domain.Person{}, domain.Room{})

	return &GormDB{db}
}
