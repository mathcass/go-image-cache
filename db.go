package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

const RequestKey = "Request"

type Request struct {
	Path      string
	UserAgent string
	Date      time.Time
}

func requestDataKey(ctx context.Context) *datastore.Key {
	return datastore.NewKey(ctx, RequestKey, "default_requests", 0, nil)
}

func generateFakeData(total int) []Request {
	fakeUA := [3]string{"Mozilla", "Chrome", "Bot"}

	fakeRequests := make([]Request, 0)

	for i := 0; i < total; i++ {
		fakeRequest := Request{Path: "blah", UserAgent: fakeUA[rand.Intn(len(fakeUA))], Date: time.Now()}
		fakeRequests = append(fakeRequests, fakeRequest)
	}

	return fakeRequests
}

func loadFakeData(ctx context.Context) {
	d := generateFakeData(5)
	key := datastore.NewIncompleteKey(ctx, RequestKey, nil)
	for _, item := range d {
		if _, err := datastore.Put(ctx, key, &item); err != nil {
			log.Fatalln(err)
		}
	}
}

// Given context & request, saves to datastore
func savePathRequest(ctx context.Context, r *http.Request) {

	ar := newImageAppRequest(r)
	pr := Request{
		Path:      ar.ImagePath,
		UserAgent: r.UserAgent(),
		Date:      time.Now(),
	}

	key := datastore.NewIncompleteKey(ctx, RequestKey, nil)
	if _, err := datastore.Put(ctx, key, &pr); err != nil {
		log.Fatalln(err)
	}
}
