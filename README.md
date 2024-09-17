# Golang HTTP Wrapper

Inspired in express.js, this is a HTTP wrapper written in golang. At this moment it has support for GET/POST/PATCH/DELETE request methods and the capacity to use middlewares. It maps request/response headers to structs so it can be read or written to.

```go
func main() {
	var app = gopress.Gopress();

	app.Use(middleware)

	app.Get("/", func(req gopress.Request, res gopress.Response) {
		res.Json(Hello{Message: "Hey world", Other: "Other"}, 200);
	})

	app.Post("/" , func(req gopress.Request, res gopress.Response) {
		res.Json("{ \"message\": \"hi\" }", 200);
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

```
