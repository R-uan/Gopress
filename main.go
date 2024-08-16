package main

import (
	"fmt"
	gopress "gocha/lib"
	"time"
)

func main() {
	var server = gopress.Gopress();
	
	server.Get("/", func(req gopress.Request) {
		println(req.Path);
		println(req.Method);
	})
	
	server.Listen(3000, func() {
		fmt.Println("Listening on port 3000.");
	});
	
	for { time.Sleep(1 * time.Second) }
}