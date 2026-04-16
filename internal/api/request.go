package api

type SubscribeCreateRequest struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	UserID  string `json:"user_id"`
	StartAt string `json:"start_at"`
	EndAt   string `json:"end_at"`
}

type SubscribeListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type SubscribeGetRequest struct {
	ID int `json:"id"`
}

type SubscribeUpdateRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	UserID  string `json:"user_id"`
	StartAt string `json:"start_at"`
}

type SubscribeDeleteRequest struct {
	ID int `json:"id"`
}
