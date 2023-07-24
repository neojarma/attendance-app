package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"presensi/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPresenceRoute(t *testing.T) {
	cases := []struct {
		CaseDescription         string
		ExpectedHttpRespStatus  string
		ExpectedHttpRespMessage string
		HttpMethod              string
		URL                     string
		BodyRequest             any
	}{
		{
			CaseDescription:         "Lihat Data Kehadiran",
			ExpectedHttpRespStatus:  "success",
			HttpMethod:              http.MethodGet,
			ExpectedHttpRespMessage: "success get record",
			URL:                     fmt.Sprintf("%s/%s", BASE_URL, "daily-presence?from=2023-06-12&to=2023-06-13"),
			BodyRequest:             "none",
		},
		{
			CaseDescription:         "Invalid Body Request",
			ExpectedHttpRespStatus:  "failed",
			ExpectedHttpRespMessage: "invalid date format",
			HttpMethod:              http.MethodPost,
			URL:                     fmt.Sprintf("%s/%s", BASE_URL, "presence"),
			BodyRequest: &model.PresencePayload{
				Nip:          "1238123687",
				WaktuMasuk:   "09:10:00",
				PresenceDate: "2023s-06-30",
				WorkTimeId:   "thqew",
			},
		},
		{
			CaseDescription:         "Success Tambah Data Kehadiran",
			ExpectedHttpRespStatus:  "success",
			ExpectedHttpRespMessage: "success create record",
			HttpMethod:              http.MethodPost,
			URL:                     fmt.Sprintf("%s/%s", BASE_URL, "presence"),
			BodyRequest: &model.PresencePayload{
				Nip:          "1238123687",
				WaktuMasuk:   "09:10:00",
				PresenceDate: "2023-06-30",
				WorkTimeId:   "thqew",
			},
		},
		{
			CaseDescription:         "Success Membuat Report",
			ExpectedHttpRespStatus:  "success",
			ExpectedHttpRespMessage: "success get record",
			HttpMethod:              http.MethodGet,
			URL:                     fmt.Sprintf("%s/%s", BASE_URL, "daily-presence/report?date=2023-06-12"),
			BodyRequest:             "none",
		},
	}

	for _, v := range cases {
		jsonByte, err := json.Marshal(v.BodyRequest)
		if err != nil {
			log.Println(err)
		}

		bodyReq := bytes.NewBuffer(jsonByte)

		if v.HttpMethod == http.MethodGet {
			resp, err := http.Get(v.URL)
			if err != nil {
				log.Println(err)
			}

			defer resp.Body.Close()

			bodyResp, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}

			result := new(model.Response)
			err = json.Unmarshal(bodyResp, result)
			if err != nil {
				log.Println(err)
			}

			t.Run(v.CaseDescription, func(t *testing.T) {
				assert.Equal(t, v.ExpectedHttpRespStatus, result.Status)
				assert.Equal(t, v.ExpectedHttpRespMessage, result.Message)
			})

			continue
		}

		req, err := http.NewRequest(v.HttpMethod, v.URL, bodyReq)
		if err != nil {
			log.Println(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}

		defer resp.Body.Close()

		bodyResp, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		result := new(model.Response)
		err = json.Unmarshal(bodyResp, result)
		if err != nil {
			log.Println(err)
		}

		t.Run(v.CaseDescription, func(t *testing.T) {
			assert.Equal(t, v.ExpectedHttpRespStatus, result.Status)
			assert.Equal(t, v.ExpectedHttpRespMessage, result.Message)
		})
	}
}
