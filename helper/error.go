package helper

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
