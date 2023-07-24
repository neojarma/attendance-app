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

func TestWorkTimeRoute(t *testing.T) {
	URL := fmt.Sprintf("%s/%s", BASE_URL, "work-time")

	cases := []struct {
		CaseDescription         string
		ExpectedHttpRespStatus  string
		ExpectedHttpRespMessage string
		HttpMethod              string
		BodyRequest             any
	}{
		{
			CaseDescription:         "Lihat Data Waktu kerja",
			ExpectedHttpRespStatus:  "success",
			HttpMethod:              http.MethodGet,
			ExpectedHttpRespMessage: "success get record",
			BodyRequest:             "none",
		},
		{
			CaseDescription:         "Invalid Body Request",
			ExpectedHttpRespStatus:  "failed",
			ExpectedHttpRespMessage: "invalid date or time",
			HttpMethod:              http.MethodPost,
			BodyRequest: &model.WorkTimes{
				JamMasuk:  "08ss:00:00",
				JamKeluar: "17:00:00",
			},
		},
		{
			CaseDescription:         "Success Tambah Waktu Kerja",
			ExpectedHttpRespStatus:  "success",
			ExpectedHttpRespMessage: "success create record",
			HttpMethod:              http.MethodPost,
			BodyRequest: &model.WorkTimes{
				JamMasuk:  "08:00:00",
				JamKeluar: "17:00:00",
			},
		},
		{
			CaseDescription:         "Success Update",
			ExpectedHttpRespStatus:  "success",
			ExpectedHttpRespMessage: "success update record",
			HttpMethod:              http.MethodPut,
			BodyRequest: &model.WorkTimes{
				WorkTimeId: "GVroo",
				JamMasuk:   "08:00:00",
				JamKeluar:  "17:30:00",
			},
		},
		{
			CaseDescription:         "Failed Update",
			ExpectedHttpRespStatus:  "failed",
			ExpectedHttpRespMessage: "there is no record with that id",
			HttpMethod:              http.MethodPut,
			BodyRequest: &model.WorkTimes{
				WorkTimeId: "GVroosss",
				JamMasuk:   "08:00:00",
				JamKeluar:  "17:30:00",
			},
		},
	}

	for _, v := range cases {
		jsonByte, err := json.Marshal(v.BodyRequest)
		if err != nil {
			log.Println(err)
		}

		bodyReq := bytes.NewBuffer(jsonByte)

		if v.HttpMethod == http.MethodGet {
			resp, err := http.Get(URL + "s")
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

		req, err := http.NewRequest(v.HttpMethod, URL, bodyReq)
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
