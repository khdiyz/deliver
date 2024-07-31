package models

import "time"

type Pagination struct {
	Limit     int64 `json:"-" default:"10"`
	Offset    int64 `json:"-"  default:"0"`
	Page      int64 `json:"page"  default:"1"`
	PageSize  int64 `json:"page_size"  default:"10"`
	PageTotal int64 `json:"page_total"`
	ItemTotal int64 `json:"item_total"`
}

type BaseResponse struct {
	Success      bool        `json:"success"`
	Status       string      `json:"status" `
	Description  string      `json:"description"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"message,omitempty"`
	Pagination   *Pagination `json:"pagination,omitempty"`
}

type Token struct {
	User      User      `json:"user"`
	Token     string    `json:"token"`
	Type      string    `json:"type"`
	ExpiresAt time.Time `json:"expires_at"`
}