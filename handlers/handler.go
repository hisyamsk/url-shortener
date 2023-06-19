package handlers

type V1Handlers struct {
	UserHandler
	UrlHandler
}

type ApiVersionHandlers struct {
	V1Handlers *V1Handlers
}
