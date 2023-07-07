package scheduler

import (
	"log"
	absent "presensi/service/absent_service"
	holiday "presensi/service/holiday_service"
	presence "presensi/service/presence_service"
	"time"

	"github.com/robfig/cron/v3"
)

type MainScheduler struct {
	holidayService  *holiday.HolidayService
	presenceService *presence.PresenceService
	absentService   *absent.AbsentService
}

func NewScheduler(hs *holiday.HolidayService, ps *presence.PresenceService, as *absent.AbsentService) *MainScheduler {
	return &MainScheduler{
		holidayService:  hs,
		presenceService: ps,
		absentService:   as,
	}
}

func (s *MainScheduler) Run() {
	scheduler := cron.New(cron.WithLocation(time.Local))
	defer scheduler.Stop()

	scheduler.AddFunc("0 0 1 1 *", s.AutoUpdatePublicHoliday) // At 00:00 on day-of-month 1 in January.
	scheduler.AddFunc("30 15 * * *", s.AutoUpdateJamPulang)   // At 15:30.
	scheduler.AddFunc("00 16 * * *", s.AutoAbsentEmployee)    // At 16:00.
	// scheduler.AddFunc("19 09 * * *", s.AutoUpdateJamPulang)   // At 15:30.
	// scheduler.AddFunc("20 09 * * *", s.AutoAbsentEmployee)    // At 16:00.
	// scheduler.AddFunc("11 09 19 6 *", s.AutoUpdatePublicHoliday) // At 00:00 on day-of-month 1 in January.

	log.Println("scheduler start")
	go scheduler.Start()
}

func (s *MainScheduler) AutoUpdatePublicHoliday() {
	log.Printf("scheduler try to update public holiday")

	err := s.holidayService.SeedPublicHoliday()
	if err != nil {
		log.Println("error while auto update public holiday, err: ", err)
	} else {
		log.Println("success run auto update public holiday")
	}
}

func (s *MainScheduler) AutoAbsentEmployee() {
	log.Printf("scheduler try to update absent employee")

	err := s.absentService.AutoAbsent()
	if err != nil {
		log.Println("error while auto update absent employee, err: ", err)
	} else {
		log.Println("success run auto update absent employee")
	}
}

func (s *MainScheduler) AutoUpdateJamPulang() {
	log.Printf("scheduler try to update jam pulang")

	err := s.presenceService.AutoUpdateJamPulang()
	if err != nil {
		log.Println("error while auto update jam pulang, err: ", err)
	} else {
		log.Println("success run auto update jam pulang")
	}
}
