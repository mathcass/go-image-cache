package main

import (
	"encoding/base64"
	"html/template"
	"io"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

const base64GifPixel = "R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs="

func imageLog(w http.ResponseWriter, r *http.Request, path string) {
	// TODO: Implementation
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Got log request for path: %v", path)

	q := datastore.NewQuery(RequestKey).
		Filter("Path =", path)
	var rr []*Request
	if _, err := q.GetAll(ctx, &rr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf(ctx, "Getall: %v", err)
		return
	}

	data := struct{ Requests []*Request }{Requests: rr}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tpl.ExecuteTemplate(w, "log.html", data); err != nil {
		log.Errorf(ctx, "%v", err)
	}
}

func image(w http.ResponseWriter, r *http.Request, path string) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Got image request for path: %v", path)
	savePathRequest(ctx, r)

	w.Header().Set("Content-Type", "image/gif")
	output, _ := base64.StdEncoding.DecodeString(base64GifPixel)
	io.WriteString(w, string(output))
}

func root(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Got root request")

	// savePathRequest(ctx, r)

	q := datastore.NewQuery(RequestKey).Limit(10)
	var rr []*Request
	if _, err := q.GetAll(ctx, &rr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf(ctx, "Getall: %v", err)
		return
	}

	data := struct{ Requests []*Request }{Requests: rr}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tpl.ExecuteTemplate(w, "root.html", data); err != nil {
		log.Errorf(ctx, "%v", err)
	}
}
