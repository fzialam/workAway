package exception

type DuplicatedError struct {
	Error string
}

func NewDuplicatedError(error string) DuplicatedError {
	return DuplicatedError{Error: error}
}
