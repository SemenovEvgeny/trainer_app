package domain

type Achievement struct {
	Id          int64  `json:"id" validate:"required"`
	TrainerId   int64  `json:"trainer_id" validate:"required"`
	Achievement string `json:"achievement" validate:"required"`
}

type AchievementList []Achievement
