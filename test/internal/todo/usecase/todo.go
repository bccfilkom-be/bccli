package usecase

type todoUsecaseItf interface {}

type todoUsecase struct {}

func NewtodoUsecase() todoUsecaseItf {
    return &todoUsecase{}
}
