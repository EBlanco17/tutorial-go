package models

type Character struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Species  string `json:"species"`
	Type     string `json:"type"`
	Gender   string `json:"gender"`
	Origin   Location
	Location Location
	Image    string `json:"image"`
	Episodes []Episode

	EpisodesBase []string     `json:"episode"`
	OriginBase   LocationBase `json:"origin"`
	LocationBase LocationBase `json:"location"`
}

type CharacterClean struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	Species  string    `json:"species"`
	Type     string    `json:"type"`
	Gender   string    `json:"gender"`
	Origin   Location  `json:"origin"`
	Location Location  `json:"location"`
	Image    string    `json:"image"`
	Episodes []Episode `json:"episodes"`
}
