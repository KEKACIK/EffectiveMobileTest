package validation

import (
	"TestTask/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSubscriptionIdValidate(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		result int
		err    error
	}{
		{
			name:   "Success 1",
			value:  "1",
			result: 1,
			err:    nil,
		},
		{
			name:   "Success 2",
			value:  "1025",
			result: 1025,
			err:    nil,
		},
		{
			name:   "Success 2",
			value:  "     1986325             ",
			result: 1986325,
			err:    nil,
		},
		{
			name:   "Error. Not number 2",
			value:  "",
			result: 0,
			err:    SubscriptionIDNotNumberErr,
		},
		{
			name:   "Error. Not number 2",
			value:  "test",
			result: 0,
			err:    SubscriptionIDNotNumberErr,
		},
		{
			name:   "Error. Negative number",
			value:  "-105",
			result: 0,
			err:    SubscriptionIDNegativeNumberErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SubscriptionIdValidate(tt.value)
			assert.Equal(t, tt.result, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestSubscriptionNameValidate(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		result string
		err    error
	}{
		{
			name:   "Success 1",
			value:  "Yandex Plus",
			result: "Yandex Plus",
			err:    nil,
		},
		{
			name:   "Success 2",
			value:  "Netflix",
			result: "Netflix",
			err:    nil,
		},
		{
			name:   "Success 3",
			value:  "     Amazon   ",
			result: "Amazon",
			err:    nil,
		},
		{
			name:   "Error. Empty",
			value:  "",
			result: "",
			err:    SubscriptionNameEmptyErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SubscriptionNameValidate(tt.value)
			assert.Equal(t, tt.result, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestSubscriptionPriceValidate(t *testing.T) {
	tests := []struct {
		name   string
		value  int
		result int
		err    error
	}{
		{
			name:   "Success 1",
			value:  1,
			result: 1,
			err:    nil,
		},
		{
			name:   "Success 2",
			value:  500,
			result: 500,
			err:    nil,
		},
		{
			name:   "Success 3",
			value:  5000,
			result: 5000,
			err:    nil,
		},
		{
			name:   "Error. Negative",
			value:  -100,
			result: 0,
			err:    SubscriptionPriceNegativeErr,
		},
		{
			name:   "Error. Zero",
			value:  0,
			result: 0,
			err:    SubscriptionPriceZeroErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SubscriptionPriceValidate(tt.value)
			assert.Equal(t, tt.result, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestSubscriptionUserIdValidate(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		result string
		err    error
	}{
		{
			name:   "Success 1",
			value:  "7a0fcd16-ad86-4a9e-93b0-48ea8222a341",
			result: "7a0fcd16-ad86-4a9e-93b0-48ea8222a341",
			err:    nil,
		},
		{
			name:   "Success 2",
			value:  "facc73c1-e0c2-487c-9dc8-3d09a59a9882",
			result: "facc73c1-e0c2-487c-9dc8-3d09a59a9882",
			err:    nil,
		},
		{
			name:   "Success 3",
			value:  "         9654cfa5-abfd-4e58-b5ec-712320d6142b      ",
			result: "9654cfa5-abfd-4e58-b5ec-712320d6142b",
			err:    nil,
		},
		{
			name:   "Error. Empty",
			value:  "",
			result: "",
			err:    SubscriptionUserIdEmptyErr,
		},
		{
			name:   "Error. UUID invalid",
			value:  "12345",
			result: "",
			err:    SubscriptionUserIdUUIDErr,
		},
		{
			name:   "Error. UUID invalid",
			value:  "7a0fcd16-ad86-4a9e-48ea8222a341",
			result: "",
			err:    SubscriptionUserIdUUIDErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SubscriptionUserIdValidate(tt.value)
			assert.Equal(t, tt.result, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestSubscriptionDateAtValidate(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		result time.Time
		err    error
	}{
		{
			name:   "Success 1",
			value:  "01-2026",
			result: utils.GetEmptyTime().AddDate(2026, 1, 0),
			err:    nil,
		},
		{
			name:   "Success 2",
			value:  "05-2021",
			result: utils.GetEmptyTime().AddDate(2021, 5, 0),
			err:    nil,
		},
		{
			name:   "Success 3",
			value:  "      08-2023          ",
			result: utils.GetEmptyTime().AddDate(2023, 8, 0),
			err:    nil,
		},
		{
			name:   "Error. Invalid date 1",
			value:  "15-2026",
			result: time.Time{},
			err:    SubscriptionDateAtErr,
		},
		{
			name:   "Error. Invalid date 2",
			value:  "01.2026",
			result: time.Time{},
			err:    SubscriptionDateAtErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SubscriptionDateAtValidate(tt.value)
			assert.Equal(t, tt.result, result)
			assert.Equal(t, tt.err, err)
		})
	}
}
