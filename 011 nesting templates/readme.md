# Readme

## Nesting templates

I template annidati permettono di rendere il codice modulare.

La sintassi è:

```Gohtml
{{define <nome_template>}}
... contenuto ...
{{end}}
```

In uno stesso template è anche possibile definire più template:

```Gohtml
{{define <nome_1>}}
... contenuto ...
{{end}}

{{define <nome_2>}}
... contenuto ...
{{end}}
```
