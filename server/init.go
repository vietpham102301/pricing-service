package server

import (
	handler "booking-service/handler"
	"booking-service/internal/repos"
	"booking-service/internal/services/pricing"
)

func (s *Server) initServices(repo repos.IRepo) *ServiceList {
	pricing := pricing.NewPricing(repo)
	return &ServiceList{
		pricing: pricing,
	}
}

func (s *Server) initRouters(serviceList *ServiceList) {
	handler := handler.NewHandler(serviceList.pricing)

	handler.ConfigAPIRoute(s.router)
}
