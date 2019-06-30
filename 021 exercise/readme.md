# readme

## Esercizio

Stampa l'url della request.

### Soluzione

```Go
func handle(conn net.Conn) {
	// read
	scanner := bufio.NewScanner(conn)
	lineIndex := 0
	var host string

	for scanner.Scan() {
		if lineIndex == 1 {
			host = strings.Fields(scanner.Text())[1] // 💥 url della request
		}
		fmt.Printf("%v\n\r", scanner.Text())

		lineIndex++
		if scanner.Text() == "" {
			// fine header, esci dal ciclo
			break
		}
	}
	fmt.Printf("### host: %v\n\n", host)
}
```
