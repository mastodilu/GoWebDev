# Notes

E' possibile assegnare il dato passato come argomento in `ExecuteTemplate` ad una variabile specificata nel template.

```Go
myData := "wela" // <-- il dato passato
err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", myData)
```

```HTML
<body>
    {{$myh1 := .}}
    <h1>{{$myh1}}</h1>
</body>
```

L'output è:

```HTML
<h1>wela</h1>
```