package server

import (
  "github.com/gorilla/mux"
  "html/template"
  "net/http"
  "log"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("tmpl/*.gohtml"))
}

func home(res http.ResponseWriter, req *http.Request) {
  err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
  if err != nil {
    log.Fatal("template didn't execute: ", err)
  }
}

func StartServer() {
  router := mux.NewRouter()
  router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  router.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
  router.HandleFunc("/", home)
  http.ListenAndServe(":5150", router)
}
