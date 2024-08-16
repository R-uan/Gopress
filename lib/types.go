package gopress

type HandlerCallback func(req Request)

type HttpServer struct {
	Port        int
	Middlewares []func()
}

type HttpMethodHandlers struct {
	GetMethod  map[string]HandlerCallback
	PostMethod map[string]HandlerCallback
}