package internal

import (
	"net/http"
	"rsi_indicator/internal/api"
	"rsi_indicator/internal/config"
	"rsi_indicator/internal/controller"
	"rsi_indicator/internal/feign"
	"rsi_indicator/internal/service"
)

func Start(config *config.Config) error {

	emaClient := feign.NewEmaFeignClient(config.EmaClientURL)
	indicatorService := service.NewIndicatorService(emaClient)
	indicatorController := controller.NewIndicatorController(indicatorService)
	srv := api.NewServer(indicatorController)
	bindingAddress := srv.BindingAddressFromPort(config.Port)

	return http.ListenAndServe(bindingAddress, srv)
}
