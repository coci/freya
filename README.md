# freya
Freya is lightweight golang HTTP rate limiter based on ip which used Redis ( for now ) as cache storage.

### install :
```bash
go get -U github.com/coci/freya
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
	lm := freya.NewLimiter(10, 60 * time.Second,"redis")
	
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
	lm := freya.NewLimiter(10, 60 * time.Second,"redis")
	
	// config redis
	lm.Cache = cache.Redis{Host: "localhost", Port: 6379, Db: 0}

	mux := http.NewServeMux()
	mux.HandleFunc("/", lm.HandlerFuncMiddleware(okHandler))

	http.ListenAndServe(":4000", mux)
}

```

# Todo :
- [ ] add Memcached cache handler

- [ ] add in-memory cache handler

- [ ] add database cache handler

# What Freya stands from ?
i do love Norse , i always choose name and get inspiration from Norse . Ferya in norse stands from :

Freya – Norse Mythology, Norse goddess of love, beauty, war, and death
