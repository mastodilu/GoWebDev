# Notes

## ParseGlob(..)

Parsa ogni file specificato usando regex invece di specificare a mano il nome dei template con `template.ParseFiles(..)`

```Go
const parseWholeFolder = "templates/*"
const parseOnlyHTML = "templates/*.gohtml"
tpl, err := template.ParseGlob(parseOnlyHTML)
mylib.IfErrThenPanic(err, "err in ParseGlob")
fmt.Printf("%T, %v\n", tpl, tpl)
```