package model

type Holidays struct {
	HolidayID   string `json:"holidayId"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type HolidaysAPIResponse struct {
	Response struct {
		Holidays []*HolidaysAPI `json:"holidays"`
	} `json:"response"`
}

type HolidaysAPI struct {
	Name string `json:"name"`
	Date struct {
		ISO string `json:"iso"`
	} `json:"date"`
	Type []string `json:"type"`
}
