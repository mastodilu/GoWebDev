# Store session info

Per gestire le sessioni in Go si usano i cookie.

Se l'utente non permette di scrivere i cookie allora è possibile concatenare ad ogni URL il parametro che altrimenti sarebbe stato scritto nel cookie, tipo un sessionID.

In quel caso è necessario utilizzare `https`.

Ad ogni richiesta è necessario controllare che il cookie di sessione (o il parametro) sia nell'url. Se non è presente ha senso fare una redirect (`http.Redirect`) verso un'altra pagina, tipo la pagina di login.
