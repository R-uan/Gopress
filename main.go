package main

import (
	"fmt"
	gopress "gopress/lib"
	"time"
)

func main() {
	var app = gopress.Gopress();

	app.Use(middleware);
	app.Use(middleware2);
	
	app.Get("/", func(req gopress.Request, res gopress.Response) {
		res.Send("Hello World", 200);
	})

	app.Post("/" , func(req gopress.Request, res gopress.Response) {
		res.Json("{ \"message\": \"hi\" }", 200);
	});

	app.Listen(3000, func() {
		fmt.Println("Listening on port 3000.");
	});
	
	for { time.Sleep(1 * time.Second) }
}

func middleware(req *gopress.Request, res *gopress.Response, next gopress.NextFunction) {
	fmt.Println("Middlware");
	next()
}

func middleware2(req *gopress.Request, res *gopress.Response, next gopress.NextFunction) {
	fmt.Println("Middlware 2");
	next()
}