package checkout

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return UseCase{repo: repo}
}
