package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"rsi_indicator/internal/controller"
)

type server struct {
	router              *mux.Router
	logger              *logrus.Logger
	indicatorController *controller.IndicatorController
}

func NewServer(indicatorController *controller.IndicatorController) *server {
	s := &server{
		router:              mux.NewRouter(),
		logger:              logrus.New(),
		indicatorController: indicatorController,
	}

	s.configureRouter()

	logrus.Info("starting api server")

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {

}

func (s *server) BindingAddressFromPort(port int) string {
	return fmt.Sprintf(":%d", port)
}
