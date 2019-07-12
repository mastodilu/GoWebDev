# Traccia quante volte un utente accede al sito usando i cookie.

La logica del contare il numero di login di un utente l'ho implementata cosi':

```Go
// private legge lo username e cerca il cookie con lo stesso nome.
// Se esisteva gia' allora ne prende il valore e lo incrementa di 1,
// altrimenti ne crea uno nuovo e lo inizializza a 1.
func private(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	if username == "" {
		home(w, r)
		return
	}

	loginCounter, err := r.Cookie(username)
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  username,
			Value: "1",
		})
	} else {
		count, err := strconv.Atoi(loginCounter.Value)
		if err != nil {
			count = 0
		}

		http.SetCookie(w, &http.Cookie{
			Name:  username,
			Value: strconv.Itoa(count + 1),
		})
	}

	err = tpl.ExecuteTemplate(w, "private.gohtml", username)
	if err != nil {
		log.Fatal(err)
	}
}
```

Ogni utente si porta dietro un cookie col suo stesso nome che contiene il numero di accesi effettuati.