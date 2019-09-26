# Context

Il [context](https://golang.org/pkg/context/) è come un contenitore che viene associato ad una web request al quale possono essere aggiunge delle informazioni, ad esempio sulla sessione.

E' un contenitore di informazioni strettamente legate ad una determinata richiesta.

Le richieste entranti in un server dovrebbero creare un Context.
