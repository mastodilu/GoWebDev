# readme

## Serve a file with `http.ServeFile`

> Usare questa funzione non è il metodo giusto, ma va beh.

La funzione `http.ServeFile` è così implementata:

```Go
// ServeFile replies to the request with the contents of the named
// file or directory.
//
// If the provided file or directory name is a relative path, it is
// interpreted relative to the current directory and may ascend to
// parent directories. If the provided name is constructed from user
// input, it should be sanitized before calling ServeFile.
//
// As a precaution, ServeFile will reject requests where r.URL.Path
// contains a ".." path element; this protects against callers who
// might unsafely use filepath.Join on r.URL.Path without sanitizing
// it and then use that filepath.Join result as the name argument.
//
// As another special case, ServeFile redirects any request where r.URL.Path
// ends in "/index.html" to the same path, without the final
// "index.html". To avoid such redirects either modify the path or
// use ServeContent.
//
// Outside of those two special cases, ServeFile does not use
// r.URL.Path for selecting the file or directory to serve; only the
// file or directory provided in the name argument is used.
func ServeFile(w ResponseWriter, r *Request, name string) {
	if containsDotDot(r.URL.Path) {
		// Too many programs use r.URL.Path to construct the argument to
		// serveFile. Reject the request under the assumption that happened
		// here and ".." may not be wanted.
		// Note that name might not contain "..", for example if code (still
		// incorrectly) used filepath.Join(myDir, r.URL.Path).
		Error(w, "invalid URL path", StatusBadRequest)
		return
	}
	dir, file := filepath.Split(name)
	serveFile(w, r, Dir(dir), file, false)
}
```

Vuole i seguenti parametri:

- response
- request
- file name

### Script

```Go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, nil)
}

func img(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "img.png")
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var tpl *template.Template

func main() {
	fmt.Println("Listening on port :8080")

	http.HandleFunc("/", home)
	http.HandleFunc("/img.png", img)

	http.ListenAndServe(":8080", nil)
}
```

### Template

```Gohtml
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=>, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <h1>BatGopher</h1>
    <img src="img.png" alt="BatGopher">
</body>
</html>
```

Alla pagina `localhost` viene inviato il template che contiene la riga `<img src="img.png" alt="BatGopher">`.

`img.png` è un route gestibile con `http.HandleFunc("/img.png", img)` tramite la sua apposita funzione.

Questa funzione è stata implementata scrivendo nella response con `http.ServeFile(res, req, "img.png")`.
