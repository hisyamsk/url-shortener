package handlers

type V1Handlers struct {
	MainHandler
	UserHandler
	UrlHandler
}

type ApiVersionHandlers struct {
	V1Handlers *V1Handlers
}
