package domain

type Achievement struct {
	ID        int64  `json:"id" `
	TrainerID int64  `json:"trainer_id" `
	Value     string `json:"value" `
}

type AchievementList []Achievement
