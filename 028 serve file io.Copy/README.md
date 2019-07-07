# readme

## Serve a file with `io.Copy`

> Usare questa funzione non è il metodo giusto, ma va beh.

La funzione `io.Copy` è così implementata:

```Go
// Copy copies from src to dst until either EOF is reached
// on src or an error occurs. It returns the number of bytes
// copied and the first error encountered while copying, if any.
//
// A successful Copy returns err == nil, not err == EOF.
// Because Copy is defined to read from src until EOF, it does
// not treat an EOF from Read as an error to be reported.
//
// If src implements the WriterTo interface,
// the copy is implemented by calling src.WriteTo(dst).
// Otherwise, if dst implements the ReaderFrom interface,
// the copy is implemented by calling dst.ReadFrom(src).
func Copy(dst Writer, src Reader) (written int64, err error) {
	return copyBuffer(dst, src, nil)
}
```

### Script

```Go
package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func home(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, nil)
}

func img(res http.ResponseWriter, req *http.Request) {
	file, err := os.Open("img.png")
	if err != nil {
		http.Error(res, "File not Found", 404)
		return
	}
	defer file.Close()
	io.Copy(res, file)
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

Questa funzione è stata implementata scrivendo nella response con `io.Copy(res, file)`.
