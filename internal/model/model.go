package model

type IDNote int

type Note struct {
	IDNote     `json:"id"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	ExpireTime string `json:"expire_time,omitempty"`
}
