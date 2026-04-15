package api

import (
	"TestTask/internal/subscriptions"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	CreateRequestUserIDErr  error = errors.New("Invalid validation UserID: UUID invalid")
	CreateRequestStartAtErr error = errors.New("Invalid validation StartAt: invalid date. Excepted format \"MM-YYYY\"")
)

func CreateRequestValidation(createReq *CreateRequest) (*subscriptions.CreateSubscriptionDTO, error) {
	subDTO := subscriptions.CreateSubscriptionDTO{}

	// Name
	subDTO.Name = createReq.Name

	// Price
	subDTO.Price = createReq.Price

	// UserID
	if err := uuid.Validate(createReq.UserID); err != nil {
		return nil, CreateRequestUserIDErr
	}
	subDTO.UserID = createReq.UserID

	// StartAt
	startAtSplit := strings.Split(createReq.StartAt, "-")
	if len(startAtSplit) != 2 {
		return nil, CreateRequestStartAtErr
	}

	startAt, err := time.Parse("2006-01", fmt.Sprintf("%s-%s", startAtSplit[1], startAtSplit[0]))
	if err != nil {
		return nil, CreateRequestStartAtErr
	}

	subDTO.StartAt = startAt

	// EndAt
	subDTO.EndAt = startAt.AddDate(0, 1, 0)

	return &subDTO, nil
}
