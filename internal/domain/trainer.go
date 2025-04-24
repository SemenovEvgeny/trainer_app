package domain

type Trainer struct {
	ID          int64  `json:"id" `
	LastName    string `json:"last_name" `
	FirstName   string `json:"first_name" `
	MiddleName  string `json:"middle_name" `
	Description string `json:"description" `
	IsActive    bool   `json:"is_active" `
}

type TrainerList []Trainer
