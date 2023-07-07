package router

import (
	"log"
	absent_controller "presensi/controller/absent_controller"
	auth_controller "presensi/controller/auth_controller"
	employee_controller "presensi/controller/employee_controller"
	holiday_controller "presensi/controller/holiday_controller"
	position_controller "presensi/controller/position_controller"
	presence_controller "presensi/controller/presence_controller"
	presence_status_controller "presensi/controller/presence_status_controller"
	role_controller "presensi/controller/role_controller"
	work_times_controller "presensi/controller/work_times_controller"
	mdl "presensi/middleware"
	absent_repo "presensi/repository/absent_repository"
	auth_repo "presensi/repository/auth_repository"
	employee_repo "presensi/repository/employee_repository"
	holiday_repo "presensi/repository/holiday_repository"
	position_repo "presensi/repository/position_repository"
	presence_repo "presensi/repository/presence_repository"
	presence_status_repo "presensi/repository/presence_status_repository"
	role_repo "presensi/repository/role_repository"
	work_times_repo "presensi/repository/work_times_repository"
	"presensi/scheduler"
	absent_service "presensi/service/absent_service"
	auth_service "presensi/service/auth_service"
	employee_service "presensi/service/employee_service"
	holiday_service "presensi/service/holiday_service"
	position_service "presensi/service/position_service"
	presence_service "presensi/service/presence_service"
	presence_status_service "presensi/service/presence_status_service"
	role_service "presensi/service/role_service"
	work_times_service "presensi/service/work_times_service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Setup struct {
	App        *echo.Echo
	Group      *echo.Group
	DB         *gorm.DB
	Scheduler  *scheduler.MainScheduler
	Middleware *mdl.Middleware
}

func Router(setup *Setup) {
	employeeRoute(setup)
	workTimesRoute(setup)
	roleRoute(setup)
	presenceStatusRoute(setup)
	positionRoute(setup)
	holidayRoute(setup)
	presenceRoute(setup)
	absentRoute(setup)
	authRoute(setup)
}

func employeeRoute(setup *Setup) {
	repo := employee_repo.NewEmployeeRepository(setup.DB)
	service := employee_service.NewEmployeeService(repo)
	controller := employee_controller.NewEmployeeController(service)

	setup.Group.GET("/employees", controller.GetEmployees, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/employee/:id", controller.GetEmployeeByNIP, setup.Middleware.AllRole)
	setup.Group.POST("/employee", controller.NewEmployee, setup.Middleware.OnlyAdmin)
	setup.Group.PUT("/employee", controller.UpdateEmployee, setup.Middleware.AllRole)
	setup.Group.DELETE("/employee/:id", controller.DeleteEmployee, setup.Middleware.OnlyAdmin)
}

func workTimesRoute(setup *Setup) {
	repo := work_times_repo.NewWorkTimeRepository(setup.DB)
	service := work_times_service.NewWorkTimesService(repo)
	controller := work_times_controller.NewWorkTimesController(service)

	setup.Group.GET("/work-times", controller.GetWorkTimes, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/work-time/:id", controller.GetWorkTimesByID, setup.Middleware.AllRole)
	setup.Group.POST("/work-time", controller.NewWorkTimes, setup.Middleware.OnlyAdmin)
	setup.Group.PUT("/work-time", controller.UpdateWorkTimes, setup.Middleware.OnlyAdmin)
}

func roleRoute(setup *Setup) {
	repo := role_repo.NewRolesRepository(setup.DB)
	service := role_service.NewRoleservice(repo)
	controller := role_controller.NewRolesController(service)

	setup.Group.GET("/roles", controller.GetRoles, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/role/:id", controller.GetRolesByID, setup.Middleware.AllRole)
	setup.Group.POST("/role", controller.NewRoles, setup.Middleware.OnlyAdmin)
	setup.Group.PUT("/role", controller.UpdateRoles, setup.Middleware.OnlyAdmin)
}

func presenceStatusRoute(setup *Setup) {
	repo := presence_status_repo.NewPresenceStatusRepository(setup.DB)
	service := presence_status_service.NewPresenceStatuservice(repo)
	controller := presence_status_controller.NewPresenceStatusController(service)

	setup.Group.GET("/presence-status", controller.GetPresenceStatus, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/presence-status/:id", controller.GetPresenceStatusByID, setup.Middleware.AllRole)
	setup.Group.POST("/presence-status", controller.NewPresenceStatus, setup.Middleware.OnlyAdmin)
	setup.Group.PUT("/presence-status", controller.UpdatePresenceStatus, setup.Middleware.OnlyAdmin)
}

func positionRoute(setup *Setup) {
	repo := position_repo.NewPositionRepository(setup.DB)
	service := position_service.NewPositionervice(repo)
	controller := position_controller.NewPositionController(service)

	setup.Group.GET("/positions", controller.GetPosition, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/position/:id", controller.GetPositionByID, setup.Middleware.AllRole)
	setup.Group.POST("/position", controller.NewPosition, setup.Middleware.OnlyAdmin)
	setup.Group.PUT("/position", controller.UpdatePosition, setup.Middleware.OnlyAdmin)
}

func holidayRoute(setup *Setup) {
	repo := holiday_repo.NewHolidayRepository(setup.DB)
	service := holiday_service.NewHolidayervice(repo)
	controller := holiday_controller.NewHolidayController(service)

	if err := service.SeedPublicHoliday(); err != nil {
		log.Println("failed to seed public holiday :", err)
	}

	setup.Group.GET("/holidays", controller.GetHolidays, setup.Middleware.AllRole)
	setup.Group.GET("/holiday/:id", controller.GetHolidayByID, setup.Middleware.AllRole)
	setup.Group.POST("/holiday", controller.NewHoliday, setup.Middleware.OnlyAdmin)
	setup.Group.PUT("/holiday", controller.UpdateHoliday, setup.Middleware.OnlyAdmin)
	setup.Group.DELETE("/holiday/:id", controller.DeleteHoliday, setup.Middleware.OnlyAdmin)
}

func presenceRoute(setup *Setup) {
	absentRepo := absent_repo.NewAbsentRepository(setup.DB)
	absentService := absent_service.NewAbsentService(absentRepo)

	holidayRepo := holiday_repo.NewHolidayRepository(setup.DB)
	holidayService := holiday_service.NewHolidayervice(holidayRepo)

	repo := presence_repo.NewPresenceRepository(setup.DB)
	service := presence_service.NewPresenceRepository(repo, absentService, holidayService)
	controller := presence_controller.NewPresenceRepository(service)

	setup.Group.POST("/presence", controller.InputPresence, setup.Middleware.AllRole)
	setup.Group.GET("/daily-presence/report", controller.GetPresenceDailyReport, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/daily-presence", controller.GetPresenceDaily, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/monthly-presence/report", controller.GetPresenceMonthlyReport, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/monthly-presence", controller.GetPresenceMonthly, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/yearly-presence/report", controller.GetPresenceYearlyReport, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/yearly-presence", controller.GetPresenceYearly, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/monthly-presence-detail/report", controller.GetPresenceMonthlyWithNIPReport, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/monthly-presence-detail", controller.GetPresenceMonthlyWithNIP, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/yearly-presence-detail/report", controller.GetPresenceYearlyWithNIPReport, setup.Middleware.OnlyAdmin)
	setup.Group.GET("/yearly-presence-detail", controller.GetPresenceYearlyWithNIP, setup.Middleware.OnlyAdmin)
	setup.Group.PUT("/jam-pulang", controller.UpdateJamPulang, setup.Middleware.OnlyAdmin)
}

func absentRoute(setup *Setup) {
	repo := absent_repo.NewAbsentRepository(setup.DB)
	service := absent_service.NewAbsentService(repo)
	controller := absent_controller.NewAbsentController(service)

	setup.Group.GET("/absent", controller.GetAbsentDaily, setup.Middleware.OnlyAdmin)
	setup.Group.POST("/absent", controller.InputAbsent, setup.Middleware.OnlyAdmin)
}

func authRoute(setup *Setup) {
	repo := auth_repo.NewAuthRepository(setup.DB)
	service := auth_service.NewAuthService(repo)
	controller := auth_controller.NewAuthController(service)

	setup.Group.POST("/login", controller.Auth)
}
