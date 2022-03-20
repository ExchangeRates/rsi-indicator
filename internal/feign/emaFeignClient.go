package feign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type EmaFeignClient interface {
	Calculate(prev, value float64, period int) (float64, error)
}

type emaFeignClientImpl struct {
	url string
}

func NewEmaFeignClient(url string) EmaFeignClient {
	return &emaFeignClientImpl{
		url: url,
	}
}

type calculateRequest struct {
	Prev   float64 `json:"prev"`
	Value  float64 `json:"value"`
	Period int     `json:"period"`
}

type calculateResponse struct {
	Value float64 `json:"value"`
}

func (c *emaFeignClientImpl) Calculate(prev, value float64, period int) (float64, error) {

	url := fmt.Sprintf("%s/calculate", c.url)
	var payload bytes.Buffer
	body := calculateRequest{
		Prev:   prev,
		Value:  value,
		Period: period,
	}
	if err := json.NewEncoder(&payload).Encode(body); err != nil {
		return 0, err
	}
	resp, err := http.Post(url, "application/json", &payload)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	response := &calculateResponse{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return 0, err
	}

	return response.Value, nil
}
