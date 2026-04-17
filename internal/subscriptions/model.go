package subscriptions

import "time"

type Subscription struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Price   int       `json:"price"`
	UserID  string    `json:"user_id"`
	StartAt time.Time `json:"start_date"`
	EndAt   time.Time `json:"end_date"`
}
