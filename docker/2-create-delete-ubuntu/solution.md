Lösung
======

Erstellung eines Containers durch:
```bash
docker run -it ubuntu
```


In diesem kann nun `uname -s` ausgeführt werden:
```
root@cb95fe48e7fb:/# uname -s
Linux
```

Herausfinden der Container ID durch
```bash
❯ docker ps --all
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                      PORTS               NAMES
ba2a99c85e7d        ubuntu              "/bin/bash"         x minutes ago       Exited (0) x minutes ago                        admiring_visvesvaraya
```

und finale Löschung durch

```bash
docker rm admiring_visvesvaraya
```
