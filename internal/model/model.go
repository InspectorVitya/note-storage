package model

import "time"

type IDNote int

type Note struct {
	IDNote     `json:"id"`
	Title      string    `json:"title"`
	Text       string    `json:"text"`
	CreateTime time.Time `json:"create_time"`
}
