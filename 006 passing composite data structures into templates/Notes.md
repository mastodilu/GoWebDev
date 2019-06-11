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

---

### Slice di struct

Per accedere al campo della struct si usa la sintassi `{{.nome_campo}}`

#### script

```Go
//struct (struct anonima)
//slice di struct (anonime)
slicestruct := []struct {
    Name string
    Age  int
}{
    {"matteo", 25},
    {"alberto", 28},
    {"pollo", 12},
}
tpl.ExecuteTemplate(os.Stdout, "sliceStruct.gohtml", slicestruct)
mylib.IfErrThenLogFatal(err, "can't execute sliceStruct template")
```

#### template

L'output di `{{range .}}` diventa l'input per comporre i vari `{{.Name}}` e `{{.Age}}`

```gohtml
{{range .}}
    <div>
        <p>Name: {{.Name}}</p>
        <p>Age: {{.Age}}</p>
    </div>
{{end}}
```

#### output

```html
<div>
    <p>Name: matteo</p>
    <p>Age: 25</p>
</div>

<div>
    <p>Name: alberto</p>
    <p>Age: 28</p>
</div>

<div>
    <p>Name: pollo</p>
    <p>Age: 12</p>
</div>
```

E' possibile dare un nome alla variabile sul quale si cicla usando la sintassi:
`{{range $nome_var := .}}`

```Gohtml
{{range $shish := .}}
    <div>
        <p>Name: {{$nome_var.Name}}</p>
        <p>Age: {{$nome_var.Age}}</p>
    </div>
{{end}}
```