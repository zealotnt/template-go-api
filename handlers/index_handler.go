package handlers

import (
	"encoding/json"
	. "github.com/zealotnt/template-go-api/lib"
	"net/http"
)

type Index struct {
	Version string `json:"version"`
}

func IndexHandler(app *App) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, c Context) {
		index := Index{"1.0.0"}

		json.NewEncoder(w).Encode(index)
	}
}
