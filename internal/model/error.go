package model

import "errors"

var (
	ErrNotExistNote = errors.New("not exist note")
	ErrEmptyList    = errors.New("empty list")
)
