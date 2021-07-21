package errors

type BusinessLogicError struct {
	ClientHTTPCode 	int
	ClientMessage 	string
}

func (err BusinessLogicError) IsNil() bool {
	return err.ClientMessage == ""
}