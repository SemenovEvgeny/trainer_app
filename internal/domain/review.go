package domain

type Review struct {
	Id        int64  `json:"id" validate:"required"`
	ServiceId int64  `json:"service_id" validate:"required"`
	Score     int64  `json:"score" validate:"required"`
	Comment   string `json:"comment" validate:"required"`
}

type ReviewList []Review
