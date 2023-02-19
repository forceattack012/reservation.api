package infrastructure

import (
	"fmt"
	"log"
	"time"

	"github.com/forceattack012/reservationroom/config"
	"github.com/forceattack012/reservationroom/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDB struct {
	*gorm.DB
}

func NewGormDB(dbConfig *config.Database) *GormDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Dbname, dbConfig.Dbport, dbConfig.Sslmode, dbConfig.Tz)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: true,
		Logger:            logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Bangkok")
			return time.Now().In(ti)
		},
	})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(domain.Person{}, domain.Room{}, domain.Reservation{})
	seedRoom(db)
	seedPerson(db)
	seedReservation(db)

	return &GormDB{db}
}

func seedReservation(db *gorm.DB) {
	var reservations []domain.Reservation
	err := db.Find(&reservations).Error
	if err != nil {
		log.Fatalf("error %s", err)
		return
	}

	if len(reservations) > 0 {
		return
	}

	to := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), 0, 0, 0, time.Local)
	from := to.Add(time.Hour)
	fmt.Printf("%s - %s \n", to, from)

	newReservations := []domain.Reservation{
		{
			Id:            1,
			StartDate:     to,
			EndDate:       from,
			IsReservation: true,
			PersonID:      1,
			RoomID:        1,
		},
		{
			Id:            2,
			StartDate:     to.Add(time.Hour),
			EndDate:       from.Add(time.Hour),
			IsReservation: true,
			PersonID:      1,
			RoomID:        1,
		},
		{
			Id:            3,
			StartDate:     to,
			EndDate:       from,
			IsReservation: true,
			PersonID:      1,
			RoomID:        2,
		},
	}

	db.Save(&newReservations)
	fmt.Printf("init seed reservation success!!")
}

func seedPerson(db *gorm.DB) {
	var people []domain.Person
	err := db.Find(&people).Error
	if err != nil {
		log.Fatalf("error seed person %v", err)
		return
	}

	if len(people) > 0 {
		return
	}

	newPeople := []domain.Person{
		{
			Id:    1,
			Name:  "test",
			Email: "test@gmail",
			Phone: "09122",
		},
	}
	db.Save(newPeople)
	fmt.Println("Intial seed people success!!!")
}

func seedRoom(db *gorm.DB) {
	var rooms []domain.Room
	err := db.Find(&rooms).Error

	if err != nil {
		log.Fatalf("err when seed %s", err.Error())
		return
	}

	if len(rooms) > 0 {
		return
	}

	newRooms := []domain.Room{
		{
			Id:          1,
			RoomName:    "R100",
			Kind:        "Small",
			Price:       100,
			IsAvaliable: true,
		},
		{
			Id:          2,
			RoomName:    "R200",
			Kind:        "Meduiam",
			Price:       200,
			IsAvaliable: true,
		},
		{
			Id:          3,
			RoomName:    "R300",
			Kind:        "Large",
			Price:       300,
			IsAvaliable: true,
		},
	}

	db.Save(newRooms)
	fmt.Println("Intial seed success!!!")
}
