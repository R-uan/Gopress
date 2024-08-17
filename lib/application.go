package gopress

import (
	"fmt"
	"net"
)

type HandlerCallback func(req Request, res Response)

type HttpServer struct {
	Port        int
	Middlewares []func()
}

type HttpMethodHandlers struct {
	GetMethod  map[string]HandlerCallback
	PostMethod map[string]HandlerCallback
}

//	Summary:
//		Creates the main server struct.
//
//	Returns:
//		Pointer to the server struct, allowing access to the rest of the application.
func Gopress() (*HttpServer) {
	server = HttpServer {}
	return &server;
}

//	Summary:
//		Initializes the server to listen on given port and runs the callback function.
//
//	Disclaimer:
//		Given that golang sucks and there is no optional parameters or method overloading,
//		the callback function is mandatory.
func (server *HttpServer) Listen(port int, callback func()) {
	var Listener net.Listener;
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port));
	if err != nil { panic(err) }
	Listener = listener;
	go func(){
		for {
			client, err := Listener.Accept();
			if err != nil {
				fmt.Println(err);
				continue;
			}
			go HandleRequests(client);
		}
	}();

	callback();
}

//	Summary:
//		Function that handles individual accepted requests on it s own goroutine.
func HandleRequests(client net.Conn) {
	defer client.Close();
	var rawRequest = make([]byte, 1024);
	client.Read(rawRequest);

	var req *Request = ExtractRequestData(string(rawRequest));
	var res *Response = BuildResponse(&client);

	RequestRouting(req, res);
}

//	Summary:
//		Builds basic response struct.
func BuildResponse(client *net.Conn) (*Response) {
	return &Response{
		Client: client,
		Protocol: "HTTP/1.1",
		Headers: ResponseHeaders{
			Server: "Gopress/0.1",
			Connection: "keep-alive",
			CacheControl: "no-cache",
			AccessControlAllowOrigin: "*",
			XPoweredBy: "Go(Golang)",
		},
	}
}

//	Summary:
//		Routes the Request struct to the appropriate handler.
func RequestRouting(request *Request, response *Response) {
	var method = request.Method
	if(method == "GET") {
		var route = httpMethods.GetMethod[request.Path]
		if route != nil { route(*request, *response) }
	} else if (method == "POST") {
			var route = httpMethods.GetMethod[request.Path]
			if route != nil { route(*request, *response) }
	} // TODO: Create a not found response method
}

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
