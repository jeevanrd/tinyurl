package utils

import "errors"

var (
	ErrInvalidArgument = errors.New("ErrInvalidArgument")
	JsonParseError = errors.New("Unable to parse json body")
	ErrBadRoute = errors.New("bad route")
	ErrNotFound = errors.New("entity not found")
)