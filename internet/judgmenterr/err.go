package judgmenterr

type ContError struct {
	Status  int
	Code    int
	Message string
}

func (err ContError) Error() string {
	return err.Message
}
