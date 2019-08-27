# http-p2s
Narrow the HTTP request traffic with a middleware

## Usage

```go
package main

import (
	p2s "github.com/rosbit/http-p2s"
	"github.com/urfave/negroni"
	"net/http"
	"fmt"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n := negroni.New()
	n.UseFunc(p2s.NarrowHttpRequest(5))  // use this middleware before other middleware
	n.UseHandler(mux)
	http.ListenAndServe(":8080", n)
}
```

## Status
The package is fully tested.

## Contribution
Pull requests are welcome! Also, if you want to discuss something send a pull request with proposal and changes.

