package main

import (
	"fmt"
	gopress "gocha/lib"
	"time"
)

func main() {
	var server = gopress.Gocha();
	server.Listen(3000, func() {
		fmt.Println("Listening on port 3000.");
	});

	for { time.Sleep(1 * time.Second) }
}