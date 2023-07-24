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

func TestHolidayRoute(t *testing.T) {
	URL := fmt.Sprintf("%s/%s", BASE_URL, "holiday")

	cases := []struct {
		CaseDescription         string
		ExpectedHttpRespStatus  string
		ExpectedHttpRespMessage string
		HttpMethod              string
		BodyRequest             any
	}{
		{
			CaseDescription:         "Lihat Data Hari Libur",
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
			BodyRequest: &model.Holidays{
				Date:        "akjshasdj",
				Description: "unknown",
			},
		},
		{
			CaseDescription:         "Success Tambah Hari Libur",
			ExpectedHttpRespStatus:  "success",
			ExpectedHttpRespMessage: "success create record",
			HttpMethod:              http.MethodPost,
			BodyRequest: &model.Holidays{
				Date:        "2028-08-08",
				Description: "new",
			},
		},
		{
			CaseDescription:         "Success Update",
			ExpectedHttpRespStatus:  "success",
			ExpectedHttpRespMessage: "success update record",
			HttpMethod:              http.MethodPut,
			BodyRequest: &model.Holidays{
				HolidayID:   "FBOTc",
				Date:        "2028-08-08",
				Description: "newzz",
			},
		},
		{
			CaseDescription:         "Failed Update",
			ExpectedHttpRespStatus:  "failed",
			ExpectedHttpRespMessage: "there is no record with that id",
			HttpMethod:              http.MethodPut,
			BodyRequest: &model.Holidays{
				HolidayID:   "FBOTcss",
				Date:        "2028-08-08",
				Description: "newzz",
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
