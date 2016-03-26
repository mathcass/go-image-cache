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

Once you have it, you can be able run it locally via the command
`go-image-cache`. That will serve via port `8080` so you will be able to access
it at [`http://localhost:8080`](http://localhost:8080).

The server will automatically create new links at `http://localhost/:label` that
should be tracked in the `hits.db` sqlite3 database. Logs of the activity are
available by visiting `http://localhost/:label/log`. 
