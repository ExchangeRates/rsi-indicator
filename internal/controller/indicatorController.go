package controller

import (
	"net/http"
	"rsi_indicator/internal/service"
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
	}
	type response struct {
	}

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
