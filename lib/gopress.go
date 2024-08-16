package gopress

import (
	"fmt"
	"net"
	"strings"
)

var server HttpServer;
var httpMethods HttpMethodHandlers;

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
	RequestRouting(req);
}

//	Summary:
//		Maps the http method, path, headers and body to a Request struct.
//
//	Returns:
//		Returns the pointer to said struct.
func ExtractRequestData(rawRequest string) (*Request) {
	
	var rawRequestArray = strings.Split(rawRequest, "\n");
	var requestHead = strings.SplitN(rawRequestArray[0], " ", 3);
	
	var headers = make(map[string]string);
	for line := range rawRequestArray {
		var parts = strings.SplitN(rawRequestArray[line], ":", 2);
		if len(parts) == 2 { headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1]) }
	}
	
	var Request = Request{
		Path: requestHead[1],
		Method: requestHead[0],
		Body: strings.SplitN(rawRequest, "\r\n", 2)[1],
		Headers: RequestHeaders{
			Host:            headers["Host"],
			UserAgent:       headers["User-Agent"],
			Accept:          headers["Accept"],
			AcceptLanguage:  headers["Accept-Language"],
			AcceptEncoding:  headers["Accept-Encoding"],
			Connection:      headers["Connection"],
			ContentType:     headers["Content-Type"],
			ContentLength:   ParseContentLength(headers["Content-Length"]),
			Authorization:   headers["Authorization"],
			Cookie:          headers["Cookie"],
			Referer:         headers["Referer"],
			CacheControl:    headers["Cache-Control"],
			UpgradeInsecure: headers["Upgrade-Insecure-Requests"],
			IfModifiedSince: headers["If-Modified-Since"],
			IfNoneMatch:     headers["If-None-Match"],
			Origin:          headers["Origin"],
			Pragma:          headers["Pragma"],
			XRequestedWith:  headers["X-Requested-With"],
			XForwardedFor:   headers["X-Forwarded-For"],
			XRealIP:         headers["X-Real-IP"],
			Range:           headers["Range"],
		},
	};
	
	return &Request;
}

//	Summary:
//		Routes the Request struct to the appropriate handler.
func RequestRouting(request *Request) {
	var method = request.Method
	if(method == "GET") {
		var route = httpMethods.GetMethod[request.Path]
		if route != nil { route(*request) }
	} else if (method == "POST") {
			var route = httpMethods.GetMethod[request.Path]
			if route != nil { route(*request) }
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
