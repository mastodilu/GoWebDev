# entype

E' un attributo del form html.

1. `application/x-www-form-urlencoded`: prende nomi e valori e li converte in una stringa del tipo `?nome=valore&...`. POST e GET stabiliscono se questa stringa viene passata tramite URL o tramite request body.
2. `multipart/form-data`: usato per quelle situazioni in cui il form contiene l'upload di un file.
3. `text/plain`: converte gli spazi in `+`. (usato per il debugging, non usarlo in produzione)
