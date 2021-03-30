package models

import "errors"

var (
	ErrConstraintViolationEmail    = errors.New("constraint email")
	ErrConstraintViolationNickname = errors.New("constraint nickname")
	ErrDefaultDB                   = errors.New("something wrong DB")
)
