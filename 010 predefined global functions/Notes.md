
# Notes

## Index

- [Notes](#notes)
  - [Index](#index)
    - [Script](#script)
    - [Template](#template)
    - [Output](#output)
  - [Nel caso di un array contenuto in una struct](#nel-caso-di-un-array-contenuto-in-una-struct)
    - [Script](#script-1)
    - [Template](#template-1)
  - [Range su uno slice contenuto in una struct](#range-su-uno-slice-contenuto-in-una-struct)
    - [Script](#script-2)
    - [Template](#template-2)
  - [If](#if)
    - [Template](#template-3)
    - [Output](#output-1)

Usando la keyword `index` è possibile accedere ai campi di uno slice:

### Script

```Go
words := []string{"zero", "one", "two", "three"}
err := tpl.ExecuteTemplate(os.Stdout, "globfunc.gohtml", words)
```

### Template

```Gohtml
<div>{{index . 1}}</div>
<div>{{index . 3}}</div>
<div>{{index . 0}}</div>
<div>{{index . 2}}</div>
```

### Output

```html
<div>one</div>
<div>three</div>
<div>zero</div>
<div>two</div>
```

## Nel caso di un array contenuto in una struct

### Script

```Go
myData := struct {
    Name string
    Data []string
}{
    Name: "matteo",
    Data: []string{"zero", "one", "two", "three"},
}
err := tpl.ExecuteTemplate(os.Stdout, "globfunc.gohtml", myData)
```

### Template

Si accede ai campi **esportati** così:

```gohtml
<div>{{.Name}}</div>

<h2>Accedi ad un campo dell'array usando index</h2>

<div>{{index .Data 1}}</div>
<div>{{index .Data 3}}</div>
<div>{{index .Data 0}}</div>
<div>{{index .Data 2}}</div>
```

## Range su uno slice contenuto in una struct

### Script

```Go
myData := struct {
    Name string
    Data []string
}{
    Name: "matteo",
    Data: []string{"zero", "one", "two", "three"},
}
err := tpl.ExecuteTemplate(os.Stdout, "globfunc.gohtml", myData)
```

### Template

```Gohtml
{{range .Data}}
<p>{{.}}</p>
{{end}}
```
---

---

## If

### Template

```Gohtml
{{if .}}
    <div>Non hai passato un nil data</div>
{{end}}

<h3>esiste "two"</h3>
{{range .Data}}
    {{if eq . "two"}}
        <div>"two" trovato</div>
    {{else}}
        <div>{{.}} - "two" non trovato</div>
    {{end}}
{{end}}
```

`{{if eq . "two"}}` controlla se `dato == "two"`
perchè `eq` è la funzione che rappresenta `==` .

### Output

```html
<div>Non hai passato un nil data</div>
<h3>esiste "two"</h3>
<div>zero - "two" non trovato</div>
<div>one - "two" non trovato</div>
<div>"two" trovato</div>
<div>three - "two" non trovato</div>
```

Altre funzioni globali sono:

- **eq**: Returns the boolean truth of arg1 == arg2
- **ne**: Returns the boolean truth of arg1 != arg2
- **lt**: Returns the boolean truth of arg1 < arg2
- **le**: Returns the boolean truth of arg1 <= arg2
- **gt**: Returns the boolean truth of arg1 > arg2
- **ge**: Returns the boolean truth of arg1 >= arg2
