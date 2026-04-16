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
	SubscribeCreateUserIDErr  error = errors.New("Invalid validation UserID: UUID invalid")
	SubscribeCreateStartAtErr error = errors.New("Invalid validation StartAt: invalid date. Excepted format \"MM-YYYY\"")
)

func SubscribeCreateValidation(req *SubscribeCreateRequest) (*subscriptions.CreateSubscriptionDTO, error) {
	subDTO := subscriptions.CreateSubscriptionDTO{}

	// Name
	subDTO.Name = req.Name

	// Price
	subDTO.Price = req.Price

	// UserID
	if err := uuid.Validate(req.UserID); err != nil {
		return nil, SubscribeCreateUserIDErr
	}
	subDTO.UserID = req.UserID

	// StartAt
	startAtSplit := strings.Split(req.StartAt, "-")
	if len(startAtSplit) != 2 {
		return nil, SubscribeCreateStartAtErr
	}

	startAt, err := time.Parse("2006-01", fmt.Sprintf("%s-%s", startAtSplit[1], startAtSplit[0]))
	if err != nil {
		return nil, SubscribeCreateStartAtErr
	}

	subDTO.StartAt = startAt

	// EndAt
	subDTO.EndAt = startAt.AddDate(0, 1, 0)

	return &subDTO, nil
}

var (
	SubscribeListLimitNotNumberErr error = errors.New("Invalid validation Limit: Not number")
)

func SubscribeListValidation(page, limit string) (*SubscribeListRequest, error) {
	// page - не обязательный параметр, по умолчанию 1

	req := SubscribeListRequest{Page: 1}

	reqPage, err := strconv.Atoi(page)
	if err == nil && reqPage >= 1 {
		req.Page = reqPage
	}

	reqLimit, err := strconv.Atoi(limit)
	if err != nil {
		return nil, SubscribeListLimitNotNumberErr
	}
	req.Limit = reqLimit

	return &req, nil
}
