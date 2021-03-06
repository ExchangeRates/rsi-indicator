package service

import (
	"math"
	"rsi_indicator/internal/feign"
)

type IndicatorService interface {
	Calculate(value float64, prev, prevD, prevU *float64, period int) (float64, float64, float64, error)
}

type indicatorServiceImpl struct {
	emaClient feign.EmaFeignClient
}

func NewIndicatorService(emaClient feign.EmaFeignClient) IndicatorService {
	return &indicatorServiceImpl{
		emaClient: emaClient,
	}
}

func (i *indicatorServiceImpl) Calculate(value float64, prev, prevD, prevU *float64, period int) (float64, float64, float64, error) {
	if prev == nil || prevD == nil || prevU == nil {
		return 0, 0, 0, nil
	}

	U := i.pointForMa(value, *prev)
	D := i.pointForMa(*prev, value)

	maOfU, err := i.emaClient.Calculate(*prevU, U, period)
	if err != nil {
		return 0, 0, 0, err
	}
	maOfD, err := i.emaClient.Calculate(*prevD, D, period)
	if err != nil {
		return 0, 0, 0, nil
	}

	if maOfD == 0 {
		return 100, U, D, nil
	}

	rs := maOfU / maOfD
	ratio := float64(100) / (1 + rs)

	return 100 - ratio, maOfU, maOfD, nil
}

func (i *indicatorServiceImpl) pointForMa(value, prev float64) float64 {
	if value > prev {
		return math.Abs(value - prev)
	}
	return 0
}
