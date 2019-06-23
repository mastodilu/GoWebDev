# Readme

## ParseGlob(..)

Parsa ogni file specificato usando regex invece di specificare a mano il nome dei template con `template.ParseFiles(..)`

```Go
const parseWholeFolder = "templates/*"
const parseOnlyHTML = "templates/*.gohtml"
tpl, err := template.ParseGlob(parseOnlyHTML)
mylib.IfErrThenPanic(err, "err in ParseGlob")
fmt.Printf("%T, %v\n", tpl, tpl)
```

Per applicare il template desiderato si usa ancora il metodo `tpl.ExecuteTemplate()`

```Go
tpl.ExecuteTemplate(os.Stdout, "sample.gnegnegne", nil)
```

Per stampare l'elenco di template:

```Go
// chiamata della funzione
printAllTemplates(tpl)

// funzione che stampa i nomi dei template contenuti
func printAllTemplates(t *template.Template) {
	for _, tpl := range t.Templates() {
		fmt.Println(tpl.Name())
	}
}
```