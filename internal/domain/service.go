package domain

import "time"

type Service struct {
	ID          int    `json:"id" `
	TrainerID   int    `json:"trainer_id" `
	Name        string `json:"name" `
	PriceID     int    `json:"price_id" `
	Description string `json:"description" `
	Location    string `json:"location" `
}

type ServiceList []Service

type ServicePrice struct {
	ID       int       `json:"id" `
	CreateAd time.Time `json:"create_at" `
	Price    float64   `json:"proce" `
}

type ServicePriceList []ServicePrice
