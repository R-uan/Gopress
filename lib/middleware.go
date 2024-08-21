package gopress

type NextFunction func( /* req Request, res Response, next NextFunction */ )
type MiddlewareCallback func(req *Request, res *Response, next NextFunction)

var middlewareStorage []MiddlewareCallback

func (server *HttpServer) Use(middleware MiddlewareCallback) {
	if middlewareStorage == nil {
		middlewareStorage = make([]MiddlewareCallback, 0)
	}
	middlewareStorage = append(middlewareStorage, middleware)
}

func middlewarePipeLine(req *Request, res *Response) {
	var runMiddleware func(int)
	runMiddleware = func(index int) {
		if index < len(middlewareStorage) {
			var middleware = middlewareStorage[index]
			middleware(req, res, func() { runMiddleware(index + 1) })
		}
	}
	runMiddleware(0)
}