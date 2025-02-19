package domain

type Contact struct {
	Id        int64  `json:"id" validate:"required"`
	TrainerId int64  `json:"trainer_id" validate:"required"`
	TypeId    int64  `json:"type_id" validate:"required"`
	Contact   string `json:"contact" validate:"required"`
}

type Contact_type struct {
	Id          int64  `json:"id" validate:"required"`
	ContactType string `json:"contact_type" validate:"required"`
}

type ContactList []Contact
