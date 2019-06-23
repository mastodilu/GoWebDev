# Readme

- si inizializza il progetto con `init()`
- con `ParseGlob` si salvano tutti i template in una variabile *contenitore di template*, cioè `*template.Template`
- con `Must()` si gestisce l'eventuale errore

## init()

init è una funzione "reservata" che viene sempre chiamata implicitamente avviando un progetto.

Non restituisce valori e non prende parametri.

init() viene lanciato soltanto la prima volta che si avvia il progetto, e non può essere più richiamata.

```Go
func init() {
	const parseWholeFolder = "templates/*"
	const parseOnlyHTML = "templates/*.gohtml"

	//controlla che il template parsato non sia null
	tpl = template.Must(template.ParseGlob(parseWholeFolder))
}
```

## template.Must()

Implementazione della funzione `Must()`:

```Go
func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}
```

Must esegue l'error checking, infatti prende un puntatore a template e un errore e chiama panic() se c'è errore.

Questa funzione si assicura che template.ParseGlob non fallisca. L'eventuale errore restituito viene gestito da `Must()`.
Per quello è bene usarlo nella funzione `init()`

```Go
tpl = template.Must(template.ParseGlob(parseWholeFolder))
```