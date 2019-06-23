# Readme

La sintassi per mostrare dei dati è `{{.}}` nel template.

```HTML
<head>
    <meta charset="utf-8">
    <title>{{.}}</title>
</head>
```

Dentro a `{{.}}`, il punto `.` rappresenta il valore corrente dei dati in un preciso istante.

Si può passare un solo dato alla volta (una sola variabile), ma il dato può essere di tipo aggregato, ad esempio una mappa di array.
