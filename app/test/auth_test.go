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

func TestAuthRoute(t *testing.T) {
	URL := fmt.Sprintf("%s/%s", BASE_URL, "login")

	cases := []struct {
		CaseDescription         string
		ExpectedHttpRespStatus  string
		ExpectedHttpRespMessage string
		BodyRequest             any
	}{
		{
			CaseDescription:         "Login Berhasil",
			ExpectedHttpRespStatus:  "success",
			ExpectedHttpRespMessage: "login success",
			BodyRequest: &model.LoginPayload{
				Username: "admin",
				Password: "admin",
			},
		},
		{
			CaseDescription:         "Login Gagal",
			ExpectedHttpRespStatus:  "failed",
			ExpectedHttpRespMessage: "invalid body request",
			BodyRequest: &struct {
				Username string
				Password int
			}{
				Username: "asd",
				Password: 123,
			},
		},
		{
			CaseDescription:         "Username atau Password Salah",
			ExpectedHttpRespStatus:  "failed",
			ExpectedHttpRespMessage: "invalid nip or password",
			BodyRequest: &model.LoginPayload{
				Username: "admin",
				Password: "admins",
			},
		},
	}

	for _, v := range cases {
		jsonByte, err := json.Marshal(v.BodyRequest)
		if err != nil {
			log.Println(err)
		}

		bodyReq := bytes.NewBuffer(jsonByte)
		resp, err := http.Post(URL, "application/json", bodyReq)
		if err != nil {
			log.Println(err)
		}

		defer resp.Body.Close()

		//Read the response body
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
