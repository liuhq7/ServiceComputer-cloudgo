package service

import (
	"net/http"
	"github.com/mux"
	"github.com/negroni"
	"github.com/render"
)

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/welcome/{id}", testHandler(formatter)).Methods("GET")
}

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func testHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Welcome " + id})
	}
}
