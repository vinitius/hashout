package checkout

type Writer interface {
}

type Reader interface {
}

type Repository interface {
	Writer
	Reader
}

type Service interface {
}
