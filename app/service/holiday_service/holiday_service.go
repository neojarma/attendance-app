package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"presensi/helper"
	"presensi/model"
	repo "presensi/repository/holiday_repository"
	"strconv"
	"time"
)

type HolidayService struct {
	holidayRepository *repo.HolidayRepository
}

func NewHolidayervice(hr *repo.HolidayRepository) *HolidayService {
	return &HolidayService{
		holidayRepository: hr,
	}
}

func (s *HolidayService) NewHoliday(payload *model.Holidays) error {
	payload.HolidayID = helper.GenerateRandString(5)
	return s.holidayRepository.NewHoliday(payload)
}

func (s *HolidayService) GetHolidayByID(payload string) (*model.Holidays, error) {
	Holiday, err := s.holidayRepository.GetHolidayByID(payload)
	if err != nil {
		return nil, err
	}

	if Holiday.HolidayID == "" {
		return nil, errors.New("there is no record with that id")
	}

	Holiday.Date = helper.ParseDateOnly(Holiday.Date)

	return Holiday, nil
}

func (s *HolidayService) GetHoliday() ([]*model.Holidays, error) {
	res, err := s.holidayRepository.GetHoliday()
	if err != nil {
		return nil, err
	}

	for _, v := range res {
		v.Date = helper.ParseDateOnly(v.Date)
	}

	return res, nil
}

func (s *HolidayService) UpdateHoliday(payload *model.Holidays) error {
	return s.holidayRepository.UpdateHoliday(payload)
}

func (s *HolidayService) DeleteHoliday(payload string) error {
	return s.holidayRepository.DeleteHoliday(payload)
}

func (s *HolidayService) SeedPublicHoliday() error {
	currentYear := time.Now().Year()
	if s.holidayRepository.IsHolidayExistCurrentYear(currentYear) {
		log.Printf("public holiday for %v is exist in database", currentYear)
		return nil
	}

	log.Printf("public holiday for %v is not exist in database", currentYear)
	holidays, err := s.GetPublicHolidayFromAPI(strconv.Itoa(currentYear))
	if err != nil {
		return err
	}

	log.Println("try to insert data to database...")
	err = s.holidayRepository.SeedPublicHoliday(holidays)
	if err != nil {
		return err
	}

	log.Println("finish insert data to database")
	return nil
}

func (s *HolidayService) GetPublicHolidayFromAPI(year string) ([]*model.Holidays, error) {
	key := os.Getenv("CALENDARIFIC_API_KEY")
	url := fmt.Sprintf("https://calendarific.com/api/v2/holidays?&api_key=%s&country=id&year=%s", key, year)

	log.Println("try to get list of public holiday...")
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	bodyContent := new(model.HolidaysAPIResponse)
	err = json.Unmarshal(rawBody, bodyContent)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	finalRes := make([]*model.Holidays, 0)
	for _, v := range bodyContent.Response.Holidays {
		if v.Type[0] == "National holiday" {
			if err != nil {
				log.Println(err)
				return nil, err
			}

			finalRes = append(finalRes, &model.Holidays{
				HolidayID:   helper.GenerateRandString(5),
				Date:        v.Date.ISO,
				Description: v.Name,
			})
		}
	}

	log.Println("finish get list public holiday")
	return finalRes, nil
}

func (s *HolidayService) IsTodayAHoliday(date string) (bool, error) {
	return s.holidayRepository.IsTodayAHoliday(date)
}
