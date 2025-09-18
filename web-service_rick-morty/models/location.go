package models

type Location struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Dimension string `json:"dimension"`
	// Residents []string `json:"residents"`
}

type LocationBase struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
