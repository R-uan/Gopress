package main

import (
	"fmt"
	gopress "gopress/lib"
)

type Hello struct {
	Message string `json:"message"`
	Other string `json:"Other"`
}

func main() {
	var app = gopress.Gopress();

	app.Use(middleware)

	app.Get("/", func(req gopress.Request, res gopress.Response) {
		res.Json(Hello{Message: "Hey world", Other: "This is a get request"}, 200);
	})

	app.Post("/" , func(req gopress.Request, res gopress.Response) {
		res.Send("This is a post request.", 200);
	});
	
	app.Patch("/", func(req gopress.Request, res gopress.Response) {
		res.Send("This is a patch request.", 200);
	});

	app.Delete("/",func(req gopress.Request, res gopress.Response) {
		res.Send("This is a delete request", 200);
	});

	// Start the server on the port 3000 and listen for requests 
	app.Listen(3000, func() {
		fmt.Println("Listening on port 3000.");
	});
}

func middleware(req *gopress.Request, res *gopress.Response, next gopress.NextFunction) {
	// This middleware will be executed before the request reach the actual handler.
	fmt.Println("Middleware");

	//	This will call the next middleware function.
	next()
}