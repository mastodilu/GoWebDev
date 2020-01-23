# Esempio: builda un container da ubuntu con cURL installato

## Dockerfile

    FROM ubuntu:devel
    RUN apt-get -y update && apt-get install -y curl

### Bulda il container

    docker build -t curler .

- `curler` è il nome assegnato a questa immagine
- `.` indica di buildare questo path

### Avvia il container

    docker run -it curler /bin/bash

- `-it` rende la shell interattiva
- `/bin/bash` esegue questo comando