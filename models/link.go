package models

type Link struct {
	Code      string `json:"code"`
	LongURL   string `json:"long_url"`
	ShortURL  string `json:"short_url"`
	CreatedAt string `json:"created_at"`
	Hits      int    `json:"hits"`
}