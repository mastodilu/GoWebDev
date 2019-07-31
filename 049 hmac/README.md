# HMAC

Usa una chiave segreta per creare un hash sicuro che permette di verificare un determinato input.

```Go
func getHMAC(s, key string) string {
    h := hmac.New(sha256.New, []byte(key))
    _, err := io.WriteString(h, s)
    if err != nil {
        log.Fatal(err)
    }
    return fmt.Sprintf("%x", h.Sum(nil))
}
```
