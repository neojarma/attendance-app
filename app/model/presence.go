package model

type Presence struct {
	PresenceId   string `json:"presenceId"`
	WorkTimeId   string `json:"workTimeId"`
	Nip          string `json:"nip"`
	Name         string `json:"name"`
	WaktuMasuk   string `json:"waktuMasuk"`
	StatusMasuk  string `json:"statusMasuk"`
	WaktuKeluar  string `json:"waktuKeluar"`
	StatusKeluar string `json:"statusKeluar"`
	PresenceDate string `json:"presenceDate"`
}

type PresencePayload struct {
	PresenceId   string
	Nip          string `json:"nip"`
	WaktuMasuk   string `json:"waktuMasuk"`
	WorkTimeId   string `json:"workTimeId"`
	WaktuKeluar  string `json:"waktuKeluar"`
	PresenceDate string `json:"presenceDate"`
}

type PresenceAggregate struct {
	Nip         string
	Name        string
	OnTime      int
	Late        int
	PulangTepat int
	PulangCepat int
}
