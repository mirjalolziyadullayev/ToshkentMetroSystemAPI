package models

type Model struct {
	Id                     int    `json:"Id"`
	Name                   string `json:"Name"`
	YearBuilt              string `json:"YearBuilt"`
	OrdinalNumberInItsLine int64  `json:"OrdinalNumberInItsLine"`
	LineName               string `json:"LineName"`
}
