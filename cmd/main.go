package main

import (
	"fmt"
	"time"

	"github.com/forceattack012/reservationroom/application"
	"github.com/forceattack012/reservationroom/config"
	"github.com/forceattack012/reservationroom/handler"
	"github.com/forceattack012/reservationroom/infrastructure"
	"github.com/forceattack012/reservationroom/router"
)

var appConfig config.Config

func init() {
	// err := os.Setenv("TZ", "Asia/Bangkok")
	// if err != nil {
	// 	log.Fatalf("error %s", err.Error())
	// }
	fmt.Printf("%s", time.Now().Format("2006-01-02 15:04:05"))
	config.NewConfig().ReadConfigYaml("./config/config.yaml", &appConfig)
}

func main() {
	gormDB := infrastructure.NewGormDB(&appConfig.Database)
	r := router.NewFiberRoute()
	initPerson(r, gormDB)
	initRoom(r, gormDB)
	initResavation(r, gormDB)

	err := r.Listen(":" + appConfig.Port)
	if err != nil {
		panic(err)
	}
}

func initPerson(r *router.FiberRouter, db *infrastructure.GormDB) {
	personRepo := infrastructure.NewPersonRepository(db)
	personService := application.NewPersonService(personRepo)
	personHandler := handler.NewPersonHandler(personService)

	r.Get("/api/person", personHandler.GetAll)
	r.Delete("/api/person/:id", personHandler.DeletePerson)
	r.Post("/api/person", personHandler.NewPerson)
}

func initRoom(r *router.FiberRouter, db *infrastructure.GormDB) {
	repo := infrastructure.NewRoomRepository(db)
	roomService := application.NewRoomService(repo)
	roomHandler := handler.NewRoomHandler(roomService)

	r.Get("/api/room", roomHandler.GetRoomAll)
	r.Post("/api/room", roomHandler.CreateRoom)
	r.Patch("/api/room/:roomId", roomHandler.UpdateRoom)
}

func initResavation(r *router.FiberRouter, db *infrastructure.GormDB) {
	repo := infrastructure.NewReservationRepository(db)
	service := application.NewReservationService(repo)
	handler := handler.NewReservationHandler(service)

	r.Get("/api/reservation/:roomId", handler.GetReservationByRoomId)
	r.Post("/api/reservation", handler.CreateReservation)
	r.Delete("/api/reservation/:id", handler.DeleteReservation)
}
