package api

type CreateRequest struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	UserID  string `json:"user_id"`
	StartAt string `json:"start_at"`
	EndAt   string `json:"end_at"`
}

type DeleteRequest struct {
	ID int `json:"id"`
}
