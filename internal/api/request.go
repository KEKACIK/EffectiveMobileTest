package api

type SubscriptionCreateRequest struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	UserID  string `json:"user_id"`
	StartAt string `json:"start_date"`
	EndAt   string `json:"end_date"`
}

type SubscriptionListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type SubscriptionGetRequest struct {
	ID int `json:"id"`
}

type SubscriptionUpdateRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	UserID  string `json:"user_id"`
	StartAt string `json:"start_date"`
}

type SubscriptionDeleteRequest struct {
	ID int `json:"id"`
}
