# Readme

## Pipeline

Il pipelining permette di concatenare le chiamate di funzioni nei template usando `|`.

### Template

```gohtml
<div> {{.}}</div>
<div> {{. | add1}}</div>
<div> {{. | add1 | add1}}</div>
<div> {{. | add1 | add1 | add1}}</div>
<div> {{. | add1 | add1 | add1 | subt1}}</div>
```

Un'alternativa a `<div> {{. | add1 | add1 | add1 | subt1}}</div>` è

```Gohtml
<div> {{add1 ( add1 ( add1 ( subt1 ( . ) ) ) ) }}</div>
```

### Output

```html
<div> 0</div>
<div> 1</div>
<div> 2</div>
<div> 3</div>
<div> 2</div>

<div> 2</div>
```