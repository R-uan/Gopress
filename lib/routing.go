package gopress

type HandlerCallback func(req Request, res Response)
type Routes struct {
	GetMethod    map[string]HandlerCallback
	PostMethod   map[string]HandlerCallback
	PatchMethod  map[string]HandlerCallback
	DeleteMethod map[string]HandlerCallback
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
		var route = routesMap.PostMethod[request.Path]
		if route != nil {
			route(*request, *response)
		}
	} else if method == "PATCH" {
		var route = routesMap.PatchMethod[request.Path]
		if route != nil {
			route(*request, *response)
		}
	} else if method == "DELETE" {
		var route = routesMap.DeleteMethod[request.Path]
		if route != nil {
			route(*request, *response)
		}
	}
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
//	Add the callback function to the POST Method handlers.
func (server *HttpServer) Post(path string, callback HandlerCallback) {
	if routesMap.PostMethod == nil {
		routesMap.PostMethod = make(map[string]HandlerCallback)
	}
	routesMap.PostMethod[path] = callback
}

// Summary:
//
//	Add the callback function to the PATCH Method handlers.
func (server *HttpServer) Patch(path string, callback HandlerCallback) {
	if routesMap.PatchMethod == nil {
		routesMap.PatchMethod = make(map[string]HandlerCallback)
	}
	routesMap.PatchMethod[path] = callback
}

// Summary:
//
//	Add the callback function to the DELETE Method handlers.
func (server *HttpServer) Delete(path string, callback HandlerCallback) {
	if routesMap.DeleteMethod == nil {
		routesMap.DeleteMethod = make(map[string]HandlerCallback)
	}
	routesMap.DeleteMethod[path] = callback
}
