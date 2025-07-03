package repository

// NewAlreadyExistsError creates a new AlreadyExistsError instance.
func NewAlreadyExistsError(id string, err error) *AlreadyExistsError {
	return &AlreadyExistsError{
		Msg: "record already exists. Value: " + id,
		Err: err,
	}
}

type AlreadyExistsError struct {
	Msg string
	Err error
}

func (e *AlreadyExistsError) Error() string {
	return e.Msg + ": " + e.Err.Error()
}

func (e *AlreadyExistsError) Unwrap() error {
	return e.Err
}

// NewInvalidReferenceError creates a new InvalidReferenceError instance.
func NewInvalidReferenceError(id string, err error) *InvalidReferenceError {
	return &InvalidReferenceError{
		Msg: "referred value does not exist. Value: " + id,
		Err: err,
	}
}

type InvalidReferenceError struct {
	Msg string
	Err error
}

func (e *InvalidReferenceError) Error() string {
	return e.Msg + ": " + e.Err.Error()
}

func (e *InvalidReferenceError) Unwrap() error {
	return e.Err
}
