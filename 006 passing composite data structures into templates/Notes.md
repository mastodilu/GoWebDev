# Notes

## Come passare composite data structures ad un template

### Slice

script

```Go
myData := []string{"one", "two", "three"}
err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", myData)
```

template

```gohtml
<ul>
    {{range .}}
    <li>{{.}}</li>
    {{end}}
</ul>
```

output

```html
<ul>
    <li>one</li>
    <li>two</li>
    <li>three</li>
</ul>
```

---

Ad ogni iterazione di `{{range .}} `

```gohtml
{{range .}} // <-- range over the slice
    {{.}}   // <-- current field of the slice
{{end}}
```

`{{.}}` contiene il valore del dato all'iterazione corrente.

---

### Map

```gohtml
{{range $key, $value := .}}
    <li>{{$key}} - {{$value}}</li>
{{end}}
```

Con le mappe funziona esattamente allo stesso identico modo.

**NB**: `$key`, `$value` sono due nomi scelti da me, potevo usare `$cik`, `$ciak`

Questi due sono equivalenti perchè entrambi restituiscono il valore associato alla chiave

- `{{range .}}`
- `{{range $val := .}}`

---

### Struct

Per accedere al campo della struct si usa la sintassi `{{.nome_campo}}`

script

```Go
//struct (struct anonima)
mystruct := struct {
    Name string
    Age  int
}{
    Name: "matteo",
    Age:  25,
}
err = tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", mystruct)
```

template

```gohtml
<div>
    <p>Name: {{.Name}}</p>
    <p>Age: {{.Age}}</p>
</div>
```

output

```html
<div>
    <p>Name: matteo</p>
    <p>Age: 25</p>
</div>
```
