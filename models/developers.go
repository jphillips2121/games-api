package models

// Developers contains an array of all developers.
type Developers struct {
	Developers []Developer `json:"developers"`
}

// Developer contains a breakdown of each individual developer.
type Developer struct {
	Name         string `json:"name"`
	Headquarters string `json:"headquarters"`
}
