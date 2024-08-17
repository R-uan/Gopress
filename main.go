package main

import (
	"fmt"
	gopress "gopress/lib"
	"time"
)

func main() {
	var app = gopress.Gopress();
	
	app.Get("/", func(req gopress.Request, res gopress.Response) {
		res.Send("<h1>Hello World</h1>", 200);
	})
	
	app.Listen(3000, func() {
		fmt.Println("Listening on port 3000.");
	});
	
	for { time.Sleep(1 * time.Second) }
}