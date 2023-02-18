package main

import (
	"github.com/forceattack012/reservationroom/application"
	"github.com/forceattack012/reservationroom/config"
	"github.com/forceattack012/reservationroom/handler"
	"github.com/forceattack012/reservationroom/infrastructure"
	"github.com/forceattack012/reservationroom/router"
)

var appConfig config.Config

func init() {
	config.NewConfig().ReadConfigYaml("./config/config.yaml", &appConfig)
}

func main() {
	gormDB := infrastructure.NewGormDB(&appConfig.Database)
	r := router.NewFiberRoute()
	initPerson(r, gormDB)
	initRoom(r, gormDB)

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
	r.Delete("/api/person", personHandler.DeletePerson)
	r.Post("/api/person", personHandler.NewPerson)
}

func initRoom(r *router.FiberRouter, db *infrastructure.GormDB) {
	repo := infrastructure.NewRoomRepository(db)
	roomService := application.NewRoomService(repo)
	roomHandler := handler.NewRoomHandler(roomService)

	r.Get("/api/room", roomHandler.GetRoomAll)
	r.Post("/api/room", roomHandler.CreateRoom)
}
