package gopress

type Routes struct {
	GetMethod  map[string]HandlerCallback
	PostMethod map[string]HandlerCallback
}

var routesMap Routes

// Summary:
//
//	Routes the Request struct to the appropriate handler.
func routing(request *Request, response *Response) {
	var method = request.Method
	if method == "GET" {
		var route = routesMap.GetMethod[request.Path]
		if route != nil {
			route(*request, *response)
		}
	} else if method == "POST" {
		var route = routesMap.GetMethod[request.Path]
		if route != nil {
			route(*request, *response)
		}
	} // TODO: Create a not found response method
}

// Summary:
//
//	Adds the callback function to the GET Method handlers.
func (server *HttpServer) Get(path string, callback HandlerCallback) {
	if routesMap.GetMethod == nil {
		routesMap.GetMethod = make(map[string]HandlerCallback)
	}
	routesMap.GetMethod[path] = callback
}

// Summary:
//
//	Adds the callback function to the POST Method handlers.
func (server *HttpServer) Post(path string, callback HandlerCallback) {
	if routesMap.PostMethod == nil {
		routesMap.PostMethod = make(map[string]HandlerCallback)
	}
	routesMap.PostMethod[path] = callback
}