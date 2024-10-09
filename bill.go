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
	billAPIURL = "/v3/bills"
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

// GetBill returns a bill by ID
func (bp *Billplz) GetBill(id string) (*GetBillResponse, error) {
	url := bp.config.apiURL + billAPIURL + "/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("failed to create request during GetBill")
	}

	req.Header.Set("Authorization", "Basic "+basicAuth(bp.config.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("failed to do request during GetBill")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read response during GetBill")
	}

	res := new(GetBillResponse)
	if err := json.Unmarshal(body, res); err != nil {
		return nil, errors.New("failed to unmarshal response during GetBill")
	}

	return res, nil
}

// DeleteBill deletes a bill by ID
func (bp *Billplz) DeleteBill(id string) error {
	url := bp.config.apiURL + billAPIURL + "/" + id
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return errors.New("failed to create request during DeleteBill")
	}

	req.Header.Set("Authorization", "Basic "+basicAuth(bp.config.apiKey))

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return errors.New("failed to do request during DeleteBill")
	}

	return nil
}
