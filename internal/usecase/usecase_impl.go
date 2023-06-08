package usecase

type Usecase interface {
	UserHandler() UserUsecase
}
