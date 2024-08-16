package gopress

import (
	"fmt"
	"net"
	"strings"
)

type Callback func();
type HttpServer struct {
	Port int;
	Middlewares []func();
}



type Response struct {}

var Listener net.Listener;
var HttpHeaders RequestHeaders;
func Gocha() (*HttpServer) {
	var server = HttpServer {}
	return &server;
}

func (server *HttpServer) Listen(port int, callback Callback) {
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

func HandleRequests(client net.Conn) {
	defer client.Close();	
	var rawRequest = make([]byte, 1024);
	client.Read(rawRequest);

	var _ *Request = ExtractRequestStruct(string(rawRequest));
}

func ExtractRequestStruct(rawRequest string) (*Request) {
	var Request = Request{};
	var headers = make(map[string]string);
	var rawRequestArray = strings.Split(rawRequest, "\n");
	var requestHead = strings.SplitN(rawRequestArray[0], " ", 3);

	for line := range rawRequestArray {
		var parts = strings.SplitN(rawRequestArray[line], ":", 2);
		if len(parts) == 2 { headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1]) }
	}

	Request.Path = requestHead[1];
	Request.Method = requestHead[0];
	Request.Headers = RequestHeaders{
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
	};

	return &Request;
}

type handle func(req Request, res Response);
func (server *HttpServer) Get(path string, callback handle) {


}
