package domain

type Contact struct {
	ID        int64  `json:"id" `
	TrainerID int64  `json:"trainer_id" `
	TypeID    int64  `json:"type_id" `
	Contact   string `json:"contact" `
}

type ContactType struct {
	ID    int64  `json:"id" `
	Value string `json:"value" `
}

type ContactList []Contact
