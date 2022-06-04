# freya
Golang Http Rate Limiter

### install :
```bash
go get -U github.com/coci/freya
```

### usage :

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


# what Freya stands from ?
i do love Norse , i always choose name and get inspiration from Norse . Ferya in norse stands from :

Freya – Norse Mythology, Norse goddess of love, beauty, war, and death
