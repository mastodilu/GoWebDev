# write user file

Per scrivere un file quando viene ricevuta una richiesta POST si può fare così:

1. controlla che la richiesta sia POST
2. ottieni i campi dal body della request
3. componi un path in cui scrivere
4. scrivi
5. controlla gli errori

```Go
if r.Method == http.MethodPost {
    username := r.PostFormValue("username")
    text := r.PostFormValue("text")
    path := filepath.Join("./Users/", fmt.Sprintf("%s.txt", username))
    err := ioutil.WriteFile(path, []byte(text), 0644)
    ...
}
```

## Script

```Go
package main

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
    "path/filepath"
)

func home(w http.ResponseWriter, r *http.Request) {

    if r.Method == http.MethodPost {
        username := r.PostFormValue("username")
        text := r.PostFormValue("text")
        path := filepath.Join("./Users/", fmt.Sprintf("%s.txt", username))
        err := ioutil.WriteFile(path, []byte(text), 0644)
        if err != nil {
            log.Println(err)
            tpl.Execute(w, nil)
            return
        }

        tpl.Execute(w, true)
        return
    }
    tpl.Execute(w, nil)
}

func init() {
    tpl = template.Must(template.ParseFiles("templates/tpl.gohtml"))
}

var tpl *template.Template

func main() {

    http.HandleFunc("/", home)

    http.Handle("/favicon.ico", http.NotFoundHandler())

    http.ListenAndServe(":8080", nil)
}
```

## Template

```Gohtml
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Write file in user dir</title>
</head>
<body>
    <h1>write file in user dir</h1>
    {{if .}}
        <p>Saved!</p>
    {{end}}

    <form method="post">
        <input type="text" name="username" id="username" placeholder="username" ><br>
        <input type="text" name="text" id="text" placeholder="some text"><br>
        <input type="submit" value="Write to file">
    </form>

</body>
</html>
```
