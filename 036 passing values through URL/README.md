# Passaggio di dati

Per inviare dati dal client al server ci sono due modi:

1. via **form**: metodo `POST` nel server
2. via **URL**: metodo `GET` nel server

## URL

![url explanation](img/URL.png)

- I parametri sono passati come `nome=valore`
- cominciano con il simbolo query `?`
- e si concatenano tra di loro usando `&`

Ad esempio: `http://localhost:8080/index?name=matteo&surname=dilu`

## Form

Dato un form

```Go
<form action="/" method="post">
    <input type="text" name="text" id="text">
    <input type="submit" value="OK">
</form>
```

I dati vengono passati con il comando submit e salvati in variabili il nome è quello dell'attributo `name`.

## FormValue

Il `(*http.Request) FormValue(string)` permette di recuperare la variabile ricevuta di nome `name`.

```Go
func home(w http.ResponseWriter, r *http.Request) {
    heSaid := r.FormValue("text")
    tpl.Execute(w, heSaid)
}
```

Con questa funzione è allo stesso modo possibile accedere ai valori passati come *GET*.
