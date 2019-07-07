# readme

## Serve a file with `http.ServeContent`

> Usare questa funzione non è il metodo giusto, ma va beh.

La funzione `http.ServeContent` è così implementata:

```Go
// ServeContent replies to the request using the content in the
// provided ReadSeeker. The main benefit of ServeContent over io.Copy
// is that it handles Range requests properly, sets the MIME type, and
// handles If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since,
// and If-Range requests.
//
// If the response's Content-Type header is not set, ServeContent
// first tries to deduce the type from name's file extension and,
// if that fails, falls back to reading the first block of the content
// and passing it to DetectContentType.
// The name is otherwise unused; in particular it can be empty and is
// never sent in the response.
//
// If modtime is not the zero time or Unix epoch, ServeContent
// includes it in a Last-Modified header in the response. If the
// request includes an If-Modified-Since header, ServeContent uses
// modtime to decide whether the content needs to be sent at all.
//
// The content's Seek method must work: ServeContent uses
// a seek to the end of the content to determine its size.
//
// If the caller has set w's ETag header formatted per RFC 7232, section 2.3,
// ServeContent uses it to handle requests using If-Match, If-None-Match, or If-Range.
//
// Note that *os.File implements the io.ReadSeeker interface.
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker) {
	sizeFunc := func() (int64, error) {
		size, err := content.Seek(0, io.SeekEnd)
		if err != nil {
			return 0, errSeeker
		}
		_, err = content.Seek(0, io.SeekStart)
		if err != nil {
			return 0, errSeeker
		}
		return size, nil
	}
	serveContent(w, req, name, modtime, sizeFunc, content)
}
```

Vuole i seguenti parametri:

- response
- request
- nome del file
- data di modifica
- il file da inviare

### Script

```Go
package main

import (
	"fmt"
	"html/template"
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

	stat, err := file.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}

	http.ServeContent(res, req, file.Name(), stat.ModTime(), file)
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
`file.Stat()` restituisce `FileInfo`:

```Go
// A FileInfo describes a file and is returned by Stat and Lstat.
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
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

Questa funzione è stata implementata scrivendo nella response con `http.ServeContent(res, req, file.Name(), stat.ModTime(), file)`.
