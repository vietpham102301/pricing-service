package pricing

import (
	models2 "booking-service/handler/models"
	"booking-service/internal/repos"
	"log"
	"strconv"
	"time"
)

type Holiday struct {
	date string
	name string
}

const (
	Saturday = 6
	Sunday   = 7
)

type IPricing interface {
	GetPricing(data map[string]int) (*models2.PricingResponse, error)
}

type Pricing struct {
	mgRepo repos.IRepo
}

func NewPricing(repo repos.IRepo) IPricing {
	return &Pricing{
		mgRepo: repo,
	}
}

func (p Pricing) GetPricing(data map[string]int) (*models2.PricingResponse, error) {
	jobType := data["jobType"]
	day := data["day"]
	month := data["month"]
	year := data["year"]
	record, err := p.mgRepo.Pricing().GetPrice(jobType)
	if err != nil {
		log.Printf("getting pricing fail with err %v", err)
		return nil, err
	}
	if record == nil {
		return nil, nil
	}
	res := &models2.PricingResponse{
		ID:      record.ID,
		JobType: record.JobType,
	}
	mockData := []Holiday{
		{date: "2/9", name: "Quoc khanh"},
		{date: "30/4", name: "Ngay giai phong"},
		{date: "1/5", name: "Quoc te lao dong"},
	}

	date := strconv.Itoa(day) + "/" + strconv.Itoa(month)
	for i := 0; i < len(mockData); i++ {
		if date == mockData[i].date {
			res.Price = record.HolidayPrice
			return res, nil
		}
	}

	currentDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	weekday := currentDate.Weekday()

	if weekday == time.Saturday || weekday == time.Sunday {
		res.Price = record.WeekendPrice
		return res, nil
	}

	res.Price = record.NormalPrice

	return res, nil
}
