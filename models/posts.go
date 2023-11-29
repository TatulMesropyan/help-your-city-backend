package models

type Post struct {
	id          int      `json:"id"`
	title       string   `json:"title"`
	description string   `json:"description"`
	status      uint8    `json:"status"`
	city        string   `json:"city"`
	region      string   `json:"region"`
	country     string   `json:"country"`
	createdAt   string   `json:"createdAt"`
	createdBy   int      `json:"createdBy"`
	starredBy   []int    `json:"starredBy"`
	images      []string `json:"images"`
}
