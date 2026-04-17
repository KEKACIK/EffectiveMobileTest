package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	SubscriptionIDNotNumberErr      error = errors.New("Invalid validation id: Not number")
	SubscriptionIDNegativeNumberErr error = errors.New("Invalid validation id: Negative number")
)

func SubscriptionIdValidate(id string) (int, error) {
	id = strings.TrimSpace(id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, SubscriptionIDNotNumberErr
	}

	if idInt < 0 {
		return 0, SubscriptionIDNegativeNumberErr
	}

	return idInt, nil
}

var (
	SubscriptionNameEmptyErr error = errors.New("Invalid validation name: Empty")
)

func SubscriptionNameValidate(name string) (string, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return "", SubscriptionNameEmptyErr
	}

	return name, nil
}

var (
	SubscriptionPriceNegativeErr error = errors.New("Invalid validation price: Negative number. Excepted price > 0")
	SubscriptionPriceZeroErr     error = errors.New("Invalid validation price: Zero number. Excepted price > 0")
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
	SubscriptionUserIdEmptyErr error = errors.New("Invalid validation user_id: Empty")
	SubscriptionUserIdUUIDErr  error = errors.New("Invalid validation user_id: UUID invalid")
)

func SubscriptionUserIdValidate(userID string) (string, error) {
	userID = strings.TrimSpace(userID)

	if userID == "" {
		return "", SubscriptionUserIdEmptyErr
	}
	if err := uuid.Validate(userID); err != nil {
		return "", SubscriptionUserIdUUIDErr
	}

	return userID, nil
}

var (
	SubscriptionDateAtErr error = errors.New("Invalid validation: invalid Date. Excepted format \"MM-YYYY\"")
)

func SubscriptionDateAtValidate(timeAt string) (time.Time, error) {
	// Формат ввода времени "MM-YYYY" (01.2026)
	timeAt = strings.TrimSpace(timeAt)

	timeAtSplit := strings.Split(timeAt, "-")
	if len(timeAtSplit) != 2 {
		return time.Time{}, SubscriptionDateAtErr
	}

	timeAtTime, err := time.Parse("2006-01", fmt.Sprintf("%s-%s", timeAtSplit[1], timeAtSplit[0]))
	if err != nil {
		return time.Time{}, SubscriptionDateAtErr
	}

	return timeAtTime, nil
}
