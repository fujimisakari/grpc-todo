package usecase

type Usecase struct {
	logger Logger
}

func NewUsecase(logger Logger) *Usecase {
	return &Usecase{
		logger: logger,
	}
}
