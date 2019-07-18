# bcrypt

Il package `bcrypt` permette di criptare e decriptare *cose*.

> `go get golang.org/x/crypto/bcrypt`

Va bene per criptare password.

Un esempio di gestione delle registrazioni degli utenti criptando anche la password è:

## Regista un nuovo utente

Per registrare un nuovo utente e' necessario avere ID per identificarlo e password per autenticarlo. Entrambi i campi vengono letti tramite metodo `post` del form della pagina di autenticazione, ma la password va gestita con cautela in quanto non deve essere salvata.

Con la funzione `func GenerateFromPassword(password []byte, cost int) ([]byte, error)` viene generato un hash univoco usando la password fornita che puo' essere salvato nel DB per quell'utente.

```Go
func registerUser(username, email string, password []byte) error {
    hashPwd, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
    if err != nil {
        return err
    }
    if err := db.Register(username, email, hashPwd); err != nil {
        return err
    }
    return nil
}
```

## Delete user

Per cancellare un utente (disiscriverlo) si cancella l'utente dal DB e si annulla la sua sessione rendendo negativo il campo `MaxAge` di `http.Cookie`.

```Go
func deleteUser(w http.ResponseWriter, r *http.Request) {
    if err := checkValidSession(r); err != nil {
        tpl.ExecuteTemplate(w, "error.gohtml", "non sei un utente registrato")
        return
    }

    coo, _ := r.Cookie("email")
    email := coo.Value

    err := db.DeleteUser(email)
    if err != nil {
        http.Error(w, "Something went wrong", http.StatusNotFound)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:   "email",
        Value:  "",
        MaxAge: -1,
    })
    tpl.ExecuteTemplate(w, "error.gohtml", "your user have been deleted from our database")
}
```

## Login

In fase di login e' necessario controllare il matching della password.

Per effettuare questa operazione in modo sicuro si usa la funzione `func CompareHashAndPassword(hashedPassword, password []byte) error` che confronta l'hash della nuova password con l'hash di quella salvata.

Se non c'e' matching allora si genera errore, altrimenti si crea la sessione settando il cookie di sessione.

```Go
func login(w http.ResponseWriter, r *http.Request) {
    email := r.FormValue("email")
    user, err := db.GetUser(email)
    if err != nil {
        tpl.ExecuteTemplate(w, "error.gohtml", "user not found")
        return
    }
    err = bcrypt.CompareHashAndPassword(user.HashPassword, []byte(r.FormValue("password")))
    if err != nil {
        tpl.ExecuteTemplate(w, "error.gohtml", "wrong password")
        return
    }
    http.SetCookie(w, &http.Cookie{
        Name:  "email",
        Value: email,
    })

    tpl.ExecuteTemplate(w, "home.gohtml", UEM{user.Username, email, user.Message})
}
```

## Logout

Per effettuare il logout di un utente basta cancellare il loro cookie e fare redirect ad una qualche pagina, ad esempio quella di registrazione.

```Go
func logout(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{
        Name:   "email",
        Value:  "",
        MaxAge: -1,
    })
    tpl.ExecuteTemplate(w, "error.gohtml", "you have logged out")
}
```

## Permessi

Per gestire i permessi si puo' aggiungere un campo alla struct utente in grado di memorizzare il ruolo dell'utente, ad esempio `IsAdmin bool`.

Un esempio:

```Go
func promoteToAdmin(w http.ResponseWriter, r *http.Request) {
    if err := checkValidSession(r); err != nil {
        tpl.ExecuteTemplate(w, "error.gohtml", "invalid session")
        return
    }

    email, _ := r.Cookie("email") // already checked for errors

    err := bcrypt.CompareHashAndPassword(adminPasswordHash, []byte(r.FormValue("password")))
    if err != nil {
        fmt.Fprintln(w, "wrong password")
        return
    }

    if err := db.Promote(email.Value); err != nil {
        log.Println(err)
        tpl.ExecuteTemplate(w, "error.gohtml", "you can't become an admin")
        return
    }

    fmt.Fprintln(w, "you're now an administrator")
}
```
