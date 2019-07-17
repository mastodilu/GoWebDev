# Fake Database to store sessions

Per creare id univoci in tutto il www si usa il package `"github.com/satori/go.uuid"`.

Il metodo `uuid.NewV4()` permette di creare un ID senza fornire alcun parametro.

Questo valore può essere salvato in un database per identificare un utente in modo univoco.

Se l'utente non è identificato da un ID allora non ha fatto login e ad esempio si può forzare una redirect alla pagina di registrazione.

```Go
// check cookie
sessionID, err := r.Cookie("sessionID")
if err != nil {
    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return
}
```
