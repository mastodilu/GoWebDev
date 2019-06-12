# Notes

Tenendo a mente il modello **MVC**

- **M, model**: i dati gestiti dal backend
- **V, view**: quello che viene visto dall'utilizzatore
- **C, controller**: quello che collega M e V

è possibile far eseguire delle funzioni nella view, quindi nel template, senza rompere questo pattern, ad esempio per meglio visualizzare/impaginare i dati.

Si possono passare funzioni predefinite, visibili globalmente, e funzioni definite nel progetto.

Per dare le funzioni al template si usa il tipo `FuncMap`.

## FuncMap

[FuncMap documentation](https://golang.org/pkg/text/template/#FuncMap)

```Go
type FuncMap map[string]interface{}
```

Definisce un mapping tra nomi e funzioni.

- Ogni funzione deve avere uno oppure due return value (due quando restituisce anche un errore)
- ci sono delle funzioni built in da usare definite nel package `template`

### Tipo FuncMap

```Go
// funzioni da passare al template
var funcmap = template.FuncMap{
    "sayhi":      sayHi,
    "toupper":    strings.ToUpper,
    "firstthree": firstThree,
}

// definizione delle funzioni

func sayHi() string {
    return "hi from sayHi() function"
}

func firstThree(s string) (string, error) {
    r := []rune(strings.TrimSpace(s))
    if len(r) < 3 {
        return "", fmt.Errorf("string '%s' is too short", s)
    }
    return string(r[:3]), nil
}
```

### init()

Non si usa più il pattern Must-ParseGlob-Execute:

- Inizializza un template vuoto, gli passa la mappa di funzioni, parsa il template e glielo applica.

```Go
// 💥 questa variabile va dichiarata prima di eseguire .Funcs(funcmap).
// altrimenti si incappa in nil pointers durante l'esecuzione
var funcmap = template.FuncMap{
    "sayhi":      sayHi,
    "sum":        sum,
    "toupper":    strings.ToUpper,
    "firstthree": firstThree,
}

func init() {
    tpl, err := template.New("").Funcs(funcmap).ParseFiles(".templates/functions.gohtml")
    tpl = template.Must(tpl, err)
}

func sum(n1, n2 int) int {
    return n1 + n2
}

func sayHi() string {
    return "hi from sayHi() function"
}

func firstThree(s string) (string, error) {
    r := []rune(strings.TrimSpace(s))
    if len(r) < 3 {
        return "", fmt.Errorf("string '%s' is too short", s)
    }
    return string(r[:3]), nil
}
```

#### template

```Gohtml
<h1>FuncMap hands on</h1>
    {{range .}}
        <p>{{firstthree (toupper .)}}</p>
    {{end}}
    <p>{{sayhi}}</p>
    <p>{{double 10}}</p>
    <p>aBC --> {{toupper "aBC"}}</p>
```

`{{firstthree (toupper .)}}`

- chiama toupper .
- chiama firstthree sull'output precedente

I parametri sono specificati senza parentesi.

#### output

```html
<h1>FuncMap hands on</h1>
    <p>CIA</p>
    <p>MIA</p>
    <p>BAU</p>
    <p>hi from sayHi() function</p>
    <p>1+2 = 3</p>
    <p>aBC --> ABC</p>
```

Quando viene parsato il template la mappa di funzioni deve già essere attaccata al template, altrimenti il codice fallisce perchè nel testo viene trovata una funzione "non ancora definita":

OK

```Go
tpl, err := template.New("").Funcs(funcmap).ParseFiles(".templates/functions.gohtml")
```

Errore

```Go
tpl, err := template.New("").ParseFiles(".templates/functions.gohtml").Funcs(funcmap)
```

### Parsing di una stringa

Per creare un template partendo da una stringa di usa questa sintassi:

```Go
myString := "bla bla bla..."
tpl, err := template.New("<NomeDelTemplate>").Funcs(funcmap).Parse(myString)
```

Dentro a `New("<Nome>")` viene specificato il nome col quale ci si riferisce al template perchè una stringa, a differenza di un file, non ha un base name utilizzabile!
