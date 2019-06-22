# Notes

## Passing data into nested templates

Per passare i dati ai template basta includere un template con questa sintassi:<br>
`{{template "nome_template" .}}`

Il punto finale indica di passare il dato corrente al template chiamato.

### Script

```Go
people := []People{
    {
        "matteo", "dilu",
        Classes: []ClassScore{ {"Maths", 9},{"Science", 8},{"English", 7}, },
    }, 
    ...
    ...
}
err := tpl.ExecuteTemplate(out, "main.gohtml", people)
```

### Main template

```Gohtml
<div>{{template "template_1" . }}</div>
```

### Template 1

```Gohtml
{{define "template_1"}}
    <h2>Template 1</h2>
    {{range $person := .}}
        <h3>{{$person.Name}} {{$person.Surname}}</h3>
        <table>
            <tr>
                <th>Class</th>
                <th>Score</th>
            </tr>
            {{range $Class := $person.Classes}}
                <tr>
                    <td>{{$Class.Class}}</td>
                    <td>{{$Class.Score}}</td>
                </tr>
            {{end}}
        </table>
    {{end}}
{{end}}
```

## Chiamare metodi

E' possibile chiamare i metodi sulle struct passate al template.

### Script

```Go
type Person struct {
    name string
}

func (p *Person) SayHi() string {
    return "Hi from " + p.Name()
}

...

err := tpl.ExecuteTemplate(out, "main.gohtml", people)
```

### Template

``` gohtml
...
{{range $class := $person.Classes}}
    {{$class.Class}} {{.SayHi}}
    {{$class.Score}}
{{end}}
```

Il metodo **`SayHi` non è stato passato nella funcmap**, ma è disponibile perchè è pubblico e legato al tipo `People`.
