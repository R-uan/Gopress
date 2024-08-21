package main

import (
	"fmt"
	gopress "gopress/lib"
	"time"
)

func main() {
	var app = gopress.Gopress();
	
	app.Get("/", func(req gopress.Request, res gopress.Response) {
		res.Headers.Location = "Hey";
		res.Send("<h1>Hello World</h1>", 200);
	})

	app.Use(middleware);
	app.Use(middleware2);

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