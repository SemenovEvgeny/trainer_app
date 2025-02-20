package domain

type Review struct {
	ID        int64  `json:"ID" `
	ServiceID int64  `json:"service_ID" `
	Score     int64  `json:"score" `
	Comment   string `json:"comment" `
}

type ReviewList []Review
