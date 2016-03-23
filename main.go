package main

import (
	"github.com/mathcass/go-image-cache/db"
	"github.com/mathcass/go-image-cache/web"
)

func main() {
	db.InitializeDb()
	web.Serve()
}
