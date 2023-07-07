package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func GetReport(jsonData interface{}, fileName string) (string, error) {
	stringJSON, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}

	file, err := writeTempFile(stringJSON, fileName)
	if err != nil {
		return "", err
	}

	defer func() {
		file.Close()
		os.Remove(fileName)
	}()

	body, writer, err := prepareBodyRequest(stringJSON, fileName, file)
	if err != nil {
		return "", err
	}

	client, request, err := prepareHTTPRequest(body, writer)
	if err != nil {
		return "", err
	}

	// do http request
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	parsedResponse, err := parsingBodyResponse(response.Body)
	if err != nil {
		log.Println("err6")
		return "", err
	}

	return getDownloadFilePath(parsedResponse["FolderName"].(string), fileName), nil
}

func parsingBodyResponse(body io.ReadCloser) (map[string]any, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var mapJson map[string]interface{}
	err = json.Unmarshal(data, &mapJson)
	if err != nil {
		return nil, err
	}

	return mapJson, nil
}

func prepareHTTPRequest(body *bytes.Buffer, writer *multipart.Writer) (*http.Client, *http.Request, error) {
	url := "https://api.products.aspose.app/cells/conversion/api/ConversionApi/Convert"

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		return nil, nil, err
	}

	prepareHeaderRequest(request, writer)

	return client, request, nil
}

func prepareHeaderRequest(request *http.Request, writer *multipart.Writer) {

	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("accept", "*/*")
	request.Header.Set("accept-language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7")
	request.Header.Set("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
	request.Header.Set("sec-ch-ua-mobile", "?1")
	request.Header.Set("sec-ch-ua-platform", `"Android"`)
	request.Header.Set("sec-fetch-dest", "empty")
	request.Header.Set("sec-fetch-mode", "cors")
	request.Header.Set("sec-fetch-site", "same-site")

}

func prepareBodyRequest(stringJSON []byte, fileName string, file *os.File) (*bytes.Buffer, *multipart.Writer, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	err := writer.WriteField("inputTxt", string(stringJSON))
	if err != nil {
		return nil, writer, err
	}

	part, err := writer.CreateFormFile("1708311304", fileName)
	if err != nil {
		return nil, writer, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, writer, err
	}

	err = writer.WriteField("MultipleWorksheets", "true")
	if err != nil {
		return nil, writer, err
	}

	err = writer.WriteField("UploadOptions", "JSON")
	if err != nil {
		return nil, writer, err
	}

	err = writer.WriteField("outputType", "XLSX")
	if err != nil {
		return nil, writer, err
	}

	err = writer.Close()
	if err != nil {
		return nil, writer, err
	}

	return body, writer, err
}

func writeTempFile(stringJSON []byte, fileName string) (*os.File, error) {
	err := os.WriteFile(fileName, stringJSON, 0644)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func getDownloadFilePath(folder string, file string) string {
	return fmt.Sprintf("https://api.products.aspose.app/cells/conversion/api/Download/%s?file=%s", folder, file)
}
