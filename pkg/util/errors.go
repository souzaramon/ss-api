package util

import "errors"

type ApiError struct {
	Message string
}

var ErrNotFound = errors.New("item not found")
