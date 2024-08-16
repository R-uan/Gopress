package gopress

//	Summary:
//		Adds the callback function to the GET Method handlers.
func (server *HttpServer) Get(path string, callback HandlerCallback) {
	if(httpMethods.GetMethod == nil) { httpMethods.GetMethod = make(map[string]HandlerCallback) }
	httpMethods.GetMethod[path] = callback;
}

//	Summary:
//		Adds the callback function to the POST Method handlers.
func (server *HttpServer) Post(path string, callback HandlerCallback) {
	if (httpMethods.PostMethod == nil) { httpMethods.PostMethod = make(map[string]HandlerCallback) }
	httpMethods.PostMethod[path] = callback;
}
