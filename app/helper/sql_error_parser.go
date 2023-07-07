package helper

import (
	"errors"
	"fmt"
	"strings"
)

func SQLErrorParser(err string) error {

	var resultStr string
	if strings.Contains(err, "duplicate") {
		splitted := strings.Split(err, "is")[1]
		nip := strings.ReplaceAll(splitted, ".", "")
		resultStr = fmt.Sprintf("duplicate nip%s", nip)
	} else if strings.Contains(err, "conflicted") {
		splitted := strings.Split(err, "column")[1]
		removeDot := strings.ReplaceAll(splitted, ".", "")
		removeSingleQuote := strings.ReplaceAll(removeDot, "'", "")
		removeUnderscore := strings.Split(removeSingleQuote, "_")[0]
		resultStr = fmt.Sprintf("there is no%s with that id", removeUnderscore)
	} else if strings.Contains(err, "date") {
		resultStr = "invalid date or time"
	} else {
		resultStr = err
	}

	return errors.New(resultStr)
}
