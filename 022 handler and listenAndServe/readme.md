# readme

## Handler interface

Il tipo Handler è una interfaccia che contiene il sequente metodo:


```Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Per implementare un handler basta che `qualcosa` implementi il metodo dell'interfaccia. Ad esempio:

```Go
type AkunaMatata bool

func (a AkunaMatata) ServeHTTP(http.ResponseWriter, *http.Request) {
	fmt.Println("ServeHTTP method")

}

func main() {
	var akmt AkunaMatata
	http.ListenAndServe("/hello", akmt)
}
```

