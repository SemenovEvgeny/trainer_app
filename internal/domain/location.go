package domain

type Location struct {
	ID         int64  `json:"id"`
	RegionID   int64  `json:"region_id"`
	CityID     int64  `json:"city_id"`
	DistrictID int64  `json:"district_id"`
	StreetID   int64  `json:"street_id"`
	HouseID    int64  `json:"house_id"`
	Text       string `json:"text"`
}

type LocationDetails struct {
	ID    int64  `json:"id"`
	Value string `json:"value"`
}
