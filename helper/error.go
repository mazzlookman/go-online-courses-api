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

type UnauthorizedError struct {
	Error string
}

func NewUnauthorizedError(error string) UnauthorizedError {
	return UnauthorizedError{Error: error}
}

type BadRequestError struct {
	Error string
}

func NewBadRequestError(error string) BadRequestError {
	return BadRequestError{Error: error}
}

type ResponseErrorKey struct {
	Error string `json:"errors"`
}

func NewResponseErrorKey(error string) ResponseErrorKey {
	return ResponseErrorKey{Error: error}
}
