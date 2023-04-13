package repos

import "booking-service/internal/models"

//go:generate mockgen -package=repos -destination=irepo_mock.go -source=irepo.go
type IRepo interface {
	Pricing() IPricingRepo
}

type IPricingRepo interface {
	GetPrice(jobType int) (*models.Price, error)
}
