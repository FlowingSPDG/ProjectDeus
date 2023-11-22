package database

import "errors"

var (
	ErrNotFound   = errors.New("specified resource not found")
	ErrProhibited = errors.New("prohibited")
	ErrInternal   = errors.New("internal error")
)
