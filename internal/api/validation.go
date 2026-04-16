package api

import (
	"TestTask/internal/subscriptions"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	SubscriptionIDNotNumberErr error = errors.New("Invalid validation ID: Not number")
)

func SubscriptionIdValidate(id string) (int, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, SubscriptionIDNotNumberErr
	}

	return idInt, nil
}

var (
	SubscriptionNameEmptyErr error = errors.New("Invalid validation Name: Empty")
)

func SubscriptionNameValidate(name string) (string, error) {
	if name == "" {
		return "", SubscriptionNameEmptyErr
	}

	return name, nil
}

var (
	SubscriptionPriceNegativeErr error = errors.New("Invalid validation Price: Negative number. Excepted Price > 0")
	SubscriptionPriceZeroErr     error = errors.New("Invalid validation Price: Zero number. Excepted Price > 0")
)

func SubscriptionPriceValidate(price int) (int, error) {
	if price < 0 {
		return 0, SubscriptionPriceNegativeErr
	}
	if price == 0 {
		return 0, SubscriptionPriceZeroErr
	}

	return price, nil
}

var (
	SubscriptionUserIdUUIDErr error = errors.New("Invalid validation UserID: UUID invalid")
)

func SubscriptionUserIdValidate(userID string) (string, error) {
	if err := uuid.Validate(userID); err != nil {
		return "", SubscriptionUserIdUUIDErr
	}

	return userID, nil
}

var (
	SubscriptionStartAtErr error = errors.New("Invalid validation StartAt: invalid date. Excepted format \"MM-YYYY\"")
)

func SubscriptionStartAtValidate(startAt string) (time.Time, error) {
	startAtSplit := strings.Split(startAt, "-")
	if len(startAtSplit) != 2 {
		return time.Time{}, SubscriptionStartAtErr
	}

	startAtTime, err := time.Parse("2006-01", fmt.Sprintf("%s-%s", startAtSplit[1], startAtSplit[0]))
	if err != nil {
		return time.Time{}, SubscriptionStartAtErr
	}

	return startAtTime, nil
}

func SubscriptionCreateValidation(req *SubscriptionCreateRequest) (*subscriptions.SubscriptionCreateDTO, error) {
	dto := &subscriptions.SubscriptionCreateDTO{}
	var err error = nil

	dto.Name, err = SubscriptionNameValidate(req.Name)
	if err != nil {
		return nil, err
	}

	dto.Price, err = SubscriptionPriceValidate(req.Price)
	if err != nil {
		return nil, err
	}

	dto.UserID, err = SubscriptionUserIdValidate(req.UserID)
	if err != nil {
		return nil, err
	}

	dto.StartAt, err = SubscriptionStartAtValidate(req.StartAt)
	if err != nil {
		return nil, err
	}
	dto.EndAt = dto.StartAt.AddDate(0, 1, 0) // Добавляем 1 месяц для endAt

	return dto, nil
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

func SubscriptionGetValidation(id string) (*SubscriptionGetRequest, error) {
	req := &SubscriptionGetRequest{}

	reqID, err := strconv.Atoi(id)
	if err != nil {
		return nil, SubscriptionIDNotNumberErr
	}
	req.ID = reqID

	return req, nil
}

func SubscriptionUpdateValidation(id int, req *SubscriptionUpdateRequest) (*subscriptions.SubscriptionUpdateDTO, error) {
	dto := &subscriptions.SubscriptionUpdateDTO{ID: id}
	var err error

	if req.Name != "" {
		dto.Name, err = SubscriptionNameValidate(req.Name)
		if err != nil {
			return nil, err
		}
	}

	if req.Price != 0 {
		dto.Price, err = SubscriptionPriceValidate(req.Price)
		if err != nil {
			return nil, err
		}
	}

	if req.UserID != "" {
		dto.UserID, err = SubscriptionUserIdValidate(req.UserID)
		if err != nil {
			return nil, err
		}
	}
	if req.StartAt != "" {
		dto.StartAt, err = SubscriptionStartAtValidate(req.StartAt)
		if err != nil {
			return nil, err
		}
		dto.EndAt = dto.StartAt.AddDate(0, 1, 0) // Добавляем 1 месяц для endAt
	}

	return dto, nil
}
