package api

import (
	"TestTask/internal/subscriptions"
	"TestTask/internal/validation"
	"errors"
	"strconv"
	"time"
)

var (
	SubscriptionCreateEndAtErr error = errors.New("Invalid validation: start_date is later than end_date")
)

func SubscriptionCreateValidation(req *SubscriptionCreateRequest) ([]*subscriptions.SubscriptionCreateDTO, error) {
	months := make([]time.Time, 0)

	name, err := validation.SubscriptionNameValidate(req.Name)
	if err != nil {
		return nil, err
	}

	price, err := validation.SubscriptionPriceValidate(req.Price)
	if err != nil {
		return nil, err
	}

	userID, err := validation.SubscriptionUserIdValidate(req.UserID)
	if err != nil {
		return nil, err
	}

	startAt, err := validation.SubscriptionTimeAtValidate(req.StartAt)
	if err != nil {
		return nil, err
	}

	if req.EndAt != "" {
		endAt, err := validation.SubscriptionTimeAtValidate(req.EndAt)
		if err != nil {
			return nil, err
		}
		if !endAt.After(startAt) {
			return nil, SubscriptionCreateEndAtErr
		}

		for curr := startAt; curr.Before(endAt); curr = curr.AddDate(0, 1, 0) {
			months = append(months, curr)
		}
	} else {
		months = append(months, startAt)
	}

	dtoList := make([]*subscriptions.SubscriptionCreateDTO, 0)
	for _, date := range months {
		dtoList = append(dtoList, &subscriptions.SubscriptionCreateDTO{
			Name:    name,
			Price:   price,
			UserID:  userID,
			StartAt: date,
			EndAt:   date.AddDate(0, 1, 0),
		})
	}

	return dtoList, nil
}

var (
	SubscriptionListLimitNotNumberErr error = errors.New("Invalid validation Limit: Not number")
)

func SubscriptionListValidation(page, limit string) (*SubscriptionListRequest, error) {
	// page - не обязательный параметр, по умолчанию 1
	req := SubscriptionListRequest{Page: 1}

	reqPage, err := strconv.Atoi(page)
	if err == nil && reqPage >= 1 {
		req.Page = reqPage
	}

	reqLimit, err := strconv.Atoi(limit)
	if err != nil {
		return nil, SubscriptionListLimitNotNumberErr
	}
	req.Limit = reqLimit

	return &req, nil
}

var (
	SubscriptionUpdateEmptyErr error = errors.New("Invalid validation Limit: Not number")
)

func SubscriptionUpdateValidation(id int, req *SubscriptionUpdateRequest) (*subscriptions.SubscriptionUpdateDTO, error) {
	dto := &subscriptions.SubscriptionUpdateDTO{ID: id}
	var err error
	is_empty := true

	if req.Name != "" {
		dto.Name, err = validation.SubscriptionNameValidate(req.Name)
		if err != nil {
			return nil, err
		}
		is_empty = false
	}

	if req.Price != 0 {
		dto.Price, err = validation.SubscriptionPriceValidate(req.Price)
		if err != nil {
			return nil, err
		}
		is_empty = false
	}

	if req.UserID != "" {
		dto.UserID, err = validation.SubscriptionUserIdValidate(req.UserID)
		if err != nil {
			return nil, err
		}
		is_empty = false
	}

	if is_empty {
		return nil, SubscriptionUpdateEmptyErr
	}

	return dto, nil
}
