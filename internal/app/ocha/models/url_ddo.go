package models

import (
	"time"
)

// Item - Entry struct for short url service
type Item struct {
	URL          string    `json:"url" example:"https://www.google.com"`
	Key          string    `json:"shortcode" example:"google"`
	StartDate    time.Time `json:"startDate"`
	LastSeenDate *time.Time `json:"lastSeenDate,omitempty"`
	Counter      int       `json:"redirectCount"`
}

// CreateRequest - Create Short URL Request
type CreateRequest struct {
	URL string `json:"url" example:"https://www.google.com" validate:"required,url"`
	Key string `json:"shortcode" example:"google"`
}

// CreateResponse - 200 Ok
type CreateResponse struct {
	Key string `json:"shortcode" example:"google"`
}

