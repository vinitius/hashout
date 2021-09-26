package discounts

type UseCase struct {
	cli Client
}

func NewUseCase(cli Client) UseCase {
	return UseCase{cli: cli}
}
