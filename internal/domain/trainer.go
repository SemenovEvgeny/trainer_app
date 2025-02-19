package domain

type Trainer struct {
	Id          int64  `json:"id" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	MiddleName  string `json:"middle_name" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsActive    bool   `json:"is_active" validate:"required"`
}

type TrainerList []Trainer
