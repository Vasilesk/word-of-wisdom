package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

func DoJSON[T any](d Doer, req *http.Request) (T, error) {
	var val T

	resp, err := d.Do(req)
	if err != nil {
		return val, fmt.Errorf("making request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return val, fmt.Errorf("reading body: %w", err)
	}

	if err := json.Unmarshal(body, &val); err != nil {
		return val, fmt.Errorf("unmarshalling json: %w", err)
	}

	return val, nil
}
