# Build a web server in Go using Docker

## Build the image

```docekr
docker build . -t hellogolang
```

## Run the container

```docker
docker run -p 8081:8081 -detached --rm hellogolang
```

Il container è attivo:

```powershell
λ docker ps

CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                    NAMES
e9e27ae15a5f        hellogolang         "/dist/main"        15 hours ago        Up 30 minutes       0.0.0.0:8081->8081/tcp   unruffled_gates
```

Visita [localhost:8081](http://localhost:8081/) per leggere:

```plaintext
Hello golang from docker
```
