package models

type Episode struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	AirDate string `json:"air_date"`
	Episode string `json:"episode"`
	// Characters []string `json:"characters"`
}
