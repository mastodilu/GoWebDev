# Notes

## Index

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