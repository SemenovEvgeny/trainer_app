package domain

type Title struct {
	Id        int64  `json:"id" validate:"required"`
	TrainerId int64  `json:"trainer_id" validate:"required"`
	Title     string `json:"title" validate:"required"`
}

type TitleList []Title
