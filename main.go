package main

import (
	"net/http"
	"strings"
)

type appRequest struct {
	request    *http.Request
	ImagePath  string
	IsLog      bool
	IsRootPath bool
}

func newImageAppRequest(r *http.Request) appRequest {
	const pathSep = "/"

	path := string(r.URL.Path)
	splits := strings.Split(path, pathSep)

	isLog := false
	if splits[len(splits)-1] == "log" {
		isLog = true
		path = strings.Join(splits[:len(splits)-1], pathSep)
	}

	isRootPath := false
	if path == "/" {
		isRootPath = true
	} else {
		if path[len(path)-1] != '/' {
			path = path + "/"
		}
	}

	return appRequest{
		request:    r,
		ImagePath:  path,
		IsLog:      isLog,
		IsRootPath: isRootPath,
	}
}

// Handle root application, multiplexing it out to the options of either:
// * root page
// * log page
// * image page
func handle(w http.ResponseWriter, r *http.Request) {
	ar := newImageAppRequest(r)

	// For a simple application with 3 different request types
	if ar.IsRootPath {
		root(w, r)
	} else if ar.IsLog {
		imageLog(w, r, ar.ImagePath)
	} else {
		image(w, r, ar.ImagePath)
	}
}

func init() {
	http.HandleFunc("/", handle)
}

// For vm enabled applications, use appengine.Main
// func main() {
// 	appengine.Main()
// }
