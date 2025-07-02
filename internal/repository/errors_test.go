package repository_test

import (
	"errors"
	"testing"

	"github.com/GoCodingX/gorilla/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestErrAlreadyExists(t *testing.T) {
	err := &repository.ErrAlreadyExists{
		ID:  "1234",
		Err: errors.New("something bad happened"),
	}
	expectedErrorMsg := "record already exists. ID: 1234. err: something bad happened"

	assert.EqualError(t, err, expectedErrorMsg)
}

func TestErrAlreadyExists_As(t *testing.T) {
	err := &repository.ErrAlreadyExists{
		ID: "9012",
	}

	// test that errors.As returns true when target is of the same type as err
	var targetErr *repository.ErrAlreadyExists

	assert.True(t, errors.As(err, &targetErr))
	assert.Equal(t, err.ID, targetErr.ID)

	// test that errors.As returns false when target is of a different type than err
	var otherErr *repository.ErrNotFound

	assert.False(t, errors.As(err, &otherErr))
}

func TestErrNotFound(t *testing.T) {
	err := &repository.ErrNotFound{
		ID:  "5678",
		Err: errors.New("something unexpected happened"),
	}
	expectedErrorMsg := "record not found. ID: 5678. err: something unexpected happened"

	assert.EqualError(t, err, expectedErrorMsg)
}

func TestErrNotFound_As(t *testing.T) {
	err := &repository.ErrNotFound{ID: "9012"}

	// test that errors.As returns true when target is of the same type as err
	var targetErr *repository.ErrNotFound

	assert.True(t, errors.As(err, &targetErr))
	assert.Equal(t, err.ID, targetErr.ID)

	// test that errors.As returns false when target is of a different type than err
	var otherErr *repository.ErrAlreadyExists

	assert.False(t, errors.As(err, &otherErr))
}
