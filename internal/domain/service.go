package domain

import "time"

type Service struct {
	Id          int    `json:"id" validate:"required"`
	TrainerId   int    `json:"trainer_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`
}

type ServiceList []Service

type ServicePrice struct {
	Id       int       `json:"id" validate:"required"`
	CreateAd time.Time `json:"create_at" validate:"required"`
	Price    float64   `json:"proce" validate:"required"`
}

type ServicePriceList []ServicePrice
