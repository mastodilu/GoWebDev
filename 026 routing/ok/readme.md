# readme

## Metodo decente per gestire i vari routes

Con

```Go
http.ListenAndServe(":8080", nil)
```

si specifica di usare ``http.DefaultServeMux`` al quale è possibile far gestire determinati path usando la funzione

```Go
// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
```

Il parametro `handler` vuole una funzione con la seguente firma:

```Go
func(ResponseWriter, *Request)
```

e tramite

```Go
func cat(res http.ResponseWriter, req *http.Request) {}
func dog(res http.ResponseWriter, req *http.Request) {}
func home(res http.ResponseWriter, req *http.Request) {}
```

abbiamo gli handler desiderati.

### Script

```Go
package main

import (
	"fmt"
	"net/http"
)

func cat(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Handles %v</h1>\n%v\n", "/catOnly", req.URL.Path)
}

func dog(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Handles %v</h1>\n%v\n", "/dog/", req.URL.Path)
}

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Handles %v</h1>\n%v\n", "/", req.URL.Path)
}

func main() {
	fmt.Println("server listening on port :8080")

	http.HandleFunc("/", home)
	http.HandleFunc("/catOnly", cat)
	http.HandleFunc("/dog/", dog)

	http.ListenAndServe(":8080", nil)
}
```