package models

type PricingResponse struct {
	ID      int    `json:"id"`
	JobType string `json:"job_type"`
	Price   int64  `json:"price"`
}
