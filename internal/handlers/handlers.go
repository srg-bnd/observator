// Handlers for server
package handlers

import (
	"net/http"

	"github.com/srg-bnd/observator/internal/services"
	"github.com/srg-bnd/observator/internal/storage"
)

type HttpHandler struct {
	service *services.Service
}

func NewHttpHandler(storage storage.Repositories) *HttpHandler {
	return &HttpHandler{
		service: services.NewService(storage),
	}
}

func (httpHandler *HttpHandler) UpdateMetricHandler(res http.ResponseWriter, req *http.Request) {
	metricType := req.PathValue("metricType")
	metricName := req.PathValue("metricName")
	metricValue := req.PathValue("metricValue")

	data := httpHandler.service.UpdateMetricService(metricType, metricName, metricValue)

	if data.Ok {
		res.Header().Set("content-type", "text/plain; charset=utf-8")
		res.WriteHeader(http.StatusOK)
	} else {
		switch data.Status {
		case "typeError", "valueError":
			res.WriteHeader(http.StatusBadRequest)
		case "nameError":
			res.WriteHeader(http.StatusNotFound)
		}
	}
}
