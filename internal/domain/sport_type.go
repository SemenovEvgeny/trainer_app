package domain

type SportType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type SportTypeList []SportType
