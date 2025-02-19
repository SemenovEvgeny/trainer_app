package domain

import "time"

type Client struct {
	Id          int64  `json:"id" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	MiddleName  string `json:"middle_name" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsActive    bool   `json:"is_active" validate:"required"`
}

type ClientList []Client

type ClientService struct {
	Id        int64     `json:"id" validate:"required"`
	ClientId  int64     `json:"client_id" validate:"required"`
	ServiceId int64     `json:"service_id" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	IsActive  bool      `json:"is_active" validate:"required"`
}

type ClientsServiceList []ClientService
