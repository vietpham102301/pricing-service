package pricing

import (
	"booking-service/handler/models"
	models2 "booking-service/internal/models"
	"booking-service/internal/repos"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPricing_GetPricing(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	defer time.Sleep(300 * time.Millisecond)
	mockIrepo := repos.NewMockIRepo(mockCtrl)
	mockPricingRepo := repos.NewMockIPricingRepo(mockCtrl)
	mockIrepo.EXPECT().Pricing().Return(mockPricingRepo).AnyTimes()

	mockPricingRepo.EXPECT().GetPrice(gomock.Any()).
		DoAndReturn(func(jobType int) (*models2.Price, error) {
			require.Equal(t, 2, jobType)

			return &models2.Price{
				ID:           2,
				JobType:      "take care baby",
				NormalPrice:  150000,
				WeekendPrice: 200000,
				HolidayPrice: 300000,
			}, nil
		}).AnyTimes()
	p := &Pricing{
		mgRepo: mockIrepo,
	}
	data := map[string]int{
		"jobType": 2,
		"day":     30,
		"month":   4,
		"year":    2023,
	}
	res, err := p.GetPricing(data)
	expectedRes := &models.PricingResponse{
		ID:      2,
		JobType: "take care baby",
		Price:   300000,
	}
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedRes, res)
}
