package model

type Absent struct {
	AbsentId     string
	Nip          string
	Name         string
	AbsentDate   string
	StatusAbsent string
	Keterangan   string
}

type AbsentPayload struct {
	Date         string   `json:"date"`
	DateRange    []string `json:"dateRange"`
	Nip          string   `json:"nip"`
	StatusAbsent string   `json:"statusAbsent"`
	Keterangan   string   `json:"keterangan"`
}

type AbsentAggregate struct {
	Nip   string
	Name  string
	Alpa  int
	Sakit int
	Izin  int
}
