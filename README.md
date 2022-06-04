# freya

Freya is lightweight golang HTTP rate limiter based on ip which uses Redis ( for now ) as cache storage.

### install :

```bash
go get github.com/coci/freya
```

### usage :

#### method 1:

```go
package main

import (
	"net/http"
	"time"
	"fmt"

	"github.com/coci/freya"
	"github.com/coci/freya/cache"
)

func okHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {

	// request : 10
	// duration : 60s
	// cache : redis
	// it create rate limit 10 request per minutes and store rate information on redis
	lm := freya.NewLimiter(10, 60*time.Second, "redis")

	// config redis
	lm.Cache = cache.Redis{Host: "localhost", Port: 6379, Db: 0}

	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	http.ListenAndServe(":4000", lm.HandlerMuxMiddleware(mux))
}

```

#### method 2:

```go
package main

import (
	"net/http"
	"time"
	"fmt"

	"github.com/coci/freya"
	"github.com/coci/freya/cache"
)

func okHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {

	// request : 10
	// duration : 60s
	// cache : redis
	// it create rate limit 10 request per minutes and store rate information on redis
	lm := freya.NewLimiter(10, 60*time.Second, "redis")

	// config redis
	lm.Cache = cache.Redis{Host: "localhost", Port: 6379, Db: 0}

	mux := http.NewServeMux()
	mux.HandleFunc("/", lm.HandlerFuncMiddleware(okHandler))

	http.ListenAndServe(":4000", mux)
}

```

# Recommendation :

*note* : use "redis" as cache because redis has built-in expire time for key, So we don't need infinite loop in our app to
cleaning-up ip address's limitations.


# Todo :

- [ ] add Memcached cache handler

- [ ] add in-memory cache handler

- [ ] add database cache handler


# Why :

1- I wrote this because at the time I didn't find proper package that answer my question

2- for get hand dirty with Go


# What Freya stands from ?

I do love Norse , I always choose name and get inspiration from Norse . Freya in norse stands from :

Freya – Norse Mythology, Norse goddess of love, beauty, war, and death
