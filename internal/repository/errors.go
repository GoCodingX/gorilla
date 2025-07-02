package repository

import "fmt"

type ErrNotFound struct {
	ID  string
	Err error
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("record not found. ID: %s. err: %v", e.ID, e.Err)
}

type ErrAlreadyExists struct {
	ID  string
	Err error
}

func (e *ErrAlreadyExists) Error() string {
	return fmt.Sprintf("record already exists. ID: %s. err: %v", e.ID, e.Err)
}
