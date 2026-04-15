package subscriptions

import "time"

type CreateSubscriptionDTO struct {
	Name    string    `json:"name"`
	Price   int       `json:"price"`
	UserID  string    `json:"user_id"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}
