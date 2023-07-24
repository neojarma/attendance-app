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

func TestPositionRoute(t *testing.T) {
	URL := fmt.Sprintf("%s/%s", BASE_URL, "position")

	cases := []struct {
		CaseDescription         string
		ExpectedHttpRespStatus  string
		ExpectedHttpRespMessage string
		HttpMethod              string
		BodyRequest             any
	}{
		{
			CaseDescription:         "Lihat Data Posisi Jabatan",
			ExpectedHttpRespStatus:  "success",
			HttpMethod:              http.MethodGet,
			ExpectedHttpRespMessage: "success get record",
			BodyRequest:             "none",
		},
		{
			CaseDescription:         "Invalid Body Request",
			ExpectedHttpRespStatus:  "failed",
			ExpectedHttpRespMessage: "invalid body request",
			HttpMethod:              http.MethodPost,
			BodyRequest: struct {
				PositionName int
			}{
				PositionName: 123,
			},
		},
		{
			CaseDescription:         "Success Tambah Data Posisi",
			ExpectedHttpRespStatus:  "success",
			ExpectedHttpRespMessage: "success create record",
			HttpMethod:              http.MethodPost,
			BodyRequest: &model.Position{
				PositionName: "new",
			},
		},
		{
			CaseDescription:         "Success Update",
			ExpectedHttpRespStatus:  "success",
			ExpectedHttpRespMessage: "success update record",
			HttpMethod:              http.MethodPut,
			BodyRequest: &model.Position{
				PositionID:   "bpTJq",
				PositionName: "something new",
			},
		},
		{
			CaseDescription:         "Failed Update",
			ExpectedHttpRespStatus:  "failed",
			ExpectedHttpRespMessage: "there is no record with that id",
			HttpMethod:              http.MethodPut,
			BodyRequest: &model.Position{
				PositionID:   "notfound",
				PositionName: "newnewnew",
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
