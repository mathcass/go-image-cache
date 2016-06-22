# Go Image Cache

This is a tutorial Golang project that implements a simple version of an open
beacon that's usable on websites or emails. Heavily based on
[`image-cache-logger`](https://github.com/kale/image-cache-logger) and
[`go`](https://github.com/kellegous/go). 

# Running

Get by running:

```
go get github.com/mathcass/go-image-cache
```

To run the code under an environment supported by Google App Engine, you'll
first need to download the
[GAE SDK](https://cloud.google.com/appengine/downloads) for Golang. 

Once you have that downloaded and in your `PATH`, run:

    make serve

or 

    goapp serve
    
    
to run a local development serving from `localhost:8080`
