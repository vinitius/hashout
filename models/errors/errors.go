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

type MalFormed struct {
	Err   error
	Input string
}

func (e *MalFormed) Error() string {
	return e.Input + " MAL FORMED: " + e.Err.Error()
}

type Unexpected struct {
	Err     error
	Message string
}

func (e *Unexpected) Error() string {
	return e.Message + " UNEXPECTED: " + e.Err.Error()
}
