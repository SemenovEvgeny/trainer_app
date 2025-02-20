package domain

type Title struct {
	ID        int64  `json:"ID" `
	TrainerID int64  `json:"trainer_id" `
	Value     string `json:"value" `
}

type TitleList []Title
