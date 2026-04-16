package subscriptions

import "time"

type SubscriptionCreateDTO struct {
	Name    string    `json:"name"`
	Price   int       `json:"price"`
	UserID  string    `json:"user_id"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

type SubscriptionUpdateDTO struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Price   int       `json:"price"`
	UserID  string    `json:"user_id"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}
