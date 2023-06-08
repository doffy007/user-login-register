package handler

type Handler interface {
	UserHandler() UserHandler
}
