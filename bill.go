package billplz

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	billAPIURL = "/api/v3/bills"
)

func (bp *Billplz) CreateBill(params *CreateBillParams) (*CreateBillResponse, error) {
	// convert req to json
	b, err := json.Marshal(params)
	if err != nil {
		return nil, errors.New("failed to marshal request during CreateBill")
	}

	url := bp.config.apiURL + billAPIURL
	// create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.New("failed to create request during CreateBill")
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+basicAuth(bp.config.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("failed to do request during CreateBill")
	}

	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read response during CreateBill")
	}

	fmt.Println(string(body))

	res := new(CreateBillResponse)
	if err := json.Unmarshal(body, res); err != nil {
		return nil, errors.New("failed to unmarshal response during CreateBill")
	}

	return res, nil
}
