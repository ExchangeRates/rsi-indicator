package controller

import (
	"encoding/json"
	"net/http"
	"rsi_indicator/internal/service"

	"github.com/sirupsen/logrus"
)

type IndicatorController struct {
	service service.IndicatorService
}

func NewIndicatorController(service service.IndicatorService) *IndicatorController {
	return &IndicatorController{
		service: service,
	}
}

func (c *IndicatorController) HandleCalculate() http.HandlerFunc {
	type request struct {
		Value  float64 `json:"value"`
		Prev   *float64 `json:"prev"`
		PrevD  *float64 `json:"prevD"`
		PrevU  *float64 `json:"prevU"`
		Period int     `json:"period"`
	}
	type response struct {
		Value float64 `json:"value"`
		MaOfU float64 `json:"maOfU"`
		MaOfD float64 `json:"maOfD"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		body := &request{}
		if err := json.NewDecoder(r.Body).Decode(body); err != nil {
			// TODO send error
			logrus.Warnln("decode ", err)
			return
		}

		value, maOfU, maOfD, err := c.service.Calculate(
			body.Value,
			body.Prev,
			body.PrevD,
			body.PrevU,
			body.Period,
		)
		if err != nil {
			// TODO send error
			logrus.Warnln("calculate ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response{
			Value: value,
			MaOfD: maOfD,
			MaOfU: maOfU,
		}); err != nil {
			// TODO send response
			logrus.Warnln("encode ", err)
			return
		}
	}
}
