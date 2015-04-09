package dashboard

import (
	"net/http"
	"path"
	"fmt"
	"os"
	
	"github.com/AdoHe/etcd-dashboard/third_party/github.com/gorilla/mux"
)

var pwd string

func addSlash(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, path.Join("mod", req.URL.Path)+"/", 302)
	return
}

func indexPage(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s", "serve index page")
	http.ServeFile(w, req, path.Join(pwd, "/dashboard/dist", "index.html"))
}

func HttpHandler() (handler http.Handler) {
	handler = http.FileServer(http.Dir(path.Join(pwd, "/dashboard/dist")))
	return handler
}

func HTTPHandler() http.Handler {
	pwd, _ = os.Getwd()
	r := mux.NewRouter()
	r.HandleFunc("/dashboard", addSlash)

	r.PathPrefix("/dashboard/static/").Handler(http.StripPrefix("/dashboard/static/", HttpHandler()))
	r.HandleFunc("/dashboard{path:.*}", indexPage)
	
	return r
}
