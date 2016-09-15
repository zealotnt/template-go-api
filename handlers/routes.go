package handlers

import (
	. "github.com/zealotnt/template-go-api/lib"
)

func GetRoutes() Routes {
	return Routes{
		Route{"GET", "/", IndexHandler},
	}
}
