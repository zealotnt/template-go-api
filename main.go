package main

import (
	"github.com/zealotnt/template-go-api/handlers"
	. "github.com/zealotnt/template-go-api/lib"
)

func main() {
	app := NewApp()
	app.AddRoutes(handlers.GetRoutes())
	app.Run()
	defer app.Close()
}
