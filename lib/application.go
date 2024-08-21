package gopress

import (
	"fmt"
	"net"
)

type HandlerCallback func(req Request, res Response)

var server HttpServer
type HttpServer struct {
	Port        int
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
			go handleRequests(client);
		}
	}();

	callback();
}