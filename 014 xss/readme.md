# Readme

## HTML template

[package html/template](https://golang.org/pkg/html/template/)

Fornisce le funzionalità del `package text/template` e qualcosa in più.

Permette di usare gli escape characters per bloccare i caratteri *unsafe* del web a seconda del context. Questo **evita il cross site scripting**.

## XSS

### Script

```go
import "text/template"
    ...
data := `<script>alert("wela")</script>`
err := tpl.Execute(os.Stdout, data)
```

### Template

```gohtml
<body>
    <h1>XSS con text/template</h1>
    {{.}}
</body>
```

### Output

```html
<body>
    <h1>XSS con text/template</h1>
    <script>alert("wela")</script>
</body>
```

## No XSS grazie al character escaping

### Script

```go
import "html/template"
    ...
data := `<script>alert("wela")</script>`
err := tpl.Execute(os.Stdout, data)
```

### Template

```gohtml
<body>
    <h1>No XSS con html/template</h1>
    {{.}}
</body>
```

### Output

```html
<body>
    <h1>No XSS con html/template</h1>
    &lt;script&gt;alert(&#34;wela&#34;)&lt;/script&gt;
</body>
```
