package model

type ReportAggregate struct {
	Nip         string `json:"NIP"`
	Name        string `json:"fullName"`
	OnTime      int    `json:"onTime"`
	Late        int    `json:"late"`
	PulangTepat int    `json:"pulangTepat"`
	PulangCepat int    `json:"pulangCepat"`
	Alpa        int    `json:"alpa"`
	Sakit       int    `json:"sakit"`
	Izin        int    `json:"izin"`
}
