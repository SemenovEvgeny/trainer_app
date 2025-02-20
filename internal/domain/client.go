package domain

import "time"

type Client struct {
	ID          int64  `json:"ID" `
	LastName    string `json:"last_name" `
	FirstName   string `json:"first_name" `
	MiddleName  string `json:"middle_name" `
	Description string `json:"description" `
	IsActive    bool   `json:"is_active" `
}

type ClientList []Client

type ClientService struct {
	ID        int64     `json:"id" `
	ClientID  int64     `json:"client_id" `
	ServiceID int64     `json:"service_id" `
	PriceID   int64     `json:"price_id" `
	Date      time.Time `json:"date" `
	IsActive  bool      `json:"is_active" `
}

type ClientServiceList []ClientService
