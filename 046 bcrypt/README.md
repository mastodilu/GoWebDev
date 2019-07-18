# bcrypt

Il package `bcrypt` permette di criptare e decriptare *cose*.

> `go get golang.org/x/crypto/bcrypt`

Va bene per criptare password.

Un esempio di gestione delle registrazioni degli utenti criptando anche la password è:

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

// TODO: consenti l'accesso ad un utente fornendo login e logout. Usa la funzione per comparare la password fornita con l'hash salvato associato ad un utente.
