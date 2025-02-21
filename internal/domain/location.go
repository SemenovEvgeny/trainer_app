package domain

type Location struct {
	ID       int64  `json:"id"`
	Region   int64  `json:"region"`
	City     int64  `json:"city"`
	District int64  `json:"district"`
	Street   int64  `json:"street"`
	House    int64  `json:"house"`
	Text     string `json:"text"`
}
