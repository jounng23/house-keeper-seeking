package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ParseResponse(resp *http.Response, target interface{}) error {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(bodyBytes) == 0 {
		return errors.New("response body is empty")
	}

	return json.Unmarshal(bodyBytes, target)
}
