package models

import (
	"time"
)

// Item - Entry struct for short url service
type Item struct {
	URL          string     `json:"url" example:"https://www.google.com"`
	Key          string     `json:"shortcode" example:"google"`
	StartDate    time.Time  `json:"startDate"`
	LastSeenDate *time.Time `json:"lastSeenDate,omitempty"`
	Counter      int        `json:"redirectCount"`
}

// CreateRequest - Create Short URL Request
type CreateRequest struct {
	URL string `json:"url" example:"https://www.google.com" valid:"url,required~url is not present"`
	Key string `json:"shortcode" example:"google" valid:"stringlength(6|6)~shortcode length should be 6,optional"`
}

// CreateResponse - 200 Ok
type CreateResponse struct {
	Key string `json:"shortcode" example:"google"`
}
