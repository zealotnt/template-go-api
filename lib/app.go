package lib

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
	ENV    string
	PORT   string
}

type BodyMultipart struct {
	Buff        bytes.Buffer
	ContentType string
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w, r)
}

func (app *App) Close() {
}

func (app *App) AddRoutes(routes Routes) {
	for _, route := range routes {
		handler := Logger(route.Handler(app))
		if route.Method == "OPTIONS" {
			app.router.
				// Match all url
				Methods(route.Method).
				Handler(handler)
		} else {
			app.router.
				// Name(route.Name).
				Methods(route.Method).
				Path(route.Pattern).
				Handler(handler)
		}

	}
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":"+app.PORT, app))
}

func (app *App) Request(method string, route string, body interface{}) *httptest.ResponseRecorder {
	var request = &http.Request{}
	switch t := body.(type) {
	case string:
		request, _ = http.NewRequest(method, route, strings.NewReader(t))
		if method == "POST" || method == "PUT" {
			if t != "" && t[0:1] == "{" {
				request.Header.Set("Content-Type", "application/json")
			} else {
				request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
			}
		}
	case BodyMultipart:
		request, _ = http.NewRequest(method, route, &t.Buff)
		request.Header.Set("Content-Type", t.ContentType)
	default:
		return nil
	}

	request.RemoteAddr = "127.0.0.1:9090"

	response := httptest.NewRecorder()
	app.ServeHTTP(response, request)

	return response
}

func NewApp() *App {
	return &App{
		router: newRouter(),
		PORT:   getPort(),
		ENV:    getEnv(),
	}
}

func getPort() string {
	env := os.Getenv("PORT")

	if env == "" {
		return "9090"
	}

	return env
}

func getEnv() string {
	env := os.Getenv("ENV")

	if env == "" {
		return "production"
	}

	return env
}

func newRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}
