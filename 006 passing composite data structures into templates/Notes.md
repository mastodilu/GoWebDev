# Notes

## Come passare composite data structures ad un template

script

```Go
myData := []string{"one", "two", "three"}
err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", myData)
```

template

```html
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

Con le mappe funziona esattamente allo stesso identico modo.
