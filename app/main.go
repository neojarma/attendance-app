package main

import (
	"fmt"
	"log"
	"os"
	"presensi/connection"
	"presensi/repository/absent_repository"
	holiday_repo "presensi/repository/holiday_repository"
	presence_repo "presensi/repository/presence_repository"
	"presensi/router"
	"presensi/scheduler"
	service "presensi/service/absent_service"
	holiday_service "presensi/service/holiday_service"
	presence_service "presensi/service/presence_service"

	"github.com/labstack/echo/v4"
)

func setup() (*router.Setup, error) {

	setup := new(router.Setup)
	app := echo.New()

	db, err := connection.SQLServerConn()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// scheduler setup
	holidayRepo := holiday_repo.NewHolidayRepository(db)
	holidayService := holiday_service.NewHolidayervice(holidayRepo)

	absentRepo := absent_repository.NewAbsentRepository(db)
	absentService := service.NewAbsentService(absentRepo)

	presenceRepo := presence_repo.NewPresenceRepository(db)
	presenceService := presence_service.NewPresenceRepository(presenceRepo, absentService, holidayService)

	scheduler := scheduler.NewScheduler(holidayService, presenceService, absentService)

	setup.DB = db
	setup.Scheduler = scheduler
	setup.Group = app.Group("/api")
	setup.App = app

	return setup, nil
}

func main() {
	setup, err := setup()
	if err != nil {
		log.Println(err)
		return
	}

	router.Router(setup)

	setup.Scheduler.Run()
	setup.App.Start(fmt.Sprint(":", os.Getenv("PORT")))
}
