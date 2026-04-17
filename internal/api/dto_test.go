package api

import (
	"TestTask/internal/subscriptions"
	"TestTask/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSubscriptionCreateDTO(t *testing.T) {
	tests := []struct {
		name   string
		value  *SubscriptionCreateRequest
		result []*subscriptions.SubscriptionCreateDTO
		err    error
	}{
		{
			name: "Success 1",
			value: &SubscriptionCreateRequest{
				Name:    "Yandex Plus",
				Price:   399,
				UserID:  "9654cfa5-abfd-4e58-b5ec-712320d6142b",
				StartAt: "01-2026",
			},
			result: []*subscriptions.SubscriptionCreateDTO{
				{
					Name:    "Yandex Plus",
					Price:   399,
					UserID:  "9654cfa5-abfd-4e58-b5ec-712320d6142b",
					StartAt: utils.GetEmptyTime().AddDate(2026, 1, 0),
					EndAt:   utils.GetEmptyTime().AddDate(2026, 2, 0),
				},
			},
			err: nil,
		},
		{
			name: "Success 2",
			value: &SubscriptionCreateRequest{
				Name:    "Yandex Plus",
				Price:   399,
				UserID:  "9654cfa5-abfd-4e58-b5ec-712320d6142b",
				StartAt: "01-2026",
				EndAt:   "02-2026",
			},
			result: []*subscriptions.SubscriptionCreateDTO{
				{
					Name:    "Yandex Plus",
					Price:   399,
					UserID:  "9654cfa5-abfd-4e58-b5ec-712320d6142b",
					StartAt: utils.GetEmptyTime().AddDate(2026, 1, 0),
					EndAt:   utils.GetEmptyTime().AddDate(2026, 2, 0),
				},
			},
			err: nil,
		},
		{
			name: "Success 3",
			value: &SubscriptionCreateRequest{
				Name:    "Yandex Plus",
				Price:   399,
				UserID:  "9654cfa5-abfd-4e58-b5ec-712320d6142b",
				StartAt: "01-2026",
				EndAt:   "03-2026",
			},
			result: []*subscriptions.SubscriptionCreateDTO{
				{
					Name:    "Yandex Plus",
					Price:   399,
					UserID:  "9654cfa5-abfd-4e58-b5ec-712320d6142b",
					StartAt: utils.GetEmptyTime().AddDate(2026, 1, 0),
					EndAt:   utils.GetEmptyTime().AddDate(2026, 2, 0),
				},
				{
					Name:    "Yandex Plus",
					Price:   399,
					UserID:  "9654cfa5-abfd-4e58-b5ec-712320d6142b",
					StartAt: utils.GetEmptyTime().AddDate(2026, 2, 0),
					EndAt:   utils.GetEmptyTime().AddDate(2026, 3, 0),
				},
			},
			err: nil,
		},
		{
			name: "Error. ",
			value: &SubscriptionCreateRequest{
				Name:    "Yandex Plus",
				Price:   399,
				UserID:  "9654cfa5-abfd-4e58-b5ec-712320d6142b",
				StartAt: "01-2026",
				EndAt:   "01-2026",
			},
			result: nil,
			err:    SubscriptionIntervalErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetSubscriptionCreateDTO(tt.value)
			assert.Equal(t, tt.result, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestGetSubscriptionListDTO(t *testing.T) {
	tests := []struct {
		name   string
		page   string
		limit  string
		result *SubscriptionListRequest
		err    error
	}{
		{
			name:   "Success 1",
			page:   "1",
			limit:  "10",
			result: &SubscriptionListRequest{Page: 1, Limit: 10},
			err:    nil,
		},
		{
			name:   "Success 2",
			page:   "",
			limit:  "50",
			result: &SubscriptionListRequest{Page: 1, Limit: 50},
			err:    nil,
		},
		{
			name:   "Success 3",
			page:   "-100",
			limit:  "100",
			result: &SubscriptionListRequest{Page: 1, Limit: 100},
			err:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetSubscriptionListDTO(tt.page, tt.limit)
			assert.Equal(t, tt.result, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestGetSubscriptionStatsDTO(t *testing.T) {
	tests := []struct {
		Name      string
		name      string
		userID    string
		startDate string
		stopDate  string
		result    *subscriptions.SubscriptionStatDTO
		err       error
	}{
		{
			Name:      "Success 1",
			name:      "Yandex Plus",
			userID:    "9654cfa5-abfd-4e58-b5ec-712320d6142b",
			startDate: "01-2026",
			stopDate:  "02-2026",
			result: &subscriptions.SubscriptionStatDTO{
				Name:      "Yandex Plus",
				UserID:    "9654cfa5-abfd-4e58-b5ec-712320d6142b",
				StartDate: utils.GetEmptyTime().AddDate(2026, 1, 0),
				StopDate:  utils.GetEmptyTime().AddDate(2026, 2, 0),
			},
			err: nil,
		},
		{
			Name:      "Success 2",
			name:      "Yandex Plus",
			userID:    "9654cfa5-abfd-4e58-b5ec-712320d6142b",
			startDate: "01-2026",
			stopDate:  "05-2026",
			result: &subscriptions.SubscriptionStatDTO{
				Name:      "Yandex Plus",
				UserID:    "9654cfa5-abfd-4e58-b5ec-712320d6142b",
				StartDate: utils.GetEmptyTime().AddDate(2026, 1, 0),
				StopDate:  utils.GetEmptyTime().AddDate(2026, 5, 0),
			},
			err: nil,
		},
		{
			Name:      "Success 3. Name miss",
			name:      "",
			userID:    "9654cfa5-abfd-4e58-b5ec-712320d6142b",
			startDate: "01-2026",
			stopDate:  "08-2026",
			result: &subscriptions.SubscriptionStatDTO{
				Name:      "",
				UserID:    "9654cfa5-abfd-4e58-b5ec-712320d6142b",
				StartDate: utils.GetEmptyTime().AddDate(2026, 1, 0),
				StopDate:  utils.GetEmptyTime().AddDate(2026, 8, 0),
			},
			err: nil,
		},
		{
			Name:      "Error. Interval invalid",
			name:      "",
			userID:    "9654cfa5-abfd-4e58-b5ec-712320d6142b",
			startDate: "05-2026",
			stopDate:  "01-2026",
			result:    nil,
			err:       SubscriptionIntervalErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := GetSubscriptionStatsDTO(tt.name, tt.userID, tt.startDate, tt.stopDate)
			assert.Equal(t, tt.result, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestGetSubscriptionUpdateDTO(t *testing.T) {
	tests := []struct {
		name   string
		id     int
		req    *SubscriptionUpdateRequest
		result *subscriptions.SubscriptionUpdateDTO
		err    error
	}{
		{
			name: "Success 1",
			id:   1,
			req: &SubscriptionUpdateRequest{
				Name: "Yandex Plus",
			},
			result: &subscriptions.SubscriptionUpdateDTO{
				ID:   1,
				Name: "Yandex Plus",
			},
			err: nil,
		},

		{
			name: "Success 2",
			id:   1,
			req: &SubscriptionUpdateRequest{
				Price: 399,
			},
			result: &subscriptions.SubscriptionUpdateDTO{
				ID:    1,
				Price: 399,
			},
			err: nil,
		},
		{
			name: "Success 3",
			id:   1,
			req: &SubscriptionUpdateRequest{
				UserID: "9654cfa5-abfd-4e58-b5ec-712320d6142b",
			},
			result: &subscriptions.SubscriptionUpdateDTO{
				ID:     1,
				UserID: "9654cfa5-abfd-4e58-b5ec-712320d6142b",
			},
			err: nil,
		},
		{
			name: "Success 4",
			id:   1,
			req: &SubscriptionUpdateRequest{
				Name:  "Yandex Plus",
				Price: 399,
			},
			result: &subscriptions.SubscriptionUpdateDTO{
				ID:    1,
				Name:  "Yandex Plus",
				Price: 399,
			},
			err: nil,
		},
		{
			name: "Success 5",
			id:   1,
			req: &SubscriptionUpdateRequest{
				Price:  399,
				UserID: "9654cfa5-abfd-4e58-b5ec-712320d6142b",
			},
			result: &subscriptions.SubscriptionUpdateDTO{
				ID:     1,
				Price:  399,
				UserID: "9654cfa5-abfd-4e58-b5ec-712320d6142b",
			},
			err: nil,
		},
		{
			name: "Success 6",
			id:   1,
			req: &SubscriptionUpdateRequest{
				Name:   "Yandex Plus",
				UserID: "9654cfa5-abfd-4e58-b5ec-712320d6142b",
			},
			result: &subscriptions.SubscriptionUpdateDTO{
				ID:     1,
				Name:   "Yandex Plus",
				UserID: "9654cfa5-abfd-4e58-b5ec-712320d6142b",
			},
			err: nil,
		},
		{
			name:   "Error. Empty",
			id:     1,
			req:    &SubscriptionUpdateRequest{},
			result: nil,
			err:    SubscriptionUpdateEmptyErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetSubscriptionUpdateDTO(tt.id, tt.req)
			assert.Equal(t, tt.result, result)
			assert.Equal(t, tt.err, err)
		})
	}
}
