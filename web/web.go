package web

import (
	"encoding/base64"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/mathcass/go-image-cache/db"
)

const base64GifPixel = "R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs="

func homeHandler(w http.ResponseWriter, r *http.Request) {
	pathResults := db.GetUniquePathResults()

	tpl := template.New("home.html")
	tpl, err := tpl.ParseFiles("web/home.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(w, pathResults)

	if err != nil {
		log.Fatal(err)
	}

}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]
	userAgent := r.UserAgent()
	db.InsertPath(path, userAgent)

	w.Header().Set("Content-Type", "image/gif")
	output, _ := base64.StdEncoding.DecodeString(base64GifPixel)
	io.WriteString(w, string(output))
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]

	logResults := db.GetPathResults(path)

	tpl := template.New("log.html")
	tpl, err := tpl.ParseFiles("web/log.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(w, logResults)

	if err != nil {
		log.Fatal(err)
	}
}

func Serve() {
	log.Println("Starting go-image-cache webserver")
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/image/{path}", imageHandler)
	r.HandleFunc("/image/{path}/log", logHandler)
	// r.HandleFunc("/articles", ArticlesHandler)

	// [START request_logging]
	// Delegate all of the HTTP routing and serving to the gorilla/mux router.
	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	// [END request_logging]

	log.Fatal(http.ListenAndServe(":8080", nil))
}
