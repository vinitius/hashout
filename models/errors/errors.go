package customErr

type NotFound struct {
	Err    error
	Entity string
}

func (e *NotFound) Error() string {
	return e.Entity + " NOT FOUND: " + e.Err.Error()
}

type NotValid struct {
	Err   error
	Input string
}

func (e *NotValid) Error() string {
	return e.Input + " NOT VALID: " + e.Err.Error()
}

type DiscountError struct {
	Err  error
	Type string
}

func (e *DiscountError) Error() string {
	return e.Type + " DISCOUNT ERROR: " + e.Err.Error()
}

type Unexpected struct {
	Err     error
	Message string
}

func (e *Unexpected) Error() string {
	return e.Message + " UNEXPECTED: " + e.Err.Error()
}
