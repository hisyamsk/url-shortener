package handlers

type Handlers struct {
	UserHandler
	UrlHandler
}

type ApiVersionHandlers struct {
	V1Handlers *Handlers
}

func NewHandlers(handlers *Handlers) *Handlers {
	return &Handlers{
		UserHandler: handlers.UserHandler,
		UrlHandler:  handlers.UrlHandler,
	}
}
