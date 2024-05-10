package models

type Record struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
	URL         string  `json:"url"`
	Description string  `json:"description"`
}
